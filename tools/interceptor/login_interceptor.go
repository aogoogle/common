package interceptor

import (
	"errors"
	"fmt"
	"github.com/aogoogle/common/model/request"
	"github.com/aogoogle/common/model/response"
	"github.com/aogoogle/common/tools/redis"
	"github.com/aogoogle/common/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type LoginHeader struct {
	ClientOS string `header:"g-clientOs" binding:"required"`
	Token string `header:"g-token" binding:"required"`
}

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func LoginInterceptor(redis redis.JRedis) gin.HandlerFunc {
	return func(c *gin.Context) {
		//校验头参数和接口必备参数
		var header LoginHeader
		userId := c.Request.URL.Query().Get("userId")
		if err := c.ShouldBindHeader(&header); err != nil || userId == ""  {
			fmt.Println("login interceptor-> error:", err, " userId:"+userId)
			response.FailWithDetailed("", "请求参数校验错误，请重新登陆", c)
			c.Abort()
			return
		}

		//校验token授权是否过期
		j := NewJWT()
		claims, err := j.ParseToken(header.Token)
		if err != nil {
			if err == TokenExpired {
				response.Result(response.TOKENEXPIRED, "", "登陆授权已过期，请重新登陆", c)
				c.Abort()
				return
			}
			response.FailWithDetailed("", err.Error(), c)
			c.Abort()
			return
		}

		//校验token是否合法
		appTokenKey := utils.GenarateAppTokenKey(userId, header.ClientOS)
		err, appToken := redis.GetKey(appTokenKey)
		if err != nil || appToken != header.Token {
			response.FailWithDetailed("", "非法登陆，Token异常", c)
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}

func NewJWT() *JWT {
	return &JWT{
		[]byte("jss"),
	}
}

/*
	创建一个token
 */
func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

/*
	解析 token
 */
func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid

	}

}

// 更新token
//func (j *JWT) RefreshToken(tokenString string) (string, error) {
//	jwt.TimeFunc = func() time.Time {
//		return time.Unix(0, 0)
//	}
//	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return j.SigningKey, nil
//	})
//	if err != nil {
//		return "", err
//	}
//	if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
//		jwt.TimeFunc = time.Now
//		claims.StandardClaims.ExpiresAt = time.Now().Unix() + 60*60*24*7
//		return j.CreateToken(*claims)
//	}
//	return "", TokenInvalid
//}
