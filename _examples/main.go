package main

import (
	"github.com/AllenDang/imgui-go"
	"image/color"
	"time"

	"golang.org/x/image/colornames"

	"github.com/AllenDang/giu"
	animations "github.com/gucio321/giu-animations"
)

var easingAlg = animations.EasingAlgNone

func loop() {
	//a := int32(easingAlg)
	//animations.Animator(
	//	animations.Transition(
	//		func(starter func()) {
	giu.Window("window1").Layout(
		giu.Label("I'm a window 1"),
		animations.Animator(
			animations.ColorFlow(
				giu.Button("start transition").OnClick(func() {
					//starter()
				}),
				func() color.RGBA {
					return colornames.Red
				},
				func() color.RGBA {
					return colornames.Blue
				},
				giu.StyleColorButtonHovered,
				giu.StyleColorButton,
			),
		).
			Duration(time.Second).
			FPS(60).
			Trigger(animations.TriggerOnChange, imgui.IsItemHovered),
		//animations.Animator(
		//	animations.Move(func(starter animations.StarterFunc) giu.Widget {
		//		return giu.Child().Layout(
		//			giu.Row(
		//				giu.Label("Set easing alg:"),
		//				giu.SliderInt(&a, 0, int32(animations.EasingAlgMax-1)).Size(100).OnChange(func() {
		//					easingAlg = animations.EasingAlgorithmType(a)
		//				}),
		//			),
		//			giu.Button("move me!").OnClick(func() {
		//				starter(animations.PlayForward)
		//			}),
		//		).Size(200, 80)
		//	}, imgui.Vec2{X: 20, Y: 100}).
		//		StartPos(imgui.Vec2{X: 5, Y: 80}).
		//		Bezier(imgui.Vec2{X: 20, Y: 20}, imgui.Vec2{X: 90}),
		//).Duration(time.Second*3).FPS(120).EasingAlgorithm(easingAlg).Trigger(animations.TriggerOnTrue, giu.IsItemHovered),
	)
	//},
	//func(starter func()) {
	//	giu.Window("window2").Layout(
	//		giu.Label("I'm a window 1"),
	//		giu.Button("start transition").OnClick(func() {
	//			starter()
	//		}),
	//	)
	//},
	//),
	//).Build()
}

func main() {
	wnd := giu.NewMasterWindow("Animations presentation [example]", 640, 480, 0)
	wnd.Run(loop)
}
