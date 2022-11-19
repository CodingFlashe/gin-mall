package cache

import (
	"fmt"
	"strconv"
)

const (
	// RankKey 排名
	RankKey = "rank"
)

func ProductViewKey(id uint) string {
	return fmt.Sprintf("view:product:%s", strconv.Itoa(int(id)))
}
