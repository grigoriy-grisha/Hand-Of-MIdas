package main

import (
	"awesomeProject/pkg/HOM"
	"awesomeProject/pkg/HOMF"
)

var longText = "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos  Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quosLorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos, saepe sapiente vero! Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos, saepe sapiente vero! Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos, saepe sapiente vero! Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos, saepe sapiente vero! Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos, saepe sapiente vero! Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos, saepe sapiente vero! Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos, saepe sapiente vero! Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos, saepe sapiente vero! Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos, saepe sapiente vero! Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos, saepe sapiente vero! 	Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos, saepe sapiente vero! 	Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos, saepe sapiente vero!"

func main() {
	homf := HOMF.NewHOMFramework(100, 30)

	closeHomf := homf.Init()
	defer closeHomf()

	homf.Mount(
		HOM.NewHOMElement(
			HOM.NewElementParams{
				Style: &HOM.Style{
					VerticalContent: HOM.VerticalContentBottom,
					AlignContent:    HOM.AlignContentLeft,
					PaddingRight:    1,
					Border:          true,
				},
				Text: &HOM.Text{Value: "hello world"},
			}))

	homf.Run()

}
