// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
// source: competition.proto

package web

import (
	_ "github.com/srikrsna/protoc-gen-gotag/tagger"
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

type GetCompetitionListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty" form:"page"`
	Number int32  `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty" form:"number"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty" form:"search"`
}

func (x *GetCompetitionListRequest) Reset() {
	*x = GetCompetitionListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_competition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCompetitionListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompetitionListRequest) ProtoMessage() {}

func (x *GetCompetitionListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_competition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCompetitionListRequest.ProtoReflect.Descriptor instead.
func (*GetCompetitionListRequest) Descriptor() ([]byte, []int) {
	return file_competition_proto_rawDescGZIP(), []int{0}
}

func (x *GetCompetitionListRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetCompetitionListRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *GetCompetitionListRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetCompetitionListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompetitionList []*GetCompetitionListResponse_CompetitionItem `protobuf:"bytes,1,rep,name=competition_list,json=competitionList,proto3" json:"competition_list,omitempty" binding:"required"`
	Total           int32                                         `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty" binding:"required"`
}

func (x *GetCompetitionListResponse) Reset() {
	*x = GetCompetitionListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_competition_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCompetitionListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompetitionListResponse) ProtoMessage() {}

func (x *GetCompetitionListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_competition_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCompetitionListResponse.ProtoReflect.Descriptor instead.
func (*GetCompetitionListResponse) Descriptor() ([]byte, []int) {
	return file_competition_proto_rawDescGZIP(), []int{1}
}

func (x *GetCompetitionListResponse) GetCompetitionList() []*GetCompetitionListResponse_CompetitionItem {
	if x != nil {
		return x.CompetitionList
	}
	return nil
}

func (x *GetCompetitionListResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetCompetitionDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" form:"id"`
}

func (x *GetCompetitionDetailRequest) Reset() {
	*x = GetCompetitionDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_competition_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCompetitionDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompetitionDetailRequest) ProtoMessage() {}

func (x *GetCompetitionDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_competition_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCompetitionDetailRequest.ProtoReflect.Descriptor instead.
func (*GetCompetitionDetailRequest) Descriptor() ([]byte, []int) {
	return file_competition_proto_rawDescGZIP(), []int{2}
}

func (x *GetCompetitionDetailRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCompetitionDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompetitionItem *GetCompetitionDetailResponse_Competition `protobuf:"bytes,1,opt,name=competition_item,json=competitionItem,proto3" json:"competition_item,omitempty" binding:"required"`
}

func (x *GetCompetitionDetailResponse) Reset() {
	*x = GetCompetitionDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_competition_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCompetitionDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompetitionDetailResponse) ProtoMessage() {}

func (x *GetCompetitionDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_competition_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCompetitionDetailResponse.ProtoReflect.Descriptor instead.
func (*GetCompetitionDetailResponse) Descriptor() ([]byte, []int) {
	return file_competition_proto_rawDescGZIP(), []int{3}
}

func (x *GetCompetitionDetailResponse) GetCompetitionItem() *GetCompetitionDetailResponse_Competition {
	if x != nil {
		return x.CompetitionItem
	}
	return nil
}

type GetRankingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" form:"id"`
}

func (x *GetRankingRequest) Reset() {
	*x = GetRankingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_competition_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRankingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRankingRequest) ProtoMessage() {}

func (x *GetRankingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_competition_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRankingRequest.ProtoReflect.Descriptor instead.
func (*GetRankingRequest) Descriptor() ([]byte, []int) {
	return file_competition_proto_rawDescGZIP(), []int{4}
}

func (x *GetRankingRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetRankingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserList []*GetRankingResponse_RankItem `protobuf:"bytes,1,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`
}

func (x *GetRankingResponse) Reset() {
	*x = GetRankingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_competition_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRankingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRankingResponse) ProtoMessage() {}

func (x *GetRankingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_competition_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRankingResponse.ProtoReflect.Descriptor instead.
func (*GetRankingResponse) Descriptor() ([]byte, []int) {
	return file_competition_proto_rawDescGZIP(), []int{5}
}

func (x *GetRankingResponse) GetUserList() []*GetRankingResponse_RankItem {
	if x != nil {
		return x.UserList
	}
	return nil
}

type GetCompetitionListResponse_CompetitionItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" binding:"required"`
	Name       string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" binding:"required"`
	Contestant []string `protobuf:"bytes,3,rep,name=contestant,proto3" json:"contestant,omitempty" binding:"required"`
	OwnerId    int32    `protobuf:"varint,4,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty" binding:"required"`
	StartTime  string   `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty" binding:"required"`
	Deadline   string   `protobuf:"bytes,6,opt,name=deadline,proto3" json:"deadline,omitempty" binding:"required"`
	Status     int32    `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty" binding:"required"`
	Category   int32    `protobuf:"varint,8,opt,name=category,proto3" json:"category,omitempty" binding:"required"`
	Permission int32    `protobuf:"varint,9,opt,name=permission,proto3" json:"permission,omitempty" binding:"required"`
	Quantity   int32    `protobuf:"varint,10,opt,name=quantity,proto3" json:"quantity,omitempty" binding:"required"`
}

func (x *GetCompetitionListResponse_CompetitionItem) Reset() {
	*x = GetCompetitionListResponse_CompetitionItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_competition_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCompetitionListResponse_CompetitionItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompetitionListResponse_CompetitionItem) ProtoMessage() {}

func (x *GetCompetitionListResponse_CompetitionItem) ProtoReflect() protoreflect.Message {
	mi := &file_competition_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCompetitionListResponse_CompetitionItem.ProtoReflect.Descriptor instead.
func (*GetCompetitionListResponse_CompetitionItem) Descriptor() ([]byte, []int) {
	return file_competition_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetCompetitionListResponse_CompetitionItem) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetCompetitionListResponse_CompetitionItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetCompetitionListResponse_CompetitionItem) GetContestant() []string {
	if x != nil {
		return x.Contestant
	}
	return nil
}

func (x *GetCompetitionListResponse_CompetitionItem) GetOwnerId() int32 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *GetCompetitionListResponse_CompetitionItem) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *GetCompetitionListResponse_CompetitionItem) GetDeadline() string {
	if x != nil {
		return x.Deadline
	}
	return ""
}

func (x *GetCompetitionListResponse_CompetitionItem) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetCompetitionListResponse_CompetitionItem) GetCategory() int32 {
	if x != nil {
		return x.Category
	}
	return 0
}

func (x *GetCompetitionListResponse_CompetitionItem) GetPermission() int32 {
	if x != nil {
		return x.Permission
	}
	return 0
}

func (x *GetCompetitionListResponse_CompetitionItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type GetCompetitionDetailResponse_Competition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" binding:"required"`
	Name       string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" binding:"required"`
	Contestant []string `protobuf:"bytes,3,rep,name=contestant,proto3" json:"contestant,omitempty" binding:"required"`
	OwnerId    int32    `protobuf:"varint,4,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty" binding:"required"`
	StartTime  string   `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty" binding:"required"`
	Deadline   string   `protobuf:"bytes,6,opt,name=deadline,proto3" json:"deadline,omitempty" binding:"required"`
	Status     int32    `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty" binding:"required"`
	Category   int32    `protobuf:"varint,8,opt,name=category,proto3" json:"category,omitempty" binding:"required"`
	Permission int32    `protobuf:"varint,9,opt,name=permission,proto3" json:"permission,omitempty" binding:"required"`
	Quantity   int32    `protobuf:"varint,10,opt,name=quantity,proto3" json:"quantity,omitempty" binding:"required"`
}

func (x *GetCompetitionDetailResponse_Competition) Reset() {
	*x = GetCompetitionDetailResponse_Competition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_competition_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCompetitionDetailResponse_Competition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompetitionDetailResponse_Competition) ProtoMessage() {}

func (x *GetCompetitionDetailResponse_Competition) ProtoReflect() protoreflect.Message {
	mi := &file_competition_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCompetitionDetailResponse_Competition.ProtoReflect.Descriptor instead.
func (*GetCompetitionDetailResponse_Competition) Descriptor() ([]byte, []int) {
	return file_competition_proto_rawDescGZIP(), []int{3, 0}
}

func (x *GetCompetitionDetailResponse_Competition) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetCompetitionDetailResponse_Competition) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetCompetitionDetailResponse_Competition) GetContestant() []string {
	if x != nil {
		return x.Contestant
	}
	return nil
}

