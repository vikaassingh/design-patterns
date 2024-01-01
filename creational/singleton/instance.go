package singleton

import (
	"fmt"
	"sync"
)

type singleton struct {
}

var lock = sync.Mutex{}
var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &singleton{}
			fmt.Println("Created First")
		} else {
			fmt.Println("Allready Created")
		}
	} else {
		fmt.Println("Instance Allready Available")
	}

	return instance
}
