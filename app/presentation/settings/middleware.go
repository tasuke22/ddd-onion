package settings

import (
	"errors"

	"github.com/gin-gonic/gin"
	errDomain "github.com/tasuke/go-onion/domain/error"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			var e *errDomain.Error
			switch {
			case errors.As(err.Err, &e):
				if errors.Is(err, errDomain.NotFoundErr) {
					ReturnNotFound(c, e)
				}
				ReturnStatusBadRequest(c, e)
			default:
				ReturnStatusInternalServerError(c, e)
			}
		}
	}
}
