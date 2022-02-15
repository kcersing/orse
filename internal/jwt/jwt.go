package jwt

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"orse/internal/config"
	"time"
)

const TokenExpireDuration = time.Hour //* 24 * 2 过期时间 -2天
var (
	TokenInvalid     = errors.New("Token Invalid")
	TokenMalformed   = errors.New("Token Malformed")
	TokenExpired     = errors.New("Token Expired")
	TokenNotValidYet = errors.New("Token Not Valid Yet")
)

type Claims struct {
	UserID int    `json:"user_id"`
	Role   int    `json:"role"`
	Name   string `json:"name"`
	jwt.StandardClaims
}

//JWT签名结构
type JWT struct {
	SigningKey []byte
}

func NewJWT() JWT {
	c := config.NewConfig()
	return JWT{
		SigningKey: []byte(c.Salt),
	}
}

func NewClaims(UserID int, Role int, Name string) Claims {
	return Claims{
		UserID,
		Role,
		Name,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),                   //签名生效时间
			ExpiresAt: int64(time.Now().Add(TokenExpireDuration).Unix()), //签名过期时间
			Issuer:    "orse",                                            //签名发行者
		},
	}
}

// GenToken 生成Token
func (j *JWT) GenToken(claims Claims) (string, error) {
	// 使用用于签名的算法和令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 创建JWT字符串
	singedToken, err := token.SignedString(j.SigningKey)
	if err != nil {
		return "", fmt.Errorf("生成token失败：%v", err)
	}
	return singedToken, nil
}

// VerifyToken 验证Token
func (j *JWT) VerifyToken(tokenStr string) (*Claims, error) {
	// 解析JWT字符串并将结果存储在`claims`中。
	// 请注意，我们也在此方法中传递了密钥。
	// 如果令牌无效（如果令牌已根据我们设置的登录到期时间过期）或者签名不匹配,此方法会返回错误.
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	//校验token
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// RefreshToken 更新Token
func (j *JWT) RefreshToken(tokenStr string) (string, error) {
	// 解析JWT字符串并将结果存储在`claims`中。
	// 请注意，我们也在此方法中传递了密钥。
	// 如果令牌无效（如果令牌已根据我们设置的登录到期时间过期）或者签名不匹配,此方法会返回错误.

	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	//校验token
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(TokenExpireDuration).Unix()
		return j.GenToken(*claims)
	}
	return "", TokenInvalid
}

// JWThMiddleware 中间件
func JWThMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			// 处理 没有token的时候
			c.JSON(http.StatusOK, gin.H{
				"message": "缺少token信息",
				"code":    401,
			})
			c.Abort() // 不会继续停止
			return
		}
		// 解析
		j := NewJWT()
		claims, err := j.VerifyToken(token)
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"status": 401,
					"msg":    "授权已过期",
				})
				c.Abort()
				return
			}
			// 处理 解析失败
			c.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
				"code":    401,
			})
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
