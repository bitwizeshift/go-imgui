package cimgui

// #include <stdlib.h>
// #include "cimgui.h"
import "C"

import "unsafe"

// TreeNodeFlags configures tree nodes and headers. Mirrors the public
// ImGuiTreeNodeFlags_.
type TreeNodeFlags int32

const (
	TreeNodeFlagsNone                 TreeNodeFlags = C.ImGuiTreeNodeFlags_None
	TreeNodeFlagsSelected             TreeNodeFlags = C.ImGuiTreeNodeFlags_Selected
	TreeNodeFlagsFramed               TreeNodeFlags = C.ImGuiTreeNodeFlags_Framed
	TreeNodeFlagsAllowOverlap         TreeNodeFlags = C.ImGuiTreeNodeFlags_AllowOverlap
	TreeNodeFlagsNoTreePushOnOpen     TreeNodeFlags = C.ImGuiTreeNodeFlags_NoTreePushOnOpen
	TreeNodeFlagsNoAutoOpenOnLog      TreeNodeFlags = C.ImGuiTreeNodeFlags_NoAutoOpenOnLog
	TreeNodeFlagsDefaultOpen          TreeNodeFlags = C.ImGuiTreeNodeFlags_DefaultOpen
	TreeNodeFlagsOpenOnDoubleClick    TreeNodeFlags = C.ImGuiTreeNodeFlags_OpenOnDoubleClick
	TreeNodeFlagsOpenOnArrow          TreeNodeFlags = C.ImGuiTreeNodeFlags_OpenOnArrow
	TreeNodeFlagsLeaf                 TreeNodeFlags = C.ImGuiTreeNodeFlags_Leaf
	TreeNodeFlagsBullet               TreeNodeFlags = C.ImGuiTreeNodeFlags_Bullet
	TreeNodeFlagsFramePadding         TreeNodeFlags = C.ImGuiTreeNodeFlags_FramePadding
	TreeNodeFlagsSpanAvailWidth       TreeNodeFlags = C.ImGuiTreeNodeFlags_SpanAvailWidth
	TreeNodeFlagsSpanFullWidth        TreeNodeFlags = C.ImGuiTreeNodeFlags_SpanFullWidth
	TreeNodeFlagsSpanLabelWidth       TreeNodeFlags = C.ImGuiTreeNodeFlags_SpanLabelWidth
	TreeNodeFlagsSpanAllColumns       TreeNodeFlags = C.ImGuiTreeNodeFlags_SpanAllColumns
	TreeNodeFlagsLabelSpanAllColumns  TreeNodeFlags = C.ImGuiTreeNodeFlags_LabelSpanAllColumns
	TreeNodeFlagsNavLeftJumpsToParent TreeNodeFlags = C.ImGuiTreeNodeFlags_NavLeftJumpsToParent
	TreeNodeFlagsCollapsingHeader     TreeNodeFlags = C.ImGuiTreeNodeFlags_CollapsingHeader
)

// TreeNode_Str opens a tree node labeled by label. Call [TreePop] if it returns
// true.
func TreeNode_Str(label string) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igTreeNode_Str(clabel))
}

// TreeNodeEx_Str opens a tree node with flags. Call [TreePop] if it returns true
// (unless TreeNodeFlagsNoTreePushOnOpen is set).
func TreeNodeEx_Str(label string, flags TreeNodeFlags) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igTreeNodeEx_Str(clabel, C.ImGuiTreeNodeFlags(flags)))
}

// TreePush_Str indents and pushes an ID onto the stack; pair with [TreePop].
func TreePush_Str(strID string) {
	cid := C.CString(strID)
	defer C.free(unsafe.Pointer(cid))
	C.igTreePush_Str(cid)
}

// TreePop unindents and pops the tree node ID pushed by an open tree node.
func TreePop() { C.igTreePop() }

// SetNextItemOpen sets the open state applied to the next tree node or header.
func SetNextItemOpen(isOpen bool, cond Cond) {
	C.igSetNextItemOpen(C.bool(isOpen), C.ImGuiCond(cond))
}

// GetTreeNodeToLabelSpacing returns the horizontal distance from a tree node's
// start to its label.
func GetTreeNodeToLabelSpacing() float32 { return float32(C.igGetTreeNodeToLabelSpacing()) }

// CollapsingHeader_TreeNodeFlags draws a collapsing header and reports whether
// it is open.
func CollapsingHeader_TreeNodeFlags(label string, flags TreeNodeFlags) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igCollapsingHeader_TreeNodeFlags(clabel, C.ImGuiTreeNodeFlags(flags)))
}

// CollapsingHeader_BoolPtr draws a collapsing header with a close button bound to
// pVisible. When pVisible becomes false the header is hidden. Reports whether the
// header is open.
func CollapsingHeader_BoolPtr(label string, pVisible *bool, flags TreeNodeFlags) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	var ret C.bool
	withBoolPtr(pVisible, func(p *C.bool) {
		ret = C.igCollapsingHeader_BoolPtr(clabel, p, C.ImGuiTreeNodeFlags(flags))
	})
	return bool(ret)
}
