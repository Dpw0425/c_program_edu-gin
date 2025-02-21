package v1

import (
	"c_program_edu-gin/internal/config"
	ctx "c_program_edu-gin/internal/pkg/context"
	myErr "c_program_edu-gin/pkg/error"
	"c_program_edu-gin/pkg/response"
	web "c_program_edu-gin/schema/genproto/web/v1/upload"
	"c_program_edu-gin/utils/file_system"
	"c_program_edu-gin/utils/generator"
)

type Upload struct {
	Config     *config.Config
	FileSystem file_system.IFileSystem
}

func (u *Upload) Avatar(ctx *ctx.Context) error {
	file, err := ctx.Context.FormFile("avatar")
	if err != nil {
		return myErr.BadRequest("", "上传失败！")
	}

	stream, _ := file_system.ReadMultipartStream(file)

	object := generator.GenMediaObjectName("png", 200, 200)
	if err := u.FileSystem.Write(u.FileSystem.BucketPublicName()+"/avatar", object, stream); err != nil {
		return myErr.BadRequest("", "文件上传失败！")
	}

	response.NorResponse(ctx.Context, &web.UploadAvatarResponse{
		Avatar: u.FileSystem.PublicUrl(u.FileSystem.BucketPublicName()+"/avatar", object),
	}, "上传成功！")
	return nil
}
