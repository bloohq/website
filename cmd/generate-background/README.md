# Background Image Generator

Interactive CLI tool for generating deterministic abstract background images using Blue's brand colors.

## Usage

```bash
go run main.go
```

Enter any text when prompted. The tool generates a 1920x1080 PNG with consistent abstract art based on your input.

## Features

- **Deterministic**: Same input always produces the same image
- **HD Resolution**: 1920x1080 for product screenshots and backgrounds
- **Brand Colors**: Uses Blue's official color palette
- **Abstract Art**: 3-6 gradient blobs with Gaussian blur
- **Local Output**: Saves as `{input}.png` in current directory

## Examples

```bash
$ go run main.go
Enter text for background: product-demo
✓ Generated: product-demo.png (1920x1080)
```

## Algorithm

Uses SHA256 hash of input as seed for:
- Random color selection from palette
- Elliptical gradient blob placement
- Base gradient angle
- 12px Gaussian blur for smooth finish