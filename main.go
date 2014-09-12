package main

// Simple tool which lists the entity kinds, and a sample for each, for an
// app engine app.
//
// This tool can be invoked using the goapp tool bundled with the SDK.
// $ goapp run demos/remote_api/datastore_info.go \
//   -email admin@example.com \
//   -host my-app.appspot.com \
//   -password 4vy3@@!dzc=f0

import (
	"flag"
	"log"
)

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
	getAppStats(*host, client)

	if ents, err := getData(*host, client); err != nil {
		log.Fatalf(err.Error())
	} else {
		log.Printf("all entities: %d", len(ents))
	}
}
