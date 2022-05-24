// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: internal/conf/mvn.proto

package conf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Mvn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response       *Mvn_Response       `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	ResponseHeader *Mvn_ResponseHeader `protobuf:"bytes,2,opt,name=responseHeader,proto3" json:"responseHeader,omitempty"`
	Spellcheck     *Mvn_Spellcheck     `protobuf:"bytes,3,opt,name=spellcheck,proto3" json:"spellcheck,omitempty"`
}

func (x *Mvn) Reset() {
	*x = Mvn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_mvn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mvn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mvn) ProtoMessage() {}

func (x *Mvn) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_mvn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mvn.ProtoReflect.Descriptor instead.
func (*Mvn) Descriptor() ([]byte, []int) {
	return file_internal_conf_mvn_proto_rawDescGZIP(), []int{0}
}

func (x *Mvn) GetResponse() *Mvn_Response {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *Mvn) GetResponseHeader() *Mvn_ResponseHeader {
	if x != nil {
		return x.ResponseHeader
	}
	return nil
}

func (x *Mvn) GetSpellcheck() *Mvn_Spellcheck {
	if x != nil {
		return x.Spellcheck
	}
	return nil
}

type Mvn_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Docs     []*Mvn_Response_Doc `protobuf:"bytes,1,rep,name=docs,proto3" json:"docs,omitempty"`
	NumFound int32               `protobuf:"varint,2,opt,name=numFound,proto3" json:"numFound,omitempty"`
	Start    int32               `protobuf:"varint,3,opt,name=start,proto3" json:"start,omitempty"`
}

func (x *Mvn_Response) Reset() {
	*x = Mvn_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_mvn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mvn_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mvn_Response) ProtoMessage() {}

func (x *Mvn_Response) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_mvn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mvn_Response.ProtoReflect.Descriptor instead.
func (*Mvn_Response) Descriptor() ([]byte, []int) {
	return file_internal_conf_mvn_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Mvn_Response) GetDocs() []*Mvn_Response_Doc {
	if x != nil {
		return x.Docs
	}
	return nil
}

func (x *Mvn_Response) GetNumFound() int32 {
	if x != nil {
		return x.NumFound
	}
	return 0
}

func (x *Mvn_Response) GetStart() int32 {
	if x != nil {
		return x.Start
	}
	return 0
}

type Mvn_ResponseHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QTime  int32                      `protobuf:"varint,1,opt,name=qTime,proto3" json:"qTime,omitempty"`
	Params *Mvn_ResponseHeader_Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	Status int32                      `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Mvn_ResponseHeader) Reset() {
	*x = Mvn_ResponseHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_mvn_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mvn_ResponseHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mvn_ResponseHeader) ProtoMessage() {}

func (x *Mvn_ResponseHeader) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_mvn_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mvn_ResponseHeader.ProtoReflect.Descriptor instead.
func (*Mvn_ResponseHeader) Descriptor() ([]byte, []int) {
	return file_internal_conf_mvn_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Mvn_ResponseHeader) GetQTime() int32 {
	if x != nil {
		return x.QTime
	}
	return 0
}

func (x *Mvn_ResponseHeader) GetParams() *Mvn_ResponseHeader_Params {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *Mvn_ResponseHeader) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type Mvn_Spellcheck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Suggestions []string `protobuf:"bytes,1,rep,name=suggestions,proto3" json:"suggestions,omitempty"`
}

func (x *Mvn_Spellcheck) Reset() {
	*x = Mvn_Spellcheck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_mvn_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mvn_Spellcheck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mvn_Spellcheck) ProtoMessage() {}

func (x *Mvn_Spellcheck) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_mvn_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mvn_Spellcheck.ProtoReflect.Descriptor instead.
func (*Mvn_Spellcheck) Descriptor() ([]byte, []int) {
	return file_internal_conf_mvn_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Mvn_Spellcheck) GetSuggestions() []string {
	if x != nil {
		return x.Suggestions
	}
	return nil
}

