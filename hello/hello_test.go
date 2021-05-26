package hello

import "testing"

func TestHello(t *testing.T) {
	type args struct {
		firstname string
		lastname  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "test-1", want: "Hello Jignesh Dhua", args: args{firstname: "Jignesh", lastname: "Dhua"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.args.firstname, tt.args.lastname); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
