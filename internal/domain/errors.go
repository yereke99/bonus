package domain

import "errors"

var (
	ErrCallExecution = errors.New("error call execution")
	ErrEmptyParams   = errors.New("error empty config")
	ErrExistsTable   = errors.New("error tables is exists")
	ErrStatusCode    = errors.New("error status code not 200")
	ErrADBDevices    = errors.New("error checking ADB status")
	ErrPingServers   = errors.New("error ping servers")
	ErrTimeOut       = errors.New("ping command timed out")
	ErrNoHubs        = errors.New("no hubs found")
	ErrNoDevices     = errors.New("no devices found")
	ErrUserNotFound  = errors.New("user not found")
)