func (x *GetCompetitionDetailResponse_Competition) GetOwnerId() int32 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *GetCompetitionDetailResponse_Competition) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *GetCompetitionDetailResponse_Competition) GetDeadline() string {
	if x != nil {
		return x.Deadline
	}
	return ""
}

func (x *GetCompetitionDetailResponse_Competition) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetCompetitionDetailResponse_Competition) GetCategory() int32 {
	if x != nil {
		return x.Category
	}
	return 0
}

func (x *GetCompetitionDetailResponse_Competition) GetPermission() int32 {
	if x != nil {
		return x.Permission
	}
	return 0
}

func (x *GetCompetitionDetailResponse_Competition) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type GetRankingResponse_RankItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName    string `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Score       int32  `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
	TotalCommit int32  `protobuf:"varint,3,opt,name=total_commit,json=totalCommit,proto3" json:"total_commit,omitempty"`
}

func (x *GetRankingResponse_RankItem) Reset() {
	*x = GetRankingResponse_RankItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_competition_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRankingResponse_RankItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRankingResponse_RankItem) ProtoMessage() {}

func (x *GetRankingResponse_RankItem) ProtoReflect() protoreflect.Message {
	mi := &file_competition_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRankingResponse_RankItem.ProtoReflect.Descriptor instead.
func (*GetRankingResponse_RankItem) Descriptor() ([]byte, []int) {
	return file_competition_proto_rawDescGZIP(), []int{5, 0}
}

func (x *GetRankingResponse_RankItem) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *GetRankingResponse_RankItem) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *GetRankingResponse_RankItem) GetTotalCommit() int32 {
	if x != nil {
		return x.TotalCommit
	}
	return 0
}

