package gosecurecookies

import (
	"testing"
)

func TestNewCookies(t *testing.T) {
	type args struct {
		secureKey string
	}
	tests := []struct {
		name string
		args args
		want Cookies
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				secureKey: "dsdfs6dfs61dssdfsdfdsdsfsdsdllsd",
			},
		},
		{
			name: "test 2",
			args: args{
				secureKey: "dsd",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name != "test 2"{
				if got, err := NewCookies(tt.args.secureKey); got == nil || err != nil {
					t.Errorf("NewCookies() = %v, want %v", got, tt.want)
				}
			}else {
				if got, err := NewCookies(tt.args.secureKey); got != nil || err == nil {
				t.Errorf("NewCookies() = %v, want %v", got, tt.want)
			}
			}
			// if got, err := NewCookies(tt.args.secureKey); got == nil || err != nil {
			// 	t.Errorf("NewCookies() = %v, want %v", got, tt.want)
			// }
		})
	}
}
