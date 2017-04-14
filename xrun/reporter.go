package xrun

import "fmt"

type Reporter struct {

}

func (r *Reporter)Run(){
	fmt.Println("(r *Reporter)Run()")
}