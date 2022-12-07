package middleware

import (
	"cloud-disk/core/helper"
	"net/http"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 客户端上传 Authorization：Token
		auth := r.Header.Get("Authorization")
		if auth == "" { //如果客户端传进来的Token为空
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("AuthMiddleware 里的Unauthorized"))
			return
		}
		//对客户端传进来的Token解码
		uc, err := helper.AnalyzeToken(auth)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}
		//解码后的用户ID，Identity，Name 写入请求头
		r.Header.Set("UserId", string(rune(uc.Id)))
		r.Header.Set("UserIdentity", uc.Identity)
		r.Header.Set("UserName", uc.Name)
		// fmt.Println("解密后的token", uc.Id, uc.Identity, uc.Name)
		// fmt.Println("Header:", r.Header)
		// fmt.Println("end!")
		next(w, r)
	}
}
