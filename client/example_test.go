package client_test

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/suzuki-shunsuke/flute/flute"

	"github.com/suzuki-shunsuke/go-graylog/v10/client"
)

func ExampleClient() {
	ctx := context.Background()

	// Create a client
	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	if err != nil {
		log.Fatal(err)
	}

	setExampleMock(cl)

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

func setExampleMock(cl *client.Client) {
	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			Services: []flute.Service{
				{
					Endpoint: "http://example.com",
					Routes: []flute.Route{
						{
							Tester: &flute.Tester{
								Method:       "GET",
								Path:         "/api/roles/Admin",
								PartOfHeader: getTestHeader(),
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 200,
								},
								BodyString: `{
  "name": "Admin",
  "description": "Grants all permissions for Graylog administrators (built-in)",
  "permissions": [
    "*"
  ],
  "read_only": true
}`,
							},
						},
					},
				},
			},
		},
	})
}
