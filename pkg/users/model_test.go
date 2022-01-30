package users

import (
	"reflect"
	"strings"
	"testing"
)

var testStr = "this is a test"
var testStrAsByteArr = []byte(testStr)
var checkedResult = "dGhpcyBpcyBhIHRlc3Q="

var strToEncrypt = "please encrypt here!!"
var checkedEncrypted = "GyxC4htJA/0PYd3qYz3ECSA7JaTM"

func TestEncode(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test encode",
			args: args{b: testStrAsByteArr},
			want: checkedResult,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encode(tt.args.b); got != tt.want {
				t.Errorf("encode() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestDecode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "test decode",
			args: args{s: checkedResult},
			want: []byte(testStr),
		},
		{
			name: "test decode (wrong)",
			args: args{s: checkedResult},
			want: []byte("wrong str"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !strings.Contains(tt.name, "wrong") {
				if got := decode(tt.args.s); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("decode() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestEncryptData(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "test encrypt",
			args:    args{s: strToEncrypt},
			want:    checkedEncrypted,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := encryptData(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("encryptData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("encryptData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decryptData(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "decrypt test",
			args:    args{s: checkedEncrypted},
			want:    strToEncrypt,
			wantErr: false,
		},
		{
			name:    "(wrong) decrypt test",
			args:    args{s: checkedEncrypted},
			want:    "wrong here",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decryptData(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("decryptData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if strings.Contains(tt.name, "wrong") {
				if got == tt.want {
					t.Error("got and want should be different")
				}
				return
			}

			if got != tt.want {
				t.Errorf("decryptData() got = %v, want %v", got, tt.want)
			}
		})
	}
}