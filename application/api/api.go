package api

import (
	"net/http"

	"github.com/letian0805/seckill/infrastructure/utils"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type Event struct{}

type Shop struct{}

func (e *Event) List(ctx *gin.Context) {
	resp := &utils.Response{
		Code: 0,
		Data: nil,
		Msg:  "ok",
	}
	status := http.StatusOK

	logrus.Info("event list")

	ctx.JSON(status, resp)
}

func (e *Event) Info(ctx *gin.Context) {
	resp := &utils.Response{
		Code: 0,
		Data: nil,
		Msg:  "ok",
	}
	status := http.StatusOK

	logrus.Info("event info")

	ctx.JSON(status, resp)
}

func (e *Event) Subscribe(ctx *gin.Context) {
	resp := &utils.Response{
		Code: 0,
		Data: nil,
		Msg:  "ok",
	}
	status := http.StatusOK

	logrus.Info("event subscribe")

	ctx.JSON(status, resp)
}

func (s *Shop) AddCart(ctx *gin.Context) {
	resp := &utils.Response{
		Code: 0,
		Data: nil,
		Msg:  "ok",
	}
	status := http.StatusOK

	logrus.Info("shop add cart")

	ctx.JSON(status, resp)
}
