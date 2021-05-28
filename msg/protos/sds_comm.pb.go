// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sds_comm.proto

/*
Package protos is a generated protocol buffer package.

j

It is generated from these files:
	sds_comm.proto
	sds.proto

It has these top-level messages:
	Result
	PPBaseInfo
	FileInfo
	SliceNumAddr
	SliceOffsetInfo
	SliceOffset
	FileIndexes
	DownloadSliceInfo
	SliceStorageInfo
	ShareLinkInfo
	AlbumInfo
	AlbumNumber
	ReqGetPPList
	RspGetPPList
	ReqRegister
	RspRegister
	ReqUploadFile
	RspUploadFile
	ReqUploadFileSlice
	RspUploadFileSlice
	UploadSpeedOfProgress
	ReportUploadSliceResult
	RspReportUploadSliceResult
	ReqFindMyFileList
	RspFindMyFileList
	ReqFindDirectoryTree
	RspFindDirectoryTree
	ReqFileStorageInfo
	RspFileStorageInfo
	ReqDownloadSlice
	RspDownloadSlice
	ReqDownloadSloceWrong
	RspDownloadSloceWrong
	ReqDownloadSlocePause
	RspDownloadSlocePause
	ReqReportDownloadResult
	RspReportDownloadResult
	ReqReportTaskBP
	ReqRegisterNewPP
	RspRegisterNewPP
	ReqDeleteFile
	RspDeleteFile
	ReqTransferNotice
	RspTransferNotice
	ReqValidateTransferCer
	RspValidateTransferCer
	ReqTransferDownload
	RspTransferDownload
	RspTransferDownloadResult
	ReqReportTransferResult
	RspReportTransferResult
	ReqGetHDInfo
	RspGetHDInfo
	ReqSendChatMessages
	RspSendChatMessages
	ReqDeleteSlice
	RspDeleteSlice
	ReqMakeDirectory
	RspMakeDirectory
	ReqRemoveDirectory
	RspRemoveDirectory
	ReqMoveFileDirectory
	RspMoveFileDirectory
	ReqBLSPublicKey
	RspBPBLSPublicKey
	ReqBalance
	RspBalance
	ReqTransaction
	RspTransaction
	ReqBlockInfo
	RspBlockInfo
	ReqBlockCheck
	RspBlockCheck
	BlockCheckInfo
	ReqDownloadTaskInfo
	RspDownloadTaskInfo
	ReqShareLink
	RspShareLink
	ReqShareFile
	RspShareFile
	ReqDeleteShare
	RspDeleteShare
	ReqGetShareFile
	RspGetShareFile
	ReqSaveFile
	RspSaveFile
	ReqSaveFolder
	RspSaveFolder
	ReqCreateAlbum
	RspCreateAlbum
	ReqEditAlbum
	RspEditAlbum
	ReqAlbumContent
	RspAlbumContent
	ReqSearchAlbum
	RspSearchAlbum
	ReqFindMyAlbum
	RspFindMyAlbum
	ReqCollectionAlbum
	RspCollectionAlbum
	ReqAbstractAlbum
	RspAbstractAlbum
	ReqMyCollectionAlbum
	RspMyCollectionAlbum
	ReqDeleteAlbum
	RspDeleteAlbum
	ReqConfig
	RspConfig
	ReqInvite
	RspInvite
	ReqGetReward
	RspGetReward
	ReqGetCapacity
	RspGetCapacity
	ReqFileSort
	RspFileSort
	ReqFindDirectory
	RspFindDirectory
	ReqCustomerAddVolume
	RspCustomerAddVolume
	ReqCustomerUseVolume
	RspCustomerUseVolume
*/
package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ResultState int32

const (
	ResultState_RES_SUCCESS ResultState = 0
	ResultState_RES_FAIL    ResultState = 1
)

var ResultState_name = map[int32]string{
	0: "RES_SUCCESS",
	1: "RES_FAIL",
}
var ResultState_value = map[string]int32{
	"RES_SUCCESS": 0,
	"RES_FAIL":    1,
}

