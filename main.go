package main

import "fmt"

func main() {
	var numberOfStudents int
	var id int64
	var ids []int64
	var name string
	var names []string
	var midTermScore float64
	var semesterScore float64
	var attendanceScore float64
	var finalScore float64
	var finalScores []float64
	var grade string
	var grades []string
	var passed int64
	var failed int64

	fmt.Print("Input the number of students: ")
	fmt.Scanln(&numberOfStudents)

	for i := 0; i < numberOfStudents; i++ {
		fmt.Println("===Student Data===")
		fmt.Print("Student ID: ")
		fmt.Scanln(&id)
		ids = append(ids, id)

		fmt.Print("Student Name: ")
		fmt.Scanln(&name)
		names = append(names, name)

		fmt.Print("Mid Term Test Score: ")
		fmt.Scanln(&midTermScore)

		fmt.Print("Semester Test Score: ")
		fmt.Scanln(&semesterScore)

		fmt.Print("Attendance: ")
		fmt.Scanln(&attendanceScore)

		finalScore = (0.2 * attendanceScore) + (0.4 * midTermScore) + (0.4 * semesterScore)
		finalScores = append(finalScores, finalScore)

		if finalScore >= 85 && finalScore <= 100 {
			grade = "A"
			passed++
		} else if finalScore > 75 && finalScore < 85 {
			grade = "B"
			passed++
		} else if finalScore > 60 && finalScore <= 75 {
			grade = "C"
			passed++
		} else if finalScore > 45 && finalScore <= 60 {
			grade = "D"
			failed++
		} else if finalScore >= 0 && finalScore <= 45 {
			grade = "E"
			failed++
		}
		grades = append(grades, grade)
	}

	fmt.Println("==========================================================")
	fmt.Println("No. 	Student ID 	Name 	Final Score 	Grade")
	fmt.Println("==========================================================")
	for i := 0; i < numberOfStudents; i++ {
		fmt.Printf("%v. 	%v 		%v	%v		%v", i+1, ids[i], names[i], finalScores[i], grades[i])
		fmt.Println("")
	}
	fmt.Println("==========================================================")
	fmt.Printf("Number of Students : %v", numberOfStudents)
	fmt.Println("")
	fmt.Printf("Number of Passing Students : %v", passed)
	fmt.Println("")
	fmt.Printf("Number of Failed Students : %v", failed)
	fmt.Println("")
}
