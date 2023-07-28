package main



import(
  "fmt"
  "math"
)


type Tuple struct{
  X float64
  Y float64
  Z float64
  W float64
}



func Init(x,y,z,w float64)Tuple{
  return Tuple{x,y,z,w}
}

func Vector(x,y,z float64)Tuple{
return Init(x,y,z,0)
}

func Point(x,y,z float64) Tuple{

  return Init(x,y,z,1)
}


func Add(a,b Tuple) Tuple{

return Init(a.X + b.X,a.Y+b.Y,a.Z+b.Z,a.W+b.W)
}

func Sub(a,b Tuple)Tuple{
  return Init(a.X - b.X, a.Y - b.Y, a.Z - b.Z,a.W - b.W)
}

func Magnitude(a Tuple) float64{
return math.Sqrt(a.X * a.X + a.Y* a.Y + a.Z*a.Z + a.W* a.W)
}

func Normalization(a Tuple)Tuple{
  return Init(a.X / Magnitude(a),a.Y/Magnitude(a), a.Z / Magnitude(a),a.W / Magnitude(a))
}

func DotProduct(a,b Tuple)float64{
  return a.X * b.X + a.Y*b.Y + a.Z*b.Z + a.W*b.W
}

func CrossProduct(a,b Tuple)Tuple{
return Vector(a.Y * b.Z - a.Z * b.Y,
            a.Z * b.X - a.X * b.Z,
            a.X * b.Y - a.Y * b.X)
}


type Projectile struct{
  position Tuple // Point
  velocity Tuple // Vector
}

type Enviromnent struct{
  gravity Tuple // vector
  wind Tuple // vector
}

func(a * Projectile)InitProject(p,v Tuple) Projectile{
  a.position = p
  a.velocity = v

  return *a
}

func(a* Enviromnent)InitEnv(v1,v2 Tuple)Enviromnent{
  a.gravity = v1
  a.wind = v2

  return *a
}

func Tick(proj* Projectile, env* Enviromnent) Projectile{
  position:= Add(proj.position,proj.velocity)
  velocity:= Add(Add(proj.velocity,env.gravity),env.wind)

  return proj.InitProject(position,velocity)

}


func main(){

  proj:= Projectile{}
  proj.InitProject(Point(0,1,0),Normalization(Vector(1,1,0)))
  
  

  env:= Enviromnent{}
  env.InitEnv(Vector(0,-0.1,0),Vector(-0.01,0,0))




  
  for i:= 1;i<100;i++{
    tick:= Tick(&proj,&env)
    fmt.Println(tick.position)
    if tick.position.Y <=0{break}
}


  }






