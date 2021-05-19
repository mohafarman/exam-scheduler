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
	PrintName() string
	PrintDate() string
	PrintPages() int
	PrintDaysOff() int
	PrintRevisionDays() int
	ExamDetails() (string, string)
}

type ExamDetails struct {
	Name         string
	Pages        int
	Date         string
	DaysOff      int
	RevisionDays int
}

func (eD *ExamDetails) UserInput() {
	// Initializes variables in ExamDetails() from user

	fmt.Println("Please enter the name of your exam: ")
	fmt.Scanln(&eD.Name)

	fmt.Println("Please enter the number of pages you need to go through: ")
	fmt.Scanln(&eD.Pages)

	fmt.Println("Please enter the date of the exam. (E.g. 2017-04-14): ")
	fmt.Scanln(&eD.Date)

	fmt.Println("How many days off can you afford?")
	fmt.Scanln(&eD.DaysOff)

	fmt.Println("How many days do you need for revision? (0 if none)")
	fmt.Scanln(&eD.RevisionDays)
}

func (eD ExamDetails) PrintName() string {
	// Returns name of the exam
	return eD.Name
}

func (eD ExamDetails) PrintDate() string {
	// Returns the date of the exam
	return eD.Date
}

func (eD ExamDetails) PrintPages() int {
	// Returns the date of the exam
	return eD.Pages
}

func (eD ExamDetails) PrintDaysOff() int {
	// Returns the date of the exam
	return eD.DaysOff
}

func (eD ExamDetails) PrintRevisionDays() int {
	// Returns the date of the exam
	return eD.RevisionDays
}

func (eD ExamDetails) ExamDetails() string {
	// Returns all the information about the exam
	s := fmt.Sprint("Exam: ", eD.Name, "\n", "Pages: ", eD.Pages, "\n", "Date: ", eD.Date, "\n", "Days for revision: ", eD.RevisionDays, "\n", "Days off: ", eD.DaysOff, "\n")

	return s

}

func CalculateNumOfPagesPerDay() {
	// Take in date today
	// Take in date of exam
	// Take in # of pages
	// Subtract exam - today = daysToExam
	// daysToExam - (DaysOff + RevisionDays) = daysToStudy
	// Divide: pages/daysToStudy
}

func CalculateNumOfPagesRevision() {
	// pages/RevisionDays
}

func main() {
	logger()
	WelcomeMessage()

	var exam ExamDetails

	exam.UserInput()

	fmt.Printf("Summary: \n%s", exam.ExamDetails())

	// logrus.Info("Name of the exam is %s and the date is %s.\nHere is the full exam detail: %s", exam.Name(), exam.Date(), exam.ExamDetails())

	//fmt.Printf("Name of exam is %s and the date is %s\n", exam.Name, exam.Date)
}
