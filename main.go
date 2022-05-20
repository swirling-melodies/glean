package main

import (
	"fmt"
	"glean/genericityGlean"
	"glean/typeGlean"
)

type IGleanWork interface {
	Run() error
}

type Action func(gleanWork IGleanWork) error

func main() {

	var i = selected()
	if i == nil {
		return
	}
	run(work, i)
}

func selected() IGleanWork {
	fmt.Println(" ------------------ Please  Select ------------------ ")
	fmt.Println(" Please enter the number:")
	fmt.Println(" 1.TypeFunc ")
	fmt.Println(" 2.GenericityDemo ")
	fmt.Println(" 0.Exit ")

	var num int
	fmt.Scan(&num)
	var i IGleanWork
	switch num {
	case 1:
		i = &typeGlean.TypeClass{}
		break
	case 2:
		i = &genericityGlean.GenericityDemo{}
		break
	case 0:
		return nil

	default:
		fmt.Println(" Error, please re-enter")
		return selected()
	}

	return i
}

func run(action Action, gleanWork IGleanWork) {
	fmt.Println(" -------------------- Work Start -------------------- ")
	err := action(gleanWork)
	if err != nil {
		fmt.Println(" -------------------- Work Error -------------------- ")
		fmt.Println(err)
		fmt.Println(" -------------------- Work Error -------------------- ")
	}
	fmt.Println(" --------------------- Work End --------------------- ")
}

func work(gleanWork IGleanWork) error {
	fmt.Println(" ------------------- Work Running ------------------- ")
	return gleanWork.Run()
}