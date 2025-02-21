package config

import "c_program_edu-gin/utils/file_system"

type FileSystem struct {
	Default string                        `json:"default" yaml:"default"`
	Local   file_system.LocalSystemConfig `json:"local" yaml:"local"`
	// Minio   file_system.MinioSystemConfig `json:"minio" yaml:"minio"`
}
