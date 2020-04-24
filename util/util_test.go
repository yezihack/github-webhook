package util

import (
	"strings"
	"testing"
)

func TestCallScript(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "test1",
			args:    args{path: "tool.sh"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CallScript(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("CallScript() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if strings.TrimSpace(got) != tt.want {
				t.Errorf("CallScript() got = %v, want %v", got, tt.want)
			}
		})
	}
}
