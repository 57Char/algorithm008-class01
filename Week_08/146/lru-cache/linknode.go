package lru_cache

type linkNode struct {
	key, val  int
	pre, next *linkNode
}
