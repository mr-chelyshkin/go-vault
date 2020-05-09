package vault

import (
	"fmt"
	"net/url"
	"path"
)

const (
	host = "https://srch-vault.g"
	port = "8080"

	version    = "v1"
	authLink   = "auth/approle/login"
	updateLink = "auth/token/renew-self"
	lookupLink = "auth/token/lookup-self"
)

type ClientApi struct {
	Host string
	Port string

	Version    string
	AuthLink   string
	UpdateLink string
	LookupLink string
}

func (c ClientApi) baseUrl() string {
	u, _ := url.Parse(fmt.Sprintf("%s:%s", c.Host, c.Port))
	u.Path = path.Join(u.Path, c.Version)

	return u.String()
}

func (c ClientApi) authUrl() string {
	u, _ := url.Parse(c.baseUrl())
	u.Path = path.Join(u.Path, c.AuthLink)

	return u.String()
}

func (c ClientApi) updateUrl() string {
	u, _ := url.Parse(c.baseUrl())
	u.Path = path.Join(u.Path, c.UpdateLink)

	return u.String()
}

func (c ClientApi) lookupUrl() string {
	u, _ := url.Parse(c.baseUrl())
	u.Path = path.Join(u.Path, c.LookupLink)

	return u.String()
}

func getBaseClientApi() *ClientApi {
	return &ClientApi{
		Host:       host,
		Port:       port,
		Version:    version,
		AuthLink:   authLink,
		UpdateLink: updateLink,
		LookupLink: lookupLink,
	}
}

func getHost(data string) string {
	if data != "" {
		return data
	}
	return host
}

func getPort(data string) string {
	if data != "" {
		return data
	}
	return port
}

func getVersion(data string) string {
	if data != "" {
		return data
	}
	return version
}

func getAuthLink(data string) string {
	if data != "" {
		return data
	}
	return authLink
}

func getUpdateLink(data string) string {
	if data != "" {
		return data
	}
	return updateLink
}

func getLookupLink(data string) string {
	if data != "" {
		return data
	}
	return lookupLink
}

