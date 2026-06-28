// Command demo opens a native window and drives the full internal/cimgui widget
// surface so each binding can be verified visually. The frame loop is provided by
// app.Run; every widget call goes directly through the internal cimgui package.
package main

import (
	"log"

	"rodusek.dev/pkg/imgui/app"
	"rodusek.dev/pkg/imgui/internal/cimgui"
)

func main() {
	// Persistent widget state.
	var (
		showDemo    = true
		showMetrics = false
		counter     int32

		checkA   = true
		radio    int32
		flagBits int32 = 0b001

		sliderF float32 = 0.42
		sliderI int32   = 3
		angle   float32 = 0.5
		dragF   float32 = 1.0
		vec3            = [3]float32{0.1, 0.2, 0.3}

		inputF  float32 = 1.5
		inputI  int32   = 7
		nameBuf         = makeBuf("edit me", 128)
		noteBuf         = makeBuf("multiline\ntext", 1024)
		hintBuf         = makeBuf("", 128)

		col3  = [3]float32{0.4, 0.7, 0.2}
		col4  = [4]float32{0.2, 0.4, 0.8, 1.0}
		pick3 = [3]float32{0.9, 0.3, 0.3}
		pick4 = [4]float32{0.3, 0.6, 0.9, 1.0}

		vsliderV  int32  = 5
		disabled  bool   = false
		comboStr  string = "Apple\x00Banana\x00Cherry\x00"
		comboStrI int32

		comboItems = []string{"Red", "Green", "Blue", "Alpha"}
		comboIdx   int32
		listItems  = []string{"One", "Two", "Three", "Four", "Five"}
		listIdx    int32
		selected   = []bool{false, false, false}

		plotData = []float32{0.1, 0.5, 0.2, 0.8, 0.4, 0.9, 0.3, 0.7, 0.6, 1.0}

		progress float32

		menuChecked  = true
		closeableTab = true
		showModal    = false

		// A user texture, lazily created on the first frame (a GL context must be
		// current, which it is once app.Run is running).
		userTex      cimgui.TextureRef
		userTexReady = false
	)

	err := app.Run(app.Config{Title: "go-imgui cimgui demo"}, func() {
		if !userTexReady {
			userTex = cimgui.CreateTextureRGBA(64, 64, checkerRGBA(64, 64))
			userTexReady = true
		}

		// ---- Main menu bar -------------------------------------------------
		if cimgui.BeginMainMenuBar() {
			if cimgui.BeginMenu("File", true) {
				cimgui.MenuItem_Bool("New", "Ctrl+N", false, true)
				cimgui.MenuItem_BoolPtr("Show metrics", "", &showMetrics, true)
				cimgui.EndMenu()
			}
			if cimgui.BeginMenu("View", true) {
				cimgui.MenuItem_BoolPtr("Demo window", "", &showDemo, true)
				cimgui.EndMenu()
			}
			cimgui.EndMainMenuBar()
		}

		if showDemo {
			cimgui.ShowDemoWindow(&showDemo)
		}
		if showMetrics {
			cimgui.ShowMetricsWindow(&showMetrics)
		}

		// ---- Main panel ----------------------------------------------------
		cimgui.SetNextWindowSize(cimgui.Vec2{X: 480, Y: 640}, cimgui.CondFirstUseEver)
		cimgui.Begin("cimgui widget gallery", nil, cimgui.WindowFlagsMenuBar)

		// Window menu bar.
		if cimgui.BeginMenuBar() {
			if cimgui.BeginMenu("Options", true) {
				cimgui.MenuItem_BoolPtr("Enabled", "", &menuChecked, true)
				cimgui.EndMenu()
			}
			cimgui.EndMenuBar()
		}

		// Text variants.
		cimgui.SeparatorText("Text")
		cimgui.TextUnformatted("Plain unformatted text.")
		cimgui.TextColored(cimgui.Vec4{X: 1, Y: 0.6, Z: 0.2, W: 1}, "Colored text.")
		cimgui.TextDisabled("Disabled text.")
		cimgui.TextWrapped("Wrapped text that should fold onto multiple lines when the window is narrow enough to force a break.")
		cimgui.LabelText("label", "value")
		cimgui.BulletText("A bulleted line.")

		// Buttons.
		cimgui.SeparatorText("Buttons")
		if cimgui.Button("Click me", cimgui.Vec2{}) {
			counter++
		}
		cimgui.SameLine(0, -1)
		cimgui.SmallButton("small")
		cimgui.SameLine(0, -1)
		cimgui.ArrowButton("##left", cimgui.DirLeft)
		cimgui.SameLine(0, -1)
		cimgui.ArrowButton("##right", cimgui.DirRight)
		cimgui.LabelText("clicks", itoa(counter))
		if cimgui.IsItemHovered(cimgui.HoveredFlagsNone) {
			cimgui.SetTooltip("number of clicks so far")
		}
		progress += 0.004
		if progress > 1 {
			progress = 0
		}
		cimgui.ProgressBar(progress, cimgui.Vec2{}, "")

		// Selection widgets.
		cimgui.SeparatorText("Selection")
		cimgui.Checkbox("checkbox", &checkA)
		cimgui.CheckboxFlags_IntPtr("bit 0", &flagBits, 0b001)
		cimgui.SameLine(0, -1)
		cimgui.CheckboxFlags_IntPtr("bit 1", &flagBits, 0b010)
		cimgui.RadioButton_IntPtr("A", &radio, 0)
		cimgui.SameLine(0, -1)
		cimgui.RadioButton_IntPtr("B", &radio, 1)

		// Sliders / drags / inputs.
		cimgui.SeparatorText("Sliders, drags & inputs")
		cimgui.SliderFloat("slider f", &sliderF, 0, 1, "%.3f", cimgui.SliderFlagsNone)
		cimgui.SliderInt("slider i", &sliderI, 0, 10, "%d", cimgui.SliderFlagsNone)
		cimgui.SliderAngle("angle", &angle, -180, 180, "%.0f deg", cimgui.SliderFlagsNone)
		cimgui.DragFloat("drag f", &dragF, 0.01, 0, 0, "%.3f", cimgui.SliderFlagsNone)
		cimgui.DragFloat3("vec3", &vec3, 0.01, 0, 1, "%.2f", cimgui.SliderFlagsNone)
		cimgui.InputFloat("input f", &inputF, 0.1, 1, "%.3f", cimgui.InputTextFlagsNone)
		cimgui.InputInt("input i", &inputI, 1, 10, cimgui.InputTextFlagsNone)
		cimgui.InputText("text", nameBuf, cimgui.InputTextFlagsNone)
		cimgui.InputTextWithHint("hinted", "type here...", hintBuf, cimgui.InputTextFlagsNone)
		cimgui.InputTextMultiline("multiline", noteBuf, cimgui.Vec2{X: 0, Y: 60}, cimgui.InputTextFlagsNone)

		// Color.
		cimgui.SeparatorText("Color")
		cimgui.ColorEdit3("rgb", &col3, cimgui.ColorEditFlagsNone)
		cimgui.ColorEdit4("rgba", &col4, cimgui.ColorEditFlagsAlphaBar)
		cimgui.ColorButton("swatch", cimgui.Vec4{X: col4[0], Y: col4[1], Z: col4[2], W: col4[3]}, cimgui.ColorEditFlagsNone, cimgui.Vec2{X: 40, Y: 20})
		if cimgui.TreeNode_Str("Color pickers") {
			cimgui.ColorPicker3("picker3", &pick3, cimgui.ColorEditFlagsNone)
			cimgui.ColorPicker4("picker4", &pick4, cimgui.ColorEditFlagsAlphaBar, nil)
			cimgui.TreePop()
		}

		// Combo / list box.
		cimgui.SeparatorText("Combo & list box")
		cimgui.Combo_Str_arr("combo (arr)", &comboIdx, comboItems, -1)
		cimgui.Combo_Str("combo (zero-sep)", &comboStrI, comboStr, -1)
		cimgui.ListBox_Str_arr("list", &listIdx, listItems, 4)
		if cimgui.BeginCombo("manual combo", comboItems[comboIdx], cimgui.ComboFlagsNone) {
			for i, item := range comboItems {
				if cimgui.Selectable_Bool(item, int32(i) == comboIdx, cimgui.SelectableFlagsNone, cimgui.Vec2{}) {
					comboIdx = int32(i)
				}
			}
			cimgui.EndCombo()
		}
		if cimgui.BeginListBox("manual list", cimgui.Vec2{X: 0, Y: 60}) {
			for i, item := range listItems {
				if cimgui.Selectable_Bool(item, int32(i) == listIdx, cimgui.SelectableFlagsNone, cimgui.Vec2{}) {
					listIdx = int32(i)
				}
			}
			cimgui.EndListBox()
		}

		// Selectables.
		cimgui.SeparatorText("Selectables")
		for i := range selected {
			cimgui.Selectable_BoolPtr("item "+itoa(int32(i)), &selected[i], cimgui.SelectableFlagsNone, cimgui.Vec2{})
		}

		// Trees & headers.
		cimgui.SeparatorText("Trees & headers")
		if cimgui.TreeNode_Str("Tree root") {
			cimgui.BulletText("child A")
			if cimgui.TreeNodeEx_Str("Branch", cimgui.TreeNodeFlagsDefaultOpen) {
				cimgui.BulletText("leaf")
				cimgui.TreePop()
			}
			cimgui.TreePop()
		}
		if cimgui.CollapsingHeader_TreeNodeFlags("Collapsing header", cimgui.TreeNodeFlagsNone) {
			cimgui.TextUnformatted("header body")
		}

		// Child region.
		cimgui.SeparatorText("Child region")
		if cimgui.BeginChild_Str("scroller", cimgui.Vec2{X: 0, Y: 60}, cimgui.ChildFlagsBorders, cimgui.WindowFlagsNone) {
			for i := range 10 {
				cimgui.TextUnformatted("scrolling line " + itoa(int32(i)))
			}
		}
		cimgui.EndChild()

		// Tabs. The first two tabs simply switch; the third is closeable (it has
		// a close button because it is passed a *bool) and can be reopened with
		// the trailing "+" tab button.
		cimgui.SeparatorText("Tabs")
		if cimgui.BeginTabBar("tabs", cimgui.TabBarFlagsReorderable) {
			if cimgui.BeginTabItem("First", nil, cimgui.TabItemFlagsNone) {
				cimgui.TextUnformatted("first tab body")
				cimgui.EndTabItem()
			}
			if cimgui.BeginTabItem("Settings", nil, cimgui.TabItemFlagsNone) {
				cimgui.TextUnformatted("settings tab body")
				cimgui.EndTabItem()
			}
			if closeableTab {
				if cimgui.BeginTabItem("Closeable", &closeableTab, cimgui.TabItemFlagsNone) {
					cimgui.TextUnformatted("close me with the x, reopen with +")
					cimgui.EndTabItem()
				}
			}
			if cimgui.TabItemButton("+", cimgui.TabItemFlagsTrailing) {
				closeableTab = true
			}
			cimgui.EndTabBar()
		}

		// Vertical slider and a disabled block.
		cimgui.SeparatorText("Vertical slider & disabled block")
		cimgui.VSliderInt("##v", cimgui.Vec2{X: 24, Y: 80}, &vsliderV, 0, 10, "%d", cimgui.SliderFlagsNone)
		cimgui.SameLine(0, -1)
		cimgui.Checkbox("disable widgets below", &disabled)
		cimgui.BeginDisabled(disabled)
		cimgui.Button("maybe disabled", cimgui.Vec2{})
		cimgui.TextLinkOpenURL("dearimgui.com", "https://www.dearimgui.com")
		cimgui.EndDisabled()

		// Plots.
		cimgui.SeparatorText("Plots")
		cimgui.PlotLines_FloatPtr("lines", plotData, 0, "", 0, 1, cimgui.Vec2{X: 0, Y: 50}, 4)
		cimgui.PlotHistogram_FloatPtr("histogram", plotData, 0, "", 0, 1, cimgui.Vec2{X: 0, Y: 50}, 4)

		// Images. userTex is a 64x64 RGBA texture built from raw pixels in Go and
		// uploaded with cimgui.CreateTextureRGBA. The font atlas is also a texture
		// and can be drawn the same way.
		cimgui.SeparatorText("Images")
		cimgui.TextUnformatted("user texture (created from RGBA pixels):")
		cimgui.Image(userTex, cimgui.Vec2{X: 64, Y: 64}, cimgui.Vec2{}, cimgui.Vec2{X: 1, Y: 1})
		cimgui.SameLine(0, -1)
		if cimgui.ImageButton("imgbtn", userTex, cimgui.Vec2{X: 48, Y: 48}, cimgui.Vec2{}, cimgui.Vec2{X: 1, Y: 1},
			cimgui.Vec4{}, cimgui.Vec4{X: 1, Y: 1, Z: 1, W: 1}) {
			counter++
		}
		cimgui.TextUnformatted("font atlas texture:")
		cimgui.Image(cimgui.FontAtlasTexRef(), cimgui.Vec2{X: 128, Y: 64}, cimgui.Vec2{}, cimgui.Vec2{X: 1, Y: 1})

		// Tables.
		cimgui.SeparatorText("Table")
		if cimgui.BeginTable("grid", 3, cimgui.TableFlagsBorders|cimgui.TableFlagsRowBg, cimgui.Vec2{}, 0) {
			cimgui.TableSetupColumn("Name", cimgui.TableColumnFlagsNone, 0, 0)
			cimgui.TableSetupColumn("Value", cimgui.TableColumnFlagsNone, 0, 0)
			cimgui.TableSetupColumn("Note", cimgui.TableColumnFlagsNone, 0, 0)
			cimgui.TableHeadersRow()
			for r := range int32(3) {
				cimgui.TableNextRow(cimgui.TableRowFlagsNone, 0)
				cimgui.TableSetColumnIndex(0)
				cimgui.TextUnformatted("row " + itoa(r))
				cimgui.TableSetColumnIndex(1)
				cimgui.TextUnformatted(itoa(r * 10))
				cimgui.TableSetColumnIndex(2)
				cimgui.TextUnformatted("ok")
			}
			cimgui.EndTable()
		}

		// Popups & modals.
		cimgui.SeparatorText("Popups")
		if cimgui.Button("Open modal", cimgui.Vec2{}) {
			showModal = true
			cimgui.OpenPopup_Str("Modal", cimgui.PopupFlagsNone)
		}
		if cimgui.BeginPopupModal("Modal", &showModal, cimgui.WindowFlagsAlwaysAutoResize) {
			cimgui.TextUnformatted("This is a modal popup.")
			if cimgui.Button("Close", cimgui.Vec2{}) {
				cimgui.CloseCurrentPopup()
			}
			cimgui.EndPopup()
		}
		cimgui.SameLine(0, -1)
		cimgui.SmallButton("right-click me")
		if cimgui.BeginPopupContextItem("ctx", cimgui.PopupFlagsMouseButtonRight) {
			cimgui.TextUnformatted("context menu")
			cimgui.Selectable_Bool("an action", false, cimgui.SelectableFlagsNone, cimgui.Vec2{})
			cimgui.EndPopup()
		}

		cimgui.End()
	})
	if err != nil {
		log.Fatal(err)
	}
}

