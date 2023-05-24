<div align="center">
    <img src="https://github.com/speakeasy-sdks/cribl-demo-go/assets/68016351/3c85f178-5ab2-4679-b0a7-c31ecdcce367" width="350px">
    <h1>Cribl Go SDK</h1>
   <p></p>
   <a href="https://docs.cribl.io/api/"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=000&style=for-the-badge" /></a>
   <a href="https://github.com/speakeasy-sdks/cribl-demo-go/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/crible-demo-go/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
  <a href="https://github.com/speakeasy-sdks/crible-demo-go/releases"><img src="https://img.shields.io/github/v/release/speakeasy-sdks/cribl-demo-go?sort=semver&style=for-the-badge" /></a>
</div>
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


### [Certificates](docs/certificates/README.md)

* [CreateCertificate](docs/certificates/README.md#createcertificate) - Create Certificate
* [ListCertificates](docs/certificates/README.md#listcertificates) - Get a list of Certificate objects
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
