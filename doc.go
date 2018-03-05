/*
Package graylog provides Graylog API client and mock server.

Usage:

  import "github.com/suzuki-shunsuke/go-graylog"

Construct a new Graylog client.

  client, err := NewClient("http://localhost:9000/api", "admin", "password")

Of course, you can use the access token and session token instead of password.

  client, err := NewClient("http://localhost:9000/api", "htgi84ut7jpivsrcldd6l4lmcigvfauldm99ofcb4hsfcvdgsru", "token")

And you can call various Graylog REST APIs as client methods.
For example, create a role.

	params := &graylog.Role{Name: "foo", Permissions: []string{"*"}}
	role, _, err := client.CreateRole(params)

In addition the conventional error object, our API returns "graylog.ErrorInfo" object. This object has http.Response object and Graylog API's error message.

  role, ei, err := client.CreateRole(params)
  if err != nil {
  	if ei == nil {
  		log.Fatal(err)
  	}
  	log.Fatalf("%s, %s, %s, %s", err, ei.Type, ei.Message, ei.Response.StatusCode)
  }

In addition the API client, this library provides the mock server.

  // a port number is automatically chosen
  server, err := NewMockServer("")
  if err != nil {
  	t.Error("Failed to Get Mock Server", err)
  	return
  }

You can specify the port number.

  server, err := NewMockServer(":8000")

The server doesn't start yet, so start it.

  server.Start()
  defer server.Close()

Then send a request with curl.

  $ curl localhost:8080/api/roles
  {"roles":[],"total":0}

The mock server uses logrus internally.
By default mock server's log level is "error", because debug and info logs are often noisy at unit tests. You can change log level freely.

  import log "github.com/sirupsen/logrus"

  ms.Logger.SetLevel(log.InfoLevel)

The mock server has Graylog's resources such as roles.

  server.Roles[admin.Name] = *admin

By default the mock server's data are temporary, but
you can persist the data to the json file by setting the file path to server.DataPath.

  server.DataPath = "data.json"
  // at first load data
  if err := server.Load() {
  	server.Logger.Fatal(err)
  }
  server.Start()

When the mock server's data are changed by a http request and server.DataPath is set, it's data are written to the file automatically for persistence.

Validation:

  role := &graylog.Role{}
  if err := graylog.CreateValidator.Struct(role); err != nil {
  	return "", nil, err
  }
*/
package graylog
