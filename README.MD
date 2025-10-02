## ESEP Grade Calculator

This library is a simple utility tool that takes the weighted average of assignments, exams, and essays and converts it into a letter grade.

## Grade Scale and Weights
Grade Scale:
A >=90
B >=80
C >=70
D >= 60
F <60
Passing is C or greater. Assignments are 50%, exams are 35%, and essays are 15%.

## Runtime requirements:
Golang 1.24+

## Running Tests 

#Clone Repository
git clone https://github.com/Absguayaco/esep-grade-calculator.git

Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```