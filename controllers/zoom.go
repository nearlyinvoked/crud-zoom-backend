package controllers

import (
	"crud-zoom/config"
	"crud-zoom/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ZoomController struct {
	cfg         config.Config
	logger      *zap.Logger
	zoomService services.ZoomSvc
}

func NewZoomController(cfg config.Config, logger *zap.Logger, zoomService services.ZoomSvc) *ZoomController {
	return &ZoomController{
		cfg:         cfg,
		logger:      logger,
		zoomService: zoomService,
	}
}

func (z *ZoomController) List(c *gin.Context) {
	res, err := z.zoomService.ListMeeting()
	if err != nil {
		z.logger.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)
}

func (z *ZoomController) Create(c *gin.Context) {
	var requestBody struct {
		Agenda      string `json:"agenda"`
		MeetingTime string `json:"meeting_time"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		z.logger.Error(err.Error())
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	res, err := z.zoomService.CreateMeeting(requestBody.Agenda, requestBody.MeetingTime)
	if err != nil {
		z.logger.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)
}

func (z *ZoomController) Update(c *gin.Context) {
	id := c.Param("id")

	var requestBody struct {
		Agenda      string `json:"agenda"`
		MeetingTime string `json:"meeting_time"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		z.logger.Error(err.Error())
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	res, err := z.zoomService.UpdateMeeting(id, requestBody.Agenda, requestBody.MeetingTime)
	if err != nil {
		z.logger.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if res < 300 {
		c.JSON(res, gin.H{"message": "Meeting updated successfully"})
		return
	} else {
		c.JSON(500, gin.H{"error": "Failed to update meeting"})
		return
	}
}

func (z *ZoomController) Delete(c *gin.Context) {
	id := c.Param("id")
	res, err := z.zoomService.DeleteMeeting(id)
	if err != nil {
		z.logger.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if res < 300 {
		c.JSON(res, gin.H{"message": "Meeting deleted successfully"})
		return
	} else {
		c.JSON(500, gin.H{"error": "Failed to delete meeting"})
		return
	}
}