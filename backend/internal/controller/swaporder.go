package controller

import (
	"net/http"
	"strconv"

	"cyskillswap/internal/errors"
	"cyskillswap/internal/model"
	"cyskillswap/internal/service"
	"github.com/gin-gonic/gin"
)

// SwapOrders 返回全部交换单。
func SwapOrders(c *gin.Context) { c.JSON(http.StatusOK, service.SwapOrders()) }

// ConfirmSwapOrder 交换双方之一确认交换单。
func ConfirmSwapOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respondError(c, errors.Validation("交换单 ID 不合法"))
		return
	}
	var input model.ConfirmInput
	if bindErr := c.ShouldBindJSON(&input); bindErr != nil {
		respondError(c, errors.Validation("请求体格式不正确"))
		return
	}
	order, err := service.ConfirmSwapOrder(id, input.User)
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, order)
}
