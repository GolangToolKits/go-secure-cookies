package gosecurecookies

import (
	b64 "encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSecureCookies_Write(t *testing.T) {
	var key = "dsdfs6dfs61dssdfsdfdsdsfsdsdllsd"
	enct, _ := NewEncrypt(key)
	tw := httptest.NewRecorder()
	var ck http.Cookie
	ck.Name = "test1"
	ck.Value = "this-is-a-test-and-more"

	var keyBad = "dsd"
	//enctBad, _ := NewEncrypt(keyBad)
	var enctBad GCMEncrypt
	enctBad.secretKey = []byte(keyBad)

	//var keyWrong = "dsdfs6dfs61dssdfsdfdsdsfsdsdllsa"
	//enctWrong, _ := NewEncrypt(keyWrong)

	type fields struct {
		encrypt Encrypt
	}
	type args struct {
		w      http.ResponseWriter
		cookie http.Cookie
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				encrypt: enct,
			},
			args: args{
				w:      tw,
				cookie: ck,
			},
			wantErr: false,
		},
		{
			name: "test 2",
			fields: fields{
				encrypt: &enctBad,
			},
			args: args{
				w:      tw,
				cookie: ck,
			},
			wantErr: true,
		},
		// {
		// 	name: "test 3",
		// 	fields: fields{
		// 		encrypt: enctWrong,
		// 	},
		// 	args: args{
		// 		w:      tw,
		// 		cookie: ck,
		// 	},
		// 	wantErr: false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SecureCookies{
				encrypt: tt.fields.encrypt,
			}
			if tt.name != "test 2" {
				if err := s.Write(tt.args.w, tt.args.cookie); (err != nil) != tt.wantErr {
					t.Errorf("SecureCookies.Write() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
			if tt.name == "test 2" {
				if err := s.Write(tt.args.w, tt.args.cookie); (err != nil) != tt.wantErr {
					t.Errorf("SecureCookies.Write() error = %v, wantErr %v", err, tt.wantErr)
				}
			}

		})
	}
}

func TestSecureCookies_Read(t *testing.T) {
	var key = "dsdfs6dfs61dssdfsdfdsdsfsdsdllsd"
	enct, _ := NewEncrypt(key)
	//tw := httptest.NewRecorder()
	tr, _ := http.NewRequest("POST", "/test/test1", nil)
	var txt = "test1:this-is-a-test-and-more"
	etxt, _ := enct.Encrypt(txt)
	//fmt.Println("err", err)
	cookie := &http.Cookie{
		Name: "test1",
		//Value: etxt,
		Value: b64.StdEncoding.EncodeToString([]byte(etxt)),

		MaxAge: 300,
	}
	//var rck = cookie.Name + "=" + cookie.Value
	//tr.Header.Set("Cookie", rck)
	tr.AddCookie(cookie)

	type fields struct {
		encrypt Encrypt
	}
	type args struct {
		r    *http.Request
		name string
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
				encrypt: enct,
			},
			args: args{
				r:    tr,
				name: "test1",
			},
			want:    "this-is-a-test-and-more",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SecureCookies{
				encrypt: tt.fields.encrypt,
			}
			got, err := s.Read(tt.args.r, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("SecureCookies.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SecureCookies.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}
