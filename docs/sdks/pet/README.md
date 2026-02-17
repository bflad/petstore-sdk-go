# Pet

## Overview

Everything about your Pets

Find out more
<http://swagger.io>

### Available Operations

* [UpdatePet](#updatepet) - Update an existing pet
* [AddPet](#addpet) - Add a new pet to the store
* [FindPetsByStatus](#findpetsbystatus) - Finds Pets by status
* [FindPetsByTags](#findpetsbytags) - Finds Pets by tags
* [GetPetByID](#getpetbyid) - Find pet by ID
* [DeletePet](#deletepet) - Deletes a pet
* [UploadFile](#uploadfile) - uploads an image

## UpdatePet

Update an existing pet by Id

### Example Usage

<!-- UsageSnippet language="go" operationID="updatePet" method="put" path="/pet" -->
```go
package main

import(
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
        ID: petstoresdk.Pointer[int64](10),
        Name: "doggie",
        Category: &components.Category{
            ID: petstoresdk.Pointer[int64](1),
            Name: petstoresdk.Pointer("Dogs"),
        },
        PhotoUrls: []string{
            "<value 1>",
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

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [components.Pet](../../models/components/pet.md)         | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.UpdatePetResponse](../../models/operations/updatepetresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.APIErrorInvalidInput | 400                            | application/json               |
| sdkerrors.APIErrorUnauthorized | 401                            | application/json               |
| sdkerrors.APIErrorNotFound     | 404                            | application/json               |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## AddPet

Add a new pet to the store

### Example Usage: fido

<!-- UsageSnippet language="go" operationID="addPet" method="post" path="/pet" example="fido" -->
```go
package main

import(
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

    res, err := s.Pet.AddPet(ctx, components.Pet{
        Name: "Fido",
        PhotoUrls: []string{
            "https://www.example.com/fido.jpg",
        },
        Status: components.StatusAvailable.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Pet != nil {
        // handle response
    }
}
```
### Example Usage: rover

<!-- UsageSnippet language="go" operationID="addPet" method="post" path="/pet" example="rover" -->
```go
package main

import(
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

    res, err := s.Pet.AddPet(ctx, components.Pet{
        Name: "Rover",
        PhotoUrls: []string{
            "https://www.example.com/rover1.jpg",
            "https://www.example.com/rover2.jpg",
        },
        Status: components.StatusPending.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Pet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [components.Pet](../../models/components/pet.md)         | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.AddPetResponse](../../models/operations/addpetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## FindPetsByStatus

Multiple status values can be provided with comma separated strings

### Example Usage

<!-- UsageSnippet language="go" operationID="findPetsByStatus" method="get" path="/pet/findByStatus" -->
```go
package main

import(
	"context"
	petstoresdk "github.com/bflad/petstore-sdk"
	"github.com/bflad/petstore-sdk/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := petstoresdk.New(
        petstoresdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Pet.FindPetsByStatus(ctx, operations.StatusAvailable.ToPointer())
    if err != nil {
        log.Fatal(err)
    }
    if res.Pets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `status`                                                 | [*operations.Status](../../models/operations/status.md)  | :heavy_minus_sign:                                       | Status values that need to be considered for filter      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.FindPetsByStatusResponse](../../models/operations/findpetsbystatusresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.APIErrorInvalidInput | 400                            | application/json               |
| sdkerrors.APIErrorUnauthorized | 401                            | application/json               |
| sdkerrors.APIErrorNotFound     | 404                            | application/json               |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## FindPetsByTags

Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.

### Example Usage

<!-- UsageSnippet language="go" operationID="findPetsByTags" method="get" path="/pet/findByTags" -->
```go
package main

import(
	"context"
	petstoresdk "github.com/bflad/petstore-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := petstoresdk.New(
        petstoresdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Pet.FindPetsByTags(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Pets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `tags`                                                   | []*string*                                               | :heavy_minus_sign:                                       | Tags to filter by                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.FindPetsByTagsResponse](../../models/operations/findpetsbytagsresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.APIErrorInvalidInput | 400                            | application/json               |
| sdkerrors.APIErrorUnauthorized | 401                            | application/json               |
| sdkerrors.APIErrorNotFound     | 404                            | application/json               |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## GetPetByID

Returns a single pet

### Example Usage

<!-- UsageSnippet language="go" operationID="getPetById" method="get" path="/pet/{petId}" -->
```go
package main

import(
	"context"
	petstoresdk "github.com/bflad/petstore-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := petstoresdk.New(
        petstoresdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Pet.GetPetByID(ctx, 311674)
    if err != nil {
        log.Fatal(err)
    }
    if res.Pet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `petID`                                                  | *int64*                                                  | :heavy_check_mark:                                       | ID of pet to return                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetPetByIDResponse](../../models/operations/getpetbyidresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.APIErrorInvalidInput | 400                            | application/json               |
| sdkerrors.APIErrorUnauthorized | 401                            | application/json               |
| sdkerrors.APIErrorNotFound     | 404                            | application/json               |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## DeletePet

Deletes a pet

### Example Usage

<!-- UsageSnippet language="go" operationID="deletePet" method="delete" path="/pet/{petId}" -->
```go
package main

import(
	"context"
	petstoresdk "github.com/bflad/petstore-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := petstoresdk.New(
        petstoresdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Pet.DeletePet(ctx, 818965, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Pet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `petID`                                                  | *int64*                                                  | :heavy_check_mark:                                       | Pet id to delete                                         |
| `apiKey`                                                 | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeletePetResponse](../../models/operations/deletepetresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.APIErrorInvalidInput | 400                            | application/json               |
| sdkerrors.APIErrorUnauthorized | 401                            | application/json               |
| sdkerrors.APIErrorNotFound     | 404                            | application/json               |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## UploadFile

uploads an image

### Example Usage

<!-- UsageSnippet language="go" operationID="uploadFile" method="post" path="/pet/{petId}/uploadImage" -->
```go
package main

import(
	"context"
	petstoresdk "github.com/bflad/petstore-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := petstoresdk.New(
        petstoresdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Pet.UploadFile(ctx, 150516, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.APIResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `petID`                                                  | *int64*                                                  | :heavy_check_mark:                                       | ID of pet to update                                      |
| `additionalMetadata`                                     | **string*                                                | :heavy_minus_sign:                                       | Additional Metadata                                      |
| `requestBody`                                            | **any*                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.UploadFileResponse](../../models/operations/uploadfileresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |