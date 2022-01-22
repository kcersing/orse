package jwt

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"time"
)

type Claims struct {
	UserID   int    `json:"user_id"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour  //* 24 * 2 过期时间 -2天

var Secret = []byte("1111") // Secret 用来加密解密

// GenToken 生成Token
func GenToken(userId int) (string, error) {
	var claims = Claims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "1111",
		},
	}
	// 使用用于签名的算法和令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 创建JWT字符串
	singedToken, err := token.SignedString([]byte(Secret))
	log.Println(singedToken)
	log.Println(err)
	if err != nil {
		//如果创建JWT时出错，则返回内部服务器错误
		//w.WriteHeader(http.StatusInternalServerError)
		return "", fmt.Errorf("生成token失败：%v", err)
	}

	//   // 最后，我们将客户端cookie token设置为刚刚生成的JWT
	//    // 我们还设置了与令牌本身相同的cookie到期时间
	//    http.SetCookie(w, &http.Cookie{
	//        Name:    "token",
	//        Value:   tokenString,
	//        Expires: expirationTime,
	//    })
	return singedToken, nil
}

// VerifyToken 验证Token
func VerifyToken(tokenStr string) (*Claims, error) {
	// 解析JWT字符串并将结果存储在`claims`中。
	// 请注意，我们也在此方法中传递了密钥。
	// 如果令牌无效（如果令牌已根据我们设置的登录到期时间过期）或者签名不匹配,此方法会返回错误.
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	//校验token
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// RefreshToken 更新Token
func RefreshToken(uid int,tokenStr string) (string,error) {
	// 解析JWT字符串并将结果存储在`claims`中。
	// 请注意，我们也在此方法中传递了密钥。
	// 如果令牌无效（如果令牌已根据我们设置的登录到期时间过期）或者签名不匹配,此方法会返回错误.
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return Secret, nil
	})

	if err != nil {
		return "",err
	}

	//校验token
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		log.Println(claims)

		if claims.UserID != uid {
			return "",errors.New("token信息无效")
		}else {
			// 我们确保在足够的时间之前不会发行新令牌。
			// 在这种情况下，仅当旧令牌在30秒到期时才发行新令牌。
			// 否则，返回错误的请求状态。
			//if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
			//	return "",errors.New("token未到过期时间")
			//}

			// 现在，为当前用户创建一个新令牌，并延长其到期时间
			expirationTime := time.Now().Add(TokenExpireDuration)
			claims.ExpiresAt = expirationTime.Unix()
			token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			// 创建JWT字符串
			singedToken, err := token.SignedString([]byte(Secret))

			if err != nil {
				//如果创建JWT时出错，则返回内部服务器错误
				return "", err
			}
			return singedToken, nil
		}
	}else {
		return "",errors.New("token无效")
	}
}

// JWThMiddleware 中间件
func JWThMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			// 处理 没有token的时候
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "缺少token信息",
				"code":    401,
			})
			c.Abort() // 不会继续停止
			return
		}
		// 解析
		mc, err := VerifyToken(token)
		if err != nil {
			// 处理 解析失败
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "token已过期",
				"code":    401,
			})
			c.Abort()
			return
		}
		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set("userID", mc.UserID)
		c.Next()
	}
}