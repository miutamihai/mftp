package ansi

import "testing"

func TestColorMessage(t *testing.T) {
	type args struct {
		message string
		color   Color
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "BrightCyan",
			args: args{
				message: "message",
				color:   BrightCyan,
			},
			want: "\033[96mmessage\033[0m",
		},
		{
			name: "BrightRed",
			args: args{
				message: "message",
				color:   BrightRed,
			},
			want: "\033[91mmessage\033[0m",
		},
		{
			name: "BrightWhite",
			args: args{
				message: "message",
				color:   BrightWhite,
			},
			want: "\033[97mmessage\033[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ColorMessage(tt.args.message, tt.args.color); got != tt.want {
				t.Errorf("ColorMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
