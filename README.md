# qrcodefyi-go

[![Go Reference](https://pkg.go.dev/badge/github.com/fyipedia/qrcodefyi-go.svg)](https://pkg.go.dev/github.com/fyipedia/qrcodefyi-go)

Go client for the [QRCodeFYI](https://qrcodefyi.com) API. Look up QR code types, versions, components, encoding modes, and standards. Zero dependencies — stdlib only.

## Install

```bash
go get github.com/fyipedia/qrcodefyi-go
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"

    qrcodefyi "github.com/fyipedia/qrcodefyi-go"
)

func main() {
    client := qrcodefyi.NewClient()

    result, err := client.Search("micro")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Found %d results for 'micro'\n", result.Total)

    qrType, err := client.QRType("model-2")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s: %s\n", qrType.Name, qrType.Description)
}
```

## API Methods

| Method | Description |
|--------|-------------|
| `Search(query)` | Search types, standards, and glossary |
| `QRType(slug)` | Get QR code type details |
| `Version(version)` | Get QR code version details |
| `Component(slug)` | Get component details |
| `Encoding(slug)` | Get encoding mode details |
| `Standard(slug)` | Get standard details |
| `UseCase(slug)` | Get use case details |
| `GlossaryTerm(slug)` | Get glossary term definition |
| `Compare(slugA, slugB)` | Compare two QR code types |
| `Random()` | Get a random QR code type |

## REST API

Free, no authentication required. Base URL: `https://qrcodefyi.com/api`

```bash
curl https://qrcodefyi.com/api/search/?q=micro
curl https://qrcodefyi.com/api/type/model-2/
curl https://qrcodefyi.com/api/random/
```

Full OpenAPI spec: https://qrcodefyi.com/api/openapi.json

## Also Available

| Language | Package | Install |
|----------|---------|---------|
| Python | [qrcodefyi](https://pypi.org/project/qrcodefyi/) | `pip install qrcodefyi` |
| TypeScript | [qrcodefyi](https://www.npmjs.com/package/qrcodefyi) | `npm install qrcodefyi` |
| Go | [qrcodefyi-go](https://pkg.go.dev/github.com/fyipedia/qrcodefyi-go) | `go get github.com/fyipedia/qrcodefyi-go` |
| Rust | [qrcodefyi](https://crates.io/crates/qrcodefyi) | `cargo add qrcodefyi` |
| Ruby | [qrcodefyi](https://rubygems.org/gems/qrcodefyi) | `gem install qrcodefyi` |

## Code FYI Family

| Site | Domain | Focus |
|------|--------|-------|
| BarcodeFYI | [barcodefyi.com](https://barcodefyi.com) | Barcode symbologies and standards |
| QRCodeFYI | [qrcodefyi.com](https://qrcodefyi.com) | QR code types and encoding |
| NFCFYI | [nfcfyi.com](https://nfcfyi.com) | NFC chips and standards |
| BLEFYI | [blefyi.com](https://blefyi.com) | Bluetooth Low Energy profiles |
| RFIDFYI | [rfidfyi.com](https://rfidfyi.com) | RFID tags and frequencies |
| SmartCardFYI | [smartcardfyi.com](https://smartcardfyi.com) | Smart card platforms and standards |

## License

MIT
