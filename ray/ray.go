package ray

import (
	"uka/util"
)

type Ray struct{
Orig  vec3.Vec3
Dir   vec3.Vec3

}

func RayColor(r* Ray) vec3.Vec3{
    
  ud:= r.Dir.UnitVector()
  t:= 0.5*(ud.Y + 1.0)
    
  color:=vec3.Init(1.0,1.0,1.0).Add(vec3.Init(0.5,0.7,1.0).Mult(t)).Mult(1.0-t)

  return  color
}


func(v* Ray)Init(orig, dir vec3.Vec3){
  v.Orig = orig
  v.Dir = dir
}


func(v* Ray)At(t float64) vec3.Vec3{
  
  mult:= v.Dir.Mult(t)
  add:= v.Orig.Add(mult)

  return add
}
