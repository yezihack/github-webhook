package logger

import "testing"

func TestLogger_Print(t *testing.T) {
	type fields struct {
		quiet bool
	}
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "1",
			fields: fields{quiet:false},
			args:   args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Logger{
				quiet: tt.fields.quiet,
			}
		})
	}
}