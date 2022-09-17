package main

import (
	"fmt"
	"os"
)

func main() {

	f, _ := os.Open("gg.jpeg")

	//fmt.Println()

	info, _ := f.Stat()

	fmt.Println(info.Size() / 1024)

	//b, _ := ioutil.ReadAll(f)

	//i, _, _ := image.Decode(f)

	//fmt.Println(i.)

	//image.

	//fmt.Println(http.DetectContentType(b))

}
