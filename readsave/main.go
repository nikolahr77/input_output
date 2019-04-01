package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//File contains the name of the file
type File struct {
	//FileName is the name of the file
	FileName string
}

//StringCreator takes the input from the user and writes it to a single string
//StringCreator stops filling the string when an dot (".") is entered
func StringCreator() (s string) {
	for {
		var m string
		_, err := fmt.Scan(&m)
		if err != nil {
			fmt.Println("Input expected", err)
		}
		s = s + " " + m
		if m == "." {
			break
		}
	}
	return s
}

//SaveToFile creates a new directory or/and file if not existant
//and saves a string to the file
func (f File) Save(s string) error{
	d1 := []byte(s)
	if _, err := os.Stat(f.FileName); os.IsNotExist(err) {
		fmt.Println("Directory does not exist, creating directory")
		os.Create(f.FileName)
	}
	err := ioutil.WriteFile(f.FileName, d1, 0777)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Written to file")
	return nil
}

func main() {
	s := StringCreator()
	f := File{FileName: "newFile.txt"}
	f.Save(s)
}
