package cmd

import (
	"context"
	"github.com/digital-ai/release-integration-sdk-go/task"
)

func (command *GetLatestReleaseC) FetchResult(ctx context.Context) (*task.Result, error) {
	return GetLatestRelease(command.Server, command.ReleaseName, command.FailIfNotFound)
}
