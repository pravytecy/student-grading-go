package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

type student struct {
	firstName, lastName, university                string
	test1Score, test2Score, test3Score, test4Score int
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

func main() {
	parseCSV("grades.csv")
}

func parseCSV(filePath string) []student {
	var students []student
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}
	for index, val := range data {
		if index > 0 {
			var Student student
			for i, field := range val {
				fmt.Println(i, field)
				if i == 0 {
					Student.firstName = field
				} else if i == 1 {
					Student.lastName = field
				} else if i == 2 {
					Student.university = field
				} else if i == 3 {
					Student.test1Score = StringToInt(field)
				} else if i == 4 {
					Student.test2Score = StringToInt(field)
				} else if i == 5 {
					Student.test3Score = StringToInt(field)
				} else if i == 6 {
					Student.test4Score = StringToInt(field)
				}
			}
			students = append(students, Student)
		}
	}
	calculateGrade(students)
	return students
}

func StringToInt(field string) int {
	value, err := strconv.Atoi(field)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func calculateGrade(students []student) []studentStat {
	var studentStats []studentStat
	for _, st := range students {
		var stStat studentStat
		avg := float32(st.test1Score+st.test2Score+st.test3Score+st.test4Score) / 4
		stStat.student = st
		stStat.finalScore = avg
		stStat.grade = GetGrade(avg)
		studentStats = append(studentStats, stStat)
	}
	return studentStats
}

func GetGrade(score float32) Grade {
	if score >= 70 {
		return A
	} else if score < 70 && score >= 50 {
		return B
	} else if score < 50 && score >= 35 {
		return C
	} else {
		return F
	}
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	return studentStat{}
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	return nil
}
