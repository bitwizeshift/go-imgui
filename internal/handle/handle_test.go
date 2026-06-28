package handle_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/bitwizeshift/go-imgui/internal/handle"
)

type point struct {
	X, Y int
}

func TestSaveRestore(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		value any
		want  any
	}{
		{
			name:  "string",
			value: "hello",
			want:  "hello",
		},
		{
			name:  "int",
			value: 42,
			want:  42,
		},
		{
			name:  "struct",
			value: point{X: 1, Y: 2},
			want:  point{X: 1, Y: 2},
		},
		{
			name:  "slice",
			value: []int{1, 2, 3},
			want:  []int{1, 2, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Arrange
			token := handle.Save(tc.value)
			t.Cleanup(func() { handle.Delete(token) })

			// Act
			got := handle.Restore(token)

			// Assert
			if got, want := got, tc.want; !reflect.DeepEqual(got, want) {
				t.Errorf("Restore() = %v, want %v", got, want)
			}
		})
	}
}

func TestSaveNilReturnsNilToken(t *testing.T) {
	t.Parallel()

	// Arrange
	var value any

	// Act
	token := handle.Save(value)

	// Assert
	if got, want := token == nil, true; got != want {
		t.Errorf("Save(nil) token == nil = %v, want %v", got, want)
	}
}

func TestRestoreNilReturnsNil(t *testing.T) {
	t.Parallel()

	// Arrange
	var token unsafe.Pointer

	// Act
	got := handle.Restore(token)

	// Assert
	if got, want := got, any(nil); got != want {
		t.Errorf("Restore(nil) = %v, want %v", got, want)
	}
}

func TestDeleteRemovesValue(t *testing.T) {
	t.Parallel()

	// Arrange
	token := handle.Save("value")

	// Act
	handle.Delete(token)
	got := handle.Restore(token)

	// Assert
	if got, want := got, any(nil); got != want {
		t.Errorf("Restore() after Delete = %v, want %v", got, want)
	}
}

func TestSaveReturnsDistinctTokens(t *testing.T) {
	t.Parallel()

	// Arrange
	first := handle.Save("a")
	second := handle.Save("b")
	t.Cleanup(func() {
		handle.Delete(first)
		handle.Delete(second)
	})

	// Act
	gotFirst := handle.Restore(first)
	gotSecond := handle.Restore(second)

	// Assert
	if got, want := first == second, false; got != want {
		t.Errorf("distinct Save tokens equal = %v, want %v", got, want)
	}
	if got, want := gotFirst, any("a"); got != want {
		t.Errorf("Restore(first) = %v, want %v", got, want)
	}
	if got, want := gotSecond, any("b"); got != want {
		t.Errorf("Restore(second) = %v, want %v", got, want)
	}
}
