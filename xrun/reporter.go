package xrun

import "fmt"

type Reporter struct {

}

func (r *Reporter)run(){
	fmt.Println("(r *Reporter)Run()")
}