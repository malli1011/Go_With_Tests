package iteration

import "testing"

func TestRepeat(t *testing.T) {
	type args struct {
		chars string
	}
	tests := []struct {
		name         string
		args         args
		wantRepeated string
	}{
		{name: "Should repeat 5 times",
			args:         args{"Malli"},
			wantRepeated: "MalliMalliMalliMalliMalli"},
		{name: "Should repeat 5 times",
			args:         args{""},
			wantRepeated: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRepeated := Repeat(tt.args.chars); gotRepeated != tt.wantRepeated {
				t.Errorf("Repeat() = %v, want %v", gotRepeated, tt.wantRepeated)
			}
		})
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
