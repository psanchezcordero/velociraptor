// Code generated by protoc-gen-go. DO NOT EDIT.
// source: artifact.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	_ "www.velocidex.com/golang/velociraptor/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type FieldDescriptor struct {
	FriendlyName string   `protobuf:"bytes,1,opt,name=friendly_name,json=friendlyName,proto3" json:"friendly_name,omitempty"`
	Name         string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Repeated     bool     `protobuf:"varint,3,opt,name=repeated,proto3" json:"repeated,omitempty"`
	Type         string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Doc          string   `protobuf:"bytes,5,opt,name=doc,proto3" json:"doc,omitempty"`
	Labels       []string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty"`
	Default      string   `protobuf:"bytes,7,opt,name=default,proto3" json:"default,omitempty"`
	// Same as doc field.
	Description          string   `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldDescriptor) Reset()         { *m = FieldDescriptor{} }
func (m *FieldDescriptor) String() string { return proto.CompactTextString(m) }
func (*FieldDescriptor) ProtoMessage()    {}
func (*FieldDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1932e98ed811590, []int{0}
}

func (m *FieldDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldDescriptor.Unmarshal(m, b)
}
func (m *FieldDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldDescriptor.Marshal(b, m, deterministic)
}
func (m *FieldDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldDescriptor.Merge(m, src)
}
func (m *FieldDescriptor) XXX_Size() int {
	return xxx_messageInfo_FieldDescriptor.Size(m)
}
func (m *FieldDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_FieldDescriptor proto.InternalMessageInfo

func (m *FieldDescriptor) GetFriendlyName() string {
	if m != nil {
		return m.FriendlyName
	}
	return ""
}

func (m *FieldDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FieldDescriptor) GetRepeated() bool {
	if m != nil {
		return m.Repeated
	}
	return false
}

func (m *FieldDescriptor) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *FieldDescriptor) GetDoc() string {
	if m != nil {
		return m.Doc
	}
	return ""
}

func (m *FieldDescriptor) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *FieldDescriptor) GetDefault() string {
	if m != nil {
		return m.Default
	}
	return ""
}

func (m *FieldDescriptor) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type EnumValue struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnumValue) Reset()         { *m = EnumValue{} }
func (m *EnumValue) String() string { return proto.CompactTextString(m) }
func (*EnumValue) ProtoMessage()    {}
func (*EnumValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1932e98ed811590, []int{1}
}

func (m *EnumValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnumValue.Unmarshal(m, b)
}
func (m *EnumValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnumValue.Marshal(b, m, deterministic)
}
func (m *EnumValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnumValue.Merge(m, src)
}
func (m *EnumValue) XXX_Size() int {
	return xxx_messageInfo_EnumValue.Size(m)
}
func (m *EnumValue) XXX_DiscardUnknown() {
	xxx_messageInfo_EnumValue.DiscardUnknown(m)
}

var xxx_messageInfo_EnumValue proto.InternalMessageInfo

func (m *EnumValue) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *EnumValue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TypeDescriptor struct {
	Doc string `protobuf:"bytes,1,opt,name=doc,proto3" json:"doc,omitempty"`
	// Same as doc field.
	Description  string             `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty"`
	Fields       []*FieldDescriptor `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
	Name         string             `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	FriendlyName string             `protobuf:"bytes,7,opt,name=friendly_name,json=friendlyName,proto3" json:"friendly_name,omitempty"`
	Kind         string             `protobuf:"bytes,4,opt,name=kind,proto3" json:"kind,omitempty"`
	// The fields are all part of a single one of structure. NOTE:
	// Currently we only support an all or nothing approach to oneof -
	// there can be at most a single oneof group within the proto and
	// it implies that all the fields belong to it.
	Oneof                bool         `protobuf:"varint,5,opt,name=oneof,proto3" json:"oneof,omitempty"`
	Default              string       `protobuf:"bytes,6,opt,name=default,proto3" json:"default,omitempty"`
	AllowedValues        []*EnumValue `protobuf:"bytes,8,rep,name=allowed_values,json=allowedValues,proto3" json:"allowed_values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TypeDescriptor) Reset()         { *m = TypeDescriptor{} }
func (m *TypeDescriptor) String() string { return proto.CompactTextString(m) }
func (*TypeDescriptor) ProtoMessage()    {}
func (*TypeDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1932e98ed811590, []int{2}
}

func (m *TypeDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeDescriptor.Unmarshal(m, b)
}
func (m *TypeDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeDescriptor.Marshal(b, m, deterministic)
}
func (m *TypeDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeDescriptor.Merge(m, src)
}
func (m *TypeDescriptor) XXX_Size() int {
	return xxx_messageInfo_TypeDescriptor.Size(m)
}
func (m *TypeDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_TypeDescriptor proto.InternalMessageInfo

func (m *TypeDescriptor) GetDoc() string {
	if m != nil {
		return m.Doc
	}
	return ""
}

func (m *TypeDescriptor) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *TypeDescriptor) GetFields() []*FieldDescriptor {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *TypeDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TypeDescriptor) GetFriendlyName() string {
	if m != nil {
		return m.FriendlyName
	}
	return ""
}

func (m *TypeDescriptor) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *TypeDescriptor) GetOneof() bool {
	if m != nil {
		return m.Oneof
	}
	return false
}

func (m *TypeDescriptor) GetDefault() string {
	if m != nil {
		return m.Default
	}
	return ""
}

func (m *TypeDescriptor) GetAllowedValues() []*EnumValue {
	if m != nil {
		return m.AllowedValues
	}
	return nil
}

type Types struct {
	Items                []*TypeDescriptor `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Types) Reset()         { *m = Types{} }
