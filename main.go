package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
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
	//var student []student
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
	for key, val := range data {
		fmt.Println(key, val)

	}

	return nil
}

func calculateGrade(students []student) []studentStat {
	return nil
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	return studentStat{}
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	return nil
}
