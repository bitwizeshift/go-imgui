package imgui

// ID is a widget identifier accepted by the id-overloaded scope functions
// ([Child], [OpenPopup], etc): either a string label or a precomputed uint32 id.
type ID interface {
	string | uint32
}

// SignedOrUnsigned32 is a 32-bit signed or unsigned integer, the bitset value
// type edited by [CheckboxFlags].
type SignedOrUnsigned32 interface {
	int32 | uint32
}

// StyleVarValue is a style-variable value accepted by [PushStyleVar] and
// [StyleVarScope]: a float32 or a [Vec2].
type StyleVarValue interface {
	float32 | Vec2
}

// BoolOrPtr is a selected-state value accepted by [Selectable] and [MenuItem],
// held either by value (bool) or by pointer (*bool).
type BoolOrPtr interface {
	bool | *bool
}
