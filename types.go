package imgui

// Vec2 is a 2D vector of 32-bit floats, mirroring ImGui's ImVec2.
type Vec2 struct {
	X, Y float32
}

// Vec4 is a 4D vector of 32-bit floats, mirroring ImGui's ImVec4. It also serves
// as an RGBA color.
type Vec4 struct {
	X, Y, Z, W float32
}
