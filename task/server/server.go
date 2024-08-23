package server

type Server struct {
	Url                  string `json:"url"`
	AuthenticationMethod string `json:"authenticationMethod"`
	VerifySSL            bool   `json:"insecure"`
	Username             string `json:"username"`
	Password             string `json:"password"`
	Domain               string `json:"domain"`
	ProxyHost            string `json:"proxyHost"`
	ProxyPort            string `json:"proxyPort"`
	ProxyUsername        string `json:"proxyUsername"`
	ProxyPassword        string `json:"proxyPassword"`
	ProxyDomain          string `json:"proxyDomain"`
	AccessTokenUrl       string `json:"accessTokenUrl"`
	ClientId             string `json:"clientId"`
	ClientSecret         string `json:"clientSecret"`
	Scope                string `json:"scope"`
	Certificate          string `json:"certificate"`
}
