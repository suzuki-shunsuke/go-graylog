package client_test

import (
	"context"
	"fmt"
	"log"

	"github.com/suzuki-shunsuke/go-graylog/v8/client"
	"github.com/suzuki-shunsuke/graylog-mock-server/mockserver"
)

func ExampleClient() {
	ctx := context.Background()
	// Create a mock server
	server, err := mockserver.NewServer("", nil)
	if err != nil {
		log.Fatal(err)
	}
	// Start a server
	server.Start()
	defer server.Close()

	// Create a client
	cl, err := client.NewClient(server.Endpoint(), "admin", "admin")
	if err != nil {
		log.Fatal(err)
	}

	// get a role "Admin"
	// ei.Response.Body is closed
	role, ei, err := cl.GetRole(ctx, "Admin")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ei.Response.StatusCode)
	fmt.Println(role.Name)
	// Output:
	// 200
	// Admin
}
