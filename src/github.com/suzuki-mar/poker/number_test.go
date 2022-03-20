package poker

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	num := number{}
	num.value = 2

	fmt.Println(num.displayString())
}
