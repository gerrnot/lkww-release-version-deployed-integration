package getlatestrelease

import (
	"context"
	"errors"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"github.com/digital-ai/release-integration-template-go/my-integration/api"
	"log/slog"
	"net/http"
)

func GetLatestRelease(portalBaseUrl string, releaseName string, failIfNotFound bool) (*task.Result, error) {
	// init
	ctx := context.Background()
	l := slog.Default()

	// validate inputs
	if len(portalBaseUrl) == 0 {
		return nil, errors.New("the 'portalBaseUrl' parameter cannot be empty")
	}
	if len(releaseName) == 0 {
		return nil, errors.New("the 'releaseName' parameter cannot be empty")
	}

	// query portal
	portalHttpClient, err := api.NewClientWithResponses(portalBaseUrl)
	if err != nil {
		l.Error("Failed to create http client", "err", err)
		return nil, err
	}
	l.Info("querying portal")
	res, err := portalHttpClient.GetReleasesByNameWithResponse(ctx, releaseName, func(ctx context.Context, req *http.Request) error {
		req.Header.Add("Accept", "application/json")
		return nil
	})
	if err != nil {
		l.Error("Failed to get releases", "err", err)
		return nil, err
	}
	var release *api.Release
	if res.JSON200 != nil {
		for i, currRelease := range *interface{}(res.JSON200).(*[]api.Release) {
			if i == 0 {
				release = &currRelease
			} else {
				err := errors.New("more than one releases found")
				l.Error(err.Error())
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
