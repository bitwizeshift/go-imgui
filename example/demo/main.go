// Command demo opens a window and exercises the high-level widget packages. It
// builds a fresh widget tree each frame from persistent state and displays it;
// it never touches the internal cimgui bindings.
package main

import (
	"fmt"
	"image"
	"log"
	"math"
	"unicode"

	"rodusek.dev/pkg/imgui"
	"rodusek.dev/pkg/imgui/app"
	"rodusek.dev/pkg/imgui/button"
	"rodusek.dev/pkg/imgui/canvas"
	"rodusek.dev/pkg/imgui/color"
	"rodusek.dev/pkg/imgui/combo"
	"rodusek.dev/pkg/imgui/debug"
	"rodusek.dev/pkg/imgui/input"
	"rodusek.dev/pkg/imgui/layout"
	"rodusek.dev/pkg/imgui/menu"
	"rodusek.dev/pkg/imgui/plot"
	"rodusek.dev/pkg/imgui/popup"
	"rodusek.dev/pkg/imgui/style"
	"rodusek.dev/pkg/imgui/tab"
	"rodusek.dev/pkg/imgui/table"
	"rodusek.dev/pkg/imgui/text"
	"rodusek.dev/pkg/imgui/texture"
	"rodusek.dev/pkg/imgui/tooltip"
	"rodusek.dev/pkg/imgui/tree"
	"rodusek.dev/pkg/imgui/window"
)

func main() {
	// Persistent state lives here; the widget tree below is rebuilt each frame
	// and reads/writes these values through bound pointers and callbacks.
	var (
		showDemo          = true
		counter   int32   = 0
		enabled           = true
		choice    int32   = 0
		speed     float32 = 0.4
		count     int32   = 3
		drag      float32 = 1
		amount    int32   = 5
		name              = "edit me"
		notes             = "multi\nline\ntext"
		shout             = "grows as you type"
		history           = "press Up / Down"
		editCount int32   = 0
		funcIdx   int32   = 0
		rounding  float32 = 4
		frame             = 0
	)

	rgb := imgui.Color{R: 0.4, G: 0.7, B: 0.2, A: 1}
	rgba := imgui.Color{R: 0.3, G: 0.6, B: 0.9, A: 1}
	comboItems := []string{"Red", "Green", "Blue"}
	var comboIdx, listIdx int32
	selected := [3]bool{}
	plotData := []float32{0.1, 0.5, 0.2, 0.8, 0.4, 0.9, 0.3, 0.7, 0.6, 1.0}
	tabOpen := true
	var progress float32
	var tex *texture.Texture
	defer func() { tex.Close() }() // tex is created lazily below; close whatever exists at exit

	// Persistent state for the "Widget coverage" window (see coverage).
	cov := coverageState{
		f2: [2]float32{1, 2}, f3: [3]float32{1, 2, 3}, f4: [4]float32{1, 2, 3, 4},
		i2: [2]int32{1, 2}, i3: [3]int32{1, 2, 3}, i4: [4]int32{1, 2, 3, 4},
		dbl: 3.14159, sf3: [3]float32{0.2, 0.5, 0.8}, si2: [2]int32{3, 7},
		angle: 0.5, vsf: 0.5, vsi: 5,
		df2: [2]float32{10, 20}, di3: [3]int32{1, 2, 3},
		rangeLo: 2, rangeHi: 8, irangeLo: 10, irangeHi: 90,
		scalarU64: 4096, scalarN: []float64{1, 2, 3},
		ssF: 1.5, ssN: []int32{2, 4}, dsI: 5, dsN: []float32{0.1, 0.2}, vssF: 0.5,
		flags: flagA,
	}

	err := app.Run(app.Config{Title: "go-imgui high-level demo"}, func() {
		if tex == nil { // create lazily: needs a current GL context
			tex = texture.FromImage(checker(64, 64))
		}
		progress += 0.004
		if progress > 1 {
			progress = 0
		}
		frame++

		imgui.Display(
			mainMenu(&showDemo),
			gallery(galleryState{
				counter: &counter, enabled: &enabled, choice: &choice,
				speed: &speed, count: &count, drag: &drag, amount: &amount,
				name: &name, notes: &notes, rgb: &rgb, rgba: &rgba,
				comboItems: comboItems, comboIdx: &comboIdx, listIdx: &listIdx,
				selected: selected[:], plotData: plotData, progress: progress,
				tabOpen: &tabOpen, tex: tex,
				shout: &shout, history: &history, editCount: &editCount,
				funcIdx: &funcIdx, rounding: &rounding,
			}),
			coverage(&cov),
			// A HUD overlay drawn in front of every window, in absolute screen
			// coordinates.
			canvas.Foreground(func(d *canvas.Drawer) {
				d.AddText(imgui.Vec2{X: 12, Y: 30}, imgui.Color{R: 1, G: 1, B: 0, A: 1},
					fmt.Sprintf("frame %d", frame))
			}),
		)
		if showDemo {
			(&debug.DemoWindow{Open: &showDemo}).Display()
		}
	})
	if err != nil {
		log.Fatal(err)
	}
}

