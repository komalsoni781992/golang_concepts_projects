package main

import (
	"fmt"
	"reflect"
)

/*q4. Print default values and Type names of variables from question 2 using printf
// Quick Tip, Use %v if not sure about what verb should be used,
// but don't use it in this question :)
// but generally using %v should be fine*/

func main() {
	var (
		projectName        string
		codeLines          uint8
		bugsFound          int
		isComplete         bool
		averageLinesOfCode float64
		teamLead           string
		projectDeadline    int
	)
	fmt.Printf("%T, %T, %T, %T, %T, %T, %T\n", projectName, codeLines, bugsFound, isComplete, averageLinesOfCode, teamLead, projectDeadline)
	fmt.Printf("%v, %v, %v, %v, %v, %v, %v\n", reflect.TypeOf(projectName), reflect.TypeOf(codeLines), reflect.TypeOf(bugsFound), reflect.TypeOf(isComplete), reflect.TypeOf(averageLinesOfCode), reflect.TypeOf(teamLead), reflect.TypeOf(projectDeadline))
	fmt.Printf("%v, %v, %v, %v, %v, %v, %v\n", projectName, codeLines, bugsFound, isComplete, averageLinesOfCode, teamLead, projectDeadline)
	fmt.Printf("%#v, %#v, %#v, %#v, %#v, %#v, %#v\n", projectName, codeLines, bugsFound, isComplete, averageLinesOfCode, teamLead, projectDeadline)

}