type Mvn_Response_Doc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A             string   `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	Ec            []string `protobuf:"bytes,2,rep,name=ec,proto3" json:"ec,omitempty"`
	G             string   `protobuf:"bytes,3,opt,name=g,proto3" json:"g,omitempty"`
	Id            string   `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	LatestVersion string   `protobuf:"bytes,5,opt,name=latestVersion,proto3" json:"latestVersion,omitempty"`
	P             string   `protobuf:"bytes,6,opt,name=p,proto3" json:"p,omitempty"`
	RepositoryId  string   `protobuf:"bytes,7,opt,name=repositoryId,proto3" json:"repositoryId,omitempty"`
	Text          []string `protobuf:"bytes,8,rep,name=text,proto3" json:"text,omitempty"`
	Timestamp     int64    `protobuf:"varint,9,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	VersionCount  int32    `protobuf:"varint,10,opt,name=versionCount,proto3" json:"versionCount,omitempty"`
}

func (x *Mvn_Response_Doc) Reset() {
	*x = Mvn_Response_Doc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_mvn_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mvn_Response_Doc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mvn_Response_Doc) ProtoMessage() {}

func (x *Mvn_Response_Doc) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_mvn_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mvn_Response_Doc.ProtoReflect.Descriptor instead.
func (*Mvn_Response_Doc) Descriptor() ([]byte, []int) {
	return file_internal_conf_mvn_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Mvn_Response_Doc) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

func (x *Mvn_Response_Doc) GetEc() []string {
	if x != nil {
		return x.Ec
	}
	return nil
}

func (x *Mvn_Response_Doc) GetG() string {
	if x != nil {
		return x.G
	}
	return ""
}

func (x *Mvn_Response_Doc) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Mvn_Response_Doc) GetLatestVersion() string {
	if x != nil {
		return x.LatestVersion
	}
	return ""
}

func (x *Mvn_Response_Doc) GetP() string {
	if x != nil {
		return x.P
	}
	return ""
}

func (x *Mvn_Response_Doc) GetRepositoryId() string {
	if x != nil {
		return x.RepositoryId
	}
	return ""
}

func (x *Mvn_Response_Doc) GetText() []string {
	if x != nil {
		return x.Text
	}
	return nil
}

func (x *Mvn_Response_Doc) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Mvn_Response_Doc) GetVersionCount() int32 {
	if x != nil {
		return x.VersionCount
	}
	return 0
}

type Mvn_ResponseHeader_Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Core    string `protobuf:"bytes,1,opt,name=core,proto3" json:"core,omitempty"`
	DefType string `protobuf:"bytes,2,opt,name=defType,proto3" json:"defType,omitempty"`
	Fl      string `protobuf:"bytes,3,opt,name=fl,proto3" json:"fl,omitempty"`
	Indent  string `protobuf:"bytes,4,opt,name=indent,proto3" json:"indent,omitempty"`
	Q       string `protobuf:"bytes,5,opt,name=q,proto3" json:"q,omitempty"`
	Qf      string `protobuf:"bytes,6,opt,name=qf,proto3" json:"qf,omitempty"`
	Rows    string `protobuf:"bytes,7,opt,name=rows,proto3" json:"rows,omitempty"`
	Sort    string `protobuf:"bytes,8,opt,name=sort,proto3" json:"sort,omitempty"`
	Start   string `protobuf:"bytes,11,opt,name=start,proto3" json:"start,omitempty"`
	Version string `protobuf:"bytes,12,opt,name=version,proto3" json:"version,omitempty"`
	Wt      string `protobuf:"bytes,13,opt,name=wt,proto3" json:"wt,omitempty"`
}

func (x *Mvn_ResponseHeader_Params) Reset() {
	*x = Mvn_ResponseHeader_Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_mvn_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mvn_ResponseHeader_Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mvn_ResponseHeader_Params) ProtoMessage() {}

func (x *Mvn_ResponseHeader_Params) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_mvn_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mvn_ResponseHeader_Params.ProtoReflect.Descriptor instead.
func (*Mvn_ResponseHeader_Params) Descriptor() ([]byte, []int) {
	return file_internal_conf_mvn_proto_rawDescGZIP(), []int{0, 1, 0}
}

func (x *Mvn_ResponseHeader_Params) GetCore() string {
	if x != nil {
		return x.Core
	}
	return ""
}

func (x *Mvn_ResponseHeader_Params) GetDefType() string {
	if x != nil {
		return x.DefType
	}
	return ""
}

func (x *Mvn_ResponseHeader_Params) GetFl() string {
	if x != nil {
		return x.Fl
	}
	return ""
}

func (x *Mvn_ResponseHeader_Params) GetIndent() string {
	if x != nil {
		return x.Indent
	}
	return ""
}

func (x *Mvn_ResponseHeader_Params) GetQ() string {
	if x != nil {
		return x.Q
	}
	return ""
}

func (x *Mvn_ResponseHeader_Params) GetQf() string {
	if x != nil {
		return x.Qf
	}
	return ""
}

func (x *Mvn_ResponseHeader_Params) GetRows() string {
	if x != nil {
		return x.Rows
	}
	return ""
}

func (x *Mvn_ResponseHeader_Params) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *Mvn_ResponseHeader_Params) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *Mvn_ResponseHeader_Params) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Mvn_ResponseHeader_Params) GetWt() string {
	if x != nil {
		return x.Wt
	}
	return ""
}

var File_internal_conf_mvn_proto protoreflect.FileDescriptor

var file_internal_conf_mvn_proto_rawDesc = []byte{
	0x0a, 0x17, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f,
	0x6d, 0x76, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f, 0x6e, 0x66, 0x22,
	0x9b, 0x07, 0x0a, 0x03, 0x4d, 0x76, 0x6e, 0x12, 0x2e, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x2e, 0x4d, 0x76, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x4d, 0x76, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x0a, 0x73, 0x70, 0x65,
	0x6c, 0x6c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x4d, 0x76, 0x6e, 0x2e, 0x53, 0x70, 0x65, 0x6c, 0x6c, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x0a, 0x73, 0x70, 0x65, 0x6c, 0x6c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x1a,
	0xda, 0x02, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x04,
	0x64, 0x6f, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x2e, 0x4d, 0x76, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44,
	0x6f, 0x63, 0x52, 0x04, 0x64, 0x6f, 0x63, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x75, 0x6d, 0x46,
	0x6f, 0x75, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6e, 0x75, 0x6d, 0x46,
	0x6f, 0x75, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x1a, 0xef, 0x01, 0x0a, 0x03, 0x44,
	0x6f, 0x63, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x61,
	0x12, 0x0e, 0x0a, 0x02, 0x65, 0x63, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x65, 0x63,
	0x12, 0x0c, 0x0a, 0x01, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x67, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24,
	0x0a, 0x0d, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x01, 0x70, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x22, 0x0a, 0x0c, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0xde, 0x02, 0x0a,
	0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x71, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x71, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x4d, 0x76, 0x6e,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0xe4, 0x01, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x66, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x66, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x66, 0x6c, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x69, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x71, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x01, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x71, 0x66, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x71, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x77, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x77, 0x74, 0x1a, 0x2e, 0x0a,
	0x0a, 0x53, 0x70, 0x65, 0x6c, 0x6c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x73,
	0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x08, 0x5a,
	0x06, 0x2e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_conf_mvn_proto_rawDescOnce sync.Once
	file_internal_conf_mvn_proto_rawDescData = file_internal_conf_mvn_proto_rawDesc
)

func file_internal_conf_mvn_proto_rawDescGZIP() []byte {
	file_internal_conf_mvn_proto_rawDescOnce.Do(func() {
		file_internal_conf_mvn_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_conf_mvn_proto_rawDescData)
	})
	return file_internal_conf_mvn_proto_rawDescData
}

var file_internal_conf_mvn_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_internal_conf_mvn_proto_goTypes = []interface{}{
	(*Mvn)(nil),                       // 0: conf.Mvn
	(*Mvn_Response)(nil),              // 1: conf.Mvn.Response
	(*Mvn_ResponseHeader)(nil),        // 2: conf.Mvn.ResponseHeader
	(*Mvn_Spellcheck)(nil),            // 3: conf.Mvn.Spellcheck
	(*Mvn_Response_Doc)(nil),          // 4: conf.Mvn.Response.Doc
	(*Mvn_ResponseHeader_Params)(nil), // 5: conf.Mvn.ResponseHeader.Params
}
var file_internal_conf_mvn_proto_depIdxs = []int32{
	1, // 0: conf.Mvn.response:type_name -> conf.Mvn.Response
	2, // 1: conf.Mvn.responseHeader:type_name -> conf.Mvn.ResponseHeader
	3, // 2: conf.Mvn.spellcheck:type_name -> conf.Mvn.Spellcheck
	4, // 3: conf.Mvn.Response.docs:type_name -> conf.Mvn.Response.Doc
	5, // 4: conf.Mvn.ResponseHeader.params:type_name -> conf.Mvn.ResponseHeader.Params
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_internal_conf_mvn_proto_init() }
func file_internal_conf_mvn_proto_init() {
	if File_internal_conf_mvn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_conf_mvn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mvn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_conf_mvn_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mvn_Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_conf_mvn_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mvn_ResponseHeader); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_conf_mvn_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mvn_Spellcheck); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_conf_mvn_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mvn_Response_Doc); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_conf_mvn_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mvn_ResponseHeader_Params); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_conf_mvn_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_conf_mvn_proto_goTypes,
		DependencyIndexes: file_internal_conf_mvn_proto_depIdxs,
		MessageInfos:      file_internal_conf_mvn_proto_msgTypes,
	}.Build()
	File_internal_conf_mvn_proto = out.File
	file_internal_conf_mvn_proto_rawDesc = nil
	file_internal_conf_mvn_proto_goTypes = nil
	file_internal_conf_mvn_proto_depIdxs = nil
}