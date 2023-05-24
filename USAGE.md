<!-- Start SDK Example Usage -->
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
        Ca: cribldemo.String("corrupti"),
        Cert: "provident",
        Description: cribldemo.String("distinctio"),
        ID: "d9d8d69a-674e-40f4-a7cc-8796ed151a05",
        InUse: []string{
            "sapiente",
            "quo",
            "odit",
            "at",
        },
        Passphrase: cribldemo.String("at"),
        PrivKey: "maiores",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Certificate != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->