package entity

import (
	"reflect"
	"testing"
)

func TestNewUser(t *testing.T) {
	type args struct {
		username string
		passHash string
	}
	tests := []struct {
		name string
		args args
		want *User
	}{
		{
			name: "Test New User",
			args: args{
				username: "test",
				passHash: "test",
			},
			want: &User{
				Username: "test",
				PassHash: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUser(tt.args.username, tt.args.passHash); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = %v, want %v (%v)", got, tt.want, tt.name)
			}
		})
	}
}
