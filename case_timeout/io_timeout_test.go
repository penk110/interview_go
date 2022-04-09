package case_timeout

import "testing"

func TestQ1(t *testing.T) {
	type args struct {
		writer IWriter
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "not timeout",
			args:    args{writer: timeout1{}},
			wantErr: false,
		},
		{
			name:    "timeout",
			args:    args{writer: timeout2{}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Q1(tt.args.writer); (err != nil) != tt.wantErr {
				t.Errorf("Q1() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
