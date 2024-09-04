package driver

import (
	"fmt"
	"mihaimiuta/mftp/pkg/mftp/types"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestTextFileDriver_Write(t *testing.T) {
	type fields struct {
		FilePath string
	}
	type args struct {
		logs      []types.Log
		encodeLog LogEncoder
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    string
	}{
		{
			name: "Sanity check",
			fields: fields{
				FilePath: fmt.Sprintf("./%s_testlog.txt", uuid.NewString()),
			},
			args: args{
				logs: []types.Log{
					{
						Message:    "message",
						Timestamp:  time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
						TraceId:    "traceId",
						Level:      types.Info,
						Attributes: nil,
					},
				},
				encodeLog: func(l types.Log) string { return "works" },
			},
			wantErr: false,
			want:    "works",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			driver := &TextFileDriver{
				FilePath: tt.fields.FilePath,
			}
			if err := driver.Write(tt.args.logs, tt.args.encodeLog); (err != nil) != tt.wantErr {
				t.Errorf("TextFileDriver.Write() error = %v, wantErr %v", err, tt.wantErr)
			}

			byteContent, _ := os.ReadFile(tt.fields.FilePath)
			content := string(byteContent)

			if content != tt.want {
				t.Errorf("TextFileDriver.Write() file content = %v, want %v", content, tt.want)
			}
		})
	}
}

func TestTextFileDriver_GetBufferSize(t *testing.T) {
	type fields struct {
		FilePath string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{

			name:   "Has Buffer Size 10",
			fields: fields{FilePath: ""},
			want:   10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			driver := &TextFileDriver{
				FilePath: tt.fields.FilePath,
			}
			if got := driver.GetBufferSize(); got != tt.want {
				t.Errorf("TextFileDriver.GetBufferSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTextFileDriver_SupportsANSIColors(t *testing.T) {
	type fields struct {
		FilePath string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "Does Not Support ANSI Colors",
			fields: fields{FilePath: ""},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			driver := &TextFileDriver{
				FilePath: tt.fields.FilePath,
			}
			if got := driver.SupportsANSIColors(); got != tt.want {
				t.Errorf("TextFileDriver.SupportsANSIColors() = %v, want %v", got, tt.want)
			}
		})
	}
}
