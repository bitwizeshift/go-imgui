// Command demo opens a window and exercises the high-level widget packages. It
// builds a fresh widget tree each frame from persistent state and displays it;
// it never touches the internal cimgui bindings.
package main

import (
	"fmt"
	"image"
	"log"

	"rodusek.dev/pkg/imgui"
	"rodusek.dev/pkg/imgui/app"
	"rodusek.dev/pkg/imgui/button"
	"rodusek.dev/pkg/imgui/color"
	"rodusek.dev/pkg/imgui/combo"
	"rodusek.dev/pkg/imgui/debug"
	"rodusek.dev/pkg/imgui/input"
	"rodusek.dev/pkg/imgui/layout"
	"rodusek.dev/pkg/imgui/menu"
	"rodusek.dev/pkg/imgui/plot"
	"rodusek.dev/pkg/imgui/popup"
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
		showDemo         = true
		counter  int32   = 0
		enabled          = true
		choice   int32   = 0
		speed    float32 = 0.4
		count    int32   = 3
		drag     float32 = 1
		amount   int32   = 5
		name             = "edit me"
		notes            = "multi\nline\ntext"
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

	err := app.Run(app.Config{Title: "go-imgui high-level demo"}, func() {
		if tex == nil { // create lazily: needs a current GL context
			tex = texture.FromImage(checker(64, 64))
		}
		progress += 0.004
		if progress > 1 {
			progress = 0
		}

		imgui.Display(
			mainMenu(&showDemo),
			gallery(galleryState{
				counter: &counter, enabled: &enabled, choice: &choice,
				speed: &speed, count: &count, drag: &drag, amount: &amount,
				name: &name, notes: &notes, rgb: &rgb, rgba: &rgba,
				comboItems: comboItems, comboIdx: &comboIdx, listIdx: &listIdx,
				selected: selected[:], plotData: plotData, progress: progress,
				tabOpen: &tabOpen, tex: tex,
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

	// Escape hatch.
	w.AddWidget(imgui.Custom(func() {
		// Anything not yet modelled can run here against the lower-level API.
	}))

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
