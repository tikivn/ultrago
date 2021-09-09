package u_hash

import "testing"

func TestHashSHA1(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "HashSHA1-UUID",
			args: args{
				key: "a202f638-f5a8-11eb-9a03-0242ac130003",
			},
			want: "d4f58bdb80aea969cb4925a023f5ed3512ba5e80",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HashSHA1(tt.args.key); got != tt.want {
				t.Errorf("HashSHA1() = %v, want %v", got, tt.want)
			}
		})
	}
}
