package mapper

import (
	"testing"
)
type Source struct {
	Idx int64
	Name string
}

type Dest struct {
	Idx int64
}
func TestMapper(t *testing.T) {
	type args struct {
		source interface{}
		dest   interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{name: "", args: args{ source: Source{Idx: 1, Name: "namil"}, dest: Dest{}}, want: Dest{Idx: 1}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Mapper(tt.args.source, tt.args.dest)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mapper() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == tt.want {
				t.Errorf("Mapper() got = %v, want %v", got, tt.want)
			}
		})
	}
}