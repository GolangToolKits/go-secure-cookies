package gosecurecookies

import (
	"encoding/hex"
	"testing"
)

func TestNew(t *testing.T) {
	var key = "dsdfs6dfs61dssdfsdfdsdsfsdsdllsd"
	skstr := hex.EncodeToString([]byte(key))
	sks, _ := hex.DecodeString(skstr)


	var keyLong = "dsdfs6dfs61dssdfsdfdsdsfsdsdllsddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd"
	skstLong := hex.EncodeToString([]byte(keyLong))
	sksLong, _ := hex.DecodeString(skstLong)

	type args struct {
		secretKey string
	}
	tests := []struct {
		name string
		args args
		want Encrypt
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				secretKey: string(sks),
			},
		},
		{
			name: "test 2",
			args: args{
				secretKey: string(sksLong),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := NewEncrypt(tt.args.secretKey); got == nil || err != nil {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGCMEncrypt_Encrypt(t *testing.T) {
	var theTxt = "this-is-a-test-and-more"
	var key = "dsdfs6dfs61dssdfsdfdsdsfsdsdllsd"
	skstr := hex.EncodeToString([]byte(key))
	sks, _ := hex.DecodeString(skstr)

	var keyBad = "abc"
	skstrBad := hex.EncodeToString([]byte(keyBad))
	sksBad, _ := hex.DecodeString(skstrBad)

	var keyWrong = "dsdfs6dfs61dssdfsdfdsdsfsdsdllsa"
	skstrWrong := hex.EncodeToString([]byte(keyWrong))
	sksWrong, _ := hex.DecodeString(skstrWrong)

	type fields struct {
		secretKey []byte
	}
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				secretKey: sks,
			},
			args: args{
				s: theTxt,
			},
			//want: "",
			wantErr: false,
		},
		{
			name: "test 2",
			fields: fields{
				secretKey: sksBad,
			},
			args: args{
				s: theTxt,
			},
			//want: "",
			wantErr: true,
		},
		{
			name: "test 3",
			fields: fields{
				secretKey: sks,
			},
			args: args{
				s: theTxt,
			},
			//want: "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &GCMEncrypt{
				secretKey: tt.fields.secretKey,
			}

			eWrong := &GCMEncrypt{
				secretKey: sksWrong,
			}
			got, err := e.Encrypt(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("GCMEncrypt.Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if got == tt.want {
			// 	t.Errorf("GCMEncrypt.Encrypt() = %v, want %v", got, tt.want)
			// }
			if tt.name == "test 1" && got == "" {
				t.Errorf("GCMEncrypt.Encrypt() = %v, want %v", got, tt.want)
			}
			if tt.name == "test 1" {
				ct, derr := e.Decrypt(got)
				if tt.name == "test 1" && (derr != nil || ct != theTxt) {
					t.Fail()
				}
			}

			if tt.name == "test 2" {
				ct, derr := e.Decrypt(got)
				if tt.name == "test 1" && (derr != nil || ct != theTxt) {
					t.Fail()
				}
			}
			if tt.name == "test 3" {
				ct, derr := eWrong.Decrypt(got)
				if tt.name == "test 3" && (derr == nil || ct == theTxt) {
					t.Fail()
				}
			}

			//fmt.Println("ct: ", ct)
		})
	}
}
