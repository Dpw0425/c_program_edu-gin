package file_system

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var _ IFileSystem = (*LocalFileSystem)(nil)

type LocalFileSystem struct {
	config LocalSystemConfig
}

func NewLocalFileSystem(config LocalSystemConfig) *LocalFileSystem {
	projectRoot, _ := os.Getwd()

	absRoot := filepath.Join(projectRoot, config.Root)

	return &LocalFileSystem{
		config: LocalSystemConfig{
			Root:          absRoot,
			BucketPublic:  config.BucketPublic,
			BucketPrivate: config.BucketPrivate,
			Endpoint:      config.Endpoint,
			SSL:           config.SSL,
		},
	}
}

func (l LocalFileSystem) Driver() string {
	return LocalDriver
}

func (l LocalFileSystem) BucketPublicName() string {
	return l.config.BucketPublic
}

func (l LocalFileSystem) BucketPrivateName() string {
	return l.config.BucketPrivate
}

func (l LocalFileSystem) Stat(bucketName string, objectName string) (*FileStatInfo, error) {
	info, err := os.Stat(l.Path(bucketName, objectName))
	if err != nil {
		return nil, err
	}

	return &FileStatInfo{
		Name:        filepath.Base(objectName),
		Size:        info.Size(),
		Ext:         filepath.Ext(objectName),
		MimeType:    "",
		LastModTime: info.ModTime(),
	}, nil
}

func (l LocalFileSystem) Write(bucketName string, objectName string, stream []byte) error {
	filePath := l.Path(bucketName, objectName)

	dir := filepath.Dir(filePath)
	if len(dir) > 0 && !isDirExist(dir) {
		if err := os.MkdirAll(dir, 0777); err != nil {
			return err
		}
	}

	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(stream)
	return err
}

// 本地文件写入(用于本地文件写入桶)
func (l LocalFileSystem) WriteLocal(bucketName string, localFile string, objectName string) error {
	srcFile, err := os.Open(localFile)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	objectName = l.Path(bucketName, objectName)
	dir := filepath.Dir(objectName)
	if len(dir) > 0 && !isDirExist(dir) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	dstFile, err := os.OpenFile(objectName, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

func (l LocalFileSystem) Copy(bucketName string, srcObjectName string, objectName string) error {
	return l.WriteLocal(bucketName, l.Path(bucketName, srcObjectName), objectName)
}

func (l LocalFileSystem) CopyObject(srcBucketName string, srcObjectName string, dstBucketName string, dstObjectName string) error {
	srcFile, err := os.Open(l.Path(srcBucketName, srcObjectName))
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstObjectName = l.Path(dstBucketName, dstObjectName)
	if dir := filepath.Dir(dstObjectName); len(dir) > 0 && !isDirExist(dir) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	dstFile, err := os.OpenFile(dstObjectName, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

func (l LocalFileSystem) Delete(bucketName string, objectName string) error {
	return os.Remove(l.Path(bucketName, objectName))
}

func (l LocalFileSystem) GetObject(bucketName string, objectName string) ([]byte, error) {
	return os.ReadFile(l.Path(bucketName, objectName))
}

func (l LocalFileSystem) PublicUrl(bucketName string, objectName string) string {
	// 获取基础域名
	domain := l.getDomain()

	// 生成公开 url
	return fmt.Sprintf(
		"%s/%s/%s",
		strings.TrimRight(domain, "/"),
		bucketName,
		strings.Trim(objectName, "/"),
	)
}

func (l LocalFileSystem) PrivateUrl(bucketName, objectName string, _ string, _ time.Duration) string {
	return l.PublicUrl(bucketName, objectName)
}

func (l LocalFileSystem) InitiateMultipartUpload() (string, error) {
	return uuid.New().String(), nil
}

func (l LocalFileSystem) PutObjectPart(bucketName string, uploadID string, index int, data io.Reader) (ObjectPart, error) {
	stream, _ := io.ReadAll(data)

	objectName := fmt.Sprintf("multipart/%s/%d_%s.tmp", uploadID, index, uploadID)
	if err := l.Write(bucketName, objectName, stream); err != nil {
		return ObjectPart{}, err
	}

	return ObjectPart{
		ETag:           "",
		PartNumber:     index,
		PartObjectName: objectName,
	}, nil
}

func (l LocalFileSystem) CompleteMultipartUpload(bucketName string, objectName string, parts []ObjectPart) error {
	for _, part := range parts {
		stream, err := l.GetObject(bucketName, part.PartObjectName)
		if err != nil {
			return err
		}

		if err := l.appendWrite(bucketName, objectName, stream); err != nil {
			return err
		}
	}

	return nil
}

func (l LocalFileSystem) AbortMultipartUpload(bucketName string, uploadID string, parts []ObjectPart) error {
	for _, part := range parts {
		objectName := part.PartObjectName
		if err := l.Delete(bucketName, objectName); err != nil {
			return err
		}
	}

	return nil
}

// Path 获取文件地址绝对路径
func (l LocalFileSystem) Path(bucketName string, objectName string) string {
	root := filepath.Clean(l.config.Root)

	return filepath.Join(
		root,
		bucketName,
		filepath.Clean(objectName), // 清理对象名
	)
}

func (l LocalFileSystem) getDomain() string {
	endpoint := l.config.Endpoint

	if strings.HasPrefix(endpoint, "http://") || strings.HasPrefix(endpoint, "https://") {
		return endpoint
	}
	if l.config.SSL {
		return fmt.Sprintf("https://%s", endpoint)
	}
	return fmt.Sprintf("http://%s", endpoint)
}

func (l LocalFileSystem) appendWrite(bucketName string, objectName string, stream []byte) error {
	filePath := l.Path(bucketName, objectName)

	dir := filepath.Dir(filePath)
	if len(dir) > 0 && !isDirExist(dir) {
		if err := os.MkdirAll(dir, 0766); err != nil {
			return err
		}
	}

	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0766)
	if err != nil {
		return err
	}

	_, err = f.Write(stream)
	return err
}
