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
	k1 := strings.Replace(string(resto), "", " ", -1)
	k2 := strings.Replace(k1, "\r\n", " ", -1)
	k3 := strings.Replace(k2, "ZONA", "\r\n ZONA", -1)
	CreateFile("filename", k3)
	log.Println("green")

}

func CreateFile(filename, text string) {

	file, err := os.Create(filename)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close()
	len, err := file.WriteString(text)
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", len)
}
