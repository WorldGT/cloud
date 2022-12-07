package define

import (
	"github.com/golang-jwt/jwt/v4"
)

type Userclaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var JwtKey = "cloud-disk-key"

// 验证码长度
var CodeLength = 6

// 验证码过期时间(s)
var CodeExpire = 300

// TencentSecretKey 腾讯云对象存储
var TencentSecretKey = "uhO1Y2Rky7DEB2BteREziLrGZmt0695v"
var TencentSecretID = "AKIDhTCoHyACGEGIGBxjR9nEcyMMm4JVehjh"
var CosBucket = "https://1-1255907395.cos.ap-shanghai.myqcloud.com"

// PageSize分页的默认参数
var PageSize = 20

var Datetime = "2006-01-02 15:04:05"

var TokenExpire = 3600
var RefreshTokenExpire = 7200
