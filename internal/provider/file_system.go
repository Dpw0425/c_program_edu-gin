package provider

import (
	"c_program_edu-gin/internal/config"
	"c_program_edu-gin/utils/file_system"
)

func NewFileSystem(conf *config.Config) file_system.IFileSystem {
	//if conf.FileSystem.Default == file_system.MinioDriver {
	//	return file_system.NewMinioFilesystem(conf.FileSystem.Minio)
	//}

	return file_system.NewLocalFileSystem(conf.FileSystem.Local)
}
