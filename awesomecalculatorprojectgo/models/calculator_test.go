package main

import "testing"

func Test_add(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "simple addition ",
			args: args{"5+5"},
			want: "5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculator(tt.args.expression); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func Test_bracket(t *testing.T) {
//	type args struct {
//		expression string
//	}
//	tests := []struct {
//		name string
//		args args
//		want string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := bracket(tt.args.expression); got != tt.want {
//				t.Errorf("bracket() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_calculator(t *testing.T) {
//	type args struct {
//		expression string
//	}
//	tests := []struct {
//		name string
//		args args
//		want string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := calculator(tt.args.expression); got != tt.want {
//				t.Errorf("calculator() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
