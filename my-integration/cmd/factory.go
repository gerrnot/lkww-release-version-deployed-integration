package cmd

import (
	"fmt"
	"github.com/digital-ai/release-integration-sdk-go/api/release/openapi"
	"github.com/digital-ai/release-integration-sdk-go/http"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"github.com/digital-ai/release-integration-sdk-go/task/command"
	"k8s.io/klog/v2"
)

const (
	getLatestRelease = "myPlugin.GetLatestRelease"
)

type CommandFactory struct {
	httpClient    *http.HttpClient
	releaseClient *openapi.APIClient
}

func NewCommandFactory(httpClient *http.HttpClient, releaseClient *openapi.APIClient) *CommandFactory {
	return &CommandFactory{httpClient: httpClient, releaseClient: releaseClient}
}
func (factory *CommandFactory) InitCommand(commandType command.CommandType) (command.CommandExecutor, error) {
	if spawnCommand, present := commandHatchery[commandType]; present {
		klog.Infof("Building task [%s]", commandType)
		return spawnCommand(factory), nil
	} else {
		task.Comment(fmt.Sprintf("Cannot create command of a type [%s]", commandType))
		return nil, fmt.Errorf("unknown command type [%s]", commandType)
	}
}

var commandHatchery = map[command.CommandType]func(*CommandFactory) command.CommandExecutor{
	getLatestRelease: func(factory *CommandFactory) command.CommandExecutor {
		return &GetLatestRelease{}
	},
}
