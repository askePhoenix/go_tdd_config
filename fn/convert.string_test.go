package fn

import "testing"

func TestIntToString(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"성공", args{10}, "10"},
		{"성공", args{11}, "11"},
		{"성공", args{0}, "0"},
		{"성공", args{1000}, "1000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToString(tt.args.i); got != tt.want {
				t.Errorf("IntToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64ToString(t *testing.T) {
	type args struct {
		i int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"성공", args{int64(10)},"10"},
		{"성공", args{int64(11)}, "11"},
		{"성공", args{int64(0)}, "0"},
		{"성공", args{int64(1000)}, "1000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64ToString(tt.args.i); got != tt.want {
				t.Errorf("Int64ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}