package vault

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCaseGettersApi struct {
	name   string
	expect string
	input  string
}

func TestGetHost(t *testing.T) {
	testCases := []testCaseGettersApi{
		{
			name:   "getBaseValue",
			input:  "",
			expect: host,
		},
		{
			name:   "getCustomValue1",
			input: 	"https://mail.ru",
			expect: "https://mail.ru",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expect, getHost(tc.input))
		})
	}
}

func TestGetPort(t *testing.T) {
	testCases := []testCaseGettersApi{
		{
			name:   "getBaseValue",
			input:  "",
			expect: port,
		},
		{
			name:   "getCustomValue1",
			input: 	"6666",
			expect: "6666",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expect, getPort(tc.input))
		})
	}
}

func TestGetVersion(t *testing.T) {
	testCases := []testCaseGettersApi{
		{
			name:   "getBaseValue",
			input:  "",
			expect: version,
		},
		{
			name:   "getCustomValue1",
			input: 	"v2",
			expect: "v2",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expect, getVersion(tc.input))
		})
	}
}

func TestGetAuthLink(t *testing.T) {
	testCases := []testCaseGettersApi{
		{
			name:   "getBaseValue",
			input:  "",
			expect: authLink,
		},
		{
			name:   "getCustomValue1",
			input: 	"new-auth/link",
			expect: "new-auth/link",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expect, getAuthLink(tc.input))
		})
	}
}

func TestGetUpdateLink(t *testing.T) {
	testCases := []testCaseGettersApi{
		{
			name:   "getBaseValue",
			input:  "",
			expect: updateLink,
		},
		{
			name:   "getCustomValue1",
			input: 	"new-update/link",
			expect: "new-update/link",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expect, getUpdateLink(tc.input))
		})
	}
}

func TestGetLookupLink(t *testing.T) {
	testCases := []testCaseGettersApi{
		{
			name:   "getBaseValue",
			input:  "",
			expect: lookupLink,
		},
		{
			name:   "getCustomValue1",
			input: 	"new-lookup/link",
			expect: "new-lookup/link",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expect, getLookupLink(tc.input))
		})
	}
}

func TestGetBaseClientApi(t *testing.T) {
	actual := &ClientApi{
		Host:       host,
		Port:       port,
		Version:    version,
		AuthLink:   authLink,
		UpdateLink: updateLink,
		LookupLink: lookupLink,
	}

	assert.Equal(t, actual, getBaseClientApi())
}

func TestLookupUrl(t *testing.T) {
	api := &ClientApi{
		Host:       host,
		Port:       port,
		Version:    "v2",
		AuthLink:   authLink,
		UpdateLink: updateLink,
		LookupLink: "lookup",
	}

	actual := fmt.Sprintf("%s:%s/%s/%s", api.Host, api.Port, api.Version, api.LookupLink)
	assert.Equal(t, actual, api.lookupUrl())
}

func TestUpdateUrl(t *testing.T) {
	api := &ClientApi{
		Host:       host,
		Port:       port,
		Version:    "v2",
		AuthLink:   authLink,
		UpdateLink: "update",
		LookupLink: lookupLink,
	}

	actual := fmt.Sprintf("%s:%s/%s/%s", api.Host, api.Port, api.Version, api.UpdateLink)
	assert.Equal(t, actual, api.updateUrl())
}

func TestAuthUrl(t *testing.T) {
	api := &ClientApi{
		Host:       host,
		Port:       port,
		Version:    "v2",
		AuthLink:   "auth",
		UpdateLink: updateLink,
		LookupLink: lookupLink,
	}

	actual := fmt.Sprintf("%s:%s/%s/%s", api.Host, api.Port, api.Version, api.AuthLink)
	assert.Equal(t, actual, api.authUrl())
}

func TestBaseUrl(t *testing.T) {
	api := &ClientApi{
		Host:       host,
		Port:       port,
		Version:    version,
		AuthLink:   authLink,
		UpdateLink: updateLink,
		LookupLink: lookupLink,
	}

	actual := fmt.Sprintf("%s:%s/%s", api.Host, api.Port, api.Version)
	assert.Equal(t, actual, api.baseUrl())
}
