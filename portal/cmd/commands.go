package cmd

import "github.com/digital-ai/release-integration-template-go/task/server"

type GetLatestReleaseC struct {
	Server         server.Server `json:"server"`
	ReleaseName    string        `json:"releaseName"`
	FailIfNotFound bool          `json:"failIfNotFound"`
}
