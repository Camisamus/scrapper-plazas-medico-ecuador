package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var param []string

//____________________________________________________________________________________________
func main() {
	content, err := ioutil.ReadFile("bdd.txt")
	if err != nil {
		log.Println(err.Error())
	}
	resto := string(content)
	k1 := strings.Replace(string(resto), " DISTRITO PROVINCIA CANTON PARROQUIA UNICODIGO ESTABLECIMIENTO TIPOLOGIA NÚMERO DE PLAZAS CARRERA Nro. ", " ", -1)
	k2 := strings.Replace(k1, "\r\n", " ", -1)
	k3 := strings.Replace(k2, "ZONA", "\r\n ZONA", -1)
	CreateFile("filename", k3)

}

func CreateFile(filename, text string) {

	// fmt package implements formatted I/O and
	// contains inbuilt methods like Printf
	// and Scanf
	fmt.Printf("Writing to a file in Go lang\n")

	// Creating the file using Create() method
	// with user inputted filename and err
	// variable catches any error thrown
	file, err := os.Create(filename)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	// closing the running file after the main
	// method has completed execution and
	// the writing to the file is complete
	defer file.Close()

	// writing data to the file using
	// WriteString() method and the
	// length of the string is stored
	// in len variable
	len, err := file.WriteString(text)
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", len)
}
