# Certificates

## Overview

Certificates

### Available Operations

* [CreateCertificate](#createcertificate) - Create Certificate
* [ListCertificates](#listcertificates) - Get a list of Certificate objects

## CreateCertificate

Create Certificate

### Example Usage

```go
package main

import(
	"context"
	"log"
	"cribl-demo"
	"cribl-demo/pkg/models/shared"
)

func main() {
    s := cribldemo.New()

    ctx := context.Background()
    res, err := s.Certificates.CreateCertificate(ctx, shared.Certificate{
        Ca: cribldemo.String("molestiae"),
        Cert: "quod",
        Description: cribldemo.String("quod"),
        ID: "78ca1ba9-28fc-4816-b42c-b73920592939",
        InUse: []string{
            "hic",
            "saepe",
        },
        Passphrase: cribldemo.String("fuga"),
        PrivKey: "in",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Certificate != nil {
        // handle response
    }
}
```

## ListCertificates

Get a list of Certificate objects

### Example Usage

```go
package main

import(
	"context"
	"log"
	"cribl-demo"
)

func main() {
    s := cribldemo.New()

    ctx := context.Background()
    res, err := s.Certificates.ListCertificates(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListCertificates200ApplicationJSONObject != nil {
        // handle response
    }
}
```
