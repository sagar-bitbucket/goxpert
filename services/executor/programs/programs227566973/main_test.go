package main

import "testing"

func Test_adder(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test1",
			args{
				2,
				3,
			},
			5,
		},
		{
			"Test2",
			args{
				1,
				3,
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := adder(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("adder() = %v, want %v", got, tt.want)
			}
		})
	}
}
