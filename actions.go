package vault

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"net/http"
)

type httpActions struct {
	httpClient *http.Client
}

func (h httpActions) get(url, token string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("X-Vault-Token", token)
	response, err := h.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	return body, nil
}

func (h httpActions) post(url, token string, credentials []byte) ([]byte, error) {
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(credentials))
	if err != nil {
		return nil, err
	}

	if token != "" {
		request.Header.Add("X-Vault-Token", token)
	}

	response, err := h.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	return body, nil
}

func newActions(certPath string) (*httpActions, error) {
	certPool, err := generateCertPool(certPath)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client {
		Transport: &http.Transport {
			TLSClientConfig: &tls.Config{RootCAs: certPool},
		},
	}

	return &httpActions{httpClient:httpClient}, nil
}

func generateCertPool(filepath string) (*x509.CertPool, error) {
	caCert, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	return caCertPool, nil
}

