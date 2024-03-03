package route

import (
	"github.com/gin-gonic/gin"
	"github.com/tasuke/go-onion/infrastructure/repository"
	"github.com/tasuke/go-onion/presentation/health_handler"
	"github.com/tasuke/go-onion/presentation/settings"
	userPre "github.com/tasuke/go-onion/presentation/user"
	userUse "github.com/tasuke/go-onion/usecase/user"
)

func InitRoute(api *gin.Engine) {
	api.Use(settings.ErrorHandler())
	v1 := api.Group("/v1")
	v1.GET("/health", health_handler.HealthCheck)

	{
		userRoute(v1)
	}
}

func userRoute(r *gin.RouterGroup) {
	userRepository := repository.NewUserRepository()
	tagRepository := repository.NewTagRepository()
	h := userPre.NewHandler(
		userUse.NewUserUseCase(userRepository, tagRepository),
		userUse.NewUpdateUserUseCase(userRepository, tagRepository),
	)
	group := r.Group("/users")

	group.POST("/register", h.SaveUser)
	group.POST("/update", h.UpdateUser)
}
