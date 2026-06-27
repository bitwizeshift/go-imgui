// package handle hands out opaque tokens that carry a Go value across the cgo
// boundary.
//
// cgo forbids passing a Go pointer to C when it points at memory that itself
// contains Go pointers, so a Go value such as a closure cannot be stored in a C
// "user_data" field directly. Instead, register the value with [Save] to obtain
// a token, hand that token to C, and resolve it back with [Restore] when C
// returns it. Release it with [Delete] once the C side no longer holds it.
//
// This mirrors github.com/mattn/go-pointer so the published module needs no
// third-party runtime dependency. It exists for ImGui callbacks (for example
// InputText) and is unused until those are wrapped.
package handle