func (m *Types) String() string { return proto.CompactTextString(m) }
func (*Types) ProtoMessage()    {}
func (*Types) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1932e98ed811590, []int{3}
}

func (m *Types) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Types.Unmarshal(m, b)
}
func (m *Types) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Types.Marshal(b, m, deterministic)
}
func (m *Types) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Types.Merge(m, src)
}
func (m *Types) XXX_Size() int {
	return xxx_messageInfo_Types.Size(m)
}
func (m *Types) XXX_DiscardUnknown() {
	xxx_messageInfo_Types.DiscardUnknown(m)
}

var xxx_messageInfo_Types proto.InternalMessageInfo

func (m *Types) GetItems() []*TypeDescriptor {
	if m != nil {
		return m.Items
	}
	return nil
}

type ArtifactParameter struct {
	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Default     string `protobuf:"bytes,2,opt,name=default,proto3" json:"default,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Type        string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	// For type = choice
	Choices              []string `protobuf:"bytes,6,rep,name=choices,proto3" json:"choices,omitempty"`
	FriendlyName         string   `protobuf:"bytes,5,opt,name=friendly_name,json=friendlyName,proto3" json:"friendly_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArtifactParameter) Reset()         { *m = ArtifactParameter{} }
func (m *ArtifactParameter) String() string { return proto.CompactTextString(m) }
func (*ArtifactParameter) ProtoMessage()    {}
func (*ArtifactParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1932e98ed811590, []int{4}
}

func (m *ArtifactParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactParameter.Unmarshal(m, b)
}
func (m *ArtifactParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactParameter.Marshal(b, m, deterministic)
}
func (m *ArtifactParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactParameter.Merge(m, src)
}
func (m *ArtifactParameter) XXX_Size() int {
	return xxx_messageInfo_ArtifactParameter.Size(m)
}
func (m *ArtifactParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactParameter.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactParameter proto.InternalMessageInfo

func (m *ArtifactParameter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ArtifactParameter) GetDefault() string {
	if m != nil {
		return m.Default
	}
	return ""
}

func (m *ArtifactParameter) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ArtifactParameter) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ArtifactParameter) GetChoices() []string {
	if m != nil {
		return m.Choices
	}
	return nil
}

func (m *ArtifactParameter) GetFriendlyName() string {
	if m != nil {
		return m.FriendlyName
	}
	return ""
}

type ArtifactSource struct {
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Precondition         string   `protobuf:"bytes,1,opt,name=precondition,proto3" json:"precondition,omitempty"`
	Query                string   `protobuf:"bytes,6,opt,name=query,proto3" json:"query,omitempty"`
	Queries              []string `protobuf:"bytes,2,rep,name=queries,proto3" json:"queries,omitempty"`
	PostProcess          []string `protobuf:"bytes,5,rep,name=post_process,json=postProcess,proto3" json:"post_process,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArtifactSource) Reset()         { *m = ArtifactSource{} }
func (m *ArtifactSource) String() string { return proto.CompactTextString(m) }
func (*ArtifactSource) ProtoMessage()    {}
func (*ArtifactSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1932e98ed811590, []int{5}
}

func (m *ArtifactSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactSource.Unmarshal(m, b)
}
func (m *ArtifactSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactSource.Marshal(b, m, deterministic)
}
func (m *ArtifactSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactSource.Merge(m, src)
}
func (m *ArtifactSource) XXX_Size() int {
	return xxx_messageInfo_ArtifactSource.Size(m)
}
func (m *ArtifactSource) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactSource.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactSource proto.InternalMessageInfo

func (m *ArtifactSource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ArtifactSource) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ArtifactSource) GetPrecondition() string {
	if m != nil {
		return m.Precondition
	}
	return ""
}

func (m *ArtifactSource) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *ArtifactSource) GetQueries() []string {
	if m != nil {
		return m.Queries
	}
	return nil
}

func (m *ArtifactSource) GetPostProcess() []string {
	if m != nil {
		return m.PostProcess
	}
	return nil
}

type Report struct {
	// Each report type will be handled differently. Read about the
	// different types in reporting.go
	Type                 string               `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Template             string               `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
	Parameters           []*ArtifactParameter `protobuf:"bytes,3,rep,name=parameters,proto3" json:"parameters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Report) Reset()         { *m = Report{} }
func (m *Report) String() string { return proto.CompactTextString(m) }
func (*Report) ProtoMessage()    {}
func (*Report) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1932e98ed811590, []int{6}
}

