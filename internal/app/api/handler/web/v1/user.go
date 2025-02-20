package v1

import (
	"c_program_edu-gin/internal/app/constants"
	"c_program_edu-gin/internal/app/schema"
	service "c_program_edu-gin/internal/app/service/web/v1"
	"c_program_edu-gin/internal/app/storage/model"
	"c_program_edu-gin/internal/config"
	ctx "c_program_edu-gin/internal/pkg/context"
	myErr "c_program_edu-gin/pkg/error"
	"c_program_edu-gin/pkg/jwt"
	"c_program_edu-gin/pkg/response"
	web "c_program_edu-gin/schema/genproto/web/v1/user"
	"strconv"
	"time"
)

type User struct {
	Config       *config.Config
	UserService  service.IUserService
	EmailService service.IEmailService
}

func (u *User) Register(ctx *ctx.Context) error {
	params := &web.UserRegisterRequest{}
	if err := ctx.Context.ShouldBind(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	if !u.EmailService.Verify(ctx.Ctx(), constants.EmailRegisterChannel, params.Email, params.VerifyCode) {
		return myErr.BadRequest("wrong_verification_code", "验证码错误！")
	}

	if err := u.UserService.Register(ctx.Ctx(), &schema.UserRegister{
		NickName: params.Nickname,
		Password: params.Password,
		Avatar:   params.Avatar,
		Email:    params.Email,
	}); err != nil {
		return myErr.BadRequest("", "用户注册失败: %v", err)
	}

	// 删除验证码缓存
	u.EmailService.Delete(ctx.Ctx(), constants.EmailRegisterChannel, params.Email)

	response.NorResponse(ctx.Context, &web.UserRegisterResponse{}, "注册成功！")
	return nil
}

func (u *User) Login(ctx *ctx.Context) error {
	params := &web.UserLoginRequest{}
	if err := ctx.Context.ShouldBind(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	var user *model.User
	if params.Password == "" {
		if !u.EmailService.Verify(ctx.Ctx(), constants.EmailLoginChannel, params.Email, params.VerifyCode) {
			return myErr.BadRequest("wrong_verification_code", "验证码错误！")
		}

		u.EmailService.Delete(ctx.Ctx(), constants.EmailLoginChannel, params.Email)
		result, err := u.UserService.GetUserByEmail(ctx.Ctx(), params.Email)
		if err != nil {
			return err
		}
		user = result
	} else {
		result, err := u.UserService.LoginByPassword(ctx.Ctx(), &schema.UserLogin{
			Email:    params.Email,
			Password: params.Password,
		})
		if err != nil {
			return err
		}
		user = result
	}

	response.NorResponse(ctx.Context, &web.UserLoginResponse{
		Type:        "Bearer",
		AccessToken: u.token(user.UserID),
		ExpiresIn:   strconv.FormatInt(u.Config.Jwt.ExpiresTime, 10),
	}, "登录成功！")
	return nil
}

func (u *User) token(uid int64) string {
	expiresAt := time.Now().Add(time.Second * time.Duration(u.Config.Jwt.ExpiresTime))

	token := jwt.GenerateToken("api", u.Config.Jwt.Secret, &jwt.Options{
		ExpiresAt: jwt.NewNumericData(expiresAt),
		ID:        strconv.FormatInt(uid, 10),
		Issuer:    "c_program_edu.web",
		IssuedAt:  jwt.NewNumericData(time.Now()),
	})

	return token
}

func (u *User) Info(ctx *ctx.Context) error {
	result, err := u.UserService.GetUserByID(ctx.Ctx(), ctx.UserID())
	if err != nil {
		return err
	}

	response.NorResponse(ctx.Context, &web.UserInfoResponse{
		UserId:   result.UserID,
		NickName: result.NickName,
		Email:    result.Email,
		Avatar:   result.Avatar,
		Status:   int32(result.Status),
	}, "登录成功！")
	return nil
}