func mainMenu(showDemo *bool) imgui.Widget {
	bar := menu.NewMainBar()

	view := menu.New("View")
	demoItem := menu.NewItem("Demo window")
	demoItem.Selected = showDemo
	view.AddWidget(demoItem)

	bar.AddWidget(view)
	return bar
}

type galleryState struct {
	counter           *int32
	enabled           *bool
	choice            *int32
	speed             *float32
	count             *int32
	drag              *float32
	amount            *int32
	name, notes       *string
	rgb, rgba         *imgui.Color
	comboItems        []string
	comboIdx, listIdx *int32
	selected          []bool
	plotData          []float32
	progress          float32
	tabOpen           *bool
	tex               *texture.Texture
	shout, history    *string
	editCount         *int32
	funcIdx           *int32
	rounding          *float32
}

func gallery(s galleryState) imgui.Widget {
	w := window.New("Widget gallery")
	w.SetMenuBar(true)
	w.Size = &imgui.Vec2{X: 480, Y: 660}

	// Window menu bar.
	options := menu.New("Options")
	enabledItem := menu.NewItem("Enabled")
	enabledItem.Selected = s.enabled
	options.AddWidget(enabledItem)
	bar := menu.NewBar()
	bar.AddWidget(options)
	w.AddWidget(bar)

	// Text.
	w.AddWidget(
		text.SeparatorText("Text"),
		text.New("Plain text."),
		text.Colored(imgui.Color{R: 1, G: 0.6, B: 0.2, A: 1}, "Colored text."),
		text.Disabled("Disabled text."),
		text.Wrapped("Wrapped text that folds onto several lines when the window is narrow."),
		text.Bullet("A bulleted line."),
		text.NewLinkURL("dearimgui.com", "https://www.dearimgui.com"),
	)

	// Buttons.
	click := button.New("Click me")
	click.OnClick = func() { *s.counter++ }
	w.AddWidget(
		text.SeparatorText("Buttons"),
		click,
		layout.SameLine(),
		button.NewArrow("##left", button.Left),
		layout.SameLine(),
		button.NewArrow("##right", button.Right),
		text.Labelf("clicks = %d", *s.counter),
		&button.ProgressBar{Fraction: s.progress},
	)

	// Selection.
	radioA := input.NewRadio("A", s.choice, 0)
	radioB := input.NewRadio("B", s.choice, 1)
	w.AddWidget(
		text.SeparatorText("Selection"),
		tooltip.For(input.NewCheckbox("enabled", s.enabled), "toggle the controls below"),
		radioA, layout.SameLine(), radioB,
	)

	// Inputs (disabled when "enabled" is off, to show layout.Disabled).
	notesInput := &input.Text{Label: "notes", Value: s.notes, Multiline: true, Size: imgui.Vec2{Y: 50}}
	w.AddWidget(
		text.SeparatorText("Inputs"),
		layout.NewDisabled(!*s.enabled,
			input.NewSliderFloat("speed", s.speed, 0, 1),
			input.NewSliderInt("count", s.count, 0, 10),
			input.NewDragFloat("drag", s.drag),
			input.NewInt("amount", s.amount),
			input.NewText("name", s.name),
			notesInput,
		),
	)

	// Color.
	w.AddWidget(
		text.SeparatorText("Color"),
		color.NewEdit("rgb", s.rgb),
		color.NewPicker("rgba", s.rgba),
		color.NewButton("swatch", *s.rgb),
	)

	// Combo / list / selectable.
	w.AddWidget(
		text.SeparatorText("Combo & list"),
		combo.New("combo", s.comboIdx, s.comboItems),
		combo.NewListBox("list", s.listIdx, s.comboItems),
	)
	for i := range s.selected {
		sel := combo.NewSelectable(fmt.Sprintf("item %d", i))
		sel.Selected = &s.selected[i]
		w.AddWidget(sel)
	}

	// Trees.
	node := tree.New("Tree root")
	node.AddWidget(text.Bullet("child a"), text.Bullet("child b"))
	header := tree.NewHeader("Collapsing header")
	header.AddWidget(text.New("header body"))
	w.AddWidget(text.SeparatorText("Trees"), node, header)

	// Child region.
	child := window.NewChild("scroller")
	child.Size = imgui.Vec2{Y: 60}
	child.Border = true
	for i := range 8 {
		child.AddWidget(text.Labelf("scrolling line %d", i))
	}
	w.AddWidget(text.SeparatorText("Child region"), child)

	// Tabs (the third tab is closeable and reopens via the checkbox).
	bars := tab.NewBar("tabs")
	first := tab.NewItem("First")
	first.AddWidget(text.New("first tab body"))
	second := tab.NewItem("Second")
	second.AddWidget(text.New("second tab body"))
	closeable := tab.NewItem("Closeable")
	closeable.Open = s.tabOpen
	closeable.AddWidget(text.New("close with the x"))
	bars.AddWidget(first, second, closeable)
	w.AddWidget(
		text.SeparatorText("Tabs"),
		input.NewCheckbox("closeable tab open", s.tabOpen),
		bars,
	)

	// Table.
	grid := table.New("grid")
	grid.AddColumn("name")
	grid.AddColumn("value")
	grid.AddRow(text.New("counter"), text.Labelf("%d", *s.counter))
	grid.AddRow(text.New("speed"), text.Labelf("%.2f", *s.speed))
	w.AddWidget(text.SeparatorText("Table"), grid)

	// Plots.
	w.AddWidget(
		text.SeparatorText("Plots"),
		&plot.Lines{Label: "lines", Values: s.plotData, Size: imgui.Vec2{Y: 50}},
		&plot.Histogram{Label: "histogram", Values: s.plotData, Size: imgui.Vec2{Y: 50}},
	)

	// Images (a texture built from a Go image.Image).
	w.AddWidget(
		text.SeparatorText("Images"),
		texture.NewImage(s.tex, imgui.Vec2{X: 64, Y: 64}),
		layout.SameLine(),
		texture.NewButton("imgbtn", s.tex, imgui.Vec2{X: 48, Y: 48}),
	)

	// Popups.
	openModal := button.New("Open modal")
	openModal.OnClick = func() { popup.Open("Demo modal") }
	ok := button.New("Close")
	ok.OnClick = func() { popup.CloseCurrent() }
	modal := popup.NewModal("Demo modal")
	modal.AddWidget(text.New("This is a modal dialog."), ok)

	rightClickMe := text.New("Right-click me for a context menu")
	ctx := popup.NewContextItem("ctx")
	ctx.AddWidget(combo.NewSelectable("a context action"))

	w.AddWidget(
		text.SeparatorText("Popups"),
		openModal, modal,
		rightClickMe, ctx,
	)

	// Input callbacks: an auto-growing buffer with an uppercase char filter and an
	// edit counter, plus an Up/Down history handler that rewrites the buffer.
	shoutInput := &input.Text{Label: "shout (uppercase)", Value: s.shout}
	shoutInput.OnCharFilter = func(r rune) rune { return unicode.ToUpper(r) }
	shoutInput.OnEdit = func() { *s.editCount++ }
	histInput := &input.Text{Label: "history (Up/Down)", Value: s.history}
	histInput.OnHistory = func(d *input.CallbackData, dir input.HistoryDir) {
		entry := "recalled previous"
		if dir == input.HistoryDown {
			entry = "recalled next"
		}
		d.DeleteChars(0, len(d.Buf()))
		d.InsertChars(0, entry)
	}
	w.AddWidget(
		text.SeparatorText("Input callbacks"),
		shoutInput,
		text.Labelf("edits = %d", *s.editCount),
		histInput,
	)

	// Getter-sourced combo and plot: items and samples are produced lazily.
	w.AddWidget(
		text.SeparatorText("Getter-sourced"),
		combo.NewFunc("func combo", s.funcIdx, 16, func(i int32) string {
			return fmt.Sprintf("generated #%d", i)
		}),
		&plot.Lines{
			Label:  "sin (getter)",
			Getter: func(i int32) float32 { return float32(math.Sin(float64(i) * 0.3)) },
			Count:  48,
			Size:   imgui.Vec2{Y: 50},
		},
	)

	// Custom drawing into a canvas-local draw list.
	draw := canvas.New("scratch", imgui.Vec2{Y: 120}, func(d *canvas.Drawer) {
		sz := d.Avail()
		d.AddRectFilled(imgui.Vec2{}, sz, imgui.Color{R: 0.12, G: 0.12, B: 0.16, A: 1}, 6)
		d.AddCircleFilled(imgui.Vec2{X: sz.X * 0.3, Y: sz.Y * 0.5}, 34, imgui.Color{R: 0.9, G: 0.45, B: 0.2, A: 1}, 0)
		d.AddLine(imgui.Vec2{X: 12, Y: 12}, imgui.Vec2{X: sz.X - 12, Y: sz.Y - 12}, imgui.Color{R: 0.4, G: 0.9, B: 0.5, A: 1}, 3)
		d.AddTriangleFilled(
			imgui.Vec2{X: sz.X * 0.7, Y: sz.Y * 0.25},
			imgui.Vec2{X: sz.X * 0.85, Y: sz.Y * 0.75},
			imgui.Vec2{X: sz.X * 0.55, Y: sz.Y * 0.75},
			imgui.Color{R: 0.4, G: 0.6, B: 1, A: 1},
		)
		d.AddText(imgui.Vec2{X: 10, Y: 8}, imgui.Color{R: 1, G: 1, B: 1, A: 1}, "draw list")
	})
	w.AddWidget(text.SeparatorText("Custom drawing"), draw)

	// Theme: presets, a live style variable, and a locally re-styled subtree.
	dark := button.New("Dark")
	dark.OnClick = func() { style.Dark() }
	light := button.New("Light")
	light.OnClick = func() { style.Light() }
	classic := button.New("Classic")
	classic.OnClick = func() { style.Classic() }
	roundSlider := input.NewSliderFloat("frame rounding", s.rounding, 0, 12)
	roundSlider.OnChange = func(v float32) {
		th := style.Current()
		th.FrameRounding = v
		th.Apply()
	}
	scoped := style.NewScoped(
		text.New("text recolored within style.Scoped"),
		button.New("scoped button"),
	)
	scoped.Colors = map[style.Col]imgui.Color{
		style.ColText:   {R: 1, G: 0.85, B: 0.3, A: 1},
		style.ColButton: {R: 0.3, G: 0.5, B: 0.2, A: 1},
	}
	scoped.Vars = []style.Override{style.FloatVar(style.VarFrameRounding, 10)}
	w.AddWidget(
		text.SeparatorText("Theme"),
		dark, layout.SameLine(), light, layout.SameLine(), classic,
		roundSlider,
		scoped,
	)

	// Escape hatch.
	w.AddWidget(imgui.Custom(func() {
		// Anything not yet modelled can run here against the lower-level API.
	}))

	return w
}

