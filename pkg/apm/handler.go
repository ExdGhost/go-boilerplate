package apm

import (
	"context"
	"net/http"
)

// HandlerInterface ... A wrapper interface on top of the apm to help out during testing
// A Dummy interface which needs to be implemented
type HandlerInterface interface {
	StartTransaction(name string) (transaction interface{}, err error)
	EndTransaction(transaction interface{}, er error) (err error)
	StartSegment(ctx context.Context, segmentName string) (segment interface{}, err error)
	EndSegment(segment interface{}) (err error)
	StartExternalSegment(ctx context.Context, URL string) (externalSegment interface{}, err error)
	EndExternalSegment(segment interface{}) error
	StartExternalWebSegment(ctx context.Context, req *http.Request) (externalSegment interface{}, err error)
	NoticeError(transaction interface{}, err error) error
	AddAttribute(transaction interface{}, key string, val interface{}) error
}

// Handler ...
type Handler struct {
}

// NewApmHandler ...
func NewApmHandler() HandlerInterface {
	return &Handler{}
}

// StartTransaction ... Interface function that calls the internal apm functions
func (a *Handler) StartTransaction(name string) (transaction interface{}, err error) {
	return nil, nil
}

// EndTransaction .. Interface function that calls the internal apm functions
func (a *Handler) EndTransaction(transaction interface{}, er error) (err error) {
	return nil
}

// StartSegment .. Interface function that calls the internal apm functions
func (a *Handler) StartSegment(ctx context.Context, segmentName string) (segment interface{}, err error) {
	return nil, nil
}

// EndSegment .. Interface function that calls the internal apm functions
func (a *Handler) EndSegment(segment interface{}) (err error) {
	return nil
}

// StartExternalSegment .. Interface function that calls the internal apm functions
func (a *Handler) StartExternalSegment(ctx context.Context, URL string) (externalSegment interface{}, err error) {
	return nil, nil
}

// EndExternalSegment .. Interface function that calls the internal apm functions
func (a *Handler) EndExternalSegment(segment interface{}) error {
	return nil
}

// StartExternalWebSegment .. Interface function that calls the internal apm functions
func (a *Handler) StartExternalWebSegment(ctx context.Context, req *http.Request) (externalSegment interface{}, err error) {
	return nil, nil
}

// NoticeError .. Interface function that calls the internal apm functions
func (a *Handler) NoticeError(transaction interface{}, err error) error {
	return nil
}

// AddAttribute .. Interface function that calls the internal apm functions
func (a *Handler) AddAttribute(transaction interface{}, key string, val interface{}) error {
	return nil
}
