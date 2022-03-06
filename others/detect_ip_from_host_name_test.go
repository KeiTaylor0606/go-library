package others

import (
	"testing"
)

func TestGetHostnameFromIPAddress(t *testing.T) {
	testcases := []struct {
		name string
		ip   string
		want string
	}{
		{
			name: "test_00",
			ip:   "127.0.0.1",
			want: "localhost",
		},
		{
			name: "test_01",
			ip:   "8.8.4.4",
			want: "dns.google.",
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			op, err := GetHostnameFromIPAddress(testcase.ip)
			if err != nil {
				t.Errorf("err: %v\n", err)
			}
			judge := false
			for _, res := range op {
				if res == testcase.want {
					judge = true
				}
			}
			if !judge {
				t.Errorf("want:%v\n", testcase.want)
			}
		})
	}
}
