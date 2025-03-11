package v1

import (
	"c_program_edu-gin/internal/app/schema"
	adminservice "c_program_edu-gin/internal/app/service/admin/v1"
	"c_program_edu-gin/internal/app/storage/repo"
	"c_program_edu-gin/internal/config"
	ctx "c_program_edu-gin/internal/pkg/context"
	myErr "c_program_edu-gin/pkg/error"
	"c_program_edu-gin/pkg/jwt"
	"c_program_edu-gin/pkg/response"
	admin "c_program_edu-gin/schema/genproto/admin/v1/user"
	"strconv"
	"time"
)

type Admin struct {
	Config       *config.Config
	AdminRepo    *repo.AdminRepo
	AdminService adminservice.AdminService
}

func (a *Admin) Login(ctx *ctx.Context) error {
	params := &admin.AdminLoginRequest{}
	if err := ctx.Context.ShouldBind(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	result, err := a.AdminService.Login(ctx.Ctx(), &schema.AdminLogin{
		UserName: params.UserName,
		Password: params.Password,
	})
	if err != nil {
		return err
	}

	response.NorResponse(ctx.Context, &admin.AdminLoginResponse{
		Type:        "Bearer",
		AccessToken: a.token(int64(result.ID)),
		ExpiresIn:   strconv.FormatInt(a.Config.Jwt.ExpiresTime, 10),
	}, "登录成功！")

	return nil
}

func (a *Admin) token(uid int64) string {
	expiresAt := time.Now().Add(time.Second * time.Duration(a.Config.Jwt.ExpiresTime))

	token := jwt.GenerateToken("admin", a.Config.Jwt.Secret, &jwt.Options{
		ExpiresAt: jwt.NewNumericData(expiresAt),
		ID:        strconv.FormatInt(uid, 10),
		Issuer:    "c_program_edu.admin",
		IssuedAt:  jwt.NewNumericData(time.Now()),
	})

	return token
}
