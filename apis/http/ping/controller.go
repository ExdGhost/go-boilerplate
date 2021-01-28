package ping

import (
	"go-boilerplate/shared"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response is a struct storing the response.
// Message : the message field of the response.
type Response struct {
	Message  string `json:"message"`
	App      string `json:"app"`
	CommitID string `json:"commitId"`
}

// pongResponse is a constant used for sending response message.
const pongResponse string = "PONG"

// Ping ... used as a pointer receiver
type Ping struct {
}

// NewPingService ...
func NewPingService() *Ping {
	return &Ping{}
}

// Get godoc
// @Summary Get ping response
// @Description Get ping response
// @Tags ping
// @Accept  json
// @Produce  json
// @Success 200 {object} ping.Response
// @Failure 400 {object} utils.Error
// @Router /ping [get]
// Get returns a response for the /ping request.
// ctx *gin.Context : allows us to pass variables between middleware
func (ping *Ping) Get(ctx *gin.Context) {
	response := Response{Message: pongResponse, App: "go-boilerplate", CommitID: shared.VERSION}
	ctx.JSON(http.StatusOK, response)
}
