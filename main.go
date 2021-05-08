package main

import (
	"fmt"
	"os"

	logrus "github.com/sirupsen/logrus"
)

func logger() {
	// Initialize our logger from logrus

	var filename string = "exam-scheduler.log"

	// Will create a file if it doesn't exist, will append if it exists
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	Formatter := new(logrus.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	logrus.SetFormatter(Formatter)

	if err != nil {
		fmt.Println(err)
	} else {
		logrus.SetOutput(f)
	}
}

func WelcomeMessage() {
	fmt.Println("This is an application that will help you schedule for your exams.")
}

type ExamOutput interface {
	ExamName() string
	ExamDate() string
	ExamDetails() (string, string)
}

type ExamDetails struct {
	Name string
	Date string
}

func (eD *ExamDetails) UserInput() {
	// Initializes variables in ExamDetails() from user

	fmt.Println("Please enter the name of your exam: ")
	fmt.Scanln(&eD.Name)
	fmt.Println("Please enter the date of the exam. (E.g. 2017-04-14): ")
	fmt.Scanln(&eD.Date)
}

func (eD ExamDetails) ExamName() string {
	// Returns name of the exam

	return eD.Name
}

func (eD ExamDetails) ExamDate() string {
	// Returns the date of the exam

	return eD.Date
}

func (eD ExamDetails) ExamDetails() string {
	// Returns all the information about the exam

	s := eD.Name + " " + eD.Date
	return s
}

func main() {
	logger()
	WelcomeMessage()

	var exam ExamDetails

	exam.UserInput()

	fmt.Printf("Name of the exam is %s and the date is %s.\nHere is the full exam detail: %s", exam.ExamName(), exam.ExamDate(), exam.ExamDetails())

	// logrus.Info("Name of the exam is %s and the date is %s.\nHere is the full exam detail: %s", exam.ExamName(), exam.ExamDate(), exam.ExamDetails())

	//fmt.Printf("Name of exam is %s and the date is %s\n", exam.Name, exam.Date)
}
