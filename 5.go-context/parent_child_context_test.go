package go_context

import (
	"context"
	"fmt"
	"testing"
)

func TestParentChildContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)
}
