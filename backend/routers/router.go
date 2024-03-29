package router

import (
	"net/http"

	commonCtrl "github.com/dosReady/dlog/backend/controllers/common"
	postCtrl "github.com/dosReady/dlog/backend/controllers/post"
	userCtrl "github.com/dosReady/dlog/backend/controllers/user"
	middleware "github.com/dosReady/dlog/backend/modules/middleware"
	"github.com/gin-gonic/gin"
)

func SettingRouters(r *gin.Engine) {
	r.Use(middleware.VerificationURL())
	r.Use(middleware.BodyParser())
	r.Use(gin.Recovery())

	r.POST("/api/user/login", commonCtrl.UserLogin)

	apir1 := r.Group("/api/user", middleware.CertifiedMdlw())
	{
		apir1.POST("/create", userCtrl.UserCreate)
		apir1.POST("/delete/:email", userCtrl.UserDelete)
		apir1.POST("/logout", commonCtrl.UserLogout)
	}

	apir2 := r.Group("/api/post", middleware.CertifiedMdlw())
	{
		apir2.POST("/register", postCtrl.PostRegsiter)
		apir2.POST("/list", postCtrl.PostList)
	}

	apitest := r.Group("/api/test")
	apitest.Use(middleware.CertifiedMdlw())
	{
		apitest.POST("/echo", func(c *gin.Context) {

			c.JSON(http.StatusOK, gin.H{"name": "Hi"})
		})
	}
}
