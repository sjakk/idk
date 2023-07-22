package vec3

import (
	"fmt"
	"os"
)


func(color* Vec3)WriteColor(){
  fmt.Fprintf(os.Stdout, "%d %d %d\n",int(255.999*color.X),int(color.Y*255.999),int(color.Z*255.999))
}
