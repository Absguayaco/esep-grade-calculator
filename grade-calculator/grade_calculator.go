package esepunittests

type GradeCalculator struct {
	//	assignments []Grade
	//	exams       []Grade
	//	essays      []Grade
	grades []Grade
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		//assignments: make([]Grade, 0),
		//exams:       make([]Grade, 0),
		//essays:      make([]Grade, 0),
		grades: make([]Grade, 0),
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if numericalGrade >= 90 {
		return "A, Pass"
	} else if numericalGrade >= 80 {
		return "B, Pass"
	} else if numericalGrade >= 70 {
		return "C, Pass"
	} else if numericalGrade >= 60 {
		return "D, Fail"
	}

	return "F, Fail"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	gc.grades = append(gc.grades, Grade{
		Name:  name,
		Grade: grade,
		Type:  gradeType,
	})

}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignment_average := computeAverage(gc.grades, Assignment)
	exam_average := computeAverage(gc.grades, Exam)
	essay_average := computeAverage(gc.grades, Essay)

	weighted_grade := float64(assignment_average)*.5 + float64(exam_average)*.35 + float64(essay_average)*.15

	return int(weighted_grade)
}

func computeAverage(grades []Grade, types GradeType) int {
	sum := 0
	count := 0

	for _, g := range grades {
		if g.Type == types {
			sum += g.Grade
			count += 1
		}
	}

	return sum / count
}
