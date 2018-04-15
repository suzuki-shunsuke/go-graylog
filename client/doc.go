/*
Package client provides Graylog API client.

Construct a new Graylog client.

  cl, err := client.NewClient("http://localhost:9000/api", "admin", "password")

Of course, you can use the access token and session token instead of password.

  cl, err := NewClient("http://localhost:9000/api", "htgi84ut7jpivsrcldd6l4lmcigvfauldm99ofcb4hsfcvdgsru", "token")

And you can call various Graylog REST APIs as client methods.
For example, create a role.

	role := &graylog.Role{Name: "foo", Permissions: set.NewStrSet("*")}
	_, err := cl.CreateRole(role)

In addition the conventional error object, client API returns "ErrorInfo" object. This object has http.Response object and Graylog API's error message.

  ei, err := cl.CreateRole(role)
  if err != nil {
  	if ei == nil {
  		log.Fatal(err)
  	}
  	log.Fatalf("%s, %s, %s, %s", err, ei.Type, ei.Message, ei.Response.StatusCode)
  }

Note that the Response's Body has been closed.
*/
package client
