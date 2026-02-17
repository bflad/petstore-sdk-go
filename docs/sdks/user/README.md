# User

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

<!-- UsageSnippet language="go" operationID="createUser" method="post" path="/user" -->
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

    res, err := s.User.CreateUser(ctx, &components.User{
        ID: petstoresdk.Pointer[int64](10),
        Username: petstoresdk.Pointer("theUser"),
        FirstName: petstoresdk.Pointer("John"),
        LastName: petstoresdk.Pointer("James"),
        Email: petstoresdk.Pointer("john@email.com"),
        Password: petstoresdk.Pointer("12345"),
        Phone: petstoresdk.Pointer("12345"),
        UserStatus: petstoresdk.Pointer[int](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [components.User](../../models/components/user.md)       | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CreateUserResponse](../../models/operations/createuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateUsersWithListInput

Creates list of users with given input array

### Example Usage

<!-- UsageSnippet language="go" operationID="createUsersWithListInput" method="post" path="/user/createWithList" -->
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

    res, err := s.User.CreateUsersWithListInput(ctx, []components.User{
        components.User{
            ID: petstoresdk.Pointer[int64](10),
            Username: petstoresdk.Pointer("theUser"),
            FirstName: petstoresdk.Pointer("John"),
            LastName: petstoresdk.Pointer("James"),
            Email: petstoresdk.Pointer("john@email.com"),
            Password: petstoresdk.Pointer("12345"),
            Phone: petstoresdk.Pointer("12345"),
            UserStatus: petstoresdk.Pointer[int](1),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [[]components.User](../../.md)                           | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CreateUsersWithListInputResponse](../../models/operations/createuserswithlistinputresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## LoginUser

Logs user into the system

### Example Usage

<!-- UsageSnippet language="go" operationID="loginUser" method="get" path="/user/login" -->
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

    res, err := s.User.LoginUser(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.String != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `username`                                               | **string*                                                | :heavy_minus_sign:                                       | The user name for login                                  |
| `password`                                               | **string*                                                | :heavy_minus_sign:                                       | The password for login in clear text                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.LoginUserResponse](../../models/operations/loginuserresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.APIErrorInvalidInput | 400                            | application/json               |
| sdkerrors.APIErrorUnauthorized | 401                            | application/json               |
| sdkerrors.APIErrorNotFound     | 404                            | application/json               |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## LogoutUser

Logs out current logged in user session

### Example Usage

<!-- UsageSnippet language="go" operationID="logoutUser" method="get" path="/user/logout" -->
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

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.LogoutUserResponse](../../models/operations/logoutuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetUserByName

Get user by user name

### Example Usage

<!-- UsageSnippet language="go" operationID="getUserByName" method="get" path="/user/{username}" -->
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

    res, err := s.User.GetUserByName(ctx, "Edyth10")
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
| `opts`                                                     | [][operations.Option](../../models/operations/option.md)   | :heavy_minus_sign:                                         | The options for this request.                              |

### Response

**[*operations.GetUserByNameResponse](../../models/operations/getuserbynameresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.APIErrorInvalidInput | 400                            | application/json               |
| sdkerrors.APIErrorUnauthorized | 401                            | application/json               |
| sdkerrors.APIErrorNotFound     | 404                            | application/json               |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## UpdateUser

This can only be done by the logged in user.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateUser" method="put" path="/user/{username}" -->
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

    res, err := s.User.UpdateUser(ctx, "Alison.Cassin", &components.User{
        ID: petstoresdk.Pointer[int64](10),
        Username: petstoresdk.Pointer("theUser"),
        FirstName: petstoresdk.Pointer("John"),
        LastName: petstoresdk.Pointer("James"),
        Email: petstoresdk.Pointer("john@email.com"),
        Password: petstoresdk.Pointer("12345"),
        Phone: petstoresdk.Pointer("12345"),
        UserStatus: petstoresdk.Pointer[int](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `username`                                               | *string*                                                 | :heavy_check_mark:                                       | name that needs to be updated                            |
| `user`                                                   | [*components.User](../../models/components/user.md)      | :heavy_minus_sign:                                       | Update an existent user in the store                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.UpdateUserResponse](../../models/operations/updateuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteUser

This can only be done by the logged in user.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteUser" method="delete" path="/user/{username}" -->
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

    res, err := s.User.DeleteUser(ctx, "Rita_Schuppe")
    if err != nil {
        log.Fatal(err)
    }
    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `username`                                               | *string*                                                 | :heavy_check_mark:                                       | The name that needs to be deleted                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteUserResponse](../../models/operations/deleteuserresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.APIErrorInvalidInput | 400                            | application/json               |
| sdkerrors.APIErrorUnauthorized | 401                            | application/json               |
| sdkerrors.APIErrorNotFound     | 404                            | application/json               |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |