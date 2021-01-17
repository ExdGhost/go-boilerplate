package utils

import (
	log "go-boilerplate-api/pkg/utils/logger"

	"github.com/gin-gonic/gin"
	"github.com/ralstan-vaz/go-errors"
	"github.com/ralstan-vaz/go-errors/http"
)

// Error ... struct used for sending errors
type Error struct {
	Code        string `json:"code" example:"MODULE.ERROR_CODE"`
	Message     string `json:"message" example:"Oops! Something went wrong"` // mainly used for frontend
	Description string `json:"description" example:"Error description"`
}

// HandleError formats, logs and sets a http response for the error
func HandleError(c *gin.Context, errObj *error, reference ...interface{}) {

	if *errObj == nil {
		return
	}

	err := errors.Get(*errObj)

	if err.Message == "" {
		err.Message = "Something Went Wrong"
	}

	if len(reference) != 0 && reference[0] != nil {
		log.Error(err.Code, err.Description, log.Priority1, err.Source, reference[0])
	} else {
		log.Error(err.Code, err.Description, log.Priority1, err.Source)
	}

	statusCode := http.StatusCode(err)

	// Forms error response
	var errResp Error

	errResp.Code = err.Code
	errResp.Message = err.Message
	errResp.Description = err.Description

	c.JSON(statusCode, errResp)
}
