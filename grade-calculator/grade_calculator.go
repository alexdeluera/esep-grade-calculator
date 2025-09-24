package esepunittests

type GradeCalculator struct {
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
		grades: make([]Grade, 0),
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	} else {
		return "F"
	}

}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	gc.grades = append(gc.grades, Grade{
		Name:  name,
		Grade: grade,
		Type:  gradeType,
	})

}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignment_average := gc.computeAverage(Assignment)
	exam_average := gc.computeAverage(Exam)
	essay_average := gc.computeAverage(Essay)

	weighted_grade := assignment_average*0.5 + exam_average*0.35 + essay_average*0.15

	return int(weighted_grade + 0.5)
}

func (gc *GradeCalculator) computeAverage(gradeType GradeType) float64 {
	sum := 0
	count := 0

	for _, g := range gc.grades {
		if g.Type == gradeType {
			sum += g.Grade
			count++
		}
	}

	return float64(sum) / float64(count)
}
