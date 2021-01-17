package helpers

import (
	"log"

	errors "go-boilerplate-api/pkg/utils/go-errors"
)

func init() {
	er := errors.Parse(GetErrors())
	if er != nil {
		log.Panicln(er)
	}
}
