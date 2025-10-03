package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := string(gradeCalculator.GetFinalGrade()[0])

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := string(gradeCalculator.GetFinalGrade()[0])

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 70, Assignment)
	gradeCalculator.AddGrade("exam 1", 72, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 75, Essay)

	actual_value := string(gradeCalculator.GetFinalGrade()[0])

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 60, Assignment)
	gradeCalculator.AddGrade("exam 1", 61, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 65, Essay)

	actual_value := string(gradeCalculator.GetFinalGrade()[0])

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}
func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 0, Assignment)
	gradeCalculator.AddGrade("exam 1", 20, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 30, Essay)

	actual_value := string(gradeCalculator.GetFinalGrade()[0])

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeTypeString_All(t *testing.T) {
	expected_assignment := "assignment"
	expected_exam := "exam"
	expected_essay := "essay"

	actual_assignment := Assignment.String()
	actual_exam := Exam.String()
	actual_essay := Essay.String()

	if expected_assignment != actual_assignment {
		t.Errorf("Expected assignment")
	}
	if expected_exam != actual_exam {
		t.Errorf("Expected exam")
	}
	if expected_essay != actual_essay {
		t.Errorf("Expected essay")
	}
}

func TestGetFinalGrade_A_Pass(t *testing.T) {
	expected_value := "Pass"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("open source assignment", 90, Assignment)
	gradeCalculator.AddGrade("exam 1", 92, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 95, Essay)

	actual_value := string(gradeCalculator.GetFinalGrade()[3:7])

	if expected_value != actual_value {
		t.Errorf("Expected GetFinalGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetFinalGrade_C_Pass(t *testing.T) {
	expected_value := "Pass"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("open source assignment", 70, Assignment)
	gradeCalculator.AddGrade("exam 1", 72, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 75, Essay)

	actual_value := string(gradeCalculator.GetFinalGrade()[3:7])

	if expected_value != actual_value {
		t.Errorf("Expected GetFinalGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}
func TestGetFinalGrade_B_Pass(t *testing.T) {
	expected_value := "Pass"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 82, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := string(gradeCalculator.GetFinalGrade()[3:7])

	if expected_value != actual_value {
		t.Errorf("Expected GetFinalGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetFinalGrade_D_Fail(t *testing.T) {
	expected_value := "Fail"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("open source assignment", 60, Assignment)
	gradeCalculator.AddGrade("exam 1", 62, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 65, Essay)

	actual_value := string(gradeCalculator.GetFinalGrade()[3:7])

	if expected_value != actual_value {
		t.Errorf("Expected GetFinalGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetFinalGrade_F_Fail(t *testing.T) {
	expected_value := "Fail"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("open source assignment", 40, Assignment)
	gradeCalculator.AddGrade("exam 1", 42, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 45, Essay)

	actual_value := string(gradeCalculator.GetFinalGrade()[3:7])

	if expected_value != actual_value {
		t.Errorf("Expected GetFinalGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}
