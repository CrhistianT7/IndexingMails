package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
}

func main() {
	// args := os.Args
	// fmt.Println("hello from backend")
	// fmt.Println(args[1])
	// _, err := os.ReadDir("../../" + args[1] + "/maildir")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// -------------------------------------------
	// HERE IT BEGINS
	// set application config
	var app application
	// read from command line

	// connect to the database
	app.Domain = "example.com"
	log.Println("Starting application on port", port)

	http.HandleFunc("/", Hello)

	// start a web server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}

	// -------------------------------------------

	// fmt.Println(entries)

	// for _, e := range entries {
	// 	fmt.Println(e.Name())
	// 	fmt.Println("----")
	// 	ent, err := os.ReadDir("../../" + args[1] + "/maildir/" + e.Name())
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	for _, f := range ent {
	// 		fmt.Println(f.Name())

	// 		dat, err := os.ReadFile("../../enron_mail_20110402/maildir/" + e.Name() + "/" + f.Name() + "/1.")
	// 		if err != nil {

	// 		}
	// 		fmt.Println("\n begin *****")
	// 		fmt.Print(string(dat))
	// 		fmt.Println("\n end *****")
	// 	}
	// }
}
