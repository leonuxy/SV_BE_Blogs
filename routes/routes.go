package routes

import (
	"sv_be_posts/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/article")
	{
		api.POST("", controllers.CreatePost)
		api.GET("/:id", controllers.GetPostByID)
		api.GET("pagination/:limit/:offset", controllers.GetAllPostsWithPagination)
		api.PUT("/:id", controllers.UpdatePost)
		api.DELETE("/:id", controllers.DeletePost)
	}
}
