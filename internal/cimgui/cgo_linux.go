//go:build linux

package cimgui

/*
#cgo CFLAGS: -I${SRCDIR}/csources/cimgui -I${SRCDIR}/csources/glfw/include -DCIMGUI_DEFINE_ENUMS_AND_STRUCTS -D_GLFW_X11 -Wno-deprecated-declarations
#cgo CXXFLAGS: -std=c++17 -I${SRCDIR} -I${SRCDIR}/csources/cimgui -I${SRCDIR}/csources/cimgui/imgui -I${SRCDIR}/csources/backends -I${SRCDIR}/csources/glfw/include -DGLFW_INCLUDE_NONE -Wno-deprecated-declarations
#cgo LDFLAGS: -lGL -lX11 -lXrandr -lXi -lXcursor -lXinerama -ldl -lpthread -lm
*/
import "C"
