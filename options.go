package vault

import (
	"os/user"
	"path/filepath"
	"strings"
)

const (
	baseTokenFile    = "~/.vault_token"
	baseCertFile     = "/etc/ssl/search/ca.pem"
)

type ClientOptions struct {
	TokenFilePath string
	CertFilePath  string
}

func getBaseClientOptions() *ClientOptions {
	return &ClientOptions {
		TokenFilePath: getTokenFilePath(""),
		CertFilePath:  getCertFilePath(""),
	}
}

func getTokenFilePath(data string) string {
	var tokenPath string

	if data != "" {
		tokenPath = data
	} else {
		tokenPath = baseTokenFile
	}

	if strings.HasPrefix(tokenPath,"~/") {
		usr, _ := user.Current()
		tokenPath = filepath.Join(usr.HomeDir, tokenPath[2:])
	}

	return tokenPath
}

func getCertFilePath(data string) string {
	var certPath string

	if data != "" {
		certPath = data
	} else {
		certPath = baseCertFile
	}

	if strings.HasPrefix(certPath, "~/") {
		usr, _ := user.Current()
		certPath = filepath.Join(usr.HomeDir, certPath[2:])
	}

	return certPath
}

