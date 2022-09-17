package main

import (
	"fmt"
	"github.com/phpisfirstofworld/image"
)

func main() {

	i := image.NewImage()

	_, err := i.LoadImage("img.png")

	if err != nil {

		fmt.Printf("err :%+v\n", err)

		return
	}

	//i.GetSize().

	fmt.Println(i.GetSize().Kb())

	//fmt.Println(res.GetSourceImageResource().Bounds().)

	//fmt.Println(res.ResizeWidth(200).Save("img_save.png"))
	//fmt.Println(res.ResizePercent(50).Save("img_save.png"))

}
