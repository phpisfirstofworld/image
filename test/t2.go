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

	//fmt.Println(res.ResizeWidth(200).Save("img_save.png"))
	fmt.Println(res.ResizeHeight(200).Save("img_save.png"))

}
