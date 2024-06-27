<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	petstoresdk "github.com/bflad/petstore-sdk"
	"github.com/bflad/petstore-sdk/models/components"
	"log"
)

func main() {
	s := petstoresdk.New(
		petstoresdk.WithSecurity("<YOUR_API_KEY_HERE>"),
	)
	request := components.Pet{
		ID:   petstoresdk.Int64(10),
		Name: "doggie",
		Category: &components.Category{
			ID:   petstoresdk.Int64(1),
			Name: petstoresdk.String("Dogs"),
		},
		PhotoUrls: []string{
			"<value>",
		},
	}
	ctx := context.Background()
	res, err := s.Pet.UpdatePet(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res.Pet != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->