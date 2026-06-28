// Command demo opens a native window and renders the Dear ImGui demo window
// alongside a small custom panel, using the public imgui and app packages.
package main

import (
	"log"

	"rodusek.dev/pkg/imgui"
	"rodusek.dev/pkg/imgui/app"
)

func main() {
	showDemo := true
	counter := 0
	slider := float32(0.5)

	err := app.Run(app.Config{Title: "go-imgui demo"}, func() {
		if showDemo {
			imgui.ShowDemoWindow(&showDemo)
		}

		imgui.Begin("Hello, go-imgui")
		imgui.Text("A thin Go wrapper around Dear ImGui.")
		imgui.Separator()
		if imgui.Button("Click me") {
			counter++
		}
		imgui.SameLine()
		imgui.Text("counter = %d", counter)
		imgui.SliderFloat("value", &slider, 0, 1)
		imgui.Checkbox("Show demo window", &showDemo)
		imgui.End()
	})
	if err != nil {
		log.Fatal(err)
	}
}
