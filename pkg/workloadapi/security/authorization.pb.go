// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: workloadapi/security/authorization.proto

package security

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Scope int32

const (
	// ALL means that the authorization policy will be applied to all workloads
	// in the mesh (any namespace).
	Scope_GLOBAL Scope = 0
	// NAMESPACE means that the policy will only be applied to workloads in a
	// specific namespace.
	Scope_NAMESPACE Scope = 1
	// WORKLOAD_SELECTOR means that the policy will only be applied to specific
	// workloads that were selected by their labels.
	Scope_WORKLOAD_SELECTOR Scope = 2
)

// Enum value maps for Scope.
var (
	Scope_name = map[int32]string{
		0: "GLOBAL",
		1: "NAMESPACE",
		2: "WORKLOAD_SELECTOR",
	}
	Scope_value = map[string]int32{
		"GLOBAL":            0,
		"NAMESPACE":         1,
		"WORKLOAD_SELECTOR": 2,
	}
)

func (x Scope) Enum() *Scope {
	p := new(Scope)
	*p = x
	return p
}

func (x Scope) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Scope) Descriptor() protoreflect.EnumDescriptor {
	return file_workloadapi_security_authorization_proto_enumTypes[0].Descriptor()
}

func (Scope) Type() protoreflect.EnumType {
	return &file_workloadapi_security_authorization_proto_enumTypes[0]
}

func (x Scope) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Scope.Descriptor instead.
func (Scope) EnumDescriptor() ([]byte, []int) {
	return file_workloadapi_security_authorization_proto_rawDescGZIP(), []int{0}
}

type Action int32

const (
	// Allow the request if it matches with the rules.
	Action_ALLOW Action = 0
	// Deny the request if it matches with the rules.
	Action_DENY Action = 1
)

// Enum value maps for Action.
var (
	Action_name = map[int32]string{
		0: "ALLOW",
		1: "DENY",
	}
	Action_value = map[string]int32{
		"ALLOW": 0,
		"DENY":  1,
	}
)

func (x Action) Enum() *Action {
	p := new(Action)
	*p = x
	return p
}

func (x Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Action) Descriptor() protoreflect.EnumDescriptor {
	return file_workloadapi_security_authorization_proto_enumTypes[1].Descriptor()
}

func (Action) Type() protoreflect.EnumType {
	return &file_workloadapi_security_authorization_proto_enumTypes[1]
}

func (x Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Action.Descriptor instead.
func (Action) EnumDescriptor() ([]byte, []int) {
	return file_workloadapi_security_authorization_proto_rawDescGZIP(), []int{1}
}

type Authorization struct {
	state     protoimpl.MessageState `protogen:"open.v1"`
	Name      string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string                 `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Determine the scope of this RBAC policy.
	// If set to NAMESPACE, the 'namespace' field value will be used.
	Scope Scope `protobuf:"varint,3,opt,name=scope,proto3,enum=istio.security.Scope" json:"scope,omitempty"`
	// The action to take if the request is matched with the rules.
	// Default is ALLOW if not specified.
	Action Action `protobuf:"varint,4,opt,name=action,proto3,enum=istio.security.Action" json:"action,omitempty"`
	// Set of RBAC policy groups each containing its rules.
	// If at least one of the groups is matched the policy action will
	// take place.
	// Groups are OR-ed.
	Groups        []*Group `protobuf:"bytes,5,rep,name=groups,proto3" json:"groups,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Authorization) Reset() {
	*x = Authorization{}
	mi := &file_workloadapi_security_authorization_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Authorization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Authorization) ProtoMessage() {}

func (x *Authorization) ProtoReflect() protoreflect.Message {
	mi := &file_workloadapi_security_authorization_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Authorization.ProtoReflect.Descriptor instead.
func (*Authorization) Descriptor() ([]byte, []int) {
	return file_workloadapi_security_authorization_proto_rawDescGZIP(), []int{0}
}

func (x *Authorization) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Authorization) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Authorization) GetScope() Scope {
	if x != nil {
		return x.Scope
	}
	return Scope_GLOBAL
}

