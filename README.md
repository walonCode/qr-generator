# qr-generator

A simple command-line tool written in Go that generates a QR code PNG from any URL or text.

## Installation

**Prerequisites:** [Go 1.18+](https://go.dev/dl/)

```bash
git clone https://github.com/your-username/qr-generator.git
cd qr-generator
go build -o qr-generator .
```

Or install directly with `go install`:

```bash
go install github.com/your-username/qr-generator@latest
```

## Usage

```
qr-generator <link> <filename.png>
```

| Argument        | Description                              |
|-----------------|------------------------------------------|
| `<link>`        | The URL or text to encode into the QR code |
| `<filename.png>`| Output file path (must end in `.png`)    |

## Examples

Generate a QR code for a website:

```bash
qr-generator https://example.com output.png
```

Encode plain text:

```bash
qr-generator "Hello, World!" hello.png
```

## Dependencies

- [`github.com/skip2/go-qrcode`](https://github.com/skip2/go-qrcode) — QR code generation library

## License

This project is licensed under the [MIT License](LICENSE).
