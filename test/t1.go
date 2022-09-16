package main

import (
	"fmt"
	"github.com/phpisfirstofworld/image"
)

func main() {

	i := image.NewImage()

	res, err := i.LoadImage("img.png")

	if err != nil {

		fmt.Printf("err :%+v\n", err)

		return
	}

	//fmt.Println(res.ResizePercent(50).Save("img_save.png"))

	fmt.Printf("err :%+v\n", res.ResizePercent(50).OverSave())

}
