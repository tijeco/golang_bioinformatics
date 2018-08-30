package modules

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var world = []byte("world")

// TODO
//  this function should take as input the file path and which of three files to make(
//  the output will simply be the name of <input>.db, NICE AND SIMPLE!!!
func MakeDB(fileName string, option string) {
	line1 := true
	// outFile := s + ".db"
	if option == "accession_id" {
		file, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			if line1 {
				line1 = false
				continue
			}

			line := scanner.Text()
			row := strings.Fields(line)

			key := row[0]
			value := row[2]
			fmt.Println(key, value)

			if err != nil {
				log.Fatal(err)
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println(option, "not supported yet")
	}

} // end MakeDB
