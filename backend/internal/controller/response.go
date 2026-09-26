package controller

import (
	"net/http"
	"strconv"

	"cyskillswap/internal/errors"
	"cyskillswap/internal/model"
	"cyskillswap/internal/service"
	"github.com/gin-gonic/gin"
)

// ListNeedResponses 查看某条需求下的全部响应。
func ListNeedResponses(c *gin.Context) {
	needID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respondError(c, errors.Validation("需求 ID 不合法"))
		return
	}
	responses, err := service.ListResponses(needID)
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, responses)
}

// CreateNeedResponse 响应者针对需求提交响应。
func CreateNeedResponse(c *gin.Context) {
	needID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respondError(c, errors.Validation("需求 ID 不合法"))
		return
	}
	var input model.ResponseInput
	if bindErr := c.ShouldBindJSON(&input); bindErr != nil {
		respondError(c, errors.Validation("请求体格式不正确"))
		return
	}
	resp, err := service.CreateResponse(needID, input)
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// AcceptNeedResponse 发布者选中一条等待中的响应，生成交换单。
func AcceptNeedResponse(c *gin.Context) {
	needID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respondError(c, errors.Validation("需求 ID 不合法"))
		return
	}
	responseID, err := strconv.Atoi(c.Param("responseId"))
	if err != nil {
		respondError(c, errors.Validation("响应 ID 不合法"))
		return
	}
	order, err := service.AcceptResponse(needID, responseID)
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, order)
}
