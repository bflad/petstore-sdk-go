# Pet
(*Pet*)

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

```go
package main

import(
	petstoresdk "github.com/bflad/petstore-sdk"
	"github.com/bflad/petstore-sdk/models/components"
	"context"
	"log"
)

func main() {
    s := petstoresdk.New(
        petstoresdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )
    request := components.Pet{
        ID: petstoresdk.Int64(10),
        Name: "doggie",
        Category: &components.Category{
            ID: petstoresdk.Int64(1),
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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [components.Pet](../../models/components/pet.md)      | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.UpdatePetResponse](../../models/operations/updatepetresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.APIErrorInvalidInput | 400                            | application/json               |
| sdkerrors.APIErrorUnauthorized | 401                            | application/json               |
| sdkerrors.APIErrorNotFound     | 404                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## AddPet

Add a new pet to the store

### Example Usage

```go
package main

import(
	petstoresdk "github.com/bflad/petstore-sdk"
	"github.com/bflad/petstore-sdk/models/components"
	"context"
	"log"
)

func main() {
    s := petstoresdk.New(
        petstoresdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )
    request := components.Pet{
        ID: petstoresdk.Int64(10),
        Name: "doggie",
        Category: &components.Category{
            ID: petstoresdk.Int64(1),
            Name: petstoresdk.String("Dogs"),
        },
        PhotoUrls: []string{
            "<value>",
        },
    }
    ctx := context.Background()
    res, err := s.Pet.AddPet(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Pet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [components.Pet](../../models/components/pet.md)      | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.AddPetResponse](../../models/operations/addpetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## FindPetsByStatus

Multiple status values can be provided with comma separated strings

### Example Usage

```go
package main

import(
	petstoresdk "github.com/bflad/petstore-sdk"
	"github.com/bflad/petstore-sdk/models/operations"
	"context"
	"log"
)

func main() {
    s := petstoresdk.New(
        petstoresdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )
    var status *operations.Status = operations.StatusAvailable.ToPointer()
    ctx := context.Background()
    res, err := s.Pet.FindPetsByStatus(ctx, status)
    if err != nil {
        log.Fatal(err)
    }
    if res.Pets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                               | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `ctx`                                                   | [context.Context](https://pkg.go.dev/context#Context)   | :heavy_check_mark:                                      | The context to use for the request.                     |
| `status`                                                | [*operations.Status](../../models/operations/status.md) | :heavy_minus_sign:                                      | Status values that need to be considered for filter     |


### Response

**[*operations.FindPetsByStatusResponse](../../models/operations/findpetsbystatusresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.APIErrorInvalidInput | 400                            | application/json               |
| sdkerrors.APIErrorUnauthorized | 401                            | application/json               |
| sdkerrors.APIErrorNotFound     | 404                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## FindPetsByTags

Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.

### Example Usage

```go
package main

import(
	petstoresdk "github.com/bflad/petstore-sdk"
	"context"
	"log"
)

func main() {
    s := petstoresdk.New(
        petstoresdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )
    var tags []string = []string{
        "<value>",
    }
    ctx := context.Background()
    res, err := s.Pet.FindPetsByTags(ctx, tags)
    if err != nil {
        log.Fatal(err)
    }
    if res.Pets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `tags`                                                | []*string*                                            | :heavy_minus_sign:                                    | Tags to filter by                                     |


### Response

**[*operations.FindPetsByTagsResponse](../../models/operations/findpetsbytagsresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.APIErrorInvalidInput | 400                            | application/json               |
| sdkerrors.APIErrorUnauthorized | 401                            | application/json               |
| sdkerrors.APIErrorNotFound     | 404                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## GetPetByID

Returns a single pet

### Example Usage

```go
package main

import(
	petstoresdk "github.com/bflad/petstore-sdk"
	"context"
	"log"
)

func main() {
    s := petstoresdk.New(
        petstoresdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )
    var petID int64 = 504151
    ctx := context.Background()
    res, err := s.Pet.GetPetByID(ctx, petID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Pet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `petID`                                               | *int64*                                               | :heavy_check_mark:                                    | ID of pet to return                                   |


### Response

**[*operations.GetPetByIDResponse](../../models/operations/getpetbyidresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.APIErrorInvalidInput | 400                            | application/json               |
| sdkerrors.APIErrorUnauthorized | 401                            | application/json               |
| sdkerrors.APIErrorNotFound     | 404                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## DeletePet

Deletes a pet

### Example Usage

```go
package main

import(
	petstoresdk "github.com/bflad/petstore-sdk"
	"context"
	"log"
)

func main() {
    s := petstoresdk.New(
        petstoresdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )
    var petID int64 = 441876

    var apiKey *string = petstoresdk.String("<value>")
    ctx := context.Background()
    res, err := s.Pet.DeletePet(ctx, petID, apiKey)
    if err != nil {
        log.Fatal(err)
    }
    if res.Pet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `petID`                                               | *int64*                                               | :heavy_check_mark:                                    | Pet id to delete                                      |
| `apiKey`                                              | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   |


### Response

**[*operations.DeletePetResponse](../../models/operations/deletepetresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.APIErrorInvalidInput | 400                            | application/json               |
| sdkerrors.APIErrorUnauthorized | 401                            | application/json               |
| sdkerrors.APIErrorNotFound     | 404                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## UploadFile

uploads an image

### Example Usage

```go
package main

import(
	petstoresdk "github.com/bflad/petstore-sdk"
	"context"
	"log"
)

func main() {
    s := petstoresdk.New(
        petstoresdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )
    var petID int64 = 565380

    var additionalMetadata *string = petstoresdk.String("<value>")

    var requestBody []byte = []byte("0x7cca7F47Dd")
    ctx := context.Background()
    res, err := s.Pet.UploadFile(ctx, petID, additionalMetadata, requestBody)
    if err != nil {
        log.Fatal(err)
    }
    if res.APIResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `petID`                                               | *int64*                                               | :heavy_check_mark:                                    | ID of pet to update                                   |
| `additionalMetadata`                                  | **string*                                             | :heavy_minus_sign:                                    | Additional Metadata                                   |
| `requestBody`                                         | *[]byte*                                              | :heavy_minus_sign:                                    | N/A                                                   |


### Response

**[*operations.UploadFileResponse](../../models/operations/uploadfileresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
