package main

import (
	"strings"
	"testing"
)

// func Test_command_printDescription(t *testing.T) {
// 	type fields struct {
// 		Cmd         string
// 		Args        []string
// 		Description string
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 	}{
// 	// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			c := command{
// 				Cmd:         tt.fields.Cmd,
// 				Args:        tt.fields.Args,
// 				Description: tt.fields.Description,
// 			}
// 			c.printDescription()
// 		})
// 	}
// }

// func Test_command_printFullCommand(t *testing.T) {
// 	type fields struct {
// 		Cmd         string
// 		Args        []string
// 		Description string
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 	}{
// 	// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			c := command{
// 				Cmd:         tt.fields.Cmd,
// 				Args:        tt.fields.Args,
// 				Description: tt.fields.Description,
// 			}
// 			c.printFullCommand()
// 		})
// 	}
// }

func Test_command_exec(t *testing.T) {
	type fields struct {
		Cmd         string
		Args        string
		Description string
	}
	type args struct {
		debug   bool
		verbose bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
		want1  string
	}{
		{
			name: "echo",
			fields: fields{
				Cmd:         "echo",
				Args:        "this is fun",
				Description: "A command execution test with echo.",
			},
			args:  args{debug: false, verbose: false},
			want:  0,
			want1: "this is fun",
		}, {
			name: "exitCode",
			fields: fields{
				Cmd:         "sh",
				Args:        "-c 'echo $?'",
				Description: "A sh command execution test with exitCode.",
			},
			args:  args{debug: false},
			want:  0,
			want1: "0",
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := command{
				Cmd:         tt.fields.Cmd,
				Args:        tt.fields.Args,
				Description: tt.fields.Description,
			}
			got, got1 := c.exec(tt.args.debug, tt.args.verbose)
			if got != tt.want {
				t.Errorf("command.exec() got = %v, want %v", got, tt.want)
			}
			if strings.TrimSpace(got1) != tt.want1 {
				t.Errorf("command.exec() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
