package main

import (
	"embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed lang/*.json
var f embed.FS

func main() {
	langPtr := flag.String("lang", "fr", "Language to translate to (fr == french, es == Spanish, etc.)")
	fPtr := flag.String("file", "tasklist_en.json", "Location of the English input file")
	flag.Parse()
	// dates for all
	// fmt.Printf("Translating dates to %s\n", *langPtr)
	// dates := DateLocales{}
	// dates.TranslateDateLocales(*langPtr)
	// task list
	fmt.Printf("Translating task list to %s\n", *langPtr)
	t := TaskListInterface{}
	t.LoadFile(*fPtr, f)
	t.TranslateTaskListInterface(*langPtr)
	// t.Dates = dates
	t.SaveTaskListInterface(*fPtr, *langPtr)
  // Cockpit
	fmt.Printf("Translating cockpit to %s\n", *langPtr)
	cockpitFile := strings.Replace(*fPtr, "tasklist", "cockpit", -1)
	c := CockpitInterface{}
	c.LoadFile(cockpitFile)
	c.TranslateCockpitInterface(*langPtr)
	// c.Dates = dates
	c.SaveCockpitInterface(cockpitFile, *langPtr)
  // Admin
	fmt.Printf("Translating admin to %s\n", *langPtr)
	adminFile := strings.Replace(*fPtr, "tasklist", "admin", -1)
	a := AdminInterface{}
	a.LoadFile(adminFile)
	a.TranslateAdminInterface(*langPtr)
	a.SaveAdminInterface(adminFile, *langPtr)
	// Welcome
	fmt.Printf("Translating welcome to %s\n", *langPtr)
	welcomeFile := strings.Replace(*fPtr, "tasklist", "welcome", -1)
	w := WelcomeInterface{}
	w.LoadFile(welcomeFile)
	w.TranslateWelcomeInterface(*langPtr)
	w.SaveWelcomeInterface(welcomeFile, *langPtr)
}