func (x *Authorization) GetAction() Action {
	if x != nil {
		return x.Action
	}
	return Action_ALLOW
}

func (x *Authorization) GetGroups() []*Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

type Group struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Rules are AND-ed
	// This is a generic form of the authz policy's to, from and when
	Rules         []*Rules `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Group) Reset() {
	*x = Group{}
	mi := &file_workloadapi_security_authorization_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_workloadapi_security_authorization_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_workloadapi_security_authorization_proto_rawDescGZIP(), []int{1}
}

func (x *Group) GetRules() []*Rules {
	if x != nil {
		return x.Rules
	}
	return nil
}

type Rules struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Conditions within a rule are AND-ed (e.g. ALL conditions must be true)
	Matches       []*Match `protobuf:"bytes,2,rep,name=matches,proto3" json:"matches,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Rules) Reset() {
	*x = Rules{}
	mi := &file_workloadapi_security_authorization_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Rules) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rules) ProtoMessage() {}

func (x *Rules) ProtoReflect() protoreflect.Message {
	mi := &file_workloadapi_security_authorization_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rules.ProtoReflect.Descriptor instead.
func (*Rules) Descriptor() ([]byte, []int) {
	return file_workloadapi_security_authorization_proto_rawDescGZIP(), []int{2}
}

func (x *Rules) GetMatches() []*Match {
	if x != nil {
		return x.Matches
	}
	return nil
}

type Match struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	Namespaces          []*StringMatch         `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	NotNamespaces       []*StringMatch         `protobuf:"bytes,2,rep,name=not_namespaces,json=notNamespaces,proto3" json:"not_namespaces,omitempty"`
	ServiceAccounts     []*ServiceAccountMatch `protobuf:"bytes,11,rep,name=service_accounts,json=serviceAccounts,proto3" json:"service_accounts,omitempty"`
	NotServiceAccounts  []*ServiceAccountMatch `protobuf:"bytes,12,rep,name=not_service_accounts,json=notServiceAccounts,proto3" json:"not_service_accounts,omitempty"`
	Principals          []*StringMatch         `protobuf:"bytes,3,rep,name=principals,proto3" json:"principals,omitempty"`
	NotPrincipals       []*StringMatch         `protobuf:"bytes,4,rep,name=not_principals,json=notPrincipals,proto3" json:"not_principals,omitempty"`
	SourceIps           []*Address             `protobuf:"bytes,5,rep,name=source_ips,json=sourceIps,proto3" json:"source_ips,omitempty"`
	NotSourceIps        []*Address             `protobuf:"bytes,6,rep,name=not_source_ips,json=notSourceIps,proto3" json:"not_source_ips,omitempty"`
	DestinationIps      []*Address             `protobuf:"bytes,7,rep,name=destination_ips,json=destinationIps,proto3" json:"destination_ips,omitempty"`
	NotDestinationIps   []*Address             `protobuf:"bytes,8,rep,name=not_destination_ips,json=notDestinationIps,proto3" json:"not_destination_ips,omitempty"`
	DestinationPorts    []uint32               `protobuf:"varint,9,rep,packed,name=destination_ports,json=destinationPorts,proto3" json:"destination_ports,omitempty"`
	NotDestinationPorts []uint32               `protobuf:"varint,10,rep,packed,name=not_destination_ports,json=notDestinationPorts,proto3" json:"not_destination_ports,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *Match) Reset() {
	*x = Match{}
	mi := &file_workloadapi_security_authorization_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Match) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Match) ProtoMessage() {}

