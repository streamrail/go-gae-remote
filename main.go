package main

import (
	"flag"
	"log"
)

// 	sample values:
//  -email admin@example.com \
//  -host my-app.appspot.com \
//  -password 4vy3@@!dzc=f0

var (
	host     = flag.String("host", "", "hostname of application")
	email    = flag.String("email", "", "email of an admin user for the application")
	password = flag.String("password", "", "your password for accesing the app as admin")
)

func main() {
	flag.Parse()

	if *host == "" {
		log.Fatalf("Required flag: -host")
	}
	if *email == "" {
		log.Fatalf("Required flag: -email")
	}
	if *password == "" {
		log.Fatalf("Required flag: -password")
	}

	client := clientLoginClient(*host, *email, *password)

	getData(*host, client)
}
