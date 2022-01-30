package users

import (
	"github.com/rs/zerolog/log"
	"testing"
)

var passToHash = "superHardPass___!"

func Test_hashPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test hash password",
			args: args{password: passToHash},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, _ := hashPassword(tt.args.password)
			got2, _ := hashPassword(tt.args.password)

			if got1 == got2 {
				t.Errorf("hashes with salt are the same: \n %s \n %s", got1, got2)
			}
			log.Info().Msg(got1)
		})
	}
}

var checkedHash = "$2a$04$UMUk/lyzD4vPlI4Lfrppf.OzGLHek4q8SPqTIpvsD9hMOtaR/c83m"
var checkedWrongHash = "yc8tttUiOW12TMTpYr6dyXf1zIATJHCBLNzECCO2FQArvRYFDCV8aHxNojEmV_rXF224n_sa7EdPva7sxKUbHQ=="

func Test_doPasswordsMatch(t *testing.T) {
	type args struct {
		hashedPassword string
		currPassword   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "do pass match test",
			args: args{
				hashedPassword: checkedHash,
				currPassword:   passToHash,
			},
			want: true,
		},
		{
			name: "(wrong) do pass match test",
			args: args{
				hashedPassword: checkedWrongHash,
				currPassword:   passToHash,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doPasswordsMatch(tt.args.hashedPassword, tt.args.currPassword); got != tt.want {
				t.Errorf("doPasswordsMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
