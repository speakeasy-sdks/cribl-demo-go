# cribl-demo

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/cribl-demo-go
```
<!-- End SDK Installation -->

## SDK Example Usage
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

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Certificates](docs/sdks/certificates/README.md)

* [CreateCertificate](docs/sdks/certificates/README.md#createcertificate) - Create Certificate
* [ListCertificates](docs/sdks/certificates/README.md#listcertificates) - Get a list of Certificate objects
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
