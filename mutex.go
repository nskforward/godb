package godb

import "sync"

type TwoLevelMutex struct {
	x sync.Mutex
	m map[string]*sync.Mutex
}

func NewTwoLevelMutex() *TwoLevelMutex {
	return &TwoLevelMutex{
		m: make(map[string]*sync.Mutex),
	}
}

func (tmtx *TwoLevelMutex) Get(bucket string) *sync.Mutex {
	tmtx.x.Lock()
	defer tmtx.x.Unlock()
	mtx, ok := tmtx.m[bucket]
	if !ok {
		mtx = &sync.Mutex{}
		tmtx.m[bucket] = mtx
	}
	return mtx
}
