package vec3

import (
	"math"
)



type Vec3 struct{
  X float64
  Y float64
  Z float64
}

func Init(x,y,z float64)Vec3{
  return Vec3{X:x, Y:y, Z:z}
}








func(v* Vec3)Length() float64{
  //a^2 + b^2 + c^2 = d^2 pyt
  len:= math.Pow(v.X,2)+ math.Pow(v.Y,2) + math.Pow(v.Z,2)
  return math.Sqrt(len)
}

func(v* Vec3)Normalize(){
 len:=v.Length()
  v.X = v.X / len
  v.Y = v.Y / len
  v.Z = v.Z / len
}

func(v* Vec3)UnitVector() Vec3{
 len:=v.Length()
  v.X = v.X / len
  v.Y = v.Y / len
  v.Z = v.Z / len

  return *v
}



func(v* Vec3)Dot(b Vec3) float64{
  return (v.X * b.X) + (v.Y * b.Y) + (v.Z * b.Z)
}

func(a* Vec3)Cross(b Vec3)Vec3{
  cross:=Init(
    a.Y * b.Z - a.Z * b.Y,
    a.Z * b.X - a.X * b.Z,
    a.X * b.Y - a.Y * b.X,
    )

  return cross
}


func(a Vec3)Add(b Vec3)Vec3{
  a.X += b.X
  a.Y += b.Y
  a.Z += b.Z

  return a
}

func(a Vec3)Mult(t float64)Vec3{
  a.X *= t
  a.Y *= t
  a.Z *= t

  return a
}

func(a Vec3)Div(t float64)Vec3{
  return a.Mult(1/t)
}

func(a Vec3)Sub(b Vec3) Vec3{

  a.X -= b.X
  a.Y -= b.Y
  a.Z -= b.Z

  return a
}


