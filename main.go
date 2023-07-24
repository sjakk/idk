package main

import (
	"fmt"
	"os"
  "uka/util"
  "uka/ray"
)

const(
  //image
  image_width = 400
  image_height = int(image_width/aspect_ratio)
  aspect_ratio = 16.0/9.0
)



func main(){

  viewport_height:= 2.0
  viewport_width:= aspect_ratio * viewport_height
  focal_length:= 1.0

  origin:= vec3.Init(0,0,0)
  horizontal:=vec3.Init(viewport_width,0,0)
  vertical:= vec3.Init(0,viewport_height,0)
    


  llc:= origin.Sub(horizontal.Div(2)).Sub(vertical.Div(2).Sub(vec3.Init(0,0,focal_length)))


fmt.Fprintf(os.Stdout,"P3\n%d %d\n255\n",image_width,image_height)
  for j:= image_height-1;j>=0;j--{
    for i:=0;i<image_width;i++{
      
      u:= float64(i)/(image_width-1)
      v:= float64(j)/(float64(image_height-1))  
      r:= ray.Ray{} 
      r.Init(origin,llc.Add(horizontal.Mult(u)).Add(vertical.Mult(v)).Sub(origin))
      pixelcolor:= ray.RayColor(&r)
      
      pixelcolor.WriteColor()
    }
  }



}
