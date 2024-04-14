package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	//...
}

// breaking the SRP by adding functions which deal with another concern
func (j *Journal) Save(filename string) {
	_ = os.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {
	//...
}

func (j *Journal) LoadFromWeb(url url.URL) {
	//...
}

// the responsability of Journal is to manage entries, NOT to handle persistence concerns.
// a separate component/struct/package should handle it.

// one solution that can work
var LineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = os.WriteFile(filename, []byte(strings.Join(j.entries, LineSeparator)), 0644)
}

// a better approach: separation of concerns 🧠
type Persistence struct {
	lineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = os.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("I cried today")
	j.AddEntry("I ate a bug")
	fmt.Println(j.String())

	// using separate function
	SaveToFile(&j, "myjournal.txt")

	//using Persistence struct
	p := Persistence{"/n"}
	p.SaveToFile(&j, "myjournal.txt")
}