// makeBuf returns a byte buffer of the given capacity seeded with s and a
// trailing NUL, suitable for the cimgui InputText* widgets.
func makeBuf(s string, size int) []byte {
	b := make([]byte, size)
	copy(b, s)
	return b
}

// checkerRGBA generates a w*h tightly-packed RGBA8 image: a colored checkerboard
// over a magenta/teal gradient, just to have real pixels to upload as a texture.
func checkerRGBA(w, h int) []byte {
	px := make([]byte, w*h*4)
	for y := range h {
		for x := range w {
			i := (y*w + x) * 4
			if (x/8+y/8)%2 == 0 {
				px[i+0] = byte(255 * x / w)
				px[i+1] = byte(255 * y / h)
				px[i+2] = 64
			} else {
				px[i+0] = 64
				px[i+1] = byte(255 * x / w)
				px[i+2] = byte(255 * y / h)
			}
			px[i+3] = 255
		}
	}
	return px
}

// itoa formats a signed 32-bit integer without importing strconv at the call site.
func itoa(v int32) string {
	if v == 0 {
		return "0"
	}
	neg := v < 0
	if neg {
		v = -v
	}
	var digits [12]byte
	i := len(digits)
	for v > 0 {
		i--
		digits[i] = byte('0' + v%10)
		v /= 10
	}
	if neg {
		i--
		digits[i] = '-'
	}
	return string(digits[i:])
}
