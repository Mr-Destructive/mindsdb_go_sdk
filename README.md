## Mindsdb Golang SDK

An unoffical golang sdk for [Mindsdb](https://mindsdb.com/). 


### Installation

Install the golang package from GitHub.

```bash
go get github.com/Mr-Destructive/mindsdb_go_sdk
```

### Usage

#### Basic Authentication

- Create a .env file for storing your credentials to authenticate to the mindsdb server.

```env
email=abc@def.com
password=secret_password
```

- Access the credentials from .env file using the helper functions `LoadEnvFromFile` with parameter as the key name in the file. Here we have `email` and `password`.

- Use `connectors.Login` method to log in into the server. 

```go
package main

import (
	"fmt"
	"os"

	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb"
	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb/connectors"
)

func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// read email and password from the .env file
	err := mindsdb.LoadEnvFromFile(".env")
	PanicError(err)
	email := os.Getenv("email")
	password := os.Getenv("password")

	// Login in with an email and password
	api, err := connectors.Login(email, password)
	PanicError(err)
}
```

- Tha `api` variable will be used to access the sdk methods.

- Refer the [examples](https://github.com/Mr-Destructive/mindsdb_go_sdk/tree/main/examples) for further reference, till all the methods in the sdk are functional and tested properly.

### References
- Also this sdk is refered from the [mindsdb-python-sdk](https://github.com/mindsdb/mindsdb_python_sdk) for the inspiration and development.
- [Mindsdb Docs](https://docs.mindsdb.com/what-is-mindsdb) for creating REST API adapters.
