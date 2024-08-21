package getlatestrelease

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

const validSemVerRegex = `^\d+\.\d+\.\d+$`

func TestGetLatestRelease(t *testing.T) {
	type args struct {
		portalBaseUrl   string
		helmReleaseName string
		failIfNotFound  bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "happy path", args: struct {
			portalBaseUrl   string
			helmReleaseName string
			failIfNotFound  bool
		}{
			portalBaseUrl:   "https://portal.test.lkw-walter.com/api",
			helmReleaseName: "aggtier-aufenthalte",
			failIfNotFound:  true,
		},
			want: validSemVerRegex, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetLatestRelease(tt.args.portalBaseUrl, tt.args.helmReleaseName, tt.args.failIfNotFound)
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
