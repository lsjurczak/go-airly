go-airly
=======
[![Test](https://github.com/lsjurczak/go-airly/workflows/Test/badge.svg?branch=master)](https://github.com/lsjurczak/go-airly/actions?query=workflow%3ATest)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/lsjurczak/go-airly?)
[![Go Report Card](https://goreportcard.com/badge/github.com/lsjurczak/go-airly)](https://goreportcard.com/report/github.com/lsjurczak/go-airly)

go-airly is a Go client library for the Airly API.

Installation
------------
This package can be installed using:

	go get github.com/lsjurczak/go-airly

Usage
-----

Import the package using:

```go
import "github.com/lsjurczak/go-airly"
```

To use this library you have to get key from https://developer.airly.eu/

Construct a new client and pass apiKey:

```go
client, err := airly.NewClient(nil, "apiKey")
if err != nil {
    log.Fatalf("airly.NewClient: %v", err)
}
```

You can also pass custom HTTP client:
```go
customClient := &http.Client{Timeout: 5 * time.Second}
```

Then use one of the client's services (Installation, Measurement, or Meta) to access the
different Airly API methods.

For example, to get the nearest installation:

```go
opt := airly.NewNearestInstallationOpts(52.2872, 21.1087).
    MaxResults(1).
    MaxDistance(10)
installation, err := client.Installation.Nearest(opt)
if err != nil {
    log.Fatal(err)
}
```

License
-----

This library is distributed under the MIT license found in the [LICENSE](./LICENSE) file.
