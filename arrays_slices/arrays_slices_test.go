package arrays_slices

import "testing"

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
