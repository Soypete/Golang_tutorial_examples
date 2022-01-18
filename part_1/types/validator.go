package types

import (
	"encoding/json"
	"fmt"
	"time"
)

type apiConnector interface {
	NotifyServer(string) error
}

type dbConnector interface {
	LogResponse(string) error
}

type Responder struct {
	dbConn  dbConnector
	apiConn apiConnector
}

// sendResponse recieves a json request, checks it is valid, and returns the response.
// The json is of the form:
// {
// 	"user": "johndoe",
// 	"message": "Hello World",
// time: "2019-01-01T00:00:00Z"
// }
func (r *Responder) sendResponse(request []byte) error {
	// check that all parts of the request are present
	strRequest, err := validateRequest(request)
	if err != nil {
		return fmt.Errorf("failed to validate request: %w", err)
	}

	// log response
	err = r.dbConn.LogResponse(strRequest)
	if err != nil {
		// TODO: update the error
		return fmt.Errorf("fail log response: %w", err)
	}

	err = r.apiConn.NotifyServer(strRequest)
	if err != nil {
		// TODO: update the error
		return fmt.Errorf("fail to notify server response: %w", err)
	}

	return nil
}

// Request is the json request that is sent to the server.
type Request struct {
	User    string    `json:"user"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

// validateRequest checks that all the required fields are present in the request.
func validateRequest(request []byte) (string, error) {
	var r Request
	err := json.Unmarshal(request, &r)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal request: %w", err)
	}
	// TODO: add code here to check that all the required fields are present

	return r.Message, nil
}
