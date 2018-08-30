package modules

import (
	"bufio"
	"fmt"
	"os"
)

var world = []byte("world")

// TODO
//  this function should take as input the file path and which of three files to make(
//  the output will simply be the name of <input>.db, NICE AND SIMPLE!!!
func MakeDB(fileName string, option string) {
	// line1 := true
	// outFile := s + ".db"
	if option == "accession_id" {
		file, err := os.Open(fileName)
		if err != nil {
			println(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		buf := make([]byte, 0, 1024*1024)
		scanner.Buffer(buf, 10*1024*1024)

		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	} else {
		fmt.Println(option, "not supported yet")
	}

} // end MakeDB
