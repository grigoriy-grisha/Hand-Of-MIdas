package main

import (
	"awesomeProject/internal"
	"awesomeProject/pkg/HOMF"
)

func main() {
	homf := HOMF.NewHOMFramework(100, 30)

	closeHomf := homf.Init()
	defer closeHomf()

	homf.Mount(internal.AppElements)

	homf.Run()

}
