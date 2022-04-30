package grade

import "testing"

func TestLetterGrade(t *testing.T) {
	tests := []struct {
		name  string
		score int
		want  string
	}{
		{"over A", 300, "A"},
		{"max A", 100, "A"},
		{"min A", 80, "A"},
		{"max B", 79, "B"},
		{"min B", 70, "B"},
		{"max C", 69, "C"},
		{"min C", 60, "C"},
		{"max D", 59, "D"},
		{"min D", 50, "D"},
		{"max F", 49, "F"},
		{"min F", 0, "F"},
		{"under F", -100, "F"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LetterGrade(tt.score); got != tt.want {
				t.Errorf("LetterGrade() = %v, want %v", got, tt.want)
			}
		})
	}
}
