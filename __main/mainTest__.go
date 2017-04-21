package main

		import (
			"github.com/johnmcdnl/xrun/carrot"
			_i0 "github.com/johnmcdnl/xrun/internal"
			_i1 "github.com/johnmcdnl/xrun/internal/features"
			_i2 "github.com/johnmcdnl/xrun/internal/features/basic"
			
		)

		var (
			_ = _i0.Imported
			_ = _i1.Imported
			_ = _i2.Imported
			

		)

		func main(){
			new(carrot.TestSuite).Run()
		}
		