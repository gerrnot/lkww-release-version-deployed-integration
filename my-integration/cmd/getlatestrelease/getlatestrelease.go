package getlatestrelease

import (
	"context"
	"errors"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"github.com/digital-ai/release-integration-template-go/my-integration/api"
	"google.golang.org/appengine/log"
	"net/http"
)

func GetLatestRelease(portalBaseUrl string, helmReleaseName string, failIfNotFound bool) (*task.Result, error) {
	// validate inputs
	if len(portalBaseUrl) == 0 {
		return nil, errors.New("the 'portalBaseUrl' parameter cannot be empty")
	}
	if len(helmReleaseName) == 0 {
		return nil, errors.New("the 'helmReleaseName' parameter cannot be empty")
	}

	// init
	ctx := context.Background()

	// query portal
	portalHttpClient, err := api.NewClientWithResponses(portalBaseUrl)
	if err != nil {
		log.Errorf(ctx, "Failed to create http client: %v", err)
		return nil, err
	}
	res, err := portalHttpClient.GetReleasesByNameWithResponse(ctx, helmReleaseName, func(ctx context.Context, req *http.Request) error {
		req.Header.Add("Accept", "application/json")
		return nil
	})
	if err != nil {
		log.Errorf(ctx, "Failed to get releases: %v", err)
		return nil, err
	}
	var release *api.Release
	if res.JSON200 != nil {
		for i, currRelease := range *interface{}(res.JSON200).(*[]api.Release) {
			if i == 0 {
				release = &currRelease
			} else {
				err := errors.New("more than one releases found")
				log.Errorf(ctx, "%v", err)
				return nil, err
			}
		}
	} else if release == nil && failIfNotFound {
		return nil, errors.New("release not found")

	} else if release == nil && !failIfNotFound {
		return task.NewResult().String("version", ""), nil
	}
	return task.NewResult().String("version", *release.Version), nil
}
