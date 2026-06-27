package handle

// #include <stdlib.h>
import "C"

import (
	"sync"
	"unsafe"
)

var (
	mu    sync.RWMutex
	store = make(map[unsafe.Pointer]any)
)

// Save registers v and returns an opaque token for it suitable for passing to C
// as a void*. The token stays valid until passed to [Delete]. Save returns nil
// when v is nil.
func Save(v any) unsafe.Pointer {
	if v == nil {
		return nil
	}
	// A one-byte C allocation gives a unique, stable address that is not a Go
	// pointer, so it is safe to hand to C and use as a map key.
	token := C.malloc(C.size_t(1))
	if token == nil {
		panic("handle: malloc returned nil")
	}
	mu.Lock()
	store[token] = v
	mu.Unlock()
	return token
}

// Restore returns the value previously registered for token, or nil when token
// is nil or unknown.
func Restore(token unsafe.Pointer) any {
	if token == nil {
		return nil
	}
	mu.RLock()
	defer mu.RUnlock()
	return store[token]
}

// Delete releases the value registered for token and frees the token. It is a
// no-op when token is nil.
func Delete(token unsafe.Pointer) {
	if token == nil {
		return
	}
	mu.Lock()
	delete(store, token)
	mu.Unlock()
	C.free(token)
}