var File_competition_proto protoreflect.FileDescriptor

var file_competition_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x77, 0x65, 0x62, 0x1a, 0x1f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f,
	0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x67,
	0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x01, 0x0a, 0x19, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x10, 0x9a, 0x84, 0x9e, 0x03, 0x0b, 0x66, 0x6f, 0x72, 0x6d,
	0x3a, 0x22, 0x70, 0x61, 0x67, 0x65, 0x22, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x12, 0x9a,
	0x84, 0x9e, 0x03, 0x0d, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x22, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x12, 0x9a, 0x84, 0x9e, 0x03, 0x0d,
	0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x52, 0x06, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0xd8, 0x05, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d,
	0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x73, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f,
	0x2e, 0x77, 0x65, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x42,
	0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12, 0x62,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x22, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x1a, 0x95, 0x04, 0x0a, 0x0f, 0x43, 0x6f, 0x6d,
	0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x27, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12, 0x62,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x42, 0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12, 0x62, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52,
	0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x08, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x17, 0x9a,
	0x84, 0x9e, 0x03, 0x12, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x36, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12,
	0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x22, 0x52, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x2f, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x42, 0x17, 0x9a, 0x84,
	0x9e, 0x03, 0x12, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a,
	0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x37, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x42, 0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12, 0x62, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52,
	0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x42, 0x17, 0x9a,
	0x84, 0x9e, 0x03, 0x12, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x22, 0x3d, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0e, 0x9a, 0x84, 0x9e,
	0x03, 0x09, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x22,
	0xa5, 0x05, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x71, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x77, 0x65, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12,
	0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x22, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x74, 0x65, 0x6d, 0x1a, 0x91, 0x04, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0x9a, 0x84, 0x9e, 0x03,
	0x12, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x22, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x42, 0x17, 0x9a,
	0x84, 0x9e, 0x03, 0x12, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x12, 0x32, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12, 0x62, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x07, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0x9a, 0x84, 0x9e, 0x03,
	0x12, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x22, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x33,
	0x0a, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c,
	0x69, 0x6e, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x42, 0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12, 0x62, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52,
	0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x37, 0x0a, 0x0a, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x42, 0x17, 0x9a,
	0x84, 0x9e, 0x03, 0x12, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12, 0x62, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x33, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x61,
	0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0e, 0x9a, 0x84, 0x9e, 0x03, 0x09, 0x66,
	0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x22, 0xb5, 0x01, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x52, 0x61, 0x6e, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x1a, 0x60, 0x0a, 0x08, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1b,
	0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x42, 0x18, 0x5a, 0x16, 0x77, 0x65, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x77, 0x65, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_competition_proto_rawDescOnce sync.Once
	file_competition_proto_rawDescData = file_competition_proto_rawDesc
)

