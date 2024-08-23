package cmd

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"github.com/digital-ai/release-integration-template-go/portal/api"
	"github.com/digital-ai/release-integration-template-go/task/server"
	"log/slog"
	"net/http"
)

func GetLatestRelease(server server.Server, releaseName string, failIfNotFound bool) (*task.Result, error) {
	// validate inputs
	task.Comment(fmt.Sprintf("Release name: %s", releaseName))
	task.Comment(fmt.Sprintf("Intializing server with params: %s, %s, %s", server.Url, server.Username, server.Certificate))
	if len(server.Url) == 0 {
		return nil, errors.New("the 'server.Url' parameter cannot be empty")
	}
	if len(releaseName) == 0 {
		err := errors.New("the 'releaseName' parameter cannot be empty")
		return nil, err
	}

	// init
	ctx := context.Background()
	l := slog.Default()
	portalHttpClient, err := api.NewClientWithResponses(server.Url, func(client *api.Client) error {
		certPool := x509.NewCertPool()
		certPool.AppendCertsFromPEM([]byte(server.Certificate))
		client.Client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					ClientCAs: certPool,
				},
			},
		}
		return nil
	})
	if err != nil {
		l.Error("Failed to create http client", "err", err)
		return nil, err
	}

	// query portal
	l.Info("querying portal", "portalUrl", server.Url)
	res, err := portalHttpClient.GetReleasesByNameWithResponse(ctx, releaseName, func(ctx context.Context, req *http.Request) error {
		req.Header.Add("Accept", "application/json")
		req.SetBasicAuth(server.Username, server.Password)
		return nil
	})
	if err != nil {
		l.Error("failed to get releases", "err", err)
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
