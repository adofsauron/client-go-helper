package kubectl

import "testing"

	func Test_InitClient(t *testing.T) {
		for i := 0; i < 10; i++ {
			go InitClient()
		}
}