// Example bitfield for the flag checkboxes.
const (
	flagA int32 = 1 << 0
	flagB int32 = 1 << 1
	flagC int32 = 1 << 2
)

// coverageState holds the persistent values bound by the widgets in coverage.
type coverageState struct {
	f2                 [2]float32
	f3                 [3]float32
	f4                 [4]float32
	i2                 [2]int32
	i3                 [3]int32
	i4                 [4]int32
	dbl                float64
	sf3                [3]float32
	si2                [2]int32
	angle              float32
	vsf                float32
	vsi                int32
	df2                [2]float32
	di3                [3]int32
	rangeLo, rangeHi   float32
	irangeLo, irangeHi int32
	scalarU64          uint64
	scalarN            []float64
	ssF                float32
	ssN                []int32
	dsI                int32
	dsN                []float32
	vssF               float32
	flags              int32
	uflags             uint32
	customSel          [3]bool
	listSel            [4]bool
	invisibleClicks    int32
	tabButtonClicks    int32
	showPanel          bool
}

// coverage builds a window exercising the multi-component, scalar, and container
// widgets, using only the public packages.
func coverage(s *coverageState) imgui.Widget {
	w := window.New("Widget coverage")
	w.Pos = &imgui.Vec2{X: 540, Y: 20}
	w.Size = &imgui.Vec2{X: 460, Y: 720}

	// Multi-component numeric inputs.
	w.AddWidget(
		text.SeparatorText("Vector inputs"),
		input.NewFloat2("float2", &s.f2),
		input.NewFloat3("float3", &s.f3),
		input.NewFloat4("float4", &s.f4),
		input.NewInt2("int2", &s.i2),
		input.NewInt3("int3", &s.i3),
		input.NewInt4("int4", &s.i4),
		input.NewDouble("double", &s.dbl),
	)

	// Multi-component and special sliders.
	w.AddWidget(
		text.SeparatorText("Sliders"),
		input.NewSliderFloat3("slider float3", &s.sf3, 0, 1),
		input.NewSliderInt2("slider int2", &s.si2, 0, 10),
		input.NewSliderAngle("angle", &s.angle),
		input.NewVSliderFloat("vfloat", imgui.Vec2{X: 24, Y: 100}, &s.vsf, 0, 1),
		layout.SameLine(),
		input.NewVSliderInt("vint", imgui.Vec2{X: 24, Y: 100}, &s.vsi, 0, 10),
	)

	// Multi-component and range drags.
	w.AddWidget(
		text.SeparatorText("Drags"),
		input.NewDragFloat2("drag float2", &s.df2),
		input.NewDragInt3("drag int3", &s.di3),
		input.NewDragFloatRange("float range", &s.rangeLo, &s.rangeHi),
		input.NewDragIntRange("int range", &s.irangeLo, &s.irangeHi),
	)

	// Generic scalar widgets (one of each kind).
	u64 := input.NewScalar[uint64]("u64", &s.scalarU64)
	u64.Format = "%llu"
	u64.Step = 16
	w.AddWidget(
		text.SeparatorText("Scalar (generic)"),
		u64,
		input.NewScalarN[float64]("doubles", s.scalarN),
		input.NewSliderScalar[float32]("slider scalar", &s.ssF, 0, 10),
		input.NewSliderScalarN[int32]("slider scalarN", s.ssN, 0, 10),
		input.NewDragScalar[int32]("drag scalar", &s.dsI),
		input.NewDragScalarN[float32]("drag scalarN", s.dsN),
		input.NewVSliderScalar[float32]("vscalar", imgui.Vec2{X: 24, Y: 100}, &s.vssF, 0, 1),
	)

	// Flag checkboxes over a shared bitfield.
	w.AddWidget(
		text.SeparatorText("Flag checkboxes"),
		input.NewCheckboxFlags("flag A", &s.flags, flagA),
		input.NewCheckboxFlags("flag B", &s.flags, flagB),
		input.NewCheckboxFlags("flag C", &s.flags, flagC),
		input.NewCheckboxFlags("A | C", &s.flags, flagA|flagC),
		input.NewCheckboxFlagsUint("unsigned bit 0", &s.uflags, 1),
		text.NewLabelTextf("flags", "0x%X", s.flags),
	)

	// Buttons: invisible region and standalone bullet.
	invisible := button.NewInvisible("##canvasbtn", imgui.Vec2{X: 120, Y: 32})
	invisible.OnClick = func() { s.invisibleClicks++ }
	w.AddWidget(
		text.SeparatorText("Buttons"),
		text.New("invisible button (120x32) below:"),
		invisible,
		text.Labelf("invisible clicks = %d", s.invisibleClicks),
		button.NewBullet(), layout.SameLine(), text.New("beside a bullet glyph"),
	)

	// Container combo and list with arbitrary bodies.
	customCombo := combo.NewCustom("custom combo", "pick options")
	for i := range s.customSel {
		sel := combo.NewSelectable(fmt.Sprintf("option %c", 'A'+i))
		sel.Selected = &s.customSel[i]
		customCombo.AddWidget(sel)
	}
	customList := combo.NewCustomList("custom list")
	customList.Size = imgui.Vec2{Y: 80}
	for i := range s.listSel {
		sel := combo.NewSelectable(fmt.Sprintf("row %d", i))
		sel.Selected = &s.listSel[i]
		customList.AddWidget(sel)
	}
	w.AddWidget(text.SeparatorText("Container combo / list"), customCombo, customList)

	// Context popups and the label/value text widget.
	winCtx := popup.NewContextWindow("##winctx")
	winCtx.AddWidget(combo.NewSelectable("window context action"))
	voidCtx := popup.NewContextVoid("##voidctx")
	voidCtx.AddWidget(combo.NewSelectable("void context action"))
	w.AddWidget(
		text.SeparatorText("Context popups"),
		text.New("right-click this window for its context menu"),
		text.NewLabelText("label", "value pair"),
		winCtx, voidCtx,
	)

	// Tab bar with a non-tab button, and rich tooltips.
	bars := tab.NewBar("covtabs")
	first := tab.NewItem("Tab")
	first.AddWidget(text.New("a normal tab body"))
	btn := tab.NewItemButton("+")
	btn.OnClick = func() { s.tabButtonClicks++ }
	bars.AddWidget(first, btn)
	rich := tooltip.ForWidgets(
		button.New("hover for rich tooltip"),
		text.New("tooltips can hold"),
		text.Colored(imgui.Color{R: 1, G: 0.7, B: 0.2, A: 1}, "arbitrary widgets"),
		&button.ProgressBar{Fraction: 0.6},
	)
	tabButtonClicks := text.Labelf("tab-button clicks = %d", s.tabButtonClicks)
	togglePanel := input.NewCheckbox("show raw tooltip panel", &s.showPanel)
	coverWidgets := []imgui.Widget{
		text.SeparatorText("Tabs & tooltips"),
		bars, tabButtonClicks, rich, togglePanel,
	}
	if s.showPanel {
		coverWidgets = append(coverWidgets, tooltip.NewPanel(text.New("raw tooltip panel (follows the cursor)")))
	}
	w.AddWidget(coverWidgets...)

	// Table with frozen header/column and angled headers.
	grid := table.New("freeze grid")
	grid.Size = imgui.Vec2{Y: 140}
	grid.FreezeCols = 1
	grid.FreezeRows = 1
	grid.SetVerticalScroll(true)
	name := grid.AddColumn("name")
	name.SetSizing(table.SizingFixed)
	name.Width = 80
	for _, h := range []string{"q1", "q2", "q3", "q4"} {
		c := grid.AddColumn(h)
		c.SetAngledHeader(true)
	}
	for r := range 20 {
		grid.AddRow(
			text.Labelf("row %d", r),
			text.Labelf("%d", r*1),
			text.Labelf("%d", r*2),
			text.Labelf("%d", r*3),
			text.Labelf("%d", r*4),
		)
	}
	w.AddWidget(text.SeparatorText("Table freeze + angled headers"), grid)

	return w
}

// checker builds a small RGBA image to upload as a texture.
func checker(w, h int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := range h {
		for x := range w {
			i := img.PixOffset(x, y)
			if (x/8+y/8)%2 == 0 {
				img.Pix[i+0] = byte(255 * x / w)
				img.Pix[i+1] = byte(255 * y / h)
				img.Pix[i+2] = 64
			} else {
				img.Pix[i+0] = 64
				img.Pix[i+1] = byte(255 * x / w)
				img.Pix[i+2] = byte(255 * y / h)
			}
			img.Pix[i+3] = 255
		}
	}
	return img
}
