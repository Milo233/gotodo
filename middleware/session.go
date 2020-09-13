package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Session 初始化session
func Session(secret string) gin.HandlerFunc {
	store := cookie.NewStore([]byte(secret))
	//Also set Secure: true if using SSL, you should though
	// 7 天
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 3 * 24 * 3600, Path: "/"})
	return sessions.Sessions("gin-session", store)
}