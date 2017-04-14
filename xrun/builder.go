package xrun

import "fmt"

type Builder struct {

}

func (b *Builder)Run() {
	fmt.Println("(b *Builder)Run()")
}