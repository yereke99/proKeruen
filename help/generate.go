package help

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GenerateRandomID(l int) (lens int) {
	s := ""
	rand.Seed(time.Now().Unix())

	for i := 0; i < l; i++ {
		s += (string)(rand.Intn(10) + 48)
	}

	lens, _ = strconv.Atoi(s)

	return
}

func GenerateRandomId(l int) int {
	var buff bytes.Buffer
	rand.Seed(time.Now().Unix())

	for i := 0; i < l; i++ {
		fmt.Fprintf(&buff, "%d", rand.Intn(10))
	}

	res, _ := strconv.Atoi(buff.String())

	if len(buff.String()) != 4 {
		return GenerateRandomId(l)
	}
	return res
}
