package cmd

import (
	"context"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"github.com/digital-ai/release-integration-template-go/my-integration/cmd/getlatestrelease"
)

func (command *GetLatestRelease) FetchResult(ctx context.Context) (*task.Result, error) {
	return getlatestrelease.GetLatestRelease(command.PortalBaseUrl, command.HelmReleaseName, command.FailIfNotFound)
}
