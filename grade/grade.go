package grade

// function for calculate letter grade from score.
// Grade A: score 80-100
// Grade B: score 70-79
// Grade C: score 60-69
// Grade D: score 50-59
// Grade F: score 0-49
func LetterGrade(score int) string {
	if score >= 80 {
		return "A"
	}
	if score >= 70 {
		return "B"
	}
	if score >= 60 {
		return "C"
	}
	if score >= 50 {
		return "D"
	}
	return "F"
}
