# Otto-Client

Repository used to stored my client to interact with my API. Purpose is to used this client in different project that need to interact with the API

Struture of the client look like https://github.com/google/go-github project, liked there structure and try to follow it!

# Installations

_Otto-client_ is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/Vico1993/Otto-client
```

will resolve and add the package to the current development module, along with its dependencies.

Alternatively the same can be achieved if you use import in a package:

```go
import "github.com/Vico1993/Otto-client/otto"
```

and run go get without parameters.

Finally, to use the top-of-trunk version of this repo, use the following command:

```bash
go get github.com/Vico1993/Otto-client@main
```

# Usages

Construct a new Otto client then use the various services on the client to access different parts of the Otto API. For example:

```go
// Create a new client with http.Client and otto api base url
client := otto.NewClient(nil, "https://otto-api.com/api")

tags := client.Tag.List("chatId", "threadId")
```

More example in the [example](/example/) folder

## Running Tests

To run tests, use the following command:

```sh
make test
```

## Running Lint

To run lint, use the following command:

```sh
make lint
```

# Contributing

Contributions are welcome! Please see the [CONTRIBUTING.md](./CONTRIBUTING.md) file for more information.

# License

This project is licensed under the [LICENSE](./LICENSE) file in the root directory of this repository.
