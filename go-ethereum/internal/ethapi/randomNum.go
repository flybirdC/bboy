package ethapi

import (
	"math/rand"
	"time"
	"strconv"
)

//随机函数例子
func RandNum() []byte {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(100)
	numString :=strconv.Itoa(randNum)
	numbyte := []byte(numString)

	return numbyte
}

