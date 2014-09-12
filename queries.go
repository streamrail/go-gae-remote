package main

import (
	"appengine/datastore"
	"appengine/remote_api"
	"net/http"
)

func getData(host string, client *http.Client) ([]SampleKind, error) {
	c, err := remote_api.NewRemoteContext(host, client)
	if err != nil {
		return nil, err
	}

	// The Query type and its methods are used to construct a query.
	q := datastore.NewQuery("SampleKind").
		Filter("Category =", "food")

	// To retrieve the results,
	// you must execute the Query using its GetAll or Run methods.
	var evts []SampleKind
	_, err2 := q.GetAll(c, &evts)
	return evts, err2
}
