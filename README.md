# Go-GAE-Remote

This is a golang app that can be used to query your GAE datastore using [GAE Remote API](https://developers.google.com/appengine/articles/remote_api).

It's useful if you are looking to grab some data (for reporting and such) from your local machine, without having to deploy anything
to app engine. 

## Getting Started

###run the app:

``` bash
goapp run *.go
```

###config values:

main.go

```go

var (
  host     = flag.String("host", "", "hostname of application")
  email    = flag.String("email", "", "email of an admin user for the application")
  password = flag.String("password", "", "your password for accesing the app as admin")
)
```

### write queries

use queries.go to write queries and model.go to declare entity types that you are storing on datastore.

### examples

two examples are implemented as getAppStats (dsInfo.go) - taken from a Google example, and getData (queries.go).

### Fetch limitation issue workaround

The app works around a known issue, following advise from [willhorn](https://groups.google.com/forum/#!topic/google-appengine-go/fA0NptlpHNE) of GAE Remote API that does not allow you to fetch large quantities of entities by using iterators in a loop (see queries.go):

```go

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

```

## License 
The MIT License (MIT)
Copyright (c) StreamRail 2014




