package main

import (
	"fmt"
	"os"
  "uka/util"
)

const(
  image_width = 256
  image_height = 256
)

func main(){

fmt.Fprintf(os.Stdout,"P3\n%d %d\n255\n",image_width,image_height)
  for j:= image_height-1;j>=0;j--{
    for i:=0;i<image_width;i++{
      
      color:=vec3.Init(float64(i)/(image_width-1),float64(j)/(image_height-1),0.25)
      
        color.WriteColor()
    }
  }



}
