# image

图片操作包

## 安装
```shell
go get github.com/phpisfirstofworld/image
```

## 设置图片尺寸
```go
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
        
	//尺寸缩放50%
	fmt.Println(res.ResizePercent(50).Save("img_save.png"))
    
	//尺寸缩放50%并覆盖保存
	//fmt.Printf("err :%+v\n", res.ResizePercent(50).OverSave())

}

```