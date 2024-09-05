package logger

import (
	"mihaimiuta/mftp/pkg/mftp/types"
	"reflect"
	"testing"
	"time"
)

func TestMakeLogEncoder(t *testing.T) {
	type args struct {
		shouldUseColors bool
		log             types.Log
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Without Colors",
			args: args{
				shouldUseColors: false,
				log: types.Log{
					Message:    "message",
					Timestamp:  time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
					TraceId:    "traceId",
					Level:      types.Info,
					Attributes: nil,
				},
			},
			want: "[Level=info][Timestamp=1970-01-01 00:00:00 +0000 UTC][TraceID=traceId] message\n",
		},
		{
			name: "With Colors Info",
			args: args{
				shouldUseColors: true,
				log: types.Log{
					Message:    "message",
					Timestamp:  time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
					TraceId:    "traceId",
					Level:      types.Info,
					Attributes: nil,
				},
			},
			want: "\033[97m[Level=info][Timestamp=1970-01-01 00:00:00 +0000 UTC][TraceID=traceId] message\n\033[0m",
		},
		{
			name: "With Attributes",
			args: args{
				shouldUseColors: false,
				log: types.Log{
					Message:   "message",
					Timestamp: time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
					TraceId:   "traceId",
					Level:     types.Info,
					Attributes: map[string]string{
						"key": "value",
					},
				},
			},
			want: "[Level=info][Timestamp=1970-01-01 00:00:00 +0000 UTC][TraceID=traceId][Attributes={[key=value]}] message\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeLogEncoder(tt.args.shouldUseColors)(tt.args.log); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeLogEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}
