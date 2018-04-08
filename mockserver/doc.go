/*
Package mockserver provides Graylog API mock server.

The mock server provides the following two type of APIs: REST API and Server object's method.

  // a port number is automatically chosen
  server, err := mockserver.NewServer("")
  if err != nil {
  	t.Error("Failed to Get Mock Server", err)
  	return
  }
  fmt.Println(server.Endpoint())

You can specify the port number.

  server, err := mockserver.NewServer(":8000")

The server has not started yet, so start it.

  server.Start()
  defer server.Close()

Then send a request with curl.

  $ curl localhost:8080/api/roles
  {"roles":[],"total":0}

The mock server uses logrus internally.

https://godoc.org/github.com/sirupsen/logrus

By default mock server's log level is "error", because debug and info logs are often noisy at unit tests. You can change log level freely.

  import log "github.com/sirupsen/logrus"

  ms.Logger().SetLevel(log.InfoLevel)

The mock server has Graylog's resources such as roles.

  role := &graylog.Role{Name: "foo", Permissions: set.NewStrSet("*")}
  sc, err := server.AddRole(role)

By default the mock server's data are temporary, but
you can persist the data to the json file.

  import "github.com/suzuki-shunsuke/go-graylog/mockserver/store/inmemory"

	s := inmemory.NewStore("data.json")
	server.SetStore(s)
  // at first load data
  if err := server.Load() {
  	server.Logger().Fatal(err)
  }
  server.Start()
  defer server.Close()

When the mock server's data are changed by a http request, it's data are written to the file automatically for persistence.

Authentication and Authorization:

The mock server supports the authentication and authorization.
If you want to make the authentication and authorization disabled, call the SetAuth method.

	server.SetAuth(false)
*/
package mockserver
