package ray

import (
	"math"
	"uka/util"
)

type Ray struct{
Orig  vec3.Vec3
Dir   vec3.Vec3

}

func hitSphere(center vec3.Vec3, radius float64,r* Ray)float64{
    
  oc:= r.Orig.Sub(center)
  a:= r.Dir.LengthWithoutRoot()
  b:= oc.Dot(r.Dir)
  c:= oc.LengthWithoutRoot() - radius * radius
  discriminant:=b*b - a*c
  if discriminant<0{
    return -1.0
  }else{
    calc:= ((0-b) - math.Sqrt(discriminant)) / (2.0*a)
    return  calc
  }
  
 }

func RayColor(r* Ray) vec3.Vec3{
  center:=vec3.Init(0,0,-1)
  //mid circle
  t:=hitSphere(center,0.5,r)
    if(t>0.0){
      x:=r.At(t).Sub(vec3.Init(0,0,-1)).UnitVector()
          //return *x.Mult(0.5).Add(vec3.Init(1,1,1))
          color:=vec3.Init(x.X+1,x.Y+1,x.Z+1).Mult(0.5)
          return *color
  }
    

  //ud := vec3.UnitVector(r.Dir)
  ud:=r.Dir.UnitVector()
  t = 0.5*(ud.Y + 1.0)
   
  color:=vec3.Init(1,1,1).Mult(1.0-t).Add(*vec3.Init(0.5,0.7,1.0).Mult(t))

    return  *color
}


func(v* Ray)Init(orig, dir vec3.Vec3){
  v.Orig = orig
  v.Dir = dir
}


func(v Ray)At(t float64) *vec3.Vec3{
  
   return v.Orig.Add(*v.Dir.Mult(t))
}
