package main

import (
	"appengine/datastore"
	"appengine/remote_api"
	"fmt"
	"log"
	"net/http"
)

func getData(host string, client *http.Client) []MyEntityKind {
	c, err := remote_api.NewRemoteContext(host, client)
	if err != nil {
		return nil
	}

	q := datastore.NewQuery("MyEntityKind").
		Filter("Category =", "Food")

	var i int
	t := q.Limit(100).Run(c)
	var entities []MyEntityKind
	for readBatch(&i, t, &entities) {
		cursor, err := t.Cursor()
		if err != nil {
			c.Errorf("error getting cursor: %v", err)
			return nil
		}
		t = q.Limit(100).Start(cursor).Run(c)
	}

	fmt.Printf("there are %d entities in the response\n", len(entities))
	return entities
}

func readBatch(i *int, t *datastore.Iterator, entities *[]MyEntityKind) bool {
	start := *i
	for {
		var entity MyEntityKind
		key, err := t.Next(&entity)
		*entities = append(*entities, entity)
		fmt.Printf("read entity: %v\n", key)
		if err == datastore.Done {
			if start == *i {
				return false
			}
			return true
		}
		if err != nil {
			log.Fatalf("error fetching next record: %v", err)
			return false
		}
		*i++
	}
}
