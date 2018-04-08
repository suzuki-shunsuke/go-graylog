package client_test

import (
	"fmt"
	"log"

	"github.com/suzuki-shunsuke/go-graylog/client"
	"github.com/suzuki-shunsuke/go-graylog/mockserver"
)

func ExampleClient() {
	server, err := mockserver.NewServer("", nil)
	if err != nil {
		log.Fatal(err)
	}
	server.Start()
	defer server.Close()

	cl, err := client.NewClient(server.Endpoint(), "admin", "admin")
	if err != nil {
		log.Fatal(err)
	}

	role, ei, err := cl.GetRole("Admin")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ei.Response.StatusCode)
	fmt.Println(role.Name)
	// Output:
	// 200
	// Admin
}
