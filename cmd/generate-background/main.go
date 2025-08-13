package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

const (
	// HD resolution for backgrounds
	width      = 1920
	height     = 1080
	blurRadius = 12 // Scaled up for larger image
)

// Blue brand color palette
var colorPalette = []color.RGBA{
	{0x00, 0xA0, 0xD2, 0xFF}, // Blue Blue
	{0xF0, 0x6C, 0x9B, 0xFF}, // Cyclamen
	{0xFE, 0xD7, 0x66, 0xFF}, // Mustard
	{0xA9, 0xF0, 0xD1, 0xFF}, // Aquamarine
	{0xF8, 0x66, 0x24, 0xFF}, // Giants Orange
}

type blob struct {
	centerX, centerY float64
	radiusX, radiusY float64
	color            color.RGBA
	angle            float64
}

func main() {
	// Interactive prompt
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text for background: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}

	// Clean input
	input = strings.TrimSpace(input)
	if input == "" {
		fmt.Println("Error: Empty input")
		os.Exit(1)
	}

	// Generate filename
	filename := textToFilename(input)
	outputPath := filename + ".png"

	// Generate image
	fmt.Printf("Generating background for \"%s\"...\n", input)
	if err := generateBackground(input, outputPath); err != nil {
		fmt.Printf("Error generating image: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ Generated: %s (%dx%d)\n", outputPath, width, height)
}

func generateBackground(text, outputPath string) error {
	// Create deterministic seed from text
	hash := sha256.Sum256([]byte(text))
	seed := int64(binary.BigEndian.Uint64(hash[:8]))

	// Create seeded random generator for consistent results
	r := rand.New(rand.NewSource(seed))

	// Create image
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Generate abstract background
	generateAbstractArt(img, r)

	// Apply Gaussian blur
	blurredImg := gaussianBlur(img, blurRadius)

	// Save PNG file
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	if err := png.Encode(file, blurredImg); err != nil {
		return fmt.Errorf("failed to encode PNG: %v", err)
	}

	return nil
}

func generateAbstractArt(img *image.RGBA, r *rand.Rand) {
	// Select 2-4 random colors from the palette
	selectedColors := selectRandomColors(r)

	// Create random gradient centers and directions
	numBlobs := 3 + r.Intn(4) // 3-6 blobs
	blobs := make([]blob, numBlobs)

	for i := 0; i < numBlobs; i++ {
		// Scale blob sizes for HD resolution
		blobs[i] = blob{
			centerX: r.Float64() * float64(width),
			centerY: r.Float64() * float64(height),
			radiusX: 400 + r.Float64()*800, // Scaled up from original
			radiusY: 400 + r.Float64()*800,
			color:   selectedColors[r.Intn(len(selectedColors))],
			angle:   r.Float64() * 2 * math.Pi,
		}
	}

	// Create base gradient
	baseColor1 := selectedColors[0]
	baseColor2 := selectedColors[len(selectedColors)-1]
	gradientAngle := r.Float64() * 2 * math.Pi

	// Fill the image
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Base gradient with random direction
			gradientPos := (float64(x)*math.Cos(gradientAngle) + float64(y)*math.Sin(gradientAngle)) /
				(float64(width)*math.Abs(math.Cos(gradientAngle)) + float64(height)*math.Abs(math.Sin(gradientAngle)))
			gradientPos = math.Max(0, math.Min(1, gradientPos))

			baseColor := interpolateColor(baseColor1, baseColor2, gradientPos)

			// Add blob influences with better blending
			currentColor := baseColor

			for _, b := range blobs {
				// Calculate elliptical distance with rotation
				dx := float64(x) - b.centerX
				dy := float64(y) - b.centerY

				// Rotate coordinates
				rotX := dx*math.Cos(-b.angle) - dy*math.Sin(-b.angle)
				rotY := dx*math.Sin(-b.angle) + dy*math.Cos(-b.angle)

				// Elliptical distance
				distance := math.Sqrt((rotX*rotX)/(b.radiusX*b.radiusX) + (rotY*rotY)/(b.radiusY*b.radiusY))

				// Smooth falloff with gentler curve
				weight := math.Max(0, 1-distance)
				weight = weight * weight * (3 - 2*weight) // Smoothstep
				weight = weight * 0.7                     // Reduce overall influence

				if weight > 0 {
					// Blend each blob one at a time
					currentColor = blendColors(currentColor, b.color, weight)
				}
			}

			finalColor := currentColor

			// Add subtle noise
			noise := (r.Float64() - 0.5) * 0.05
			finalColor = adjustBrightness(finalColor, 1.0+noise)

			img.Set(x, y, finalColor)
		}
	}
}

