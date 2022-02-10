package crypto
import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
)
// BcryptPassEncrypt 密码加密
func BcryptPassEncrypt(pwd []byte) string {
	pass, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	return string(pass)
}

// BcryptPassVerify 验证密码
func BcryptPassVerify(pass string, pad []byte) bool {
	byteHash := []byte(pass)
	err := bcrypt.CompareHashAndPassword(byteHash, pad)
	if err != nil {
		return false
	}
	return true
}
// ScryptPassEncrypt 密码加密
func ScryptPassEncrypt(salt, password string) string {
	dk, _ := scrypt.Key([]byte(password), []byte(salt), 32768, 8, 1, 32)
	return fmt.Sprintf("%x", string(dk))
}
