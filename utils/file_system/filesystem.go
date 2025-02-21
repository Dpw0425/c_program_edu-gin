package file_system

import (
	"io"
	"time"
)

const (
	// MinioDriver = "Minio"
	LocalDriver = "local"
)

type IFileSystem interface {
	// Driver 驱动方式
	Driver() string
	// BucketPublicName 获取公开桶名称
	BucketPublicName() string
	// BucketPrivateName 获取私有桶名称
	BucketPrivateName() string
	// Stat 获取文件信息
	Stat(bucketName string, objectName string) (*FileStatInfo, error)
	// Write 文件写入
	Write(bucketName string, objectName string, stream []byte) error
	// Copy 文件拷贝(桶内)
	Copy(bucketName string, srcObjectName string, objectName string) error
	// CopyObject 拷贝文件(桶外)
	CopyObject(srcBucketName string, srcObjectName string, dstBucketName string, dstObjectName string) error
	// Delete 删除文件
	Delete(bucketName string, objectName string) error
	// GetObject 读取文件内容
	GetObject(bucketName string, objectName string) ([]byte, error)
	// PublicUrl 获取公开文件的访问地址
	PublicUrl(bucketName string, objectName string) string
	// PrivateUrl 获取私有文件的访问地址
	PrivateUrl(bucketName string, objectName string, filename string, expire time.Duration) string
	// InitiateMultipartUpload 初始化分片上传
	InitiateMultipartUpload() (string, error)
	// PutObjectPart 分片上传
	PutObjectPart(bucketName string, uploadID string, index int, data io.Reader) (ObjectPart, error)
	// CompleteMultipartUpload 完成分片上传
	CompleteMultipartUpload(bucketName string, objectName string, parts []ObjectPart) error
	// AbortMultipartUpload 取消分片上传
	AbortMultipartUpload(bucketName string, uploadID string, parts []ObjectPart) error
}

// 文件信息
type FileStatInfo struct {
	Name        string    // 文件名
	Size        int64     // 文件大小
	Ext         string    // 文件后缀
	MimeType    string    // 媒体类型
	LastModTime time.Time // 最后修改时间
}

// 分块上传
type ObjectPart struct {
	PartNumber     int
	ETag           string
	PartObjectName string
}

// 本地存储配置信息
type LocalSystemConfig struct {
	Root          string `json:"root" yaml:"root"`
	SSL           bool   `json:"ssl" yaml:"ssl"`
	BucketPublic  string `json:"bucket_public" yaml:"bucket_public"`
	BucketPrivate string `json:"bucket_private" yaml:"bucket_private"`
	Endpoint      string `json:"endpoint" yaml:"endpoint"`
}

// Minio 配置信息
// type MinioSystemConfig struct {}
