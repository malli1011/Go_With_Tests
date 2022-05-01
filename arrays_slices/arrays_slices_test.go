package arrays_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	type args struct {
		nums [5]int
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{name: "First Test", args: args{[5]int{1, 2, 3, 4, 5}}, wantSum: 15},
		{name: "Second Test", args: args{[5]int{1, 1, 1, 1, 1}}, wantSum: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := Sum(tt.args.nums); gotSum != tt.wantSum {
				t.Errorf("Sum() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func TestSum_v2(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum_v2(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum_v2(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func TestSum_v1(t *testing.T) {
	type args struct {
		nums [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test with two arrays",
			args: args{nums: [][]int{{1, 1}, {2, 2}}},
			want: []int{2, 4},
		},
		{
			name: "Test with 1 arrays",
			args: args{nums: [][]int{{1, 1}, {}}},
			want: []int{2, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum_v1(tt.args.nums...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sum_v1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumTails(t *testing.T) {
	type args struct {
		nums [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantSum []int
	}{
		{
			name:    "Test with two arrays",
			args:    args{nums: [][]int{{1, 1, 1}, {2, 2, 2}}},
			wantSum: []int{2, 4},
		},
		{
			name:    "Test with 1 arrays",
			args:    args{nums: [][]int{{1, 1, 2, 3}, {}}},
			wantSum: []int{6, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := sumTails(tt.args.nums...); !reflect.DeepEqual(gotSum, tt.wantSum) {
				t.Errorf("sumTails() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
