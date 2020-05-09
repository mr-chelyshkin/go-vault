package vault

import (
	"crypto/tls"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const certTestActions = `
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

func TestGenerateCertPoolNegative1(t *testing.T) {
	cert, err := generateCertPool("path_not_exit/file.txt")

	assert.Nil(t, cert)
	assert.Error(t, err, err)
}

func TestGenerateCertPoolPositive1(t *testing.T) {
	file, _ := ioutil.TempFile("", "")
	defer os.Remove(file.Name())

	b := []byte(certTestActions)
	_ = ioutil.WriteFile(file.Name(), b, 0644)

	_, err := generateCertPool(file.Name())
	assert.Nil(t, err)
}

func TestNewActionPositive1(t *testing.T) {
	file, _ := ioutil.TempFile("", "")
	defer os.Remove(file.Name())

	b := []byte(certTestActions)
	_ = ioutil.WriteFile(file.Name(), b, 0644)

	cert, _ := generateCertPool(file.Name())
	httpClient := &http.Client {
		Transport: &http.Transport {
			TLSClientConfig: &tls.Config{RootCAs: cert},
		},
	}

	expect := &httpActions{httpClient: httpClient}
	actual, err := newActions(file.Name())

	assert.Nil(t, err)
	assert.Equal(t, expect, actual)
}

func TestNewActionNegative1(t *testing.T) {
	actual, err := newActions("path_not_exist/file.txt")

	assert.Nil(t, actual)
	assert.Error(t, err, "")
}

func TestActionGetPositive1(t *testing.T) {
	file, _ := ioutil.TempFile("", "")
	defer os.Remove(file.Name())

	_, _ = generateCertPool(file.Name())

	testHandler := func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, "GET")
		assert.Equal(t, req.Header.Get("X-Vault-Token"), "test_token")
	}

	testServer := httptest.NewServer(http.HandlerFunc(testHandler))
	defer testServer.Close()

	b := []byte(certTestActions)
	_ = ioutil.WriteFile(file.Name(), b, 0644)

	action, _ := newActions(file.Name())
	_, err := action.get(testServer.URL, "test_token")
	assert.Nil(t, err)

}

func TestActionGetNegative1(t *testing.T) {
	file, _ := ioutil.TempFile("", "")
	defer os.Remove(file.Name())

	_, _ = generateCertPool(file.Name())

	testHandler := func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, "GET")
		assert.Equal(t, req.Header.Get("X-Vault-Token"), "test_token")
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	testServer := httptest.NewServer(http.HandlerFunc(testHandler))
	defer testServer.Close()

	b := []byte(certTestActions)
	_ = ioutil.WriteFile(file.Name(), b, 0644)

	action, _ := newActions(file.Name())
	resp, err := action.get(testServer.URL, "test_token")

	assert.Nil(t, resp)
	assert.Error(t, err, "")
}
