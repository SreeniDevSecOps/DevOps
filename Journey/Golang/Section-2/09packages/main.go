package main

import (
	"github.com/SreeniDevSecOps/DevOps/learningPackages/testPackage"
	"log"
)

func main() {

	log.Println("Hello")

	var myVar testPackage.Example

	myVar.TypeName = "Kunal"
	myVar.TypeNum = 19

}
