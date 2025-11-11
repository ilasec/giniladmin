package global

import (
	"sync"
)

var (
	lock      sync.RWMutex
	PAGE_SIZE = 10 // 默认的分页数量

)
