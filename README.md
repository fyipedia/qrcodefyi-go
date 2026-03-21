# qrcodefyi-go

[![Go Reference](https://pkg.go.dev/badge/github.com/fyipedia/qrcodefyi-go.svg)](https://pkg.go.dev/github.com/fyipedia/qrcodefyi-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Go client for the [QRCodeFYI](https://qrcodefyi.com) REST API. QR code types, versions, encoding. Zero external dependencies beyond stdlib.

> **Try the interactive tools at [qrcodefyi.com](https://qrcodefyi.com)** — explore, search, and discover.

## Install

```bash
go get github.com/fyipedia/qrcodefyi-go
```

## Quick Start

```go
package main

import (
    "fmt"
    qrcodefyi "github.com/fyipedia/qrcodefyi-go"
)

func main() {
    client := qrcodefyi.NewClient()

    // Search across all content
    result, err := client.Search("query")
    if err != nil {
        panic(err)
    }
    fmt.Println(result)
}
```

## API Methods

| Method | Description |
|--------|-------------|
| `Comparisons()` | List comparisons |
| `Components()` | List components |
| `EncodingModes()` | List encoding modes |
| `Faqs()` | List faqs |
| `Glossary()` | List glossary |
| `Guides()` | List guides |
| `Recipes()` | List recipes |
| `ScanScenarios()` | List scan scenarios |
| `Standards()` | List standards |
| `Tools()` | List tools |
| `Types()` | List types |
| `UseCases()` | List use cases |
| `Versions()` | List versions |
| `Search(query)` | Search across all content |

## Also Available

| Platform | Package | Link |
|----------|---------|------|
| **Python** | `pip install qrcodefyi` | [PyPI](https://pypi.org/project/qrcodefyi/) |
| **npm** | `npm install qrcodefyi` | [npm](https://www.npmjs.com/package/qrcodefyi) |
| **Go** | `go get github.com/fyipedia/qrcodefyi-go` | [pkg.go.dev](https://pkg.go.dev/github.com/fyipedia/qrcodefyi-go) |
| **Rust** | `cargo add qrcodefyi` | [crates.io](https://crates.io/crates/qrcodefyi) |
| **Ruby** | `gem install qrcodefyi` | [rubygems](https://rubygems.org/gems/qrcodefyi) |

## Tag FYI Family

Part of the [FYIPedia](https://fyipedia.com) open-source developer tools ecosystem.

| Site | Domain | Focus |
|------|--------|-------|
| BarcodeFYI | [barcodefyi.com](https://barcodefyi.com) | Barcode formats, EAN, UPC, ISBN standards |
| BLEFYI | [blefyi.com](https://blefyi.com) | Bluetooth Low Energy, GATT, beacons |
| NFCFYI | [nfcfyi.com](https://nfcfyi.com) | NFC chips, tag types, NDEF, standards |
| **QRCodeFYI** | [qrcodefyi.com](https://qrcodefyi.com) | QR code types, versions, encoding modes |
| RFIDFYI | [rfidfyi.com](https://rfidfyi.com) | RFID tags, frequency bands, standards |
| SmartCardFYI | [smartcardfyi.com](https://smartcardfyi.com) | Smart cards, EMV, APDU, Java Card |

## License

MIT
