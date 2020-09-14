package prototest3

import "github.com/2767mr/prototest2/sub"

//Bar calls github.com/2767mr/prototest2/sub.Test.
func Bar() {
	sub.Test()
}
