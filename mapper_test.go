package mapper

import (
	"reflect"
	"testing"
	"time"
)

type (
	Class struct {
		Name  string
		Level string
	}

	Person struct {
		Id    string
		Name  string
		Age   int
		Birth time.Time
		Class Class
	}

	Student struct {
		No    int
		Name  string
		Age   int
		Class interface{}
	}
)

func BenchmarkMapper(b *testing.B) {
	tests := []struct {
		source interface{}
		dest interface{}
	}{
		{source: nil, dest: nil},
		{source: Person{}, dest: nil},
		{source: Person{Id: "nih", Name: "chris han", Age: 30, Class: Class{Name: "Deep learning", Level: "basic"}}, dest: Student{}},
		{source: Person{Name: "chris han", Age: 30, Birth: time.Now().Add(time.Hour*24*365*30*-1)}, dest: Student{}},
	}
	for _, tt := range tests {
		_, _ = Mapper(tt.source, tt.dest)
	}
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
		{name: "Zero value source", args: args{ source: nil, dest: Student{}}, want: nil, wantErr: true},
		{name: "Zero value dest", args: args{ source: Person{}, dest: nil}, want: nil, wantErr: true},
		{name: "Zero value source/dest", args: args{ source: nil, dest: nil}, want: nil, wantErr: true},
		{name: "Default value source", args: args{ source: Person{}, dest: Student{}}, want: Student{Class: Class{}}, wantErr: false},
		{name: "Default value dest", args: args{ source: Person{Name: "chris han", Age: 30, Birth: time.Now().Add(time.Hour*24*365*30*-1)}, dest: Student{}}, want: Student{Name: "chris han", Age: 30, Class: Class{}}, wantErr: false},
		{name: "Normal", args: args{ source: Person{Id: "nih", Name: "chris han", Age: 30, Class: Class{Name: "Deep learning", Level: "basic"}}, dest: Student{}}, want: Student{Name: "chris han", Age: 30, Class: Class{Name: "Deep learning", Level: "basic"}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Mapper(tt.args.source, tt.args.dest)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mapper() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !IsEqual(got, tt.want) {
				t.Errorf("Mapper() got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func IsEqual(dest, want interface{}) bool {
	return reflect.DeepEqual(dest.(reflect.Value).Interface().(Student), want)
}