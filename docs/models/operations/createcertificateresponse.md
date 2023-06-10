# CreateCertificateResponse


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `Certificate`                                             | [*shared.Certificate](../../models/shared/certificate.md) | :heavy_minus_sign:                                        | Created                                                   |
| `ContentType`                                             | *string*                                                  | :heavy_check_mark:                                        | N/A                                                       |
| `Error`                                                   | [*shared.Error](../../models/shared/error.md)             | :heavy_minus_sign:                                        | Unexpected error                                          |
| `StatusCode`                                              | *int*                                                     | :heavy_check_mark:                                        | N/A                                                       |
| `RawResponse`                                             | [*http.Response](https://pkg.go.dev/net/http#Response)    | :heavy_minus_sign:                                        | N/A                                                       |