func (x ResultState) String() string {
	return proto.EnumName(ResultState_name, int32(x))
}
func (ResultState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type DownloadWrongType int32

const (
	DownloadWrongType_LOSESLICE DownloadWrongType = 0
	DownloadWrongType_OVERTIME  DownloadWrongType = 1
)

var DownloadWrongType_name = map[int32]string{
	0: "LOSESLICE",
	1: "OVERTIME",
}
var DownloadWrongType_value = map[string]int32{
	"LOSESLICE": 0,
	"OVERTIME":  1,
}

func (x DownloadWrongType) String() string {
	return proto.EnumName(DownloadWrongType_name, int32(x))
}
func (DownloadWrongType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type AlbumType int32

const (
	AlbumType_ALL   AlbumType = 0
	AlbumType_VIDEO AlbumType = 1
	AlbumType_MUSIC AlbumType = 2
	AlbumType_OTHER AlbumType = 3
)

var AlbumType_name = map[int32]string{
	0: "ALL",
	1: "VIDEO",
	2: "MUSIC",
	3: "OTHER",
}
var AlbumType_value = map[string]int32{
	"ALL":   0,
	"VIDEO": 1,
	"MUSIC": 2,
	"OTHER": 3,
}

func (x AlbumType) String() string {
	return proto.EnumName(AlbumType_name, int32(x))
}
func (AlbumType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type AlbumSortType int32

const (
	AlbumSortType_LATEST AlbumSortType = 0
	AlbumSortType_VISITS AlbumSortType = 1
)

var AlbumSortType_name = map[int32]string{
	0: "LATEST",
	1: "VISITS",
}
var AlbumSortType_value = map[string]int32{
	"LATEST": 0,
	"VISITS": 1,
}

func (x AlbumSortType) String() string {
	return proto.EnumName(AlbumSortType_name, int32(x))
}
func (AlbumSortType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type FileSortType int32

const (
	FileSortType_DEF  FileSortType = 0
	FileSortType_TIME FileSortType = 1
	FileSortType_SIZE FileSortType = 2
	FileSortType_NAME FileSortType = 3
)

var FileSortType_name = map[int32]string{
	0: "DEF",
	1: "TIME",
	2: "SIZE",
	3: "NAME",
}
var FileSortType_value = map[string]int32{
	"DEF":  0,
	"TIME": 1,
	"SIZE": 2,
	"NAME": 3,
}

func (x FileSortType) String() string {
	return proto.EnumName(FileSortType_name, int32(x))
}
func (FileSortType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type MsgType int32

const (
	MsgType_DEFAULT MsgType = 0
)

var MsgType_name = map[int32]string{
	0: "DEFAULT",
}
var MsgType_value = map[string]int32{
	"DEFAULT": 0,
}

func (x MsgType) String() string {
	return proto.EnumName(MsgType_name, int32(x))
}
func (MsgType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type Result struct {
	State ResultState `protobuf:"varint,1,opt,name=state,enum=protos.ResultState" json:"state,omitempty"`
	Msg   string      `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Result) GetState() ResultState {
	if m != nil {
		return m.State
	}
	return ResultState_RES_SUCCESS
}

func (m *Result) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type PPBaseInfo struct {
	WalletAddress  string `protobuf:"bytes,1,opt,name=wallet_address,json=walletAddress" json:"wallet_address,omitempty"`
	NetworkAddress string `protobuf:"bytes,2,opt,name=network_address,json=networkAddress" json:"network_address,omitempty"`
}

func (m *PPBaseInfo) Reset()                    { *m = PPBaseInfo{} }
func (m *PPBaseInfo) String() string            { return proto.CompactTextString(m) }
func (*PPBaseInfo) ProtoMessage()               {}
func (*PPBaseInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PPBaseInfo) GetWalletAddress() string {
	if m != nil {
		return m.WalletAddress
	}
	return ""
}

func (m *PPBaseInfo) GetNetworkAddress() string {
	if m != nil {
		return m.NetworkAddress
	}
	return ""
}

type FileInfo struct {
	FileSize           uint64 `protobuf:"varint,1,opt,name=file_size,json=fileSize" json:"file_size,omitempty"`
	FileHash           string `protobuf:"bytes,2,opt,name=file_hash,json=fileHash" json:"file_hash,omitempty"`
	FileName           string `protobuf:"bytes,3,opt,name=file_name,json=fileName" json:"file_name,omitempty"`
	CreateTime         uint64 `protobuf:"varint,4,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	IsDirectory        bool   `protobuf:"varint,5,opt,name=is_directory,json=isDirectory" json:"is_directory,omitempty"`
	StoragePath        string `protobuf:"bytes,6,opt,name=storage_path,json=storagePath" json:"storage_path,omitempty"`
	IsPrivate          bool   `protobuf:"varint,7,opt,name=is_private,json=isPrivate" json:"is_private,omitempty"`
	OwnerWalletAddress string `protobuf:"bytes,8,opt,name=owner_wallet_address,json=ownerWalletAddress" json:"owner_wallet_address,omitempty"`
	ShareLink          string `protobuf:"bytes,9,opt,name=share_link,json=shareLink" json:"share_link,omitempty"`
	SortId             uint64 `protobuf:"varint,10,opt,name=sort_id,json=sortId" json:"sort_id,omitempty"`
}

func (m *FileInfo) Reset()                    { *m = FileInfo{} }
func (m *FileInfo) String() string            { return proto.CompactTextString(m) }
func (*FileInfo) ProtoMessage()               {}
func (*FileInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FileInfo) GetFileSize() uint64 {
	if m != nil {
		return m.FileSize
	}
	return 0
}

func (m *FileInfo) GetFileHash() string {
	if m != nil {
		return m.FileHash
	}
	return ""
}

func (m *FileInfo) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *FileInfo) GetCreateTime() uint64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *FileInfo) GetIsDirectory() bool {
	if m != nil {
		return m.IsDirectory
	}
	return false
}

func (m *FileInfo) GetStoragePath() string {
	if m != nil {
		return m.StoragePath
	}
	return ""
}

func (m *FileInfo) GetIsPrivate() bool {
	if m != nil {
		return m.IsPrivate
	}
	return false
}

func (m *FileInfo) GetOwnerWalletAddress() string {
	if m != nil {
		return m.OwnerWalletAddress
	}
	return ""
}

func (m *FileInfo) GetShareLink() string {
	if m != nil {
		return m.ShareLink
	}
	return ""
}

func (m *FileInfo) GetSortId() uint64 {
	if m != nil {
		return m.SortId
	}
	return 0
}

// slice number and pp address
type SliceNumAddr struct {
	SliceNumber uint64       `protobuf:"varint,1,opt,name=slice_number,json=sliceNumber" json:"slice_number,omitempty"`
	SliceOffset *SliceOffset `protobuf:"bytes,2,opt,name=slice_offset,json=sliceOffset" json:"slice_offset,omitempty"`
	PpInfo      *PPBaseInfo  `protobuf:"bytes,3,opt,name=pp_info,json=ppInfo" json:"pp_info,omitempty"`
}

func (m *SliceNumAddr) Reset()                    { *m = SliceNumAddr{} }
func (m *SliceNumAddr) String() string            { return proto.CompactTextString(m) }
func (*SliceNumAddr) ProtoMessage()               {}
func (*SliceNumAddr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SliceNumAddr) GetSliceNumber() uint64 {
	if m != nil {
		return m.SliceNumber
	}
	return 0
}

func (m *SliceNumAddr) GetSliceOffset() *SliceOffset {
	if m != nil {
		return m.SliceOffset
	}
	return nil
}

func (m *SliceNumAddr) GetPpInfo() *PPBaseInfo {
	if m != nil {
		return m.PpInfo
	}
	return nil
}

type SliceOffsetInfo struct {
	SliceHash   string       `protobuf:"bytes,1,opt,name=slice_hash,json=sliceHash" json:"slice_hash,omitempty"`
	SliceOffset *SliceOffset `protobuf:"bytes,2,opt,name=slice_offset,json=sliceOffset" json:"slice_offset,omitempty"`
}

func (m *SliceOffsetInfo) Reset()                    { *m = SliceOffsetInfo{} }
func (m *SliceOffsetInfo) String() string            { return proto.CompactTextString(m) }
func (*SliceOffsetInfo) ProtoMessage()               {}
func (*SliceOffsetInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SliceOffsetInfo) GetSliceHash() string {
	if m != nil {
		return m.SliceHash
	}
	return ""
}

func (m *SliceOffsetInfo) GetSliceOffset() *SliceOffset {
	if m != nil {
		return m.SliceOffset
	}
	return nil
}

type SliceOffset struct {
	SliceOffsetStart uint64 `protobuf:"varint,1,opt,name=slice_offset_start,json=sliceOffsetStart" json:"slice_offset_start,omitempty"`
	SliceOffsetEnd   uint64 `protobuf:"varint,2,opt,name=slice_offset_end,json=sliceOffsetEnd" json:"slice_offset_end,omitempty"`
}

func (m *SliceOffset) Reset()                    { *m = SliceOffset{} }
func (m *SliceOffset) String() string            { return proto.CompactTextString(m) }
func (*SliceOffset) ProtoMessage()               {}
func (*SliceOffset) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SliceOffset) GetSliceOffsetStart() uint64 {
	if m != nil {
		return m.SliceOffsetStart
	}
	return 0
}

func (m *SliceOffset) GetSliceOffsetEnd() uint64 {
	if m != nil {
		return m.SliceOffsetEnd
	}
	return 0
}

type FileIndexes struct {
	//  string file_hash = 1; // file hash
	FilePath      string `protobuf:"bytes,1,opt,name=file_path,json=filePath" json:"file_path,omitempty"`
	WalletAddress string `protobuf:"bytes,2,opt,name=wallet_address,json=walletAddress" json:"wallet_address,omitempty"`
	SavePath      string `protobuf:"bytes,3,opt,name=save_path,json=savePath" json:"save_path,omitempty"`
}

func (m *FileIndexes) Reset()                    { *m = FileIndexes{} }
func (m *FileIndexes) String() string            { return proto.CompactTextString(m) }
func (*FileIndexes) ProtoMessage()               {}
func (*FileIndexes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *FileIndexes) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

func (m *FileIndexes) GetWalletAddress() string {
	if m != nil {
		return m.WalletAddress
	}
	return ""
}

func (m *FileIndexes) GetSavePath() string {
	if m != nil {
		return m.SavePath
	}
	return ""
}

type DownloadSliceInfo struct {
	SliceStorageInfo *SliceStorageInfo `protobuf:"bytes,1,opt,name=slice_storage_info,json=sliceStorageInfo" json:"slice_storage_info,omitempty"`
	SliceNumber      uint64            `protobuf:"varint,2,opt,name=slice_number,json=sliceNumber" json:"slice_number,omitempty"`
	StoragePpInfo    *PPBaseInfo       `protobuf:"bytes,3,opt,name=storage_pp_info,json=storagePpInfo" json:"storage_pp_info,omitempty"`
	BackupsPpInfo    *PPBaseInfo       `protobuf:"bytes,4,opt,name=backups_pp_info,json=backupsPpInfo" json:"backups_pp_info,omitempty"`
	VisitResult      bool              `protobuf:"varint,5,opt,name=visit_result,json=visitResult" json:"visit_result,omitempty"`
	TaskId           string            `protobuf:"bytes,6,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	SliceOffset      *SliceOffset      `protobuf:"bytes,7,opt,name=slice_offset,json=sliceOffset" json:"slice_offset,omitempty"`
}

func (m *DownloadSliceInfo) Reset()                    { *m = DownloadSliceInfo{} }
func (m *DownloadSliceInfo) String() string            { return proto.CompactTextString(m) }
func (*DownloadSliceInfo) ProtoMessage()               {}
func (*DownloadSliceInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DownloadSliceInfo) GetSliceStorageInfo() *SliceStorageInfo {
	if m != nil {
		return m.SliceStorageInfo
	}
	return nil
}

func (m *DownloadSliceInfo) GetSliceNumber() uint64 {
	if m != nil {
		return m.SliceNumber
	}
	return 0
}

func (m *DownloadSliceInfo) GetStoragePpInfo() *PPBaseInfo {
	if m != nil {
		return m.StoragePpInfo
	}
	return nil
}

func (m *DownloadSliceInfo) GetBackupsPpInfo() *PPBaseInfo {
	if m != nil {
		return m.BackupsPpInfo
	}
	return nil
}

func (m *DownloadSliceInfo) GetVisitResult() bool {
	if m != nil {
		return m.VisitResult
	}
	return false
}

func (m *DownloadSliceInfo) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *DownloadSliceInfo) GetSliceOffset() *SliceOffset {
	if m != nil {
		return m.SliceOffset
	}
	return nil
}

type SliceStorageInfo struct {
	SliceSize uint64 `protobuf:"varint,1,opt,name=slice_size,json=sliceSize" json:"slice_size,omitempty"`
	SliceHash string `protobuf:"bytes,2,opt,name=slice_hash,json=sliceHash" json:"slice_hash,omitempty"`
}

func (m *SliceStorageInfo) Reset()                    { *m = SliceStorageInfo{} }
func (m *SliceStorageInfo) String() string            { return proto.CompactTextString(m) }
func (*SliceStorageInfo) ProtoMessage()               {}
func (*SliceStorageInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *SliceStorageInfo) GetSliceSize() uint64 {
	if m != nil {
		return m.SliceSize
	}
	return 0
}

func (m *SliceStorageInfo) GetSliceHash() string {
	if m != nil {
		return m.SliceHash
	}
	return ""
}

type ShareLinkInfo struct {
	Name               string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	FileHash           string `protobuf:"bytes,2,opt,name=file_hash,json=fileHash" json:"file_hash,omitempty"`
	LinkTime           uint64 `protobuf:"varint,3,opt,name=link_time,json=linkTime" json:"link_time,omitempty"`
	FileSize           uint64 `protobuf:"varint,4,opt,name=file_size,json=fileSize" json:"file_size,omitempty"`
	LinkTimeExp        uint64 `protobuf:"varint,5,opt,name=link_time_exp,json=linkTimeExp" json:"link_time_exp,omitempty"`
	ShareId            string `protobuf:"bytes,6,opt,name=share_id,json=shareId" json:"share_id,omitempty"`
	IsDirectory        bool   `protobuf:"varint,7,opt,name=is_directory,json=isDirectory" json:"is_directory,omitempty"`
	ShareLink          string `protobuf:"bytes,8,opt,name=share_link,json=shareLink" json:"share_link,omitempty"`
	IsPrivate          bool   `protobuf:"varint,9,opt,name=is_private,json=isPrivate" json:"is_private,omitempty"`
	ShareLinkPassword  string `protobuf:"bytes,10,opt,name=share_link_password,json=shareLinkPassword" json:"share_link_password,omitempty"`
	OwnerWalletAddress string `protobuf:"bytes,11,opt,name=owner_wallet_address,json=ownerWalletAddress" json:"owner_wallet_address,omitempty"`
}

func (m *ShareLinkInfo) Reset()                    { *m = ShareLinkInfo{} }
func (m *ShareLinkInfo) String() string            { return proto.CompactTextString(m) }
func (*ShareLinkInfo) ProtoMessage()               {}
func (*ShareLinkInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ShareLinkInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ShareLinkInfo) GetFileHash() string {
	if m != nil {
		return m.FileHash
	}
	return ""
}

func (m *ShareLinkInfo) GetLinkTime() uint64 {
	if m != nil {
		return m.LinkTime
	}
	return 0
}

func (m *ShareLinkInfo) GetFileSize() uint64 {
	if m != nil {
		return m.FileSize
	}
	return 0
}

func (m *ShareLinkInfo) GetLinkTimeExp() uint64 {
	if m != nil {
		return m.LinkTimeExp
	}
	return 0
}

func (m *ShareLinkInfo) GetShareId() string {
	if m != nil {
		return m.ShareId
	}
	return ""
}

func (m *ShareLinkInfo) GetIsDirectory() bool {
	if m != nil {
		return m.IsDirectory
	}
	return false
}

func (m *ShareLinkInfo) GetShareLink() string {
	if m != nil {
		return m.ShareLink
	}
	return ""
}

func (m *ShareLinkInfo) GetIsPrivate() bool {
	if m != nil {
		return m.IsPrivate
	}
	return false
}

func (m *ShareLinkInfo) GetShareLinkPassword() string {
	if m != nil {
		return m.ShareLinkPassword
	}
	return ""
}

func (m *ShareLinkInfo) GetOwnerWalletAddress() string {
	if m != nil {
		return m.OwnerWalletAddress
	}
	return ""
}

type AlbumInfo struct {
	AlbumId        string    `protobuf:"bytes,1,opt,name=album_id,json=albumId" json:"album_id,omitempty"`
	AlbumName      string    `protobuf:"bytes,2,opt,name=album_name,json=albumName" json:"album_name,omitempty"`
	AlbumBlurb     string    `protobuf:"bytes,3,opt,name=album_blurb,json=albumBlurb" json:"album_blurb,omitempty"`
	AlbumVisit     int64     `protobuf:"varint,4,opt,name=album_visit,json=albumVisit" json:"album_visit,omitempty"`
	AlbumTime      int64     `protobuf:"varint,5,opt,name=album_time,json=albumTime" json:"album_time,omitempty"`
	AlbumCoverLink string    `protobuf:"bytes,6,opt,name=album_cover_link,json=albumCoverLink" json:"album_cover_link,omitempty"`
	IsPrivate      bool      `protobuf:"varint,7,opt,name=is_private,json=isPrivate" json:"is_private,omitempty"`
	AlbumType      AlbumType `protobuf:"varint,8,opt,name=album_type,json=albumType,enum=protos.AlbumType" json:"album_type,omitempty"`
}

func (m *AlbumInfo) Reset()                    { *m = AlbumInfo{} }
func (m *AlbumInfo) String() string            { return proto.CompactTextString(m) }
func (*AlbumInfo) ProtoMessage()               {}
func (*AlbumInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *AlbumInfo) GetAlbumId() string {
	if m != nil {
		return m.AlbumId
	}
	return ""
}

func (m *AlbumInfo) GetAlbumName() string {
	if m != nil {
		return m.AlbumName
	}
	return ""
}

func (m *AlbumInfo) GetAlbumBlurb() string {
	if m != nil {
		return m.AlbumBlurb
	}
	return ""
}

func (m *AlbumInfo) GetAlbumVisit() int64 {
	if m != nil {
		return m.AlbumVisit
	}
	return 0
}

func (m *AlbumInfo) GetAlbumTime() int64 {
	if m != nil {
		return m.AlbumTime
	}
	return 0
}

func (m *AlbumInfo) GetAlbumCoverLink() string {
	if m != nil {
		return m.AlbumCoverLink
	}
	return ""
}

func (m *AlbumInfo) GetIsPrivate() bool {
	if m != nil {
		return m.IsPrivate
	}
	return false
}

func (m *AlbumInfo) GetAlbumType() AlbumType {
	if m != nil {
		return m.AlbumType
	}
	return AlbumType_ALL
}

type AlbumNumber struct {
	All   int64 `protobuf:"varint,1,opt,name=all" json:"all,omitempty"`
	Video int64 `protobuf:"varint,2,opt,name=video" json:"video,omitempty"`
	Music int64 `protobuf:"varint,3,opt,name=music" json:"music,omitempty"`
	Other int64 `protobuf:"varint,4,opt,name=other" json:"other,omitempty"`
}

func (m *AlbumNumber) Reset()                    { *m = AlbumNumber{} }
func (m *AlbumNumber) String() string            { return proto.CompactTextString(m) }
func (*AlbumNumber) ProtoMessage()               {}
func (*AlbumNumber) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *AlbumNumber) GetAll() int64 {
	if m != nil {
		return m.All
	}
	return 0
}

func (m *AlbumNumber) GetVideo() int64 {
	if m != nil {
		return m.Video
	}
	return 0
}

func (m *AlbumNumber) GetMusic() int64 {
	if m != nil {
		return m.Music
	}
	return 0
}

func (m *AlbumNumber) GetOther() int64 {
	if m != nil {
		return m.Other
	}
	return 0
}

func init() {
	proto.RegisterType((*Result)(nil), "protos.Result")
	proto.RegisterType((*PPBaseInfo)(nil), "protos.PPBaseInfo")
	proto.RegisterType((*FileInfo)(nil), "protos.FileInfo")
	proto.RegisterType((*SliceNumAddr)(nil), "protos.SliceNumAddr")
	proto.RegisterType((*SliceOffsetInfo)(nil), "protos.SliceOffsetInfo")
	proto.RegisterType((*SliceOffset)(nil), "protos.SliceOffset")
	proto.RegisterType((*FileIndexes)(nil), "protos.FileIndexes")
	proto.RegisterType((*DownloadSliceInfo)(nil), "protos.DownloadSliceInfo")
	proto.RegisterType((*SliceStorageInfo)(nil), "protos.SliceStorageInfo")
	proto.RegisterType((*ShareLinkInfo)(nil), "protos.ShareLinkInfo")
	proto.RegisterType((*AlbumInfo)(nil), "protos.AlbumInfo")
	proto.RegisterType((*AlbumNumber)(nil), "protos.AlbumNumber")
	proto.RegisterEnum("protos.ResultState", ResultState_name, ResultState_value)
	proto.RegisterEnum("protos.DownloadWrongType", DownloadWrongType_name, DownloadWrongType_value)
	proto.RegisterEnum("protos.AlbumType", AlbumType_name, AlbumType_value)
	proto.RegisterEnum("protos.AlbumSortType", AlbumSortType_name, AlbumSortType_value)
	proto.RegisterEnum("protos.FileSortType", FileSortType_name, FileSortType_value)
	proto.RegisterEnum("protos.MsgType", MsgType_name, MsgType_value)
}

func init() { proto.RegisterFile("sds_comm.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1111 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xdb, 0x6e, 0xdb, 0x46,
	0x13, 0x36, 0x45, 0x59, 0x14, 0x87, 0x96, 0xcc, 0x6c, 0x82, 0x3f, 0x0a, 0x82, 0xe0, 0x4f, 0x08,
	0x14, 0x71, 0xdd, 0xc0, 0x08, 0x5c, 0xb4, 0x17, 0xbd, 0x53, 0x6c, 0x1a, 0x21, 0x20, 0xdb, 0x02,
	0x29, 0x3b, 0x40, 0x51, 0x80, 0x58, 0x89, 0x6b, 0x8b, 0x30, 0x4f, 0xe0, 0x52, 0x3e, 0xe4, 0x41,
	0x7a, 0xdf, 0x47, 0xea, 0x6b, 0xf4, 0x19, 0x7a, 0x51, 0xec, 0xec, 0x52, 0x47, 0x1b, 0x29, 0x7a,
	0xa5, 0xdd, 0x6f, 0x0e, 0xbb, 0xfb, 0xcd, 0xcc, 0x27, 0x42, 0x97, 0x47, 0x3c, 0x9c, 0xe4, 0x69,
	0x7a, 0x50, 0x94, 0x79, 0x95, 0x93, 0x16, 0xfe, 0x70, 0xc7, 0x85, 0x96, 0xcf, 0xf8, 0x2c, 0xa9,
	0xc8, 0xf7, 0xb0, 0xcd, 0x2b, 0x5a, 0xb1, 0x9e, 0xf6, 0x56, 0xdb, 0xeb, 0x1e, 0x3e, 0x97, 0x8e,
	0xfc, 0x40, 0x9a, 0x03, 0x61, 0xf2, 0xa5, 0x07, 0xb1, 0x41, 0x4f, 0xf9, 0x75, 0xaf, 0xf1, 0x56,
	0xdb, 0x33, 0x7d, 0xb1, 0x74, 0x7e, 0x03, 0x18, 0x0e, 0x3f, 0x51, 0xce, 0xbc, 0xec, 0x2a, 0x27,
	0xdf, 0x41, 0xf7, 0x8e, 0x26, 0x09, 0xab, 0x42, 0x1a, 0x45, 0x25, 0xe3, 0x1c, 0x73, 0x9a, 0x7e,
	0x47, 0xa2, 0x7d, 0x09, 0x92, 0xf7, 0xb0, 0x9b, 0xb1, 0xea, 0x2e, 0x2f, 0x6f, 0xe6, 0x7e, 0x32,
	0x65, 0x57, 0xc1, 0xca, 0xd1, 0xf9, 0xb3, 0x01, 0xed, 0x93, 0x38, 0x91, 0xc9, 0x5f, 0x83, 0x79,
	0x15, 0x27, 0x2c, 0xe4, 0xf1, 0x57, 0x79, 0xd7, 0xa6, 0xdf, 0x16, 0x40, 0x10, 0x7f, 0x65, 0x73,
	0xe3, 0x94, 0xf2, 0xa9, 0x4a, 0x86, 0xc6, 0xcf, 0x94, 0x4f, 0xe7, 0xc6, 0x8c, 0xa6, 0xac, 0xa7,
	0x2f, 0x8c, 0x67, 0x34, 0x65, 0xe4, 0xff, 0x60, 0x4d, 0x4a, 0x46, 0x2b, 0x16, 0x56, 0x71, 0xca,
	0x7a, 0x4d, 0x4c, 0x0c, 0x12, 0x1a, 0xc5, 0x29, 0x23, 0xef, 0x60, 0x27, 0xe6, 0x61, 0x14, 0x97,
	0x6c, 0x52, 0xe5, 0xe5, 0x43, 0x6f, 0xfb, 0xad, 0xb6, 0xd7, 0xf6, 0xad, 0x98, 0x1f, 0xd7, 0x90,
	0x70, 0xe1, 0x55, 0x5e, 0xd2, 0x6b, 0x16, 0x16, 0xb4, 0x9a, 0xf6, 0x5a, 0x78, 0x86, 0xa5, 0xb0,
	0x21, 0xad, 0xa6, 0xe4, 0x0d, 0x40, 0xcc, 0xc3, 0xa2, 0x8c, 0x6f, 0x05, 0xd5, 0x06, 0xe6, 0x30,
	0x63, 0x3e, 0x94, 0x00, 0xf9, 0x08, 0x2f, 0xf2, 0xbb, 0x8c, 0x95, 0xe1, 0x1a, 0x7f, 0x6d, 0xcc,
	0x44, 0xd0, 0xf6, 0x65, 0x85, 0xc4, 0x37, 0x00, 0x7c, 0x4a, 0x4b, 0x16, 0x26, 0x71, 0x76, 0xd3,
	0x33, 0xd1, 0xcf, 0x44, 0x64, 0x10, 0x67, 0x37, 0xe4, 0x25, 0x18, 0x3c, 0x2f, 0xab, 0x30, 0x8e,
	0x7a, 0x80, 0x4f, 0x6a, 0x89, 0xad, 0x17, 0x39, 0xbf, 0x6b, 0xb0, 0x13, 0x24, 0xf1, 0x84, 0x9d,
	0xcd, 0x52, 0x91, 0x0b, 0x2f, 0x2f, 0xf6, 0x61, 0x36, 0x4b, 0xc7, 0xac, 0x54, 0xd4, 0x5a, 0x5c,
	0xf9, 0x8c, 0x59, 0x49, 0x7e, 0xae, 0x5d, 0xf2, 0xab, 0x2b, 0xce, 0x2a, 0x24, 0xd8, 0x5a, 0x74,
	0x0a, 0xa6, 0x3b, 0x47, 0x93, 0x8a, 0x93, 0x1b, 0xf2, 0x03, 0x18, 0x45, 0x11, 0xc6, 0xd9, 0x55,
	0x8e, 0xb4, 0x5b, 0x87, 0xa4, 0x0e, 0x59, 0x34, 0x8d, 0xdf, 0x2a, 0x0a, 0xf1, 0xeb, 0x4c, 0x61,
	0x77, 0x29, 0x11, 0x96, 0x5c, 0xbc, 0x11, 0xcf, 0xc5, 0xb2, 0x6a, 0xea, 0x8d, 0x02, 0xc1, 0xba,
	0xfe, 0xc7, 0x6b, 0x39, 0x0c, 0xac, 0x25, 0x1b, 0xf9, 0x00, 0x64, 0x39, 0x4d, 0xc8, 0x2b, 0x5a,
	0x56, 0x8a, 0x06, 0x7b, 0x29, 0x2e, 0x10, 0x38, 0xd9, 0x03, 0x7b, 0xc5, 0x9b, 0x65, 0x11, 0x1e,
	0xdc, 0xf4, 0xbb, 0x4b, 0xbe, 0x6e, 0x16, 0x39, 0x19, 0x58, 0xb2, 0x79, 0x23, 0x76, 0xcf, 0xf8,
	0xbc, 0x0b, 0xb1, 0x43, 0xb4, 0x45, 0x17, 0x62, 0x7b, 0x6c, 0x4e, 0x4e, 0xe3, 0xb1, 0xc9, 0x79,
	0x0d, 0x26, 0xa7, 0xb7, 0x2a, 0x87, 0xea, 0x64, 0x01, 0x88, 0x1c, 0xce, 0x5f, 0x0d, 0x78, 0x76,
	0x9c, 0xdf, 0x65, 0x49, 0x4e, 0x23, 0x7c, 0x1f, 0x72, 0x78, 0x52, 0xbf, 0xae, 0xee, 0x50, 0x2c,
	0x87, 0x86, 0x54, 0xf5, 0x56, 0xa8, 0x0a, 0xa4, 0x03, 0x16, 0x45, 0xbe, 0x71, 0x09, 0xd9, 0x68,
	0x93, 0xc6, 0x66, 0x9b, 0xfc, 0x02, 0xbb, 0xf3, 0x31, 0xf8, 0x66, 0xd9, 0x3b, 0xf5, 0x74, 0x60,
	0xf5, 0x45, 0xec, 0x98, 0x4e, 0x6e, 0x66, 0x05, 0x9f, 0xc7, 0x36, 0x9f, 0x8e, 0x55, 0xae, 0x2a,
	0xf6, 0x1d, 0xec, 0xdc, 0xc6, 0x3c, 0xae, 0xc2, 0x12, 0x25, 0xab, 0x9e, 0x50, 0xc4, 0x94, 0xc8,
	0xbd, 0x04, 0xa3, 0xa2, 0xfc, 0x46, 0x8c, 0x83, 0x1c, 0xce, 0x96, 0xd8, 0x7a, 0xd1, 0x46, 0x0f,
	0x19, 0xff, 0xb2, 0x87, 0x86, 0x60, 0xaf, 0x93, 0xb6, 0x68, 0xd7, 0x25, 0x89, 0x92, 0xed, 0x8a,
	0x1a, 0xb5, 0xda, 0xcd, 0x8d, 0xb5, 0x6e, 0x76, 0xfe, 0x6e, 0x40, 0x27, 0xa8, 0xe7, 0x17, 0xf3,
	0x11, 0x68, 0xa2, 0x64, 0xc9, 0x66, 0xc1, 0xf5, 0x37, 0x85, 0x4e, 0xa8, 0x81, 0x54, 0x32, 0x5d,
	0x4a, 0xa4, 0x00, 0x50, 0xc7, 0x56, 0xf4, 0xb3, 0xb9, 0xa6, 0x9f, 0x0e, 0x74, 0xe6, 0x91, 0x21,
	0xbb, 0x2f, 0x90, 0xc3, 0xa6, 0x6f, 0xd5, 0xd1, 0xee, 0x7d, 0x41, 0x5e, 0x41, 0x5b, 0x2a, 0xce,
	0x9c, 0x44, 0x03, 0xf7, 0x5e, 0xb4, 0xa1, 0x91, 0xc6, 0xa6, 0x46, 0xae, 0xea, 0x55, 0x7b, 0x5d,
	0xaf, 0x56, 0xf5, 0xd1, 0x5c, 0xd7, 0xc7, 0x03, 0x78, 0xbe, 0x88, 0x0e, 0x0b, 0xca, 0xf9, 0x5d,
	0x5e, 0x4a, 0x69, 0x33, 0xfd, 0x67, 0xf3, 0x34, 0x43, 0x65, 0x78, 0x52, 0x4f, 0xad, 0xa7, 0xf4,
	0xd4, 0xf9, 0xa3, 0x01, 0x66, 0x3f, 0x19, 0xcf, 0x52, 0xa4, 0xfe, 0x15, 0xb4, 0xa9, 0xd8, 0x88,
	0xb7, 0x4a, 0xfa, 0x0d, 0xdc, 0x7b, 0x91, 0xb8, 0xa9, 0x34, 0x61, 0x6d, 0x54, 0x19, 0x11, 0xa9,
	0xff, 0x4f, 0xa4, 0x79, 0x9c, 0xcc, 0xca, 0xb1, 0x1a, 0x52, 0x19, 0xf1, 0x49, 0x20, 0x0b, 0x07,
	0xec, 0x4f, 0xac, 0x84, 0xae, 0x1c, 0x2e, 0x05, 0xb2, 0x38, 0x00, 0xcb, 0xb8, 0x8d, 0x76, 0x79,
	0x00, 0xd6, 0x71, 0x0f, 0x6c, 0x69, 0x9e, 0xe4, 0xb7, 0xac, 0x94, 0x74, 0xca, 0x72, 0x74, 0x11,
	0x3f, 0x12, 0xf0, 0x23, 0x9c, 0x3e, 0xf2, 0x9f, 0x53, 0x9f, 0xf3, 0x50, 0x30, 0xac, 0x48, 0xf7,
	0xf0, 0x59, 0xdd, 0xf8, 0x48, 0xc5, 0xe8, 0xa1, 0x60, 0xf5, 0xd1, 0x0f, 0x05, 0x73, 0x28, 0x58,
	0x88, 0xab, 0x79, 0xb7, 0x41, 0xa7, 0x49, 0x82, 0xfc, 0xe8, 0xbe, 0x58, 0x92, 0x17, 0xb0, 0x7d,
	0x1b, 0x47, 0x2c, 0x47, 0x5a, 0x74, 0x5f, 0x6e, 0x04, 0x9a, 0xce, 0x78, 0x3c, 0x41, 0x32, 0x74,
	0x5f, 0x6e, 0x04, 0x9a, 0x57, 0x53, 0x56, 0x2a, 0x06, 0xe4, 0x66, 0xff, 0x03, 0x58, 0x4b, 0x1f,
	0x1e, 0x64, 0x17, 0x2c, 0xdf, 0x0d, 0xc2, 0xe0, 0xe2, 0xe8, 0xc8, 0x0d, 0x02, 0x7b, 0x8b, 0xec,
	0x40, 0x5b, 0x00, 0x27, 0x7d, 0x6f, 0x60, 0x6b, 0xfb, 0x1f, 0x17, 0x8a, 0xf7, 0xa5, 0xcc, 0xb3,
	0x6b, 0x71, 0x4b, 0xd2, 0x01, 0x73, 0x70, 0x1e, 0xb8, 0xc1, 0xc0, 0x3b, 0x72, 0x65, 0xc4, 0xf9,
	0xa5, 0xeb, 0x8f, 0xbc, 0x53, 0xd7, 0xd6, 0xf6, 0x7f, 0x52, 0x55, 0x46, 0x4f, 0x03, 0xf4, 0xfe,
	0x60, 0x60, 0x6f, 0x11, 0x13, 0xb6, 0x2f, 0xbd, 0x63, 0xf7, 0xdc, 0xd6, 0xc4, 0xf2, 0xf4, 0x22,
	0xf0, 0x8e, 0xec, 0x86, 0x58, 0x9e, 0x8f, 0x3e, 0xbb, 0xbe, 0xad, 0xef, 0xbf, 0x87, 0x0e, 0x86,
	0x05, 0x79, 0x59, 0x61, 0x28, 0x40, 0x6b, 0xd0, 0x1f, 0xb9, 0xc1, 0xc8, 0xde, 0x12, 0xeb, 0x4b,
	0x2f, 0xf0, 0x46, 0x01, 0xe6, 0xdf, 0x11, 0xa2, 0x3f, 0xf7, 0x33, 0x40, 0x3f, 0x76, 0x4f, 0xec,
	0x2d, 0xd2, 0x86, 0xa6, 0xbc, 0x82, 0x58, 0x05, 0xde, 0xaf, 0xae, 0xdd, 0x10, 0xab, 0xb3, 0xfe,
	0xa9, 0x6b, 0xeb, 0xfb, 0xff, 0x03, 0xe3, 0x94, 0xcb, 0xeb, 0x5b, 0x60, 0x1c, 0xbb, 0x27, 0xfd,
	0x8b, 0xc1, 0xc8, 0xde, 0x1a, 0xcb, 0xcf, 0xb5, 0x1f, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xd0,
	0x55, 0x59, 0x58, 0xc7, 0x09, 0x00, 0x00,
}