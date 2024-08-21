package cmd

type GetLatestRelease struct {
	PortalBaseUrl   string `json:"portalBaseUrl"`
	HelmReleaseName string `json:"helmReleaseName"`
	FailIfNotFound  bool   `json:"failIfNotFound"`
}
