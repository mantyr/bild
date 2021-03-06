package bild

// RGBAF64 represents an RGBA color using the range 0.0 to 1.0 with a float64 for each channel.
type RGBAF64 struct {
	R, G, B, A float64
}

// NewRGBAF64 returns a new RGBAF64 color based on the provided uint8 values.
func NewRGBAF64(r, g, b, a uint8) RGBAF64 {
	outR := float64(r) / 255
	outG := float64(g) / 255
	outB := float64(b) / 255
	outA := float64(a) / 255

	return RGBAF64{outR, outG, outB, outA}
}

// Clamp limits the channel values of the RGBAF64 color to the range 0.0 to 1.0.
func (c *RGBAF64) Clamp() {
	c.R = clampFloat64(c.R, 0, 1)
	c.G = clampFloat64(c.G, 0, 1)
	c.B = clampFloat64(c.B, 0, 1)
	c.A = clampFloat64(c.A, 0, 1)
}
