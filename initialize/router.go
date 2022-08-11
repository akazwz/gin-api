package initialize

import (
	"github.com/akazwz/gin-api/api"
	"github.com/akazwz/gin-api/api/auth"
	"github.com/akazwz/gin-api/api/file"
	"github.com/akazwz/gin-api/api/image/classify"
	"github.com/akazwz/gin-api/api/posts"
	"github.com/akazwz/gin-api/api/projects"
	"github.com/akazwz/gin-api/api/s3/r2"
	"github.com/akazwz/gin-api/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	//  cors 跨域
	r.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
	}))

	r.NoRoute(api.NotFound)
	r.GET("/healthz", api.Healthz)
	r.GET("", api.Endpoints)
	r.GET("teapot", api.Teapot)

	// 文件路由组
	fileGroup := r.Group("/file").Use(middleware.LimitByRequest(3))
	{
		// 简单上传
		fileGroup.POST("", file.UploadFile)
		// 分块上传
		fileGroup.POST("/chunk", file.UploadChunk)
		fileGroup.POST("/chunk/merge", file.MergeChunk)
		fileGroup.GET("/chunk/state", file.ChunkState)
	}
	// auth 路由组
	authGroup := r.Group("/auth").Use(middleware.LimitByRequest(3))
	{
		authGroup.POST("/signup", auth.SignupByUsernamePwd)
		authGroup.POST("/login", auth.LoginByUsernamePwd)
		/* me jwt auth  */
		authGroup.GET("/me", middleware.JWTAuth(), auth.Me)
	}

	// s3
	s3Group := r.Group("/s3").Use(middleware.LimitByRequest(3))
	{
		// 直传
		s3Group.POST("/r2/upload",
			middleware.FileSizeLimit(100*1024*1024),
			r2.Upload,
		)
		// https://docs.aws.amazon.com/amazonglacier/latest/dev/uploading-an-archive-mpu-using-rest.html
		s3Group.POST("/r2/upload/:key", r2.CreateMultipartUpload)
		s3Group.PUT("/r2/upload/:key", r2.UploadPart)
		s3Group.POST("/r2/upload/complete", r2.CompleteMultipartUpload)
		s3Group.DELETE("/r2/upload/:key", r2.AbortMultipartUpload)
		s3Group.GET("/r2/upload/:key", r2.ListParts)
		s3Group.GET("/r2/upload", r2.ListMultipartUploads)
	}

	imagesTypes := []string{"image/jpeg", "image/png", "image/gif", "image/webp"}
	imagesGroup := r.Group("/image")
	{
		imagesGroup.POST("",
			middleware.FileSizeLimit(10*1024*1024),
			middleware.FileMimeTypeLimit(imagesTypes),
			r2.Upload)

		imagesGroup.POST("/classify/file",
			middleware.FileSizeLimit(10*1024*1024),
			middleware.FileMimeTypeLimit(imagesTypes),
			classify.ImageFile)
	}

	postsGroup := r.Group("/posts").Use(middleware.LimitByRequest(3))
	{
		postsGroup.GET("/:id", posts.GetPostById)
		postsGroup.DELETE("/:id", middleware.JWTAuth(), posts.DeletePostById)
		postsGroup.POST("", middleware.JWTAuth(), posts.CreatePost)
		postsGroup.GET("", posts.FindPosts)
	}

	projectsGroup := r.Group("/projects").Use(middleware.LimitByRequest(3))
	{
		projectsGroup.GET("", projects.FindProjects)
		projectsGroup.POST("", middleware.JWTAuth(), projects.CreateProject)
		projectsGroup.GET("/:id", projects.FindProjectByID)
		projectsGroup.DELETE("/:id", middleware.JWTAuth(), projects.DeleteProjectByID)
	}

	return r
}
