//package main
//
//import (
//	"fmt"
//	"github.com/johnmcdnl/xrun/xrun"
//	"os"
//)

//func main() {
//	fmt.Println("XRUN v0.1")
//
//	new(xrun.Runner).Build()
//}
//

package main

import (
"os"
"github.com/ajstarks/svgo"
)
func main() {
	width := 500
	height := 500
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height)
	canvas.Ellipse(width/2, height, width/2,height/3,"fill:rgb(44,77,232)")
	canvas.Text(width/2, height/2, "Hello, World", "fill:white;font-size:48pt;text-anchor:middle")
	canvas.End()


}