package lib

import (
	"math/rand"
	"time"
)

var RandSource = rand.NewSource(time.Now().UnixNano())
