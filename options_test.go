package vault

import (
	"github.com/stretchr/testify/assert"
	"os/user"
	"path/filepath"
	"testing"
)

type testCaseGettersOptions struct {
	name   string
	expect string
	input  string
}

func TestGetTokenFilePath(t *testing.T) {
	usr, _ := user.Current()

	testCases := []testCaseGettersOptions{
		{
			name:   "getBaseValue",
			input:  "",
			expect: filepath.Join(usr.HomeDir, baseTokenFile[2:]),
		},
		{
			name:   "getCustomValue1",
			input: 	"/tmp/.token",
			expect: "/tmp/.token",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expect, getTokenFilePath(tc.input))
		})
	}
}

func TestGetCertFilePath(t *testing.T) {
	testCases := []testCaseGettersOptions{
		{
			name:   "getBaseValue",
			input:  "",
			expect: getCertFilePath(baseCertFile),
		},
		{
			name:   "getCustomValue",
			input: 	"/tmp/cert/crt.pem",
			expect: "/tmp/cert/crt.pem",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expect, getCertFilePath(tc.input))
		})
	}
}

func TestGetBaseClientOptions(t *testing.T) {
	usr, _ := user.Current()

	expect := &ClientOptions{
		CertFilePath:  baseCertFile,
		TokenFilePath: filepath.Join(usr.HomeDir, baseTokenFile[2:]),
	}

	assert.Equal(t, expect, getBaseClientOptions())
}
