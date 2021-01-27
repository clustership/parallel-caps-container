package main
 
import (
	"bufio"
	"fmt"
	"log"
	"os"
  "strings"
)
 
func main() {
  scmd := buildCmdArray("test.txt")
	for _, eachline := range scmd {
		// fmt.Println(strings.Join(eachline[4:len(eachline)], "|"))
		fmt.Println(strings.Join(eachline, "|"))
	}
}

func buildCmdArray(filename string) [][]string {
	file, err := os.Open(filename)
	defer file.Close()
  
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
 
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
  var commands [][]string
 
	for scanner.Scan() {
    var txtline = scanner.Text()

		commands = append(commands, strings.Fields(txtline))
	}
 
 

  return commands
}
