package utils

import (
	log "go-boilerplate/pkg/utils/logger"

	"github.com/ralstan-vaz/go-errors"
	"github.com/ralstan-vaz/go-errors/grpc"
	"google.golang.org/grpc/status"
)

// HandleError formats, logs and sets a GRPC response for the error
func HandleError(errObj *error, reference ...interface{}) error {
	if *errObj == nil {
		return nil
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

	statusCode := grpc.StatusCode(err)

	return status.New(statusCode, err.Description).Err()
}
