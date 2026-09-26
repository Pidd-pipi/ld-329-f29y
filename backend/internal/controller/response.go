package controller

import (
	"net/http"

	bizerr "cyskillswap/internal/errors"
	"cyskillswap/internal/logger"
	"cyskillswap/internal/model"
	"cyskillswap/internal/service"
	"github.com/gin-gonic/gin"
)

// CreateNeedResponse POST /api/needs/:id/responses
func CreateNeedResponse(c *gin.Context) {
	needID, ok := parseID(c)
	if !ok {
		return
	}
	var in model.CreateResponseInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, bizerr.ErrValidation)
		return
	}
	resp, err := service.CreateResponse(needID, in)
	if err != nil {
		writeError(c, err)
		return
	}
	logger.Info("response created:", resp.ID, "need:", needID, "responder:", resp.Responder)
	c.JSON(http.StatusCreated, resp)
}

// NeedResponses GET /api/needs/:id/responses
func NeedResponses(c *gin.Context) {
	needID, ok := parseID(c)
	if !ok {
		return
	}
	responses, err := service.ListNeedResponses(needID)
	if err != nil {
		writeError(c, err)
		return
	}
	c.JSON(http.StatusOK, responses)
}
