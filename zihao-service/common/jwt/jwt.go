package jwt

import (
"github.com/dgrijalva/jwt-go"
"github.com/kataras/golog"
"github.com/kataras/iris/v12"
"github.com/kataras/iris/v12/context"
"goiris/common"
"goiris/common/support"
"time"
)

// Middleware the middleware for JSON Web tokens authentication method
type Jwt struct {
	Config Config
}

// CheckJWT the main functionality, checks for token
func (j *Jwt) Check(ctx context.Context) (*jwt.Token, error) {
	var (
		code        support.Code
		err         error
		token       string
		parsedToken *jwt.Token
	)
	// cors不处理token验证
	if !j.Config.EnableAuthOnOptions {
		if ctx.Method() == iris.MethodOptions {
			return nil, nil
		}
	}
	// Use the specified token extractor to extract a token from the request
	token, err = j.Config.Extractor(ctx)
	// If an error occurs, call the error handler and return an error
	if err != nil {
		code = support.CODE_TOKEN_INVALID
		j.Config.ErrorHandler(ctx, code)
		goto FAIL
	}
	if parsedToken, code, err = j.CheckTokenString(token); err != nil {
		goto FAIL
	}
	ctx.Values().Set(j.Config.ContextKey, parsedToken) // 设置mapclaims到request的上下文中
	return parsedToken, nil
FAIL:
	j.Config.ErrorHandler(ctx, support.CODE_TOKEN_EXPIRE)
	return nil, err
}

func (j *Jwt) CheckTokenString(token string) (*jwt.Token, support.Code, error) {
	var (
		code        support.Code
		err         error
		parsedToken *jwt.Token
		//message     string
	)
	// If the token is empty...
	if token == "" {
		code = support.CODE_TOKEN_EMPTY
		// Check if it was required
		if j.Config.CredentialsOptional {
			// No error, just no token (and that is ok given that CredentialsOptional is true)
			return nil, code, nil
		}
		golog.Error("Error: No credentials found (CredentialsOptional=false)")
		goto FAIL
	}

	// Now parse the token
	parsedToken, err = jwt.Parse(token, j.Config.ValidationKeyGetter)
	// Check if there was an error in parsing...
	if err != nil {
		code = support.CODE_TOKEN_EXPIRE
		goto FAIL
	}

	if j.Config.SigningMethod != nil && j.Config.SigningMethod.Alg() != parsedToken.Header["alg"] { // 算法错误
		//message = fmt.Sprintf("Expected %s signing method but token specified is %s", j.Config.SigningMethod.Alg(), parsedToken.Header["alg"])
		code = support.CODE_TOKEN_INVALID
		goto FAIL
	}

	// Check if the parsed token is valid...
	if !parsedToken.Valid {
		code = support.CODE_TOKEN_INVALID
		goto FAIL
	}

	if j.Config.Expiration {
		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
			if expired := claims.VerifyExpiresAt(time.Now().Unix(), true); !expired {
				code = support.CODE_TOKEN_EXPIRE
				goto FAIL
			}
		}
	}
	return parsedToken, code, nil
FAIL:
	return nil, code, err
}

// Get returns the user (&token) information for this client/request
func (j *Jwt) GetToken(ctx context.Context) *jwt.Token {
	return ctx.Values().Get(j.Config.ContextKey).(*jwt.Token)
}

// ---------------------------------------------------------
func (j *Jwt) InitJwtConfig() *Jwt {
	c := Config{
		ContextKey: DefaultContextKey,
		//这个方法将验证jwt的token
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			//自己加密的秘钥或者说盐值
			return []byte(common.G_AppConfig.Secret), nil
		},
		//设置后，中间件会验证令牌是否使用特定的签名算法进行签名
		//如果签名方法不是常量，则可以使用ValidationKeyGetter回调来实现其他检查
		//重要的是要避免此处的安全问题：https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		//加密的方式
		SigningMethod: jwt.SigningMethodHS256,
		//验证未通过错误处理方式
		ErrorHandler: func(ctx context.Context, code support.Code) {
			golog.Error("token进入错误handler，原因：%s", code.String())
			support.Unauthorized(ctx, code)
		},
		// 指定func用于提取请求中的token
		Extractor: fromAuthHeader,
		// if the token was expired, expiration error will be returned
		Expiration:          true,
		Debug:               true,
		EnableAuthOnOptions: false,
	}
	return &Jwt{Config: c}
}

