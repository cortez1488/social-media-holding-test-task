package rest

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) getUsers(c *gin.Context) {
	result, err := h.userInfoService.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "server error")
		logrus.Fatal(errors.New("h.userInfoService.GetUsers: " + err.Error()))
	}
	c.JSON(200, result)
}

func (h *Handler) getUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "server error")
		logrus.Fatal(errors.New("parse param" + err.Error()))
	}
	result, err := h.userInfoService.GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "server error")
		logrus.Fatal(errors.New("h.userInfoService.GetUser: " + err.Error()))
	}
	c.JSON(200, result)
}

func (h *Handler) getHistoryByTgID(c *gin.Context) {
	chatID, err := strconv.ParseInt(c.Query("chatid"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "server error")
		logrus.Fatal(errors.New("parse param" + err.Error()))
	}
	history, err := h.userInfoService.GetHistoryByTgID(chatID)
	c.JSON(200, history)
}

func (h *Handler) deleteIp(c *gin.Context) {
	ip := c.Query("ip")
	err := h.userInfoService.DeleteIp(ip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "server error")
		logrus.Fatal("h.userInfoService.DeleteIp: " + err.Error())
	}
	c.JSON(200, "deleted")
}
