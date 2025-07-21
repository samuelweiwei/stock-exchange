package order

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// GenerateOderNo 生成订单号
func GenerateOderNo(prefix string, maxPos int, timeFormat string) string {
	randStr := fmt.Sprintf("%0*s", maxPos, strconv.Itoa(rand.Intn(int(math.Pow10(maxPos)))))
	return fmt.Sprintf("%s%s%s", prefix, time.Now().Format(timeFormat), randStr)
}
