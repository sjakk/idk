package testing

import (
	"math"
	"testing"
	"uka/ray"
	_ "uka/ray"
	"uka/util"
)


func RoundPlaces(v float64)float64{
  return math.Round(v*1000)/1000
}


func TestArithmetic(t* testing.T){
t.Run("Unit Vector",func(t *testing.T) {

    input:= vec3.Init(10.0,10.0,0.0)
    got:= input.UnitVector()


    expected:=vec3.Init(input.X/input.Length(),input.Y/input.Length(),input.Z/input.Length())
    
    got.X = RoundPlaces(got.X)
    got.Y = RoundPlaces(got.Y)
    got.Z = RoundPlaces(got.Z)
    
    expected.X = RoundPlaces(expected.X)
    expected.Y = RoundPlaces(expected.Y)
    expected.Z = RoundPlaces(expected.Z)


  if got!=expected{
      t.Errorf("Unit Vector Error: Got:\t%v\n\tWant:\t%v\n",got,expected)
    }

})

t.Run("Length",func(t *testing.T) {

    input:=vec3.Init(10,10,10)
    got:=input.Length()
    want:=math.Sqrt(300)
    
      if got!=want{
      t.Errorf("Length\n Got:\t%f\nWant: \t%f",got,want)
    }


     })
}


func TestLLC(t* testing.T){
  focal_length:= 1.0
  vpx:= (16.0/9.0)*2.0
  horizontal:= vec3.Init(vpx,0.0,0.0) 
  vertical:= vec3.Init(0.0,2.0,0)
  origin:=vec3.Init(0.0,0.0,0.0)
  
  llcgot:= origin.Sub(*horizontal.Div(2)).Sub(*vertical.Div(2).Sub(vec3.Init(0,0,focal_length)))
  
  llcwant:=vec3.Init(1.775,1.0,0.5)
  
  if *llcgot!=llcwant{
    t.Errorf("LLC\nGot:\t%v\nWant:\t%v\n",llcgot,llcwant)
  }
  
t.Run("color draw",func(t *testing.T) {
    
  raygot:=ray.Ray{}
  llc1:=llcgot.Add(*horizontal.Mult(1))
  llc2:= vertical.Mult(254).Sub(origin)
  
  raygot.Init(origin,*llcgot.Add(*horizontal.Mult(1)).Add(*vertical.Mult(254)).Sub(origin))
  raywant:= llc1.Add(*llc2)
  


    if raygot.Dir!=*raywant{
      t.Errorf("Color Draw:\nGot:\t%v\nWant:\t%v\n",raygot,raywant)
    }

  })
}