func (m *Report) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Report.Unmarshal(m, b)
}
func (m *Report) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Report.Marshal(b, m, deterministic)
}
func (m *Report) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Report.Merge(m, src)
}
func (m *Report) XXX_Size() int {
	return xxx_messageInfo_Report.Size(m)
}
func (m *Report) XXX_DiscardUnknown() {
	xxx_messageInfo_Report.DiscardUnknown(m)
}

var xxx_messageInfo_Report proto.InternalMessageInfo

func (m *Report) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Report) GetTemplate() string {
	if m != nil {
		return m.Template
	}
	return ""
}

func (m *Report) GetParameters() []*ArtifactParameter {
	if m != nil {
		return m.Parameters
	}
	return nil
}

type Artifact struct {
	Name                string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description         string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Author              string   `protobuf:"bytes,12,opt,name=author,proto3" json:"author,omitempty"`
	Reference           []string `protobuf:"bytes,5,rep,name=reference,proto3" json:"reference,omitempty"`
	RequiredPermissions []string `protobuf:"bytes,13,rep,name=required_permissions,json=requiredPermissions,proto3" json:"required_permissions,omitempty"`
	// If set here this applies to all sources.
	Precondition string               `protobuf:"bytes,8,opt,name=precondition,proto3" json:"precondition,omitempty"`
	Parameters   []*ArtifactParameter `protobuf:"bytes,3,rep,name=parameters,proto3" json:"parameters,omitempty"`
	// Unfortunately we can not use enum due to limitations in
	// yaml->json->protobuf conversion.
	Type     string            `protobuf:"bytes,10,opt,name=type,proto3" json:"type,omitempty"`
	Sources  []*ArtifactSource `protobuf:"bytes,4,rep,name=sources,proto3" json:"sources,omitempty"`
	Includes []string          `protobuf:"bytes,9,rep,name=includes,proto3" json:"includes,omitempty"`
	Reports  []*Report         `protobuf:"bytes,11,rep,name=reports,proto3" json:"reports,omitempty"`
	// Internal use only
	Raw                  string   `protobuf:"bytes,7,opt,name=raw,proto3" json:"raw,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Artifact) Reset()         { *m = Artifact{} }
func (m *Artifact) String() string { return proto.CompactTextString(m) }
func (*Artifact) ProtoMessage()    {}
func (*Artifact) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1932e98ed811590, []int{7}
}

func (m *Artifact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Artifact.Unmarshal(m, b)
}
func (m *Artifact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Artifact.Marshal(b, m, deterministic)
}
func (m *Artifact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Artifact.Merge(m, src)
}
func (m *Artifact) XXX_Size() int {
	return xxx_messageInfo_Artifact.Size(m)
}
func (m *Artifact) XXX_DiscardUnknown() {
	xxx_messageInfo_Artifact.DiscardUnknown(m)
}

var xxx_messageInfo_Artifact proto.InternalMessageInfo

func (m *Artifact) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Artifact) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Artifact) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Artifact) GetReference() []string {
	if m != nil {
		return m.Reference
	}
	return nil
}

func (m *Artifact) GetRequiredPermissions() []string {
	if m != nil {
		return m.RequiredPermissions
	}
	return nil
}

func (m *Artifact) GetPrecondition() string {
	if m != nil {
		return m.Precondition
	}
	return ""
}

func (m *Artifact) GetParameters() []*ArtifactParameter {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *Artifact) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Artifact) GetSources() []*ArtifactSource {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *Artifact) GetIncludes() []string {
	if m != nil {
		return m.Includes
	}
	return nil
}

func (m *Artifact) GetReports() []*Report {
	if m != nil {
		return m.Reports
	}
	return nil
}

func (m *Artifact) GetRaw() string {
	if m != nil {
		return m.Raw
	}
	return ""
}

type ArtifactDescriptors struct {
	Items                []*Artifact `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ArtifactDescriptors) Reset()         { *m = ArtifactDescriptors{} }
func (m *ArtifactDescriptors) String() string { return proto.CompactTextString(m) }
func (*ArtifactDescriptors) ProtoMessage()    {}
func (*ArtifactDescriptors) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1932e98ed811590, []int{8}
}

func (m *ArtifactDescriptors) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactDescriptors.Unmarshal(m, b)
}
func (m *ArtifactDescriptors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactDescriptors.Marshal(b, m, deterministic)
}
func (m *ArtifactDescriptors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactDescriptors.Merge(m, src)
}
func (m *ArtifactDescriptors) XXX_Size() int {
	return xxx_messageInfo_ArtifactDescriptors.Size(m)
}
func (m *ArtifactDescriptors) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactDescriptors.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactDescriptors proto.InternalMessageInfo

