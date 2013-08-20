/*
Simple tool to approximatively count the progress of go-gtk library.
Counts line starting by "func" (binding done) or starting by "//" and containing "gtk_" (binding to do) for each section (starting by a "// text" comment preceded by a "//------------")
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	list := make([]string, 0, 100)
	alphabeticalOrder := flag.Bool("a", false, "alphabetical order")
	githubCode := flag.Bool("g", false, "prepend spaces to each line so github will format them as code")
	flag.Parse()
	fname := flag.Arg(0)
	if fname == "" {
		fmt.Println("Usage: gogtkinfo [-ag] file")
		return
	}
	file, err := os.Open(fname)
	if err != nil {
		log.Print(err.Error())
		return
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	var currentSection, sectionAlertName string
	var cDone, cTodo, p, tDone, tTodo int
	var sectionAlert, falseAlarm bool
	for {
		var l string
		if falseAlarm {
			falseAlarm = false
		} else {
			line, isPrefix, err := rd.ReadLine()
			if err != nil {
				break
			}
			if isPrefix {
				return
			}
			l = string(line)
		}
		if strings.HasPrefix(l, "//---------------") {
			if sectionAlertName == "" {
				sectionAlert = true
			} else {
				if currentSection != "" {
					cTot := cDone + cTodo
					if cTot == 0 {
						p = 100
					} else {
						p = 100 * cDone / cTot
					}
					s := fmt.Sprintf("%-30s: %3d%% (%3d/%3d)\n", currentSection, p, cDone, cTot)
					list = append(list, s)
				}
				currentSection = sectionAlertName
				tDone += cDone
				tTodo += cTodo
				cDone = 0
				cTodo = 0
				sectionAlertName = ""
			}
		} else if sectionAlert {
			if strings.HasPrefix(l, "//") && len(l) > 3 && !strings.Contains(l, "gtk_") {
				sectionAlertName = strings.TrimSpace(l[2:len(l)])
			} else {
				falseAlarm = true
			}
			sectionAlert = false
		} else if strings.HasPrefix(l, "func") {
			cDone++
		} else if strings.HasPrefix(l, "//") && strings.Contains(l, "gtk_") {
			cTodo++
		}
	}
	if *alphabeticalOrder {
		sort.StringSlice(list).Sort()
	}
	for _, s := range list {
		if *githubCode {
			fmt.Print("    ")
		}
		fmt.Print(s)
	}
	tTot := tDone + tTodo
	if tTot == 0 {
		p = 0
	} else {
		p = 100 * tDone / tTot
	}
	if *githubCode {
		fmt.Print("\n    ")
	} else {
		fmt.Print("\n")
	}
	fmt.Printf("Total progress : %18d%% (%d/%d)\n", p, tDone, tTot)
}
