package types

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type DB struct{}

func (db *DB) LogResponse(message string) error {
	return nil
}

type API struct{}

func (api *API) NotifyServer(message string) error {
	return nil
}

type badDB struct{}

func (db *badDB) LogResponse(message string) error {
	return fmt.Errorf("failed to write response to DB")
}

type badAPI struct{}

func (api *badAPI) NotifyServer(message string) error {
	return fmt.Errorf("failed to notify server")
}

var rGood = Request{
	User:    "1",
	Message: "Hello",
	Time:    time.Now(),
}

var removingUser = Request{
	Message: "Hello",
	Time:    time.Now(),
}
var removingMessage = Request{
	Time: time.Now(),
}
var removingTime = Request{
	User:    "1",
	Message: "Hello",
}

func TestResponder_sendResponse(t *testing.T) {
	payloadGood, err := json.Marshal(rGood)
	if err != nil {
		t.Error(err)
	}
	payloadUser, err := json.Marshal(removingUser)
	if err != nil {
		t.Error(err)
	}
	payloadMessage, err := json.Marshal(removingMessage)
	if err != nil {
		t.Error(err)
	}
	payloadTime, err := json.Marshal(removingTime)
	if err != nil {
		t.Error(err)
	}
	tests := []struct {
		name      string
		r         *Responder
		request   []byte
		wantFail  bool
		wantError error
	}{
		{
			name: "success",
			r: &Responder{
				dbConn:  &DB{},
				apiConn: &API{},
			},
			request:  payloadGood,
			wantFail: false,
		},
		{
			name: "fail, removing user",
			r: &Responder{
				dbConn:  &DB{},
				apiConn: &API{},
			},
			request:   payloadUser,
			wantFail:  true,
			wantError: fmt.Errorf("user not found"),
		},
		{
			name: "fail, removing message",
			r: &Responder{
				dbConn:  &DB{},
				apiConn: &API{},
			},
			request:   payloadMessage,
			wantFail:  true,
			wantError: fmt.Errorf("message not found"),
		},
		{
			name: "fail, removing time",
			r: &Responder{
				dbConn:  &DB{},
				apiConn: &API{},
			},
			request:   payloadTime,
			wantFail:  true,
			wantError: fmt.Errorf("message not time"),
		},
		{
			name: "fail, DB fail",
			r: &Responder{
				dbConn:  &badDB{},
				apiConn: &API{},
			},
			request:   payloadGood,
			wantFail:  true,
			wantError: fmt.Errorf("fail log response: failed to write response to DB"),
		},
		{
			name: "fail, API fail",
			r: &Responder{
				dbConn:  &DB{},
				apiConn: &badAPI{},
			},
			request:   payloadGood,
			wantFail:  true,
			wantError: fmt.Errorf("fail to notify server response: failed to notify server"),
		},
		{
			name: "fail, removing time",
			r: &Responder{
				dbConn:  &DB{},
				apiConn: &API{},
			},
			request:   payloadTime,
			wantFail:  true,
			wantError: fmt.Errorf("message not time"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.r.sendResponse(tt.request)
			if err != nil {
				if !tt.wantFail {
					t.Errorf("Responder.sendResponse() error = %v, wantFail %v", err, tt.wantFail)
				}
				if err.Error() != tt.wantError.Error() {
					t.Errorf("Responder.sendResponse() error = %v, wantError %v", err, tt.wantError)
				}
			}
		})
	}
}

func Test_validateRequest(t *testing.T) {
	payloadGood, err := json.Marshal(rGood)
	if err != nil {
		t.Error(err)
	}
	payloadUser, err := json.Marshal(removingUser)
	if err != nil {
		t.Error(err)
	}
	payloadMessage, err := json.Marshal(removingMessage)
	if err != nil {
		t.Error(err)
	}
	payloadTime, err := json.Marshal(removingTime)
	if err != nil {
		t.Error(err)
	}
	tests := []struct {
		name      string
		request   []byte
		want      string
		wantFail  bool
		wantError error
	}{
		{
			name:     "success",
			request:  payloadGood,
			want:     "Hello",
			wantFail: false,
		},
		{
			name:      "fail, removing user",
			request:   payloadUser,
			wantFail:  true,
			wantError: fmt.Errorf("user not found"),
		},
		{
			name:      "fail, removing message",
			request:   payloadMessage,
			wantFail:  true,
			wantError: fmt.Errorf("message not found"),
		},
		{
			name:      "fail, removing time",
			request:   payloadTime,
			wantFail:  true,
			wantError: fmt.Errorf("message not time"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validateRequest(tt.request)
			if (err != nil) != tt.wantFail {
				t.Errorf("validateRequest() error = %v, wantFail %v", err, tt.wantFail)
				return
			}
			if got != tt.want {
				t.Errorf("validateRequest() = %v, want %v", got, tt.wantError)
			}
		})
	}
}
