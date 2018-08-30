package modules

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/boltdb/bolt"
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
		// make database
		accession_id_db, err := bolt.Open(fileName+".db", 0644, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer accession_id_db.Close()
		for scanner.Scan() {
			if line1 {
				line1 = false
				continue
			}

			line := scanner.Text()
			row := strings.Fields(line)
			// fmt.Println(row[0], row[2])

			key := []byte(row[0])
			value := []byte(row[2])

			// store some data
			err = accession_id_db.Update(func(tx *bolt.Tx) error {
				bucket, err := tx.CreateBucketIfNotExists([]byte("MYBUCKET"))
				if err != nil {
					return err
				}

				err = bucket.Put(key, value)
				if err != nil {
					return err
				}
				return nil
			})
			if err != nil {
				log.Fatal(err)
			}

			// TODO
			//  split into space delimited row
			// skip first line
			// make map, {assesion:taxid}, 1st and 3rd column
			//  add to fileName +".db"
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println(option, "not supported yet")
	}
	// the path does not exist or some error occurred.

	// if id_nodes
	// do things
	// if id_parent
	// do other things
	// if accession_id
	// do slow things

	//
	// db, err := bolt.Open("/Users/jeffcole/bolt.db", 0644, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()
	//
	// key := []byte("hello")
	// value := []byte("Hello World!")
	//
	// // store some data
	// err = db.Update(func(tx *bolt.Tx) error {
	// 	bucket, err := tx.CreateBucketIfNotExists(world)
	// 	if err != nil {
	// 		return err
	// 	}
	//
	// 	err = bucket.Put(key, value)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return nil
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// retrieve the data

	// worlds := []byte("world")
	// err = db.View(func(tx *bolt.Tx) error {
	// 	bucket := tx.Bucket(worlds)
	// 	if bucket == nil {
	// 		return fmt.Errorf("Bucket %q not found!", worlds)
	// 	}
	// 	keyOld := []byte("helloo")
	// 	val := bucket.Get(keyOld)
	// 	fmt.Println(string(val))
	// 	return nil
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

} // end MakeDB
