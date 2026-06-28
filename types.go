package imgui

import "rodusek.dev/pkg/imgui/internal/cimgui"

// Vec2 is a 2D vector of 32-bit floats, mirroring Dear ImGui's ImVec2. It is an
// alias of the binding type so it can be passed straight through with no copy.
type Vec2 = cimgui.Vec2

// Vec4 is a 4D vector of 32-bit floats, mirroring Dear ImGui's ImVec4.
type Vec4 = cimgui.Vec4

// Color is an RGBA color with components in the 0..1 range.
type Color struct {
	R, G, B, A float32
}

// Vec4 converts the color to a [Vec4] (the form Dear ImGui consumes).
func (c Color) Vec4() Vec4 { return Vec4{X: c.R, Y: c.G, Z: c.B, W: c.A} }

// ColorFromVec4 builds a [Color] from a [Vec4].
func ColorFromVec4(v Vec4) Color { return Color{R: v.X, G: v.Y, B: v.Z, A: v.W} }
