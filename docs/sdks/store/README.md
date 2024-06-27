# Store
(*Store*)

## Overview

Access to Petstore orders

Find out more about our store
<http://swagger.io>
### Available Operations

* [GetInventory](#getinventory) - Returns pet inventories by status
* [PlaceOrder](#placeorder) - Place an order for a pet
* [GetOrderByID](#getorderbyid) - Find purchase order by ID
* [DeleteOrder](#deleteorder) - Delete purchase order by ID

## GetInventory

Returns a map of status codes to quantities

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

    ctx := context.Background()
    res, err := s.Store.GetInventory(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetInventoryResponse](../../models/operations/getinventoryresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.APIErrorUnauthorized | 401                            | application/json               |
| sdkerrors.APIErrorNotFound     | 404                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## PlaceOrder

Place a new order in the store

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
    var request *components.Order = &components.Order{
        ID: petstoresdk.Int64(10),
        PetID: petstoresdk.Int64(198772),
        Quantity: petstoresdk.Int(7),
        Status: components.OrderStatusApproved.ToPointer(),
    }
    ctx := context.Background()
    res, err := s.Store.PlaceOrder(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Order != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [components.Order](../../models/components/order.md)  | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PlaceOrderResponse](../../models/operations/placeorderresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.APIErrorUnauthorized | 401                            | application/json               |
| sdkerrors.APIErrorNotFound     | 404                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## GetOrderByID

For valid response try integer IDs with value <= 5 or > 10. Other values will generate exceptions.

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
    var orderID int64 = 614993
    ctx := context.Background()
    res, err := s.Store.GetOrderByID(ctx, orderID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Order != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `orderID`                                             | *int64*                                               | :heavy_check_mark:                                    | ID of order that needs to be fetched                  |


### Response

**[*operations.GetOrderByIDResponse](../../models/operations/getorderbyidresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.APIErrorInvalidInput | 400                            | application/json               |
| sdkerrors.APIErrorUnauthorized | 401                            | application/json               |
| sdkerrors.APIErrorNotFound     | 404                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## DeleteOrder

For valid response try integer IDs with value < 1000. Anything above 1000 or nonintegers will generate API errors

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
    var orderID int64 = 127902
    ctx := context.Background()
    res, err := s.Store.DeleteOrder(ctx, orderID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Order != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `orderID`                                             | *int64*                                               | :heavy_check_mark:                                    | ID of the order that needs to be deleted              |


### Response

**[*operations.DeleteOrderResponse](../../models/operations/deleteorderresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.APIErrorInvalidInput | 400                            | application/json               |
| sdkerrors.APIErrorUnauthorized | 401                            | application/json               |
| sdkerrors.APIErrorNotFound     | 404                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |
