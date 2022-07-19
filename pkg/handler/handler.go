package handler

import (
	"app/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		groups := api.Group("/groups")
		{
			groups.GET("/", h.listGroup)
			groups.POST("/", h.createGroup)
			groups.GET("/:id", h.getItemGroup)
			groups.PUT("/:id", h.updateItemGroup)
			groups.DELETE("/:id", h.deleteItemGroup)
			groupsInfo := groups.Group("/info")
			{
				groupsInfo.GET("/:id", h.getItemGroupInfo)
			}
		}
		users := api.Group("/users")
		{
			users.POST("/", h.createUser)
		}
	}

	return router
}
