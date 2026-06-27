//go:build darwin

package cimgui

/*
#cgo CFLAGS: -I${SRCDIR}/csources/cimgui -I${SRCDIR}/csources/glfw/include -DCIMGUI_DEFINE_ENUMS_AND_STRUCTS -D_GLFW_COCOA -Wno-deprecated-declarations
#cgo CXXFLAGS: -std=c++17 -I${SRCDIR} -I${SRCDIR}/csources/cimgui -I${SRCDIR}/csources/cimgui/imgui -I${SRCDIR}/csources/backends -I${SRCDIR}/csources/glfw/include -DGLFW_INCLUDE_NONE -DGL_SILENCE_DEPRECATION -Wno-deprecated-declarations
#cgo LDFLAGS: -framework Cocoa -framework IOKit -framework CoreFoundation -framework OpenGL
*/
import "C"
