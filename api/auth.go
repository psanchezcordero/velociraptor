/*
   Velociraptor - Hunting Evil
   Copyright (C) 2019 Velocidex Innovations.

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/
package api

import (
	"context"
	"encoding/json"
	"net/http"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"www.velocidex.com/golang/velociraptor/acls"
	api_proto "www.velocidex.com/golang/velociraptor/api/proto"
	"www.velocidex.com/golang/velociraptor/config"
	config_proto "www.velocidex.com/golang/velociraptor/config/proto"
	"www.velocidex.com/golang/velociraptor/users"
)

var (
	contextKeyUser = "USER"
)

func checkUserCredentialsHandler(
	config_obj *config_proto.Config,
	parent http.Handler) http.Handler {

	// We are supposed to do the oauth thing.
	if config.GoogleAuthEnabled(config_obj) {
		return authenticateOAUTHCookie(config_obj, parent)
	} else if config.SAMLEnabled(config_obj) {
		return authenticateSAML(config_obj, parent)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

		username, password, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		user_record, err := users.GetUser(config_obj, username)
		if err != nil {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		// Must have at least reader.
		perm, err := acls.CheckAccess(config_obj, username, acls.READ_RESULTS)
		if !perm || err != nil || user_record.Locked || user_record.Name != username ||
			!user_record.VerifyPassword(password) {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		// Checking is successfull - user authorized. Here we
		// build a token to pass to the underlying GRPC
		// service with metadata about the user.
		user_info := &api_proto.VelociraptorUser{
			Name: username,
		}

		// Must use json encoding because grpc can not handle
		// binary data in metadata.
		serialized, _ := json.Marshal(user_info)
		ctx := context.WithValue(
			r.Context(), "USER", string(serialized))

		// Need to call logging after auth so it can access
		// the USER value in the context.
		GetLoggingHandler(config_obj)(parent).ServeHTTP(
			w, r.WithContext(ctx))
	})
}

// TODO: Implement this properly.
func IsUserApprovedForClient(
	config_obj *config_proto.Config,
	md *metadata.MD,
	client_id string) bool {
	return true
}

func GetGRPCUserInfo(
	config_obj *config_proto.Config,
	ctx context.Context) *api_proto.VelociraptorUser {
	result := &api_proto.VelociraptorUser{}

	peer, ok := peer.FromContext(ctx)
	if ok {
		tlsInfo, ok := peer.AuthInfo.(credentials.TLSInfo)
		if ok {
			v := tlsInfo.State.PeerCertificates[0].Subject.CommonName

			// Calls from the gRPC gateway are allowed to
			// embed the authenticated web user in the
			// metadata. This allows the API gateway to
			// impersonate anyone - it must be trusted to
			// convert web side authentication to a valid
			// user name which it may pass in the call
			// context.
			if v == config_obj.API.PinnedGwName {
				md, ok := metadata.FromIncomingContext(ctx)
				if ok {
					userinfo := md.Get("USER")
					if len(userinfo) > 0 {
						data := []byte(userinfo[0])
						json.Unmarshal(data, result)
					}
				}
			}

			// Other callers will return the name on their
			// certificate.
			if result.Name == "" {
				result.Name = v
			}
		}
	}

	return result
}

func NewDefaultUserObject(config_obj *config_proto.Config) *api_proto.ApiGrrUser {
	result := &api_proto.ApiGrrUser{
		InterfaceTraits: &api_proto.ApiGrrUserInterfaceTraits{
			AuthUsingGoogle: config_obj.GUI.GoogleOauthClientId != "",
			Links:           []*api_proto.UILink{},
		},
		UserType: api_proto.ApiGrrUser_USER_TYPE_ADMIN,
	}

	for _, link := range config_obj.GUI.Links {
		result.InterfaceTraits.Links = append(result.InterfaceTraits.Links,
			&api_proto.UILink{Text: link.Text, Url: link.Url})
	}

	return result
}
