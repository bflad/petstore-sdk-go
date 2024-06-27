# User
(*User*)

## Overview

Operations about user

### Available Operations

* [CreateUser](#createuser) - Create user
* [CreateUsersWithListInput](#createuserswithlistinput) - Creates list of users with given input array
* [LoginUser](#loginuser) - Logs user into the system
* [LogoutUser](#logoutuser) - Logs out current logged in user session
* [GetUserByName](#getuserbyname) - Get user by user name
* [UpdateUser](#updateuser) - Update user
* [DeleteUser](#deleteuser) - Delete user

## CreateUser

This can only be done by the logged in user.

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
    var request *components.User = &components.User{
        ID: petstoresdk.Int64(10),
        Username: petstoresdk.String("theUser"),
        FirstName: petstoresdk.String("John"),
        LastName: petstoresdk.String("James"),
        Email: petstoresdk.String("john@email.com"),
        Password: petstoresdk.String("12345"),
        Phone: petstoresdk.String("12345"),
        UserStatus: petstoresdk.Int(1),
    }
    ctx := context.Background()
    res, err := s.User.CreateUser(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [components.User](../../models/components/user.md)    | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.CreateUserResponse](../../models/operations/createuserresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateUsersWithListInput

Creates list of users with given input array

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
    var request []components.User = []components.User{
        components.User{
            ID: petstoresdk.Int64(10),
            Username: petstoresdk.String("theUser"),
            FirstName: petstoresdk.String("John"),
            LastName: petstoresdk.String("James"),
            Email: petstoresdk.String("john@email.com"),
            Password: petstoresdk.String("12345"),
            Phone: petstoresdk.String("12345"),
            UserStatus: petstoresdk.Int(1),
        },
    }
    ctx := context.Background()
    res, err := s.User.CreateUsersWithListInput(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [[]components.User](../../.md)                        | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.CreateUsersWithListInputResponse](../../models/operations/createuserswithlistinputresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## LoginUser

Logs user into the system

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
    var username *string = petstoresdk.String("<value>")

    var password *string = petstoresdk.String("<value>")
    ctx := context.Background()
    res, err := s.User.LoginUser(ctx, username, password)
    if err != nil {
        log.Fatal(err)
    }
    if res.String != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `username`                                            | **string*                                             | :heavy_minus_sign:                                    | The user name for login                               |
| `password`                                            | **string*                                             | :heavy_minus_sign:                                    | The password for login in clear text                  |


### Response

**[*operations.LoginUserResponse](../../models/operations/loginuserresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.APIErrorInvalidInput | 400                            | application/json               |
| sdkerrors.APIErrorUnauthorized | 401                            | application/json               |
| sdkerrors.APIErrorNotFound     | 404                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## LogoutUser

Logs out current logged in user session

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
    res, err := s.User.LogoutUser(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.LogoutUserResponse](../../models/operations/logoutuserresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetUserByName

Get user by user name

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
    var username string = "<value>"
    ctx := context.Background()
    res, err := s.User.GetUserByName(ctx, username)
    if err != nil {
        log.Fatal(err)
    }
    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `username`                                                 | *string*                                                   | :heavy_check_mark:                                         | The name that needs to be fetched. Use user1 for testing.  |


### Response

**[*operations.GetUserByNameResponse](../../models/operations/getuserbynameresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.APIErrorInvalidInput | 400                            | application/json               |
| sdkerrors.APIErrorUnauthorized | 401                            | application/json               |
| sdkerrors.APIErrorNotFound     | 404                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## UpdateUser

This can only be done by the logged in user.

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
    var username string = "<value>"

    var user *components.User = &components.User{
        ID: petstoresdk.Int64(10),
        Username: petstoresdk.String("theUser"),
        FirstName: petstoresdk.String("John"),
        LastName: petstoresdk.String("James"),
        Email: petstoresdk.String("john@email.com"),
        Password: petstoresdk.String("12345"),
        Phone: petstoresdk.String("12345"),
        UserStatus: petstoresdk.Int(1),
    }
    ctx := context.Background()
    res, err := s.User.UpdateUser(ctx, username, user)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `username`                                            | *string*                                              | :heavy_check_mark:                                    | name that needs to be updated                         |
| `user`                                                | [*components.User](../../models/components/user.md)   | :heavy_minus_sign:                                    | Update an existent user in the store                  |


### Response

**[*operations.UpdateUserResponse](../../models/operations/updateuserresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteUser

This can only be done by the logged in user.

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
    var username string = "<value>"
    ctx := context.Background()
    res, err := s.User.DeleteUser(ctx, username)
    if err != nil {
        log.Fatal(err)
    }
    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `username`                                            | *string*                                              | :heavy_check_mark:                                    | The name that needs to be deleted                     |


### Response

**[*operations.DeleteUserResponse](../../models/operations/deleteuserresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.APIErrorInvalidInput | 400                            | application/json               |
| sdkerrors.APIErrorUnauthorized | 401                            | application/json               |
| sdkerrors.APIErrorNotFound     | 404                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |
