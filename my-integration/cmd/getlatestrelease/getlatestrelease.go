package getlatestrelease

import (
	"errors"
	"github.com/digital-ai/release-integration-sdk-go/task"
)

func GetLatestRelease(portalBaseUrl string, helmReleaseName string, failIfNotFound bool) (*task.Result, error) {
	// validate inputs
	if len(portalBaseUrl) == 0 {
		return nil, errors.New("the 'portalBaseUrl' parameter cannot be empty")
	}
	if len(helmReleaseName) == 0 {
		return nil, errors.New("the 'helmReleaseName' parameter cannot be empty")
	}

	// query portal
	// TODO: Gernot
	version := "1.2.3"

	// return output
	return task.NewResult().String("version", version), nil
}