func selectRandomColors(r *rand.Rand) []color.RGBA {
	// Select 2-4 random colors from the palette
	numColors := 2 + r.Intn(3) // 2, 3, or 4 colors

	// Shuffle the palette
	shuffled := make([]color.RGBA, len(colorPalette))
	copy(shuffled, colorPalette)

	for i := len(shuffled) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}

	// Return the first numColors colors
	return shuffled[:numColors]
}

func interpolateColor(c1, c2 color.RGBA, t float64) color.RGBA {
	return color.RGBA{
		R: uint8(float64(c1.R)*(1-t) + float64(c2.R)*t),
		G: uint8(float64(c1.G)*(1-t) + float64(c2.G)*t),
		B: uint8(float64(c1.B)*(1-t) + float64(c2.B)*t),
		A: 255,
	}
}

func blendColors(c1, c2 color.RGBA, factor float64) color.RGBA {
	factor = math.Max(0, math.Min(1, factor))
	return color.RGBA{
		R: uint8(float64(c1.R)*(1-factor) + float64(c2.R)*factor),
		G: uint8(float64(c1.G)*(1-factor) + float64(c2.G)*factor),
		B: uint8(float64(c1.B)*(1-factor) + float64(c2.B)*factor),
		A: 255,
	}
}

func adjustBrightness(c color.RGBA, factor float64) color.RGBA {
	return color.RGBA{
		R: uint8(math.Min(255, float64(c.R)*factor)),
		G: uint8(math.Min(255, float64(c.G)*factor)),
		B: uint8(math.Min(255, float64(c.B)*factor)),
		A: c.A,
	}
}

func gaussianBlur(src *image.RGBA, radius int) *image.RGBA {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)

	// Generate Gaussian kernel
	kernel := make([]float64, radius*2+1)
	sum := 0.0
	sigma := float64(radius) / 3.0

	for i := 0; i < len(kernel); i++ {
		x := float64(i - radius)
		kernel[i] = math.Exp(-(x*x)/(2*sigma*sigma)) / (sigma * math.Sqrt(2*math.Pi))
		sum += kernel[i]
	}

	// Normalize kernel
	for i := range kernel {
		kernel[i] /= sum
	}

	// Horizontal pass
	temp := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			var r, g, b float64

			for i := 0; i < len(kernel); i++ {
				px := x + i - radius
				if px < bounds.Min.X {
					px = bounds.Min.X
				} else if px >= bounds.Max.X {
					px = bounds.Max.X - 1
				}

				srcColor := src.RGBAAt(px, y)
				weight := kernel[i]

				r += float64(srcColor.R) * weight
				g += float64(srcColor.G) * weight
				b += float64(srcColor.B) * weight
			}

			temp.Set(x, y, color.RGBA{
				R: uint8(r),
				G: uint8(g),
				B: uint8(b),
				A: 255,
			})
		}
	}

	// Vertical pass
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			var r, g, b float64

			for i := 0; i < len(kernel); i++ {
				py := y + i - radius
				if py < bounds.Min.Y {
					py = bounds.Min.Y
				} else if py >= bounds.Max.Y {
					py = bounds.Max.Y - 1
				}

				srcColor := temp.RGBAAt(x, py)
				weight := kernel[i]

				r += float64(srcColor.R) * weight
				g += float64(srcColor.G) * weight
				b += float64(srcColor.B) * weight
			}

			dst.Set(x, y, color.RGBA{
				R: uint8(r),
				G: uint8(g),
				B: uint8(b),
				A: 255,
			})
		}
	}

	return dst
}

func textToFilename(text string) string {
	// Convert to lowercase
	filename := strings.ToLower(text)

	// Replace spaces with hyphens
	filename = strings.ReplaceAll(filename, " ", "-")

	// Remove any characters that aren't alphanumeric or hyphens
	reg := regexp.MustCompile(`[^a-z0-9\-]+`)
	filename = reg.ReplaceAllString(filename, "")

	// Remove multiple consecutive hyphens
	reg = regexp.MustCompile(`-+`)
	filename = reg.ReplaceAllString(filename, "-")

	// Trim hyphens from start and end
	filename = strings.Trim(filename, "-")

	// If filename is empty after cleaning, use a hash
	if filename == "" {
		hash := sha256.Sum256([]byte(text))
		filename = fmt.Sprintf("image-%x", hash[:8])
	}

	// Limit length to prevent filesystem issues
	if len(filename) > 100 {
		filename = filename[:100]
		filename = strings.TrimRight(filename, "-")
	}

	return filename
}