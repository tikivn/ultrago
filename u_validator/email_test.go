package u_validator

import (
	"testing"
)

func TestVerifyEmail(t *testing.T) {
	type args struct {
		e string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "user_001@gmail.com",
			args: args{
				e: "user_001@gmail.com",
			},
			want: true,
		},
		{
			name: "User.+-_002@yahoo.com",
			args: args{
				e: "User.+-_002@yahoo.com",
			},
			want: true,
		},
		{
			name: "user@gmail",
			args: args{
				e: "user@gmail",
			},
			want: false,
		},
		{
			name: "user",
			args: args{
				e: "user",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VerifyEmail(tt.args.e); got != tt.want {
				t.Errorf("VerifyEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
