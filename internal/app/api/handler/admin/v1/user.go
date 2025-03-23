package v1

import (
	"c_program_edu-gin/internal/app/schema"
	adminservice "c_program_edu-gin/internal/app/service/admin/v1"
	"c_program_edu-gin/internal/app/storage/cache"
	"c_program_edu-gin/internal/config"
	ctx "c_program_edu-gin/internal/pkg/context"
	myErr "c_program_edu-gin/pkg/error"
	"c_program_edu-gin/pkg/jwt"
	"c_program_edu-gin/pkg/response"
	admin "c_program_edu-gin/schema/genproto/admin/v1/user"
	"strconv"
	"strings"
	"time"
)

type Admin struct {
	Config          *config.Config
	JwtTokenStorage *cache.JwtTokenStorage
	AdminService    adminservice.AdminService
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

func (a *Admin) Logout(ctx *ctx.Context) error {
	a.toBlackList(ctx)

	response.NorResponse(ctx.Context, &admin.AdminLogoutResponse{}, "退出成功！")
	return nil
}

func (a *Admin) toBlackList(ctx *ctx.Context) {
	session := ctx.JWTSession()
	if session != nil {
		if ex := session.ExpiresAt - time.Now().Unix(); ex > 0 {
			_ = a.JwtTokenStorage.SetBlackList(ctx.Ctx(), session.Token, time.Duration(ex)*time.Second)
		}
	}
}

func (a *Admin) Info(ctx *ctx.Context) error {
	result, err := a.AdminService.GetAdminByID(ctx.Ctx(), uint(ctx.UserID()))
	if err != nil {
		return err
	}

	response.NorResponse(ctx.Context, &admin.AdminInfoResponse{
		UserName:   result.UserName,
		TeacherId:  result.TeacherID,
		Permission: strings.Split(result.Permission, ","),
		Status:     int32(result.Status),
	}, "登录成功！")
	return nil
}
