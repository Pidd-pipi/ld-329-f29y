package controller

import (
	"net/http"

	bizerr "cyskillswap/internal/errors"
	"cyskillswap/internal/logger"
	"cyskillswap/internal/model"
	"cyskillswap/internal/service"
	"github.com/gin-gonic/gin"
)

// SelectNeedResponse POST /api/needs/:id/select 发布者选中响应，生成交换单
func SelectNeedResponse(c *gin.Context) {
	needID, ok := parseID(c)
	if !ok {
		return
	}
	var in model.SelectResponseInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, bizerr.ErrValidation)
		return
	}
	order, err := service.SelectResponse(needID, in)
	if err != nil {
		writeError(c, err)
		return
	}
	logger.Info("order created:", order.ID, "need:", needID, "response:", in.ResponseID)
	c.JSON(http.StatusCreated, order)
}

// Orders GET /api/orders
func Orders(c *gin.Context) { c.JSON(http.StatusOK, service.ListOrders()) }

// ConfirmOrder POST /api/orders/:id/confirm 参与者确认交换单
func ConfirmOrder(c *gin.Context) {
	orderID, ok := parseID(c)
	if !ok {
		return
	}
	var in model.ConfirmOrderInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, bizerr.ErrValidation)
		return
	}
	order, err := service.ConfirmOrder(orderID, in.User)
	if err != nil {
		writeError(c, err)
		return
	}
	logger.Info("order confirmed by:", in.User, "order:", orderID, "status:", order.Status)
	c.JSON(http.StatusOK, order)
}
