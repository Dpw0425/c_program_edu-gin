package file_system

import (
	"io"
	"mime/multipart"
	"os"
)

func isDirExist(fileAddr string) bool {
	s, err := os.Stat(fileAddr)

	return err == nil && s.IsDir()
}

func ReadMultipartStream(file *multipart.FileHeader) ([]byte, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	return io.ReadAll(src)
}
