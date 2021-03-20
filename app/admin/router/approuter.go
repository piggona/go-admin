package router

import (
	"go-admin/app/admin/apis"
	"go-admin/app/admin/middleware"
	jwt "go-admin/pkg/jwtauth"

	"github.com/gin-gonic/gin"
)

// 需认证的路由代码
func registerAppRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	r := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/articleList", apis.GetArticleList)
	}

}
