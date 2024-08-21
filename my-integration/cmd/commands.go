package cmd

type GetLatestRelease struct {
	PortalBaseUrl  string `json:"portalBaseUrl"`
	ReleaseName    string `json:"releaseName"`
	FailIfNotFound bool   `json:"failIfNotFound"`
}
