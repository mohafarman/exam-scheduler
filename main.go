package main

import (
	"fmt"
)

func WelcomeMessage() {
	fmt.Println("This is an application that will help you schedule for your exams.")
}

type examOutput interface {
	ExamName()
	ExamDate()
	ExamDetails()
}

type examDetails struct {
	Name string
	Date string
}

func (eD examDetails) UserInput() {
	fmt.Println("What exam are you taking?")
	fmt.Scanln(&eD.Name)
	fmt.Println("When is your exam? (E.g 2017-05-14)")
	fmt.Scanln(&eD.Date)
}

func (eD examDetails) ExamName() string {
	return eD.Name
}

func (eD examDetails) ExamDate() string {
	return eD.Date
}

func (eD examDetails) ExamDetails() (string, string) {
	return eD.Name, eD.Date
}

func main() {
	WelcomeMessage()

	var exam examDetails

	exam.UserInput()

	fmt.Println(exam.ExamName())
}
