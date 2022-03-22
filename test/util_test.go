package test

import (
	"fmt"
	"testing"

	"sabariram.com/goserverbase/utils"
)

func TestGetHash(t *testing.T) {
	fmt.Println(utils.GetHash("3edcRFV5tgb"))
}
