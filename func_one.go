package func_one

import (
    "sync"
)

type mFunc struct {
    key       string
    f         func() interface{}
    listeners []chan interface{}
}

var (
    funcMap = make(map[string]*mFunc)
    lock    = &sync.Mutex{}
)

func Run(key string, f func() interface{}) interface{} {
    ch := insert(key, f)
    return <-ch
}

func insert(key string, f func() interface{}) chan interface{} {
    ch := make(chan interface{})
    lock.Lock()
    if mfunc, ok := funcMap[key]; ok {
        mfunc.listeners = append(mfunc.listeners, ch)
        lock.Unlock()
        return ch
    }
    //没有的话创建一个
    mfunc := &mFunc{
        key:       key,
        f:         f,
        listeners: []chan interface{}{ch},
    }
    funcMap[key] = mfunc
    lock.Unlock()
    // 添加之后开一个新的协程去运行
    go func() {
        result := mfunc.f()
        lock.Lock()
        for _, c := range mfunc.listeners {
            c <- result
        }
        delete(funcMap, key)
        lock.Unlock()
    }()
    return ch
}
