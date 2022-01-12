package pkg

import "sync"

// singleton is type of single instance
type singleton struct {
}

var instance *singleton
var mu sync.Mutex

/*
	传统方式，每次获取实例都会先加锁，无论是否实例化
	优点：简单、清洗
	缺点：代价高，每次读写都加锁
*/
func GetNormalIns() *singleton {
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

/*
	双重锁：获取前先判断是否实例化，是则直接返回，否则加锁再判断是否实例化
           再决定是直接返回示例还是实例化
	优点：避免每次加锁，提高代码执行效率
*/
func DoubleLockIns() *singleton {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			instance = &singleton{}
		}
	}

	return instance
}

/*
	内置sync库实现，内部实现逻辑看源码

*/
var once sync.Once

func SynOncecIns() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
