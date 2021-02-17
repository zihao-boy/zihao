package jwt

// 提供jwt的基础工具方法
import (
"github.com/kataras/iris/v12/context"
	"github.com/zihao-boy/zihao/zihao-service/common/sysError"
	"strings"
)

type (
	// TokenExtractor is a function that takes a context as input and returns
	// either a token or an error.  An error should only be returned if an attempt
	// to specify a token was found, but the information was somehow incorrectly
	// formed.  In the case where a token is simply not present, this should not
	// be treated as an error.  An empty string should be returned in that case.
	TokenExtractor func(context.Context) (string, error)
)

// below 3 method is get token from url
// FromAuthHeader is a "TokenExtractor" that takes a give context and extracts
// the JWT token from the Authorization header.
func fromAuthHeader(ctx context.Context) (string, error) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return "", nil // No error, just no token
	}

	// TODO: Make this a bit more robust, parsing-wise
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", sysError.ERR_HEADER_NON_BEARER
	}

	return authHeaderParts[1], nil
}

// below 3 method is get token from url
// FromParameter returns a function that extracts the token from the specified
// query string parameter
func fromParameter(param string) TokenExtractor {
	return func(ctx context.Context) (string, error) {
		return ctx.URLParam(param), nil
	}
}

// below 3 method is get token from url
// FromFirst returns a function that runs multiple token extractors and takes the
// first token it finds
func fromFirst(extractors ...TokenExtractor) TokenExtractor {
	return func(ctx context.Context) (string, error) {
		for _, ex := range extractors {
			token, err := ex(ctx)
			if err != nil {
				return "", err
			}
			if token != "" {
				return token, nil
			}
		}
		return "", nil
	}
}