func file_competition_proto_rawDescGZIP() []byte {
	file_competition_proto_rawDescOnce.Do(func() {
		file_competition_proto_rawDescData = protoimpl.X.CompressGZIP(file_competition_proto_rawDescData)
	})
	return file_competition_proto_rawDescData
}

var file_competition_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_competition_proto_goTypes = []interface{}{
	(*GetCompetitionListRequest)(nil),                  // 0: web.GetCompetitionListRequest
	(*GetCompetitionListResponse)(nil),                 // 1: web.GetCompetitionListResponse
	(*GetCompetitionDetailRequest)(nil),                // 2: web.GetCompetitionDetailRequest
	(*GetCompetitionDetailResponse)(nil),               // 3: web.GetCompetitionDetailResponse
	(*GetRankingRequest)(nil),                          // 4: web.GetRankingRequest
	(*GetRankingResponse)(nil),                         // 5: web.GetRankingResponse
	(*GetCompetitionListResponse_CompetitionItem)(nil), // 6: web.GetCompetitionListResponse.CompetitionItem
	(*GetCompetitionDetailResponse_Competition)(nil),   // 7: web.GetCompetitionDetailResponse.Competition
	(*GetRankingResponse_RankItem)(nil),                // 8: web.GetRankingResponse.RankItem
}
var file_competition_proto_depIdxs = []int32{
	6, // 0: web.GetCompetitionListResponse.competition_list:type_name -> web.GetCompetitionListResponse.CompetitionItem
	7, // 1: web.GetCompetitionDetailResponse.competition_item:type_name -> web.GetCompetitionDetailResponse.Competition
	8, // 2: web.GetRankingResponse.user_list:type_name -> web.GetRankingResponse.RankItem
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_competition_proto_init() }
func file_competition_proto_init() {
	if File_competition_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_competition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCompetitionListRequest); i {
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
		file_competition_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCompetitionListResponse); i {
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
		file_competition_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCompetitionDetailRequest); i {
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
		file_competition_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCompetitionDetailResponse); i {
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
		file_competition_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRankingRequest); i {
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
		file_competition_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRankingResponse); i {
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
		file_competition_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCompetitionListResponse_CompetitionItem); i {
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
		file_competition_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCompetitionDetailResponse_Competition); i {
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
		file_competition_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRankingResponse_RankItem); i {
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
			RawDescriptor: file_competition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_competition_proto_goTypes,
		DependencyIndexes: file_competition_proto_depIdxs,
		MessageInfos:      file_competition_proto_msgTypes,
	}.Build()
	File_competition_proto = out.File
	file_competition_proto_rawDesc = nil
	file_competition_proto_goTypes = nil
	file_competition_proto_depIdxs = nil
}
