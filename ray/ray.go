package ray

import (
	"uka/util"
)

type Ray struct{
Orig  vec3.Vec3
Dir   vec3.Vec3

}

func hitSphere(center vec3.Vec3,radius float64,r* Ray)bool{
    
  oc:= r.Orig.Sub(center)
  a:= r.Dir.Dot(r.Dir)
  b:= 2.0 * oc.Dot(r.Dir)
  c:= oc.Dot(*oc) - radius*radius
  discriminant:=b*b - 4*a*c

  return discriminant>0
}

func RayColor(r* Ray) vec3.Vec3{
    
  if(hitSphere(vec3.Init(0,0,-1),0.5,r)){
    return vec3.Init(1,0,0)
  }
  


  ud:= r.Dir.UnitVector()
  t:= 0.5*(ud.Y + 1.0)
    
  ad:=vec3.Init(0.5,0.5,0.25)

  color:=vec3.Init(1.0,1.0,1.0).Add(*ad.Mult(t)).Mult(1.0-t)

  return  *color
}


func(v* Ray)Init(orig, dir vec3.Vec3){
  v.Orig = orig
  v.Dir = dir
}


func(v* Ray)At(t float64) vec3.Vec3{
  
  mult:= v.Dir.Mult(t)
  add:= v.Orig.Add(*mult)

  return *add
}
