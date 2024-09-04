package driver

import (
	"mihaimiuta/mftp/pkg/mftp/types"
	"testing"
)

func TestStandardOutputDriver_Write(t *testing.T) {
	type args struct {
		logs      []types.Log
		encodeLog LogEncoder
	}
	tests := []struct {
		name    string
		driver  *StandardOutputDriver
		args    args
		wantErr bool
	}{
		{
			name:   "Sanity Check",
			driver: nil,
			args: args{
				logs:      []types.Log{},
				encodeLog: func(l types.Log) string { return "" },
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			driver := &StandardOutputDriver{}
			if err := driver.Write(tt.args.logs, tt.args.encodeLog); (err != nil) != tt.wantErr {
				t.Errorf("StandardOutputDriver.Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStandardOutputDriver_GetBufferSize(t *testing.T) {
	tests := []struct {
		name   string
		driver *StandardOutputDriver
		want   int
	}{
		{
			name:   "Has Buffer Size 1",
			driver: nil,
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			driver := &StandardOutputDriver{}
			if got := driver.GetBufferSize(); got != tt.want {
				t.Errorf("StandardOutputDriver.GetBufferSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStandardOutputDriver_SupportsANSIColors(t *testing.T) {
	tests := []struct {
		name   string
		driver *StandardOutputDriver
		want   bool
	}{
		{
			name:   "Supports ANSI Colors",
			driver: nil,
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			driver := &StandardOutputDriver{}
			if got := driver.SupportsANSIColors(); got != tt.want {
				t.Errorf("StandardOutputDriver.SupportsANSIColors() = %v, want %v", got, tt.want)
			}
		})
	}
}
