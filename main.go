//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=portal/api/oapi-codegen.yaml ./portal/api/portal-api.yaml
package main

import (
	"context"
	"github.com/digital-ai/release-integration-sdk-go/api/release"
	"github.com/digital-ai/release-integration-sdk-go/runner"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"github.com/digital-ai/release-integration-sdk-go/task/command"
	"github.com/digital-ai/release-integration-template-go/portal/cmd"
	"os"
)

var PluginVersion = os.Getenv("VERSION")
var BuildDate = os.Getenv("BUILD_DATE")

func prepareCommandFactory(input task.InputContext) (command.CommandFactory, error) {
	ctx := task.ReleaseContext{
		Id: input.Release.Id,
		AutomatedTaskAsUser: task.AutomatedTaskAsUserContext{
			Username: input.Release.AutomatedTaskAsUser.Username,
			Password: input.Release.AutomatedTaskAsUser.Password,
		},
		Url: input.Release.Url,
	}
	releaseClient, err := release.NewReleaseApiClient(ctx)
	if err != nil {
		return nil, err
	}

	return cmd.NewCommandFactory(releaseClient), nil
}

var commandRunner = runner.NewCommandRunner(prepareCommandFactory)

func main() {
	context.Background()
	runner.Execute(PluginVersion, BuildDate, commandRunner)
}
