package vault

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"os"
	"path"
)

type Client struct {
	credentials

	options *ClientOptions
	actions *httpActions
	api     *ClientApi
}

type credentials struct {
	RoleId   string `json:"role_id"`
	SecretId string `json:"secret_id"`
}

func NewBasicClient(roleId, secretId string, options *ClientOptions) (*Client, error) {
	var cliOpt *ClientOptions
	if options == nil {
		cliOpt = getBaseClientOptions()
	} else {
		cliOpt = &ClientOptions{
			TokenFilePath: getTokenFilePath(options.TokenFilePath),
			CertFilePath:  getCertFilePath(options.CertFilePath),
		}
	}

	actions, err := newActions(cliOpt.CertFilePath)
	if err != nil {
		return nil, err
	}

	return &Client{
		credentials: credentials{RoleId: roleId, SecretId: secretId},
		options:     cliOpt,
		actions:     actions,
		api:         getBaseClientApi(),
	}, nil
}

func NewCustomClient(roleId, secretId string, options *ClientOptions, api *ClientApi) (*Client, error) {
	var cliOpt *ClientOptions
	if options == nil {
		cliOpt = getBaseClientOptions()
	} else {
		cliOpt = &ClientOptions{
			TokenFilePath: getTokenFilePath(options.TokenFilePath),
			CertFilePath:  getCertFilePath(options.CertFilePath),
		}
	}

	var cliApi *ClientApi
	if api == nil {
		cliApi = getBaseClientApi()
	} else {
		cliApi = &ClientApi{
			Host:       getHost(api.Host),
			Port:       getPort(api.Port),
			Version:    getVersion(api.Version),
			AuthLink:   getAuthLink(api.AuthLink),
			UpdateLink: getUpdateLink(api.UpdateLink),
			LookupLink: getLookupLink(api.LookupLink),
		}
	}

	actions, err := newActions(cliOpt.CertFilePath)
	if err != nil {
		return nil, err
	}

	return &Client{
		credentials: credentials{RoleId: roleId, SecretId: secretId},
		options:     cliOpt,
		actions:     actions,
		api:         cliApi,
	}, nil
}

func (c Client) Get(dataUrl string) (interface{}, error) {
	u, _ := url.Parse(c.api.baseUrl())
	u.Path = path.Join(u.Path, dataUrl)

	token, err := c.token()
	if err != nil {
		return nil, err
	}

	response, err := c.actions.get(u.String(), *token)
	if err != nil {
		return nil, err
	}

	type vaultData struct {
		Data interface{} `json:"data"`
	}
	var data vaultData

	err = json.Unmarshal(response, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c Client) token() (*string, error) {
	requestData, _ := json.Marshal(c.credentials)

	if _, err := os.Stat(c.options.TokenFilePath); os.IsNotExist(err) {
		return c.__auth__(requestData)
	}

	bufToken, _ := ioutil.ReadFile(c.options.TokenFilePath)
	ttl, renewable, err := c.__lookup__(string(bufToken))
	if err != nil {
		return c.__auth__(requestData)
	}

	if renewable && ttl < 1500 {
		return c.__update__(requestData, string(bufToken))
	}

	token := string(bufToken)
	return &token, nil
}

func (c Client) __auth__(requestData []byte) (*string, error) {
	response, err := c.actions.post(c.api.authUrl(), "", requestData)
	if err != nil {
		return nil, err
	}

	return createTokenFile(response, c.options.TokenFilePath)
}

func (c Client) __update__(requestData []byte, token string) (*string, error) {
	response, err := c.actions.post(c.api.updateUrl(), token, requestData)

	if err != nil {
		return nil, err
	}
	return createTokenFile(response, c.options.TokenFilePath)
}

func (c Client) __lookup__(token string) (int, bool, error) {
	type jsonResponseData struct {
		Ttl       int  `json:"ttl"`
		Renewable bool `json:"renewable"`
	}
	type jsonResponse struct {Data jsonResponseData `json:"data"`}

	response, err := c.actions.get(c.api.lookupUrl(), token)
	if err != nil {
		return 0, false, err
	}

	var lookupResponse jsonResponse
	err = json.Unmarshal(response, &lookupResponse)
	if err != nil {
		return 0, false, err
	}

	return lookupResponse.Data.Ttl, lookupResponse.Data.Renewable, nil
}

func createTokenFile(response []byte, tokenPath string) (*string, error) {
	type authJson struct {Token string   `json:"client_token"`}
	type respJson struct {Auth  authJson `json:"auth"`}

	var respJsonData respJson
	err := json.Unmarshal(response, &respJsonData)
	if err != nil {
		return nil, err
	}

	b := []byte(respJsonData.Auth.Token)
	err = ioutil.WriteFile(tokenPath, b, 0644)
	if err != nil {
		return nil, err
	}

	return &respJsonData.Auth.Token, nil
}

