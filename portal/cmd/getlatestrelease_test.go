package cmd

import (
	"github.com/digital-ai/release-integration-template-go/task/server"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

const validSemVerRegex = `^\d+\.\d+\.\d+$`

func TestGetLatestRelease(t *testing.T) {
	type args struct {
		server         server.Server
		releaseName    string
		failIfNotFound bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "happy path", args: struct {
			server         server.Server
			releaseName    string
			failIfNotFound bool
		}{
			server: server.Server{
				Url:         "https://portal.test.lkw-walter.com/api",
				Certificate: "",
				Username:    "",
				Password:    "",
			},
			releaseName:    "aggtier-aufenthalte",
			failIfNotFound: true,
		},
			want: validSemVerRegex, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetLatestRelease(tt.args.server, tt.args.releaseName, tt.args.failIfNotFound)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLatestRelease() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			resultMap, err := got.Get()
			if err != nil {
				t.Errorf("Could not get result map: %v", err)
			}
			version := resultMap["version"].(string)
			matched, err := regexp.MatchString(tt.want, version)
			assert.True(t, matched)
		})
	}
}
