package Routes

import (
	"appdirs/cns-parser/Controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/api/v1")
	{
		grp1.GET("consumer", Controllers.GetConsumer)
		grp1.POST("consumer", Controllers.CreateConsumer)
		grp1.GET("consumer/:id", Controllers.GetConsumerByID)
		grp1.PUT("consumer/:id", Controllers.UpdateConsumer)
		grp1.DELETE("consumer/:id", Controllers.DeleteConsumer)
		grp1.POST("webhook/:id", Controllers.WebHook)
		grp1.GET("pullhook/:id", Controllers.PullHook)

	}

	return r
}
