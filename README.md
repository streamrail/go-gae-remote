# Go-GAE-Remote

This is a golang app that can be used to query your GAE datastore using [GAE Remote API](https://developers.google.com/appengine/articles/remote_api).

It's useful if you are looking to grab some data (for reporting and such) from your local machine, without having to deploy anything
to app engine. 

## Getting Started

#run the app:

``` bash
goapp run *.go
```

#config values:

main.go

```go

var (
  host     = flag.String("host", "", "hostname of application")
  email    = flag.String("email", "", "email of an admin user for the application")
  password = flag.String("password", "", "your password for accesing the app as admin")
)
```

#write queries

use queries.go to write queries and model.go to declare entity types that you are storing on datastore.

two examples are implemented as getAppStats (dsInfo.go) and getData (queries.go).

# License 
2014 StreamRail all rights reserved




