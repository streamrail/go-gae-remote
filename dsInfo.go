package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"appengine"
	"appengine/datastore"
	"appengine/remote_api"
)

const DatastoreKindName = "__Stat_Kind__"

type DatastoreKind struct {
	KindName            string    `datastore:"kind_name"`
	EntityBytes         int       `datastore:"entity_bytes"`
	BuiltinIndexBytes   int       `datastore:"builtin_index_bytes"`
	BuiltinIndexCount   int       `datastore:"builtin_index_count"`
	CompositeIndexBytes int       `datastore:"composite_index_bytes"`
	CompositeIndexCount int       `datastore:"composite_index_count"`
	Timestamp           time.Time `datastore:"timestamp"`
	Count               int       `datastore:"count"`
	Bytes               int       `datastore:"bytes"`
}

func getAppStats(host string, client *http.Client) {
	c, err := remote_api.NewRemoteContext(host, client)
	if err != nil {
		log.Fatalf("Failed to create context: %v", err)
	}
	log.Printf("App ID %q", appengine.AppID(c))

	q := datastore.NewQuery(DatastoreKindName).Order("kind_name")
	kinds := []*DatastoreKind{}
	if _, err := q.GetAll(c, &kinds); err != nil {
		log.Fatalf("Failed to fetch kind info: %v", err)
	}

	for _, k := range kinds {
		fmt.Printf("\nkind %q\t%d entries\t%d bytes\n", k.KindName, k.Count, k.Bytes)

		props := datastore.PropertyList{}
		if _, err := datastore.NewQuery(k.KindName).Limit(1).Run(c).Next(&props); err != nil {
			log.Printf("Unable to fetch sample entity kind %q: %v", k.KindName, err)
			continue
		}
		for _, prop := range props {
			fmt.Printf("\t%s: %v\n", prop.Name, prop.Value)
		}
	}
}
