package ssl

import (
	"testing"
)

func TestSSLCertGen(t *testing.T) {
	testcases := []struct {
		name             string
		organization     string
		organizationUnit string
		commonName       string
		cert             string
		key              string
	}{
		{
			name:             "test_00",
			organization:     "test_00_organization",
			organizationUnit: "test_00_organizationUnit",
			commonName:       "test_00_commonName",
			cert:             "test_00_cert",
			key:              "test_00_key",
		},
		{
			name:             "test_01",
			organization:     "test_01_organization",
			organizationUnit: "test_01_organizationUnit",
			commonName:       "test_01_commonName",
			cert:             "test_01_cert",
			key:              "test_01_key",
		},
		{
			name:             "test_02",
			organization:     "test_02_organization",
			organizationUnit: "test_02_organizationUnit",
			commonName:       "test_02_commonName",
			cert:             "test_02_cert",
			key:              "test_02_key",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			if err := SSLCertGen(testcase.organization, testcase.organizationUnit, testcase.commonName, testcase.cert, testcase.key); err != nil {
				t.Errorf("err: %v\n", err)
			}
		})
	}
}