func (m *ArtifactDescriptors) GetItems() []*Artifact {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*FieldDescriptor)(nil), "proto.FieldDescriptor")
	proto.RegisterType((*EnumValue)(nil), "proto.EnumValue")
	proto.RegisterType((*TypeDescriptor)(nil), "proto.TypeDescriptor")
	proto.RegisterType((*Types)(nil), "proto.Types")
	proto.RegisterType((*ArtifactParameter)(nil), "proto.ArtifactParameter")
	proto.RegisterType((*ArtifactSource)(nil), "proto.ArtifactSource")
	proto.RegisterType((*Report)(nil), "proto.Report")
	proto.RegisterType((*Artifact)(nil), "proto.Artifact")
	proto.RegisterType((*ArtifactDescriptors)(nil), "proto.ArtifactDescriptors")
}

func init() { proto.RegisterFile("artifact.proto", fileDescriptor_a1932e98ed811590) }

var fileDescriptor_a1932e98ed811590 = []byte{
	// 1681 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x4f, 0x6f, 0x1c, 0x49,
	0x15, 0x57, 0x7b, 0x32, 0xf6, 0xb8, 0xf2, 0x67, 0x97, 0xda, 0x25, 0x34, 0x39, 0x2c, 0x8f, 0x89,
	0x00, 0x07, 0x9c, 0xb6, 0xd8, 0x6c, 0xb4, 0xc1, 0xac, 0x10, 0x33, 0x89, 0xd9, 0xf5, 0xca, 0x99,
	0x71, 0xda, 0x56, 0xa2, 0xdd, 0x8b, 0x55, 0xee, 0x7e, 0x33, 0x53, 0xa4, 0xba, 0xaa, 0x53, 0x55,
	0xed, 0xf1, 0x70, 0xe1, 0x80, 0x84, 0xb8, 0x44, 0x42, 0x20, 0x0e, 0x20, 0xbe, 0x03, 0xe2, 0x03,
	0x70, 0xe3, 0x33, 0x20, 0xc1, 0x69, 0x81, 0xaf, 0xc1, 0x01, 0x55, 0x75, 0xf7, 0x74, 0x8f, 0x27,
	0x02, 0x71, 0xe0, 0x34, 0x53, 0xd5, 0xaf, 0xde, 0xfb, 0xbd, 0xf7, 0x7e, 0xef, 0x57, 0x45, 0x6e,
	0x31, 0x6d, 0xf9, 0x84, 0x25, 0x36, 0xca, 0xb5, 0xb2, 0x8a, 0x76, 0xfd, 0xcf, 0x9d, 0xfd, 0xf9,
	0x7c, 0x1e, 0x5d, 0xa0, 0x50, 0x09, 0x4f, 0xf1, 0x32, 0x4a, 0x54, 0xb6, 0x37, 0x55, 0x82, 0xc9,
	0xe9, 0x5e, 0xb9, 0xa9, 0x59, 0x6e, 0x95, 0xde, 0xf3, 0xc6, 0x7b, 0x06, 0x33, 0x26, 0x2d, 0x4f,
	0x4a, 0x17, 0xfd, 0x2f, 0x02, 0xf2, 0xd6, 0x8f, 0x38, 0x8a, 0xf4, 0x09, 0x9a, 0x44, 0x73, 0x67,
	0x48, 0xef, 0x92, 0x9b, 0x13, 0xcd, 0x51, 0xa6, 0x62, 0x71, 0x26, 0x59, 0x86, 0x61, 0x00, 0xc1,
	0xce, 0x76, 0x7c, 0xa3, 0xde, 0x1c, 0xb1, 0x0c, 0x29, 0x25, 0xd7, 0xfc, 0xb7, 0x0d, 0xff, 0xcd,
	0xff, 0xa7, 0x77, 0x48, 0x4f, 0x63, 0x8e, 0xcc, 0x62, 0x1a, 0x76, 0x20, 0xd8, 0xe9, 0xc5, 0xcb,
	0xb5, 0xb3, 0xb7, 0x8b, 0x1c, 0xc3, 0x6b, 0xa5, 0xbd, 0xfb, 0x4f, 0xdf, 0x26, 0x9d, 0x54, 0x25,
	0x61, 0xd7, 0x6f, 0xb9, 0xbf, 0xf4, 0x36, 0xd9, 0x14, 0xec, 0x1c, 0x85, 0x09, 0x37, 0xa1, 0xb3,
	0xb3, 0x1d, 0x57, 0x2b, 0x1a, 0x92, 0xad, 0x14, 0x27, 0xac, 0x10, 0x36, 0xdc, 0xf2, 0xd6, 0xf5,
	0x92, 0x02, 0xb9, 0x9e, 0x56, 0xd0, 0xb9, 0x92, 0x61, 0xcf, 0x7f, 0x6d, 0x6f, 0xf5, 0x1f, 0x92,
	0xed, 0x03, 0x59, 0x64, 0xcf, 0x99, 0x28, 0x90, 0xbe, 0x4b, 0xba, 0x17, 0xee, 0x8f, 0xcf, 0xa9,
	0x1b, 0x97, 0x8b, 0x37, 0x25, 0xd3, 0xff, 0xfb, 0x06, 0xb9, 0x75, 0xba, 0xc8, 0xb1, 0x55, 0x98,
	0x0a, 0x6f, 0xd0, 0xe0, 0xbd, 0x12, 0x7d, 0x7b, 0x2d, 0x3a, 0x8d, 0xc8, 0xe6, 0xc4, 0xd5, 0xd7,
	0x84, 0x1b, 0xd0, 0xd9, 0xb9, 0xfe, 0xfe, 0xed, 0xb2, 0xf0, 0xd1, 0x95, 0xa2, 0xc7, 0x95, 0xd5,
	0x12, 0x4a, 0xa7, 0x55, 0xd7, 0xb5, 0x86, 0x6c, 0xbd, 0xb9, 0x21, 0x2f, 0xb9, 0x4c, 0xeb, 0x02,
	0xbb, 0xff, 0x2e, 0x5b, 0x25, 0x51, 0x4d, 0x7c, 0x89, 0x7b, 0x71, 0xb9, 0xa0, 0xe3, 0xa6, 0x98,
	0x9b, 0xce, 0x78, 0xf8, 0xf0, 0x1f, 0xff, 0xfa, 0xe7, 0x9f, 0x83, 0x3d, 0x7a, 0xff, 0x74, 0x86,
	0xf0, 0x63, 0xa3, 0x24, 0xa0, 0x4c, 0x54, 0x8a, 0x29, 0x54, 0x76, 0xe0, 0xcb, 0x04, 0x13, 0xa5,
	0xc1, 0xce, 0xb8, 0x01, 0xd7, 0xbd, 0xa8, 0xe9, 0xc1, 0x87, 0xe4, 0x16, 0x13, 0x42, 0xcd, 0x31,
	0x3d, 0xf3, 0x86, 0x26, 0xec, 0xf9, 0x5c, 0xdf, 0xae, 0x72, 0x5d, 0x96, 0x3f, 0xbe, 0x59, 0xd9,
	0xf9, 0x95, 0xe9, 0x7f, 0x40, 0xba, 0xae, 0xc4, 0x86, 0x7e, 0x87, 0x74, 0xb9, 0xc5, 0xcc, 0x84,
	0x81, 0x3f, 0xf8, 0xe5, 0xea, 0xe0, 0x6a, 0xfd, 0xe3, 0xd2, 0xa6, 0xff, 0xdb, 0x0e, 0xf9, 0xd2,
	0xa0, 0x9a, 0x84, 0x63, 0xa6, 0x59, 0x86, 0x16, 0xf5, 0xb2, 0x70, 0x41, 0xab, 0x70, 0x2d, 0xda,
	0x6c, 0xfc, 0x47, 0xda, 0x74, 0xd6, 0x1b, 0xf7, 0x26, 0xc2, 0x86, 0x64, 0x2b, 0x99, 0x29, 0x9e,
	0x60, 0xcd, 0xcf, 0x7a, 0xb9, 0xde, 0xa2, 0xee, 0x7a, 0x8b, 0xf6, 0xbf, 0x08, 0xfe, 0xe6, 0x0a,
	0xfd, 0xd7, 0x80, 0xfc, 0x25, 0xa8, 0x13, 0x30, 0x90, 0xb1, 0x05, 0xb0, 0x24, 0xc1, 0xdc, 0x42,
	0x5e, 0x67, 0x63, 0x60, 0x3e, 0xe3, 0xc9, 0x0c, 0x98, 0x46, 0x60, 0xa9, 0xeb, 0x84, 0x55, 0x60,
	0x67, 0x08, 0x26, 0x51, 0x39, 0x42, 0xae, 0xb9, 0x6b, 0x85, 0x02, 0xbc, 0xc4, 0xa4, 0x70, 0x70,
	0x23, 0x18, 0x8d, 0x4f, 0x0f, 0xf6, 0x81, 0x09, 0xd1, 0xf6, 0xe2, 0xce, 0x1b, 0xab, 0xb9, 0x9c,
	0x1a, 0xb8, 0x0f, 0x7c, 0x02, 0x0b, 0x55, 0x80, 0x44, 0x4c, 0xc1, 0xa8, 0x0c, 0xed, 0x8c, 0xcb,
	0x29, 0xa0, 0x30, 0xe8, 0x7d, 0xbf, 0x2a, 0x50, 0x2f, 0x20, 0x61, 0x12, 0x0a, 0x99, 0xb3, 0xe4,
	0x25, 0x60, 0x34, 0x8d, 0x60, 0xa2, 0x55, 0x06, 0x9f, 0x9e, 0x8c, 0x47, 0x50, 0x18, 0x67, 0xee,
	0x2c, 0xdd, 0xf2, 0x98, 0x69, 0x83, 0x3b, 0xf7, 0xe0, 0xf9, 0xb3, 0x23, 0x98, 0x14, 0x32, 0xf1,
	0x28, 0xfa, 0xbf, 0xd8, 0x24, 0xb7, 0xea, 0xd4, 0x4e, 0x54, 0xa1, 0x13, 0xa4, 0x7f, 0x08, 0xda,
	0x94, 0x1e, 0xfe, 0x2e, 0xf0, 0x6c, 0xfb, 0x75, 0x40, 0x7f, 0x19, 0x38, 0xbe, 0xb9, 0x4f, 0xa0,
	0x26, 0x25, 0xb1, 0x6a, 0x81, 0x03, 0xe3, 0xcf, 0x46, 0x70, 0x38, 0x01, 0xa9, 0x2c, 0x18, 0xb4,
	0x30, 0x47, 0x28, 0x2a, 0xa4, 0x66, 0x79, 0x06, 0x9b, 0x23, 0xdc, 0x1a, 0x14, 0x93, 0x08, 0x4e,
	0xdb, 0x9b, 0x89, 0xca, 0x72, 0x2e, 0x50, 0xc3, 0x9c, 0x0b, 0x01, 0x53, 0x94, 0xa8, 0x99, 0x45,
	0x60, 0x55, 0xb2, 0x73, 0x6e, 0x67, 0x65, 0x64, 0x07, 0x23, 0xaa, 0x68, 0xf3, 0x3a, 0x58, 0x65,
	0x87, 0xa7, 0xc0, 0xf0, 0xa5, 0xc7, 0x8d, 0x34, 0x19, 0x40, 0xeb, 0x63, 0x55, 0xde, 0x66, 0x36,
	0x6a, 0xe4, 0x23, 0x65, 0x11, 0xb8, 0xf5, 0xf5, 0x3c, 0x47, 0xe0, 0xd2, 0xa2, 0xce, 0x95, 0x70,
	0x62, 0x58, 0x86, 0x55, 0x76, 0x86, 0xba, 0x41, 0xda, 0xf2, 0x69, 0xa2, 0x55, 0x2a, 0xce, 0xc8,
	0x8d, 0x5c, 0x63, 0xa2, 0x64, 0xca, 0x3d, 0x1e, 0x4f, 0xf1, 0xe1, 0x13, 0x8f, 0xe7, 0x07, 0xf4,
	0xa3, 0x81, 0xef, 0x01, 0x5e, 0xe6, 0x1a, 0x8d, 0x71, 0x90, 0xac, 0x72, 0x11, 0xd1, 0xcd, 0xa2,
	0x0f, 0xb7, 0xa4, 0x4b, 0xdd, 0xc1, 0x06, 0x66, 0xbc, 0xe2, 0x99, 0x7e, 0x40, 0xba, 0xbe, 0x30,
	0x95, 0x30, 0xbc, 0xe7, 0x43, 0x84, 0xf4, 0xf6, 0x00, 0xb2, 0x42, 0x58, 0x7e, 0x5f, 0x70, 0x89,
	0x3e, 0x9a, 0xb7, 0x8a, 0x4b, 0x63, 0x6a, 0xc9, 0x96, 0xfb, 0xc3, 0xb1, 0x14, 0xb9, 0xed, 0xe1,
	0xe7, 0xfe, 0xdc, 0x29, 0x8d, 0x9f, 0x95, 0xdb, 0x60, 0x67, 0xcc, 0x96, 0x2d, 0xd0, 0x85, 0x04,
	0x2e, 0x41, 0xe9, 0x14, 0x75, 0x04, 0x63, 0x29, 0x16, 0xa0, 0x0a, 0x9b, 0x17, 0xb6, 0x24, 0x9a,
	0xeb, 0xa8, 0x60, 0xc6, 0x2e, 0xfb, 0x23, 0x84, 0xcb, 0x25, 0x51, 0x42, 0x60, 0x62, 0x31, 0x8d,
	0xe2, 0x3a, 0x14, 0x35, 0xe4, 0x46, 0xae, 0x8c, 0x3d, 0xcb, 0xb5, 0x4a, 0xd0, 0x98, 0xb0, 0xeb,
	0x43, 0x1f, 0xfb, 0xd0, 0x9f, 0xd2, 0x4f, 0x06, 0x20, 0xb8, 0xb1, 0x8e, 0x25, 0xaf, 0xd6, 0x40,
	0x9c, 0xa3, 0xc7, 0xa1, 0x2e, 0x50, 0xfb, 0x88, 0x1a, 0x4d, 0x21, 0xac, 0xf1, 0x4d, 0x74, 0x3e,
	0xa1, 0xf2, 0xc9, 0xe5, 0x34, 0x8a, 0xaf, 0xbb, 0x9d, 0xe3, 0x72, 0x63, 0xff, 0x9e, 0x9f, 0xe0,
	0xbb, 0xe4, 0xeb, 0x2f, 0x66, 0xa8, 0x71, 0x95, 0x80, 0x53, 0xb4, 0xc6, 0xb1, 0x10, 0x52, 0x66,
	0x59, 0xd4, 0xff, 0xd3, 0x06, 0xd9, 0x8c, 0x31, 0x57, 0xda, 0xd2, 0xa7, 0x95, 0x96, 0x94, 0x8d,
	0xfb, 0x9e, 0x87, 0xf8, 0x80, 0x7e, 0xd7, 0xc9, 0x9b, 0x03, 0xa8, 0xbd, 0xd5, 0x3e, 0x3c, 0x1d,
	0x8f, 0x0e, 0x4f, 0xc7, 0xf1, 0xe1, 0xe8, 0xe3, 0xb3, 0x27, 0x83, 0xc3, 0xa3, 0xcf, 0x76, 0xe1,
	0x78, 0x7c, 0x72, 0x7a, 0x76, 0x1c, 0x8f, 0x1f, 0x1f, 0x9c, 0x9c, 0x1c, 0x8e, 0x3e, 0xae, 0x64,
	0xe8, 0x0e, 0xe9, 0x59, 0xcc, 0x72, 0xc7, 0xa5, 0x4a, 0xd7, 0x96, 0x6b, 0xfa, 0x88, 0x90, 0x66,
	0xfe, 0xc3, 0x8e, 0x97, 0xd3, 0xb0, 0x92, 0xd3, 0x35, 0xd1, 0x8c, 0x5b, 0xb6, 0xfb, 0xaf, 0x4b,
	0x75, 0xfa, 0x79, 0x40, 0x7e, 0x16, 0x0c, 0x2a, 0x4c, 0xc0, 0xcd, 0x72, 0x68, 0xd2, 0xa6, 0x43,
	0x55, 0xc7, 0xae, 0x4e, 0x60, 0xd3, 0x25, 0x37, 0x84, 0x1a, 0x6b, 0xee, 0x7b, 0xee, 0xe4, 0x02,
	0x6b, 0xa7, 0x2e, 0x89, 0xb6, 0xdf, 0x14, 0x73, 0x94, 0xa9, 0xa3, 0xa6, 0x92, 0x90, 0x28, 0x69,
	0xf1, 0xd2, 0x46, 0xfd, 0xd7, 0x84, 0xf4, 0x6a, 0xc4, 0xf4, 0x8f, 0x41, 0x5b, 0xde, 0x87, 0xbf,
	0x2f, 0x45, 0xe4, 0x37, 0x01, 0xfd, 0xd5, 0x15, 0x11, 0x69, 0xe0, 0x44, 0x70, 0x32, 0x53, 0x85,
	0x48, 0x1d, 0x82, 0x42, 0xf2, 0x57, 0x05, 0x02, 0x93, 0xa9, 0x97, 0x5b, 0x17, 0x83, 0x71, 0x09,
	0xa9, 0xb2, 0x26, 0x82, 0x81, 0xd3, 0x95, 0x49, 0x21, 0xc0, 0x24, 0x33, 0xcc, 0xd0, 0xe5, 0xec,
	0x26, 0x48, 0x23, 0x7b, 0x09, 0x09, 0xb3, 0x38, 0x55, 0x9e, 0x3e, 0x7e, 0x68, 0x53, 0x65, 0x4b,
	0x59, 0x3c, 0xe2, 0xb2, 0xb8, 0x8c, 0x86, 0x5a, 0xcd, 0x0d, 0x6a, 0x13, 0x3d, 0x9e, 0x69, 0x95,
	0xe1, 0x27, 0xdc, 0x58, 0xa5, 0x17, 0x95, 0x8c, 0x3c, 0x5b, 0x55, 0x11, 0xdf, 0xa9, 0xe1, 0x9e,
	0x07, 0x7e, 0x8f, 0x7e, 0xeb, 0x85, 0x63, 0xe3, 0xaa, 0x80, 0x19, 0xb0, 0x7a, 0xe1, 0x27, 0x54,
	0xd5, 0xb5, 0xbc, 0xa2, 0x04, 0x1f, 0x92, 0x4d, 0x56, 0xd8, 0x99, 0xd2, 0xe1, 0x0d, 0xef, 0xed,
	0x6b, 0xde, 0xdb, 0x57, 0xe9, 0x57, 0x06, 0x7e, 0x77, 0xad, 0x04, 0x71, 0x65, 0x4e, 0x9f, 0x90,
	0x6d, 0x8d, 0x13, 0xd4, 0x28, 0x13, 0xac, 0x26, 0xe5, 0x9b, 0xfe, 0x2c, 0xd0, 0xf7, 0x5c, 0xb3,
	0xab, 0x4f, 0x8d, 0x8c, 0x35, 0x2e, 0x9a, 0x83, 0xd4, 0x90, 0x77, 0x35, 0xbe, 0x2a, 0xb8, 0xc6,
	0xf4, 0x2c, 0x47, 0x9d, 0x71, 0xaf, 0x37, 0x26, 0xbc, 0xe9, 0x1d, 0xfe, 0xd0, 0x3b, 0xdc, 0xa7,
	0x8f, 0x9a, 0xd1, 0xab, 0xad, 0xa1, 0x65, 0xdd, 0xca, 0xee, 0x6a, 0xa8, 0x77, 0x6a, 0xfb, 0xe3,
	0xc6, 0x7c, 0x4d, 0xfd, 0x7a, 0xff, 0x37, 0xf5, 0x13, 0xff, 0xcb, 0xec, 0x0c, 0xdf, 0xf7, 0x08,
	0x76, 0xe9, 0xb7, 0x8f, 0x9b, 0xfb, 0xb6, 0x8c, 0x9d, 0x6b, 0x75, 0xc1, 0x5b, 0x97, 0x76, 0x93,
	0x60, 0xcb, 0x3f, 0x1d, 0x57, 0xa2, 0x40, 0x7c, 0x3e, 0xdf, 0xf7, 0xde, 0x1e, 0xd2, 0x07, 0x8e,
	0xce, 0xb6, 0x12, 0x86, 0xe5, 0xad, 0x1c, 0xc1, 0xe3, 0x72, 0x94, 0x46, 0xe3, 0xf8, 0xe9, 0xe0,
	0x68, 0x17, 0x0e, 0x9e, 0x1f, 0x8c, 0x4e, 0x77, 0xe1, 0xe4, 0x20, 0x7e, 0x7e, 0x10, 0x57, 0xb2,
	0xf0, 0x39, 0xd9, 0x2a, 0xf3, 0x32, 0xe1, 0xb5, 0x95, 0x67, 0xd4, 0xea, 0x85, 0x3c, 0xbc, 0xe7,
	0x43, 0xdd, 0xa5, 0xff, 0x5d, 0xc3, 0xe2, 0xda, 0x21, 0x7d, 0x41, 0x7a, 0x5c, 0x26, 0xa2, 0x48,
	0xd1, 0x84, 0xdb, 0xbe, 0xdb, 0x4b, 0xc0, 0x4d, 0xb7, 0x57, 0x6f, 0x34, 0x5f, 0x8d, 0x0c, 0xf5,
	0xd4, 0xe9, 0x85, 0xae, 0x4a, 0xee, 0xa7, 0x28, 0x8a, 0x97, 0xce, 0x68, 0x46, 0xb6, 0x4a, 0x55,
	0x30, 0xe1, 0x75, 0x0f, 0xfa, 0x66, 0x05, 0xba, 0x94, 0xce, 0x76, 0x9f, 0x1b, 0x52, 0x79, 0x7b,
	0xe7, 0x3e, 0x57, 0x16, 0xa5, 0xe5, 0x4c, 0x88, 0xc5, 0x8a, 0x70, 0x5f, 0x25, 0x56, 0x1d, 0x83,
	0x3e, 0x22, 0x1d, 0xcd, 0xe6, 0xe5, 0x03, 0xba, 0x99, 0x00, 0x57, 0x73, 0xcd, 0xe6, 0xf0, 0xd9,
	0xe0, 0xe9, 0xd1, 0xda, 0x5b, 0x24, 0x8a, 0xdd, 0x91, 0xfd, 0x9f, 0x7a, 0x75, 0x5c, 0x90, 0xf9,
	0x40, 0x36, 0xe5, 0x9a, 0x6b, 0x96, 0x1b, 0x60, 0xcd, 0x95, 0xe8, 0xee, 0x36, 0x8d, 0x85, 0x61,
	0xe7, 0x02, 0x77, 0x21, 0x55, 0x49, 0x91, 0xa1, 0xf4, 0x57, 0x3e, 0x5b, 0x44, 0xcd, 0x9b, 0xcf,
	0xbf, 0xeb, 0x84, 0x00, 0x76, 0xae, 0x8a, 0xa5, 0x6e, 0x56, 0xec, 0x74, 0x6f, 0x35, 0xf7, 0xe6,
	0x61, 0x92, 0x89, 0xc5, 0x4f, 0xaa, 0x27, 0x57, 0x16, 0xf5, 0x3f, 0x22, 0xef, 0xd4, 0x0e, 0x9a,
	0x37, 0xb1, 0xa1, 0xdf, 0x58, 0x7d, 0x3a, 0xbf, 0x75, 0xa5, 0xe7, 0xd5, 0xa3, 0xf9, 0x7c, 0xd3,
	0x6f, 0x3f, 0xf8, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x45, 0x46, 0xad, 0x44, 0x0e, 0x00,
	0x00,
}
