package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12/context"
	"github.com/zihao-boy/zihao/zihao-service/common/cache/redis"
	"github.com/zihao-boy/zihao/zihao-service/common/constants"
	"github.com/zihao-boy/zihao/zihao-service/config"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/user"
	"time"
)

var G_JWT *JWT

type (
	JWT struct {
		*IJwt
	}
	Claims struct {
		UserId     string  `json:"userId"`
		RealName string `json:"realName"`
		Phone  string `json:"phone"`
		Enable bool   `json:"enable"`
		jwt.StandardClaims
	}
)

func InitJWT() {
	iJwt := IJwt{}
	G_JWT = &JWT{iJwt.InitIJwtConfig()}
}

// Serve the middleware's action
func (j *JWT) ServeHTTP(ctx *context.Context) (err error) {
	var (
		token *jwt.Token
		user  *user.UserDto
	)
	if token, err = j.Check(*ctx); err != nil {
		return err
	}

	if user, err = j.Token2Model(token); err != nil {
		return err
	}
	// 检查redis缓存
	if _, err = redis.G_Redis.GetToken(constants.REDIS_ADMIN_FORMAT, user.UserId); err != nil {
		return err
	}
	// token校验通过，设置当前用户id到上下文
	ctx.Values().Set(constants.UID, user.UserId)
	return nil
}

// Serve the middleware's action
func (j *JWT) ServeWebsocket(ctx context.Context) {
	var (
		token *jwt.Token
		user  *user.UserDto
		err error
	)
	if token, err = j.Check(ctx); err != nil {
		return
	}

	if user, err = j.Token2Model(token); err != nil {
		return
	}
	if _, err = redis.G_Redis.GetToken(constants.REDIS_ADMIN_FORMAT, user.UserId); err != nil {
		return
	}
	// If everything ok then call next.
	ctx.Next()
}

// 在登录成功的时候生成token
func (j *JWT) GenerateToken(user *user.UserDto) (string, error) {
	expireTime := time.Now().Add(time.Duration(config.G_AppConfig.JWTTimeout) * time.Minute)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		user.UserId,
		user.RealName,
		user.Phone,
		true,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "zihao-jwt",
		},
	})
	return tokenClaims.SignedString([]byte(config.G_AppConfig.Secret))
}

// 解析token的信息为用户
func (j *JWT) Token2Model(token *jwt.Token) (*user.UserDto, error) {
	//mapClaims := (jwt.Get(ctx).Claims).(jwt.MapClaims)
	var (
		mapClaims, ok = token.Claims.(jwt.MapClaims)
		id            string
		phone         string
		realName      string
	)
	if !ok {
		return nil, fmt.Errorf("%s", constants.CODE_TOKEN_INVALID.String())
	}

	id = mapClaims["userId"].(string)
	phone = mapClaims["phone"].(string)
	realName = mapClaims["realName"].(string)
	return &user.UserDto{
		UserId:     id,
		Phone:  phone,
		RealName:realName,
	}, nil
}

func (j *JWT) TokenString2Model(tokenString string) (user *user.UserDto, err error) {
	var token *jwt.Token
	if token, err = jwt.Parse(tokenString, j.Config.ValidationKeyGetter); err != nil {
		return nil, err
	}
	return j.Token2Model(token)
}

