package imgui

import "github.com/bitwizeshift/go-imgui/internal/cimgui"

// TreeNodeOptions are the optional inputs to [TreeNodeEx], [CollapsingHeader] and
// [CollapsingHeaderClosable]. A nil *TreeNodeOptions uses Dear ImGui's defaults;
// each field maps to an ImGuiTreeNodeFlags_ bit.
type TreeNodeOptions struct {
	Selected             bool // ImGuiTreeNodeFlags_Selected
	Framed               bool // ImGuiTreeNodeFlags_Framed
	AllowOverlap         bool // ImGuiTreeNodeFlags_AllowOverlap
	NoTreePushOnOpen     bool // ImGuiTreeNodeFlags_NoTreePushOnOpen
	NoAutoOpenOnLog      bool // ImGuiTreeNodeFlags_NoAutoOpenOnLog
	DefaultOpen          bool // ImGuiTreeNodeFlags_DefaultOpen
	OpenOnDoubleClick    bool // ImGuiTreeNodeFlags_OpenOnDoubleClick
	OpenOnArrow          bool // ImGuiTreeNodeFlags_OpenOnArrow
	Leaf                 bool // ImGuiTreeNodeFlags_Leaf
	Bullet               bool // ImGuiTreeNodeFlags_Bullet
	FramePadding         bool // ImGuiTreeNodeFlags_FramePadding
	SpanAvailWidth       bool // ImGuiTreeNodeFlags_SpanAvailWidth
	SpanFullWidth        bool // ImGuiTreeNodeFlags_SpanFullWidth
	SpanLabelWidth       bool // ImGuiTreeNodeFlags_SpanLabelWidth
	SpanAllColumns       bool // ImGuiTreeNodeFlags_SpanAllColumns
	LabelSpanAllColumns  bool // ImGuiTreeNodeFlags_LabelSpanAllColumns
	NavLeftJumpsToParent bool // ImGuiTreeNodeFlags_NavLeftJumpsToParent
	CollapsingHeader     bool // ImGuiTreeNodeFlags_CollapsingHeader (composite)
}

// flags folds o into the cimgui flag set. A nil receiver yields no flags.
func (o *TreeNodeOptions) flags() cimgui.TreeNodeFlags {
	if o == nil {
		return cimgui.TreeNodeFlagsNone
	}
	var f cimgui.TreeNodeFlags
	if o.Selected {
		f |= cimgui.TreeNodeFlagsSelected
	}
	if o.Framed {
		f |= cimgui.TreeNodeFlagsFramed
	}
	if o.AllowOverlap {
		f |= cimgui.TreeNodeFlagsAllowOverlap
	}
	if o.NoTreePushOnOpen {
		f |= cimgui.TreeNodeFlagsNoTreePushOnOpen
	}
	if o.NoAutoOpenOnLog {
		f |= cimgui.TreeNodeFlagsNoAutoOpenOnLog
	}
	if o.DefaultOpen {
		f |= cimgui.TreeNodeFlagsDefaultOpen
	}
	if o.OpenOnDoubleClick {
		f |= cimgui.TreeNodeFlagsOpenOnDoubleClick
	}
	if o.OpenOnArrow {
		f |= cimgui.TreeNodeFlagsOpenOnArrow
	}
	if o.Leaf {
		f |= cimgui.TreeNodeFlagsLeaf
	}
	if o.Bullet {
		f |= cimgui.TreeNodeFlagsBullet
	}
	if o.FramePadding {
		f |= cimgui.TreeNodeFlagsFramePadding
	}
	if o.SpanAvailWidth {
		f |= cimgui.TreeNodeFlagsSpanAvailWidth
	}
	if o.SpanFullWidth {
		f |= cimgui.TreeNodeFlagsSpanFullWidth
	}
	if o.SpanLabelWidth {
		f |= cimgui.TreeNodeFlagsSpanLabelWidth
	}
	if o.SpanAllColumns {
		f |= cimgui.TreeNodeFlagsSpanAllColumns
	}
	if o.LabelSpanAllColumns {
		f |= cimgui.TreeNodeFlagsLabelSpanAllColumns
	}
	if o.NavLeftJumpsToParent {
		f |= cimgui.TreeNodeFlagsNavLeftJumpsToParent
	}
	if o.CollapsingHeader {
		f |= cimgui.TreeNodeFlagsCollapsingHeader
	}
	return f
}

// TreeNode opens a tree node labelled label. It models ImGui::TreeNode. open
// reports whether the node is expanded; the returned [EndFunc] (ImGui::TreePop)
// unindents it and is called only when open.
func TreeNode(label string) (open bool, end EndFunc) {
	open = cimgui.TreeNode_Str(label)
	if !open {
		return open, func() {}
	}
	return open, cimgui.TreePop
}

// TreeNodeEx opens a tree node with options. It models ImGui::TreeNodeEx. open
// reports whether the node is expanded; the returned [EndFunc] (ImGui::TreePop)
// is called only when open and the node pushed onto the tree stack (i.e. unless
// NoTreePushOnOpen is set).
func TreeNodeEx(label string, opts *TreeNodeOptions) (open bool, end EndFunc) {
	flags := opts.flags()
	open = cimgui.TreeNodeEx_Str(label, flags)
	if !open || flags&cimgui.TreeNodeFlagsNoTreePushOnOpen != 0 {
		return open, func() {}
	}
	return open, cimgui.TreePop
}

// TreePush indents and pushes strID onto the ID stack; balance it with
// [TreePop]. It models ImGui::TreePush.
func TreePush(strID string) {
	cimgui.TreePush_Str(strID)
}

// TreePop unindents and pops the ID pushed by [TreePush]. It models
// ImGui::TreePop. (A tree node opened with [TreeNode] or [TreeNodeEx] is closed
// instead by the [EndFunc] those return.)
func TreePop() {
	cimgui.TreePop()
}

// SetNextItemOpen sets the open state applied to the next tree node or header. It
// models ImGui::SetNextItemOpen.
func SetNextItemOpen(isOpen bool, cond Cond) {
	cimgui.SetNextItemOpen(isOpen, cond)
}

// GetTreeNodeToLabelSpacing returns the horizontal distance from a tree node's
// start to its label. It models ImGui::GetTreeNodeToLabelSpacing.
func GetTreeNodeToLabelSpacing() float32 {
	return cimgui.GetTreeNodeToLabelSpacing()
}

// CollapsingHeader draws a collapsing header and reports whether it is open. It
// models ImGui::CollapsingHeader (the flags overload).
func CollapsingHeader(label string, opts *TreeNodeOptions) bool {
	return cimgui.CollapsingHeader_TreeNodeFlags(label, opts.flags())
}

// CollapsingHeaderClosable draws a collapsing header with a close button bound to
// pVisible; when *pVisible becomes false the header is hidden. It reports whether
// the header is open and models ImGui::CollapsingHeader (the bool* overload).
func CollapsingHeaderClosable(label string, pVisible *bool, opts *TreeNodeOptions) bool {
	return cimgui.CollapsingHeader_BoolPtr(label, pVisible, opts.flags())
}