func (x *Match) ProtoReflect() protoreflect.Message {
	mi := &file_workloadapi_security_authorization_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Match.ProtoReflect.Descriptor instead.
func (*Match) Descriptor() ([]byte, []int) {
	return file_workloadapi_security_authorization_proto_rawDescGZIP(), []int{3}
}

func (x *Match) GetNamespaces() []*StringMatch {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

func (x *Match) GetNotNamespaces() []*StringMatch {
	if x != nil {
		return x.NotNamespaces
	}
	return nil
}

func (x *Match) GetServiceAccounts() []*ServiceAccountMatch {
	if x != nil {
		return x.ServiceAccounts
	}
	return nil
}

func (x *Match) GetNotServiceAccounts() []*ServiceAccountMatch {
	if x != nil {
		return x.NotServiceAccounts
	}
	return nil
}

func (x *Match) GetPrincipals() []*StringMatch {
	if x != nil {
		return x.Principals
	}
	return nil
}

func (x *Match) GetNotPrincipals() []*StringMatch {
	if x != nil {
		return x.NotPrincipals
	}
	return nil
}

func (x *Match) GetSourceIps() []*Address {
	if x != nil {
		return x.SourceIps
	}
	return nil
}

func (x *Match) GetNotSourceIps() []*Address {
	if x != nil {
		return x.NotSourceIps
	}
	return nil
}

func (x *Match) GetDestinationIps() []*Address {
	if x != nil {
		return x.DestinationIps
	}
	return nil
}

func (x *Match) GetNotDestinationIps() []*Address {
	if x != nil {
		return x.NotDestinationIps
	}
	return nil
}

func (x *Match) GetDestinationPorts() []uint32 {
	if x != nil {
		return x.DestinationPorts
	}
	return nil
}

func (x *Match) GetNotDestinationPorts() []uint32 {
	if x != nil {
		return x.NotDestinationPorts
	}
	return nil
}

type Address struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       []byte                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Length        uint32                 `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Address) Reset() {
	*x = Address{}
	mi := &file_workloadapi_security_authorization_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_workloadapi_security_authorization_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_workloadapi_security_authorization_proto_rawDescGZIP(), []int{4}
}

func (x *Address) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Address) GetLength() uint32 {
	if x != nil {
		return x.Length
	}
	return 0
}

type ServiceAccountMatch struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Namespace      string                 `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ServiceAccount string                 `protobuf:"bytes,2,opt,name=serviceAccount,proto3" json:"serviceAccount,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ServiceAccountMatch) Reset() {
	*x = ServiceAccountMatch{}
	mi := &file_workloadapi_security_authorization_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServiceAccountMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceAccountMatch) ProtoMessage() {}

func (x *ServiceAccountMatch) ProtoReflect() protoreflect.Message {
	mi := &file_workloadapi_security_authorization_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceAccountMatch.ProtoReflect.Descriptor instead.
func (*ServiceAccountMatch) Descriptor() ([]byte, []int) {
	return file_workloadapi_security_authorization_proto_rawDescGZIP(), []int{5}
}

func (x *ServiceAccountMatch) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ServiceAccountMatch) GetServiceAccount() string {
	if x != nil {
		return x.ServiceAccount
	}
	return ""
}

type StringMatch struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to MatchType:
	//
	//	*StringMatch_Exact
	//	*StringMatch_Prefix
	//	*StringMatch_Suffix
	//	*StringMatch_Presence
	MatchType     isStringMatch_MatchType `protobuf_oneof:"match_type"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StringMatch) Reset() {
	*x = StringMatch{}
	mi := &file_workloadapi_security_authorization_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StringMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringMatch) ProtoMessage() {}

func (x *StringMatch) ProtoReflect() protoreflect.Message {
	mi := &file_workloadapi_security_authorization_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringMatch.ProtoReflect.Descriptor instead.
func (*StringMatch) Descriptor() ([]byte, []int) {
	return file_workloadapi_security_authorization_proto_rawDescGZIP(), []int{6}
}

func (x *StringMatch) GetMatchType() isStringMatch_MatchType {
	if x != nil {
		return x.MatchType
	}
	return nil
}

func (x *StringMatch) GetExact() string {
	if x != nil {
		if x, ok := x.MatchType.(*StringMatch_Exact); ok {
			return x.Exact
		}
	}
	return ""
}

func (x *StringMatch) GetPrefix() string {
	if x != nil {
		if x, ok := x.MatchType.(*StringMatch_Prefix); ok {
			return x.Prefix
		}
	}
	return ""
}

func (x *StringMatch) GetSuffix() string {
	if x != nil {
		if x, ok := x.MatchType.(*StringMatch_Suffix); ok {
			return x.Suffix
		}
	}
	return ""
}

func (x *StringMatch) GetPresence() *emptypb.Empty {
	if x != nil {
		if x, ok := x.MatchType.(*StringMatch_Presence); ok {
			return x.Presence
		}
	}
	return nil
}

type isStringMatch_MatchType interface {
	isStringMatch_MatchType()
}

type StringMatch_Exact struct {
	// exact string match
	Exact string `protobuf:"bytes,1,opt,name=exact,proto3,oneof"`
}

type StringMatch_Prefix struct {
	// prefix-based match
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3,oneof"`
}

type StringMatch_Suffix struct {
	// suffix-based match
	Suffix string `protobuf:"bytes,3,opt,name=suffix,proto3,oneof"`
}

type StringMatch_Presence struct {
	Presence *emptypb.Empty `protobuf:"bytes,4,opt,name=presence,proto3,oneof"`
}

func (*StringMatch_Exact) isStringMatch_MatchType() {}

func (*StringMatch_Prefix) isStringMatch_MatchType() {}

func (*StringMatch_Suffix) isStringMatch_MatchType() {}

func (*StringMatch_Presence) isStringMatch_MatchType() {}

var File_workloadapi_security_authorization_proto protoreflect.FileDescriptor

var file_workloadapi_security_authorization_proto_rawDesc = string([]byte{
	0x0a, 0x28, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x69, 0x73, 0x74, 0x69,
	0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x01, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x69, 0x73, 0x74,
	0x69, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x63, 0x6f, 0x70,
	0x65, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f,
	0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f,
	0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x34, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x2b, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x2e, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x38, 0x0a,
	0x05, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x07,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x22, 0x93, 0x06, 0x0a, 0x05, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x3b, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x42,
	0x0a, 0x0e, 0x6e, 0x6f, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x0d, 0x6e, 0x6f, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x73, 0x12, 0x4e, 0x0a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x69,
	0x73, 0x74, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x12, 0x55, 0x0a, 0x14, 0x6e, 0x6f, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x12, 0x6e, 0x6f, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x3b, 0x0a, 0x0a, 0x70, 0x72, 0x69,
	0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x6e,
	0x63, 0x69, 0x70, 0x61, 0x6c, 0x73, 0x12, 0x42, 0x0a, 0x0e, 0x6e, 0x6f, 0x74, 0x5f, 0x70, 0x72,
	0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x0d, 0x6e, 0x6f, 0x74,
	0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x73, 0x12, 0x36, 0x0a, 0x0a, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x70, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x70, 0x73, 0x12, 0x3d, 0x0a, 0x0e, 0x6e, 0x6f, 0x74, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x69, 0x70, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x69, 0x73, 0x74,
	0x69, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x0c, 0x6e, 0x6f, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x70,
	0x73, 0x12, 0x40, 0x0a, 0x0f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x70, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x69, 0x73, 0x74,
	0x69, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x70, 0x73, 0x12, 0x47, 0x0a, 0x13, 0x6e, 0x6f, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x70, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x11, 0x6e, 0x6f, 0x74, 0x44, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x70, 0x73, 0x12, 0x2b, 0x0a, 0x11,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x10, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x6e, 0x6f, 0x74,
	0x5f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x13, 0x6e, 0x6f, 0x74, 0x44, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x22, 0x3b, 0x0a,
	0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x5b, 0x0a, 0x13, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x26, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x9d, 0x01, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x16, 0x0a, 0x05, 0x65, 0x78, 0x61, 0x63, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x78, 0x61, 0x63, 0x74, 0x12,
	0x18, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x75, 0x66,
	0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x75, 0x66,
	0x66, 0x69, 0x78, 0x12, 0x34, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52,
	0x08, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2a, 0x39, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x70, 0x65,
	0x12, 0x0a, 0x0a, 0x06, 0x47, 0x4c, 0x4f, 0x42, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x4e, 0x41, 0x4d, 0x45, 0x53, 0x50, 0x41, 0x43, 0x45, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x57,
	0x4f, 0x52, 0x4b, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x4f, 0x52,
	0x10, 0x02, 0x2a, 0x1d, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x09, 0x0a, 0x05,
	0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x4e, 0x59, 0x10,
	0x01, 0x42, 0x1a, 0x5a, 0x18, 0x70, 0x6b, 0x67, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_workloadapi_security_authorization_proto_rawDescOnce sync.Once
	file_workloadapi_security_authorization_proto_rawDescData []byte
)

func file_workloadapi_security_authorization_proto_rawDescGZIP() []byte {
	file_workloadapi_security_authorization_proto_rawDescOnce.Do(func() {
		file_workloadapi_security_authorization_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_workloadapi_security_authorization_proto_rawDesc), len(file_workloadapi_security_authorization_proto_rawDesc)))
	})
	return file_workloadapi_security_authorization_proto_rawDescData
}

