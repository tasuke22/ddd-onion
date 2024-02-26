package route

import (
	"github.com/gin-gonic/gin"
	"github.com/tasuke/go-onion/infrastructure/repository"
	"github.com/tasuke/go-onion/presentation/health_controller"
	userPre "github.com/tasuke/go-onion/presentation/user"
	userUse "github.com/tasuke/go-onion/usecase/user"
)

func InitRoute(api *gin.Engine) {
	v1 := api.Group("/v1")
	v1.GET("/health", health_controller.HealthCheck)

	{
		userRoute(v1)
	}
}

func userRoute(r *gin.RouterGroup) {
	userRepository := repository.NewUserRepository()
	tagRepository := repository.NewTagRepository()
	h := userPre.NewHandler(
		userUse.NewUserUseCase(userRepository, tagRepository),
	)
	group := r.Group("/users")
	group.POST("/save", h.SaveUser)
}
