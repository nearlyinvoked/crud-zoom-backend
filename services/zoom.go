package services

import (
	"bytes"
	"crud-zoom/config"
	"crud-zoom/repositories"
	"encoding/json"
	"io"
	"net/http"

	"go.uber.org/zap"
)

type ZoomSvc interface {
	ListMeeting() (map[string]interface{}, error)
	CreateMeeting(agenda string, meetingTime string) (map[string]interface{}, error)
	UpdateMeeting(meetingID string, agenda string, meetingTime string) (int, error)
	DeleteMeeting(meetingID string) (int, error)
}

type ZoomService struct {
	cfg             config.Config
	logger          *zap.Logger
	zoomRepo        repositories.ZoomRepository
}

func NewZoomService(cfg config.Config, logger *zap.Logger, zoomRepo repositories.ZoomRepository) *ZoomService {
	return &ZoomService{
		cfg:             cfg,
		logger:          logger,
		zoomRepo:        zoomRepo,
	}
}

func (z *ZoomService) ListMeeting() (map[string]interface{}, error) {
	// Create the HTTP request
	req, err := http.NewRequest("GET", "https://api.zoom.us/v2/users/kevin.test.zoom%40gmail.com/meetings", nil)
	if err != nil {
			z.logger.Error("Failed to create HTTP request", zap.Error(err))
			return nil, err
	}

	// Set the request headers
	req.Header.Set("Authorization", "Bearer "+z.cfg.ZoomAccessToken)

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
			z.logger.Error("Failed to send HTTP request", zap.Error(err))
			return nil, err
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
			z.logger.Error("Failed to list meetings", zap.Int("status", resp.StatusCode))
			return nil, err
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
			z.logger.Error("Failed to read response body", zap.Error(err))
			return nil, err
	}

	// Unmarshal the response body into a map
	var responseBody map[string]interface{}
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
			z.logger.Error("Failed to unmarshal response body", zap.Error(err))
			return nil, err
	}

	z.logger.Info("Meetings listed successfully")
	return responseBody, nil
}

func (z *ZoomService) CreateMeeting(agenda string, meetingTime string) (map[string]interface{}, error) {
	// Define the meeting details
	meetingDetails := map[string]interface{}{
			"agenda":           agenda,
			"default_password": false,
			"duration":         60,
			"password":         "123456",
			"pre_schedule":     false,
			"schedule_for":     "kevin.test.zoom@gmail.com",
			"start_time": meetingTime,
			"timezone":   "Indonesia/Jakarta",
			"topic":      agenda,
			"type":       2,
	}

	// Convert meeting details to JSON
	meetingDetailsJSON, err := json.Marshal(meetingDetails)
	if err != nil {
			z.logger.Error("Failed to marshal meeting details", zap.Error(err))
			return nil, err
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", "https://api.zoom.us/v2/users/kevin.test.zoom%40gmail.com/meetings", bytes.NewBuffer(meetingDetailsJSON))
	if err != nil {
			z.logger.Error("Failed to create HTTP request", zap.Error(err))
			return nil, err
	}

	// Set the request headers
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+z.cfg.ZoomAccessToken)

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
			z.logger.Error("Failed to send HTTP request", zap.Error(err))
			return nil, err
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusCreated {
			z.logger.Error("Failed to create meeting", zap.Int("status", resp.StatusCode))
			return nil, err
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
			z.logger.Error("Failed to read response body", zap.Error(err))
			return nil, err
	}

	// Unmarshal the response body into a map
	var responseBody map[string]interface{}
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
			z.logger.Error("Failed to unmarshal response body", zap.Error(err))
			return nil, err
	}

	z.logger.Info("Meeting created successfully")
	return responseBody, nil
}

func (z *ZoomService) UpdateMeeting(meetingID string, agenda string, meetingTime string) (int, error) {
	// Define the meeting ID and update details
	updateDetails := map[string]interface{}{
		"agenda":       agenda,
		"duration":     60,
		"password":     "123456",
		"pre_schedule": false,
		"schedule_for": "kevin.test.zoom@gmail.com",
		"start_time": meetingTime,
		"template_id": "5Cj3ceXoStO6TGOVvIOVPA==",
		"timezone":   "Indonesia/Jakarta",
		"topic":      agenda,
		"type": 2,
	}

	// Convert update details to JSON
	updateDetailsJSON, err := json.Marshal(updateDetails)
	if err != nil {
		z.logger.Error("Failed to marshal update details", zap.Error(err))
		return http.StatusInternalServerError, err
	}

	// Create the HTTP request
	req, err := http.NewRequest("PATCH", "https://api.zoom.us/v2/meetings/"+meetingID, bytes.NewBuffer(updateDetailsJSON))
	if err != nil {
		z.logger.Error("Failed to create HTTP request", zap.Error(err))
		return http.StatusInternalServerError, err
	}

	// Set the request headers
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+z.cfg.ZoomAccessToken)

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		z.logger.Error("Failed to send HTTP request", zap.Error(err))
		return http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusNoContent {
		z.logger.Error("Failed to update meeting", zap.Int("status", resp.StatusCode))
		return http.StatusNoContent, err
	}

	z.logger.Info("Meeting updated successfully")
	return resp.StatusCode, nil
}

func (z *ZoomService) DeleteMeeting(meetingID string) (int, error) {
	// Create the HTTP request
	req, err := http.NewRequest("DELETE", "https://api.zoom.us/v2/meetings/"+meetingID, nil)
	if err != nil {
		z.logger.Error("Failed to create HTTP request", zap.Error(err))
		return http.StatusInternalServerError, err
	}

	// Set the request headers
	req.Header.Set("Authorization", "Bearer "+z.cfg.ZoomAccessToken)

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		z.logger.Error("Failed to send HTTP request", zap.Error(err))
		return http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusNoContent {
		z.logger.Error("Failed to delete meeting", zap.Int("status", resp.StatusCode))
		return resp.StatusCode, err
	}

	z.logger.Info("Meeting deleted successfully")
	return resp.StatusCode, nil
}