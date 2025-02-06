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
	ctx := context.Background()

	s := petstoresdk.New(
		petstoresdk.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Pet.UpdatePet(ctx, components.Pet{
		ID:   petstoresdk.Int64(10),
		Name: "doggie",
		Category: &components.Category{
			ID:   petstoresdk.Int64(1),
			Name: petstoresdk.String("Dogs"),
		},
		PhotoUrls: []string{
			"<value>",
			"<value>",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Pet != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->