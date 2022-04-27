package rest

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	userInfoService UserInfoService
}

func NewHandler(userInfoService UserInfoService) *Handler {
	return &Handler{userInfoService: userInfoService}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/get_users", h.getUsers)
		api.GET("/get_user", h.getUser)
		api.GET("/get_history_by_tg", h.getHistoryByTgID)
		api.DELETE("/delete_ip", h.deleteIp)
	}
	return router
}
