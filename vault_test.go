package vault

import (
	"crypto/tls"
	"github.com/stretchr/testify/assert"

	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

const certTestVault = `
-----BEGIN CERTIFICATE-----
MIIFMzCCAxugAwIBAgIUFba3dcyDp9qymQCgryrpvUfUaFswDQYJKoZIhvcNAQEL
BQAwITEfMB0GA1UEAxMWU0VBUkNIIE1BSUwuUlUgUm9vdCBDQTAeFw0xOTA3MDMx
MzM5MjVaFw0yOTA2MzAxMzM5NTNaMCExHzAdBgNVBAMTFlNFQVJDSCBNQUlMLlJV
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIFhjCCA26gAwIBAgIUYswAasjzKm+hGa6uYIjAeAPnOBowDQYJKoZIhvcNAQEL
BQAwITEfMB0GA1UEAxMWU0VBUkNIIE1BSUwuUlUgUm9vdCBDQTAeFw0xOTA3MTkx
NjUxMjFaFw0yNDA3MTcxNjUxNTFaMC4xLDAqBgNVBAMTI1NFQVJDSCBNQUlMLlJV
-----END CERTIFICATE-----
`

func TestNewBasicClient_Positive1(t *testing.T) {
	file, _ := ioutil.TempFile("", "")
	defer os.Remove(file.Name())

	b := []byte(certTestVault)
	_ = ioutil.WriteFile(file.Name(), b, 0644)

	testX509, _ := generateCertPool(file.Name())
	httpClient := &http.Client {
		Transport: &http.Transport {
			TLSClientConfig: &tls.Config{RootCAs: testX509},
		},
	}

	cliOpt := &ClientOptions{
		CertFilePath:  file.Name(),
		TokenFilePath: getTokenFilePath(""),
	}
	apiOpt := &ClientApi{
		Host:       host,
		Port:       port,
		Version:    version,
		AuthLink:   authLink,
		UpdateLink: updateLink,
		LookupLink: lookupLink,
	}
	creds := &credentials{
		RoleId:   "roleId",
		SecretId: "secretId",
	}
	actions := &httpActions{
		httpClient: httpClient,
	}

	expect := &Client{
		credentials: *creds,
		options:     cliOpt,
		api:         apiOpt,
		actions:     actions,
	}

	actual, _ := NewBasicClient("roleId", "secretId", cliOpt)
	assert.Equal(t, expect.credentials, actual.credentials)
	assert.Equal(t, expect.actions, actual.actions)
	assert.Equal(t, expect.options, actual.options)
	assert.Equal(t, expect.api, actual.api)
}

func TestNewCustomClient_Positive1(t *testing.T) {
	file, _ := ioutil.TempFile("", "")
	defer os.Remove(file.Name())

	b := []byte(certTestVault)
	_ = ioutil.WriteFile(file.Name(), b, 0644)

	testX509, _ := generateCertPool(file.Name())
	httpClient := &http.Client {
		Transport: &http.Transport {
			TLSClientConfig: &tls.Config{RootCAs: testX509},
		},
	}

	cliOpt := &ClientOptions{
		CertFilePath:  file.Name(),
		TokenFilePath: getTokenFilePath(""),
	}
	apiOpt := &ClientApi{
		Host:       host,
		Port:       port,
		Version:    version,
		AuthLink:   authLink,
		UpdateLink: updateLink,
		LookupLink: lookupLink,
	}
	creds := &credentials{
		RoleId:   "roleId",
		SecretId: "secretId",
	}
	actions := &httpActions{
		httpClient: httpClient,
	}

	expect := &Client{
		credentials: *creds,
		options:     cliOpt,
		api:         apiOpt,
		actions:     actions,
	}

	actual, _ := NewCustomClient("roleId", "secretId", cliOpt, nil)
	assert.Equal(t, expect.credentials, actual.credentials)
	assert.Equal(t, expect.actions, actual.actions)
	assert.Equal(t, expect.options, actual.options)
	assert.Equal(t, expect.api, actual.api)
}

func TestNewCustomClient_Positive2(t *testing.T) {
	file, _ := ioutil.TempFile("", "")
	defer os.Remove(file.Name())

	b := []byte(certTestVault)
	_ = ioutil.WriteFile(file.Name(), b, 0644)

	testX509, _ := generateCertPool(file.Name())
	httpClient := &http.Client {
		Transport: &http.Transport {
			TLSClientConfig: &tls.Config{RootCAs: testX509},
		},
	}

	cliOpt := &ClientOptions{
		CertFilePath:  file.Name(),
		TokenFilePath: getTokenFilePath(""),
	}
	apiOpt := &ClientApi{
		Host:       "https://mail.ru",
		Port:       port,
		Version:    "v2",
		AuthLink:   authLink,
		UpdateLink: updateLink,
		LookupLink: lookupLink,
	}
	creds := &credentials{
		RoleId:   "roleId",
		SecretId: "secretId",
	}
	actions := &httpActions{
		httpClient: httpClient,
	}

	expect := &Client{
		credentials: *creds,
		options:     cliOpt,
		api:         apiOpt,
		actions:     actions,
	}

	api := &ClientApi{Host: "https://mail.ru", Version: "v2"}
	actual, _ := NewCustomClient("roleId", "secretId", cliOpt, api)
	assert.Equal(t, expect.credentials, actual.credentials)
	assert.Equal(t, expect.actions, actual.actions)
	assert.Equal(t, expect.options, actual.options)
	assert.Equal(t, expect.api, actual.api)
}

func TestNewBasicClient_Negative1(t *testing.T) {
	cliOpt := &ClientOptions{
		CertFilePath:  "not_exist_cert.pem",
		TokenFilePath: getTokenFilePath(""),
	}

	actual, err := NewBasicClient("roleId", "secretId", cliOpt)
	assert.Nil(t, actual)
	assert.Error(t, err, "")
}

func TestNewCustomClient_Negative1(t *testing.T) {
	cliOpt := &ClientOptions{
		CertFilePath:  "not_exist_cert.pem",
		TokenFilePath: getTokenFilePath(""),
	}

	actual, err := NewCustomClient("roleId", "secretId", cliOpt, nil)
	assert.Nil(t, actual)
	assert.Error(t, err, "")
}

func TestCreateTokenFilePositive1(t *testing.T) {
	file, _ := ioutil.TempFile("", "")
	defer os.Remove(file.Name())

	responseData := []byte(`{"auth":{"client_token":"response_token","other":"other"},"data":"data"}`)
	token, err := createTokenFile(responseData, file.Name())
	data, _ := ioutil.ReadFile(file.Name())

	assert.Nil(t, err)
	assert.Equal(t, "response_token", *token)
	assert.Equal(t, "response_token", string(data))
}
