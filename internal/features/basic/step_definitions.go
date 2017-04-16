package basic

import (
	. "github.com/johnmcdnl/xrun/xrun"
	"fmt"
)

func init(){
	And(`^I have a second step$`, func() {
		fmt.Println("ALMOST!")
	})

}