var file_workloadapi_security_authorization_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_workloadapi_security_authorization_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_workloadapi_security_authorization_proto_goTypes = []any{
	(Scope)(0),                  // 0: istio.security.Scope
	(Action)(0),                 // 1: istio.security.Action
	(*Authorization)(nil),       // 2: istio.security.Authorization
	(*Group)(nil),               // 3: istio.security.Group
	(*Rules)(nil),               // 4: istio.security.Rules
	(*Match)(nil),               // 5: istio.security.Match
	(*Address)(nil),             // 6: istio.security.Address
	(*ServiceAccountMatch)(nil), // 7: istio.security.ServiceAccountMatch
	(*StringMatch)(nil),         // 8: istio.security.StringMatch
	(*emptypb.Empty)(nil),       // 9: google.protobuf.Empty
}
var file_workloadapi_security_authorization_proto_depIdxs = []int32{
	0,  // 0: istio.security.Authorization.scope:type_name -> istio.security.Scope
	1,  // 1: istio.security.Authorization.action:type_name -> istio.security.Action
	3,  // 2: istio.security.Authorization.groups:type_name -> istio.security.Group
	4,  // 3: istio.security.Group.rules:type_name -> istio.security.Rules
	5,  // 4: istio.security.Rules.matches:type_name -> istio.security.Match
	8,  // 5: istio.security.Match.namespaces:type_name -> istio.security.StringMatch
	8,  // 6: istio.security.Match.not_namespaces:type_name -> istio.security.StringMatch
	7,  // 7: istio.security.Match.service_accounts:type_name -> istio.security.ServiceAccountMatch
	7,  // 8: istio.security.Match.not_service_accounts:type_name -> istio.security.ServiceAccountMatch
	8,  // 9: istio.security.Match.principals:type_name -> istio.security.StringMatch
	8,  // 10: istio.security.Match.not_principals:type_name -> istio.security.StringMatch
	6,  // 11: istio.security.Match.source_ips:type_name -> istio.security.Address
	6,  // 12: istio.security.Match.not_source_ips:type_name -> istio.security.Address
	6,  // 13: istio.security.Match.destination_ips:type_name -> istio.security.Address
	6,  // 14: istio.security.Match.not_destination_ips:type_name -> istio.security.Address
	9,  // 15: istio.security.StringMatch.presence:type_name -> google.protobuf.Empty
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_workloadapi_security_authorization_proto_init() }
func file_workloadapi_security_authorization_proto_init() {
	if File_workloadapi_security_authorization_proto != nil {
		return
	}
	file_workloadapi_security_authorization_proto_msgTypes[6].OneofWrappers = []any{
		(*StringMatch_Exact)(nil),
		(*StringMatch_Prefix)(nil),
		(*StringMatch_Suffix)(nil),
		(*StringMatch_Presence)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_workloadapi_security_authorization_proto_rawDesc), len(file_workloadapi_security_authorization_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_workloadapi_security_authorization_proto_goTypes,
		DependencyIndexes: file_workloadapi_security_authorization_proto_depIdxs,
		EnumInfos:         file_workloadapi_security_authorization_proto_enumTypes,
		MessageInfos:      file_workloadapi_security_authorization_proto_msgTypes,
	}.Build()
	File_workloadapi_security_authorization_proto = out.File
	file_workloadapi_security_authorization_proto_goTypes = nil
	file_workloadapi_security_authorization_proto_depIdxs = nil
}
