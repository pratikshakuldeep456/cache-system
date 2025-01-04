package pkg

type CacheStrategy interface {

	//func name
	funcName()
}

type LRU struct {
}

func (l *LRU) funcName() {

}

type LFU struct {
}

func (L *LFU) funcName() {

}
