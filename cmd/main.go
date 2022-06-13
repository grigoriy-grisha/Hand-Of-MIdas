package main

import (
	"awesomeProject/pkg/HOM"
	"awesomeProject/pkg/HOMF"
)

func main() {
	homf := HOMF.NewHOMFramework(100, 30)

	closeHomf := homf.Init()
	defer closeHomf()

	homf.Mount(
		HOM.NewHOMElement(
			HOM.NewElementParams{
				Style: &HOM.Style{
					VerticalContent: HOM.VerticalContentBottom,
					AlignContent:    HOM.AlignContentCenter,
					PaddingRight:    1,
					Border:          true,
				},
				Text: &HOM.Text{Value: "hello"},
			}))

	homf.Run()

}
