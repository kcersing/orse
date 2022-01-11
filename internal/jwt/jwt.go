package jwt

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"learn/demo/ent/ent"
	"net/http"
	"time"
)
//r.POST("/auth/token", GetTokenHandler)
type Claims struct {
	UserID   int `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 24 * 2 // 过期时间 -2天

var Secret = []byte("my_secret_key") // Secret 用来加密解密

// 生成 token
func GenToken(userId int,username string)(string,error)  {
	var claims = Claims{
		userId,
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer: "1111",
		},
	}
	// 使用用于签名的算法和令牌
	token :=jwt.NewWithClaims(jwt.SigningMethodES256,claims)
	// 创建JWT字符串
	singedToken,err:=token.SignedString([]byte(Secret))
	if err !=nil{
		//如果创建JWT时出错，则返回内部服务器错误
		//w.WriteHeader(http.StatusInternalServerError)
		return "",fmt.Errorf("生成token失败：%v",err)
	}

	//   // 最后，我们将客户端cookie token设置为刚刚生成的JWT
	//    // 我们还设置了与令牌本身相同的cookie到期时间
	//    http.SetCookie(w, &http.Cookie{
	//        Name:    "token",
	//        Value:   tokenString,
	//        Expires: expirationTime,
	//    })
	return singedToken,nil
}

//验证token
func VerifyToken(tokenStr string)(*Claims,error)   {
	// 解析JWT字符串并将结果存储在`claims`中。
	// 请注意，我们也在此方法中传递了密钥。
	// 如果令牌无效（如果令牌已根据我们设置的登录到期时间过期）或者签名不匹配,此方法会返回错误.
	token,err:=jwt.ParseWithClaims(tokenStr,&Claims{}, func(token *jwt.Token) (i interface{},err error) {
		return Secret,nil
	})
	if err != nil{
		return nil,err
	}
	//校验token
	if claims,ok := token.Claims.(*Claims);ok && token.Valid{
		return claims,nil
	}
	return nil,errors.New("invalid token")
}

func RefreshToken(tokenStr string)(*Claims,string,error) {
	claims := &Claims{}
	token,err:=jwt.ParseWithClaims(tokenStr,claims, func(token *jwt.Token) (i interface{},err error) {
		return Secret,nil
	})
	if err != nil{

	}
	//校验token
	if _,ok := token.Claims.(*Claims);!ok{

	}
	if !token.Valid{

	}
	// 我们确保在足够的时间之前不会发行新令牌。
	// 在这种情况下，仅当旧令牌在30秒到期时才发行新令牌。
	// 否则，返回错误的请求状态。
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {

	}
	// 现在，为当前用户创建一个新令牌，并延长其到期时间
	expirationTime := time.Now().Add(TokenExpireDuration)
	claims.ExpiresAt = expirationTime.Unix()
	token = jwt.NewWithClaims(jwt.SigningMethodES256,claims)
	// 创建JWT字符串
	singedToken,err:=token.SignedString([]byte(Secret))
	if err !=nil{
		//如果创建JWT时出错，则返回内部服务器错误
		//w.WriteHeader(http.StatusInternalServerError)
		return claims,"",fmt.Errorf("生成token失败：%v",err)
	}
	return claims,singedToken,nil

}

func GetTokenHandler(c *gin.Context)  {
	var user ent.Users
	err := c.ShouldBind(user)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":1001,
			"message":"参数错误",
		})
	}
	if user.Name =="" && user.ID == 1{
		//生成token
		tokenString,err:=GenToken(user.ID,user.Name)
		if err !=nil {
			c.JSON(http.StatusOK,gin.H{
				"code":1001,
				"message":"生成token错误",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": gin.H{"token": tokenString},
		})
		return
	}else {
		c.JSON(http.StatusOK, gin.H{
			"code": 2002,
			"msg":  "用户名或密码错误",
		})
		return
	}
	return
}

// JWThMiddleware 中间件
func JWThMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的token中
		token := c.Request.Header.Get("token")
		if token == "" {
			// 处理 没有token的时候
			c.Abort() // 不会继续停止
			return
		}
		// 解析
		mc, err := VerifyToken(token)
		if err != nil {
			// 处理 解析失败
			c.Abort()
			return
		}
		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set("userID", mc.UserID)
		c.Next()
	}
}