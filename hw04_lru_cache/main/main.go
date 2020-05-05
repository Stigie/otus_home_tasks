package main

import (
	"fmt"

	lru "github.com/Stigie/otus_home_tasks/hw04_lru_cache"
)

func main() {
	cache := lru.NewCache(3)
	ok := cache.Set("qwe", 1)
	fmt.Println(ok)
	ok = cache.Set("qwe1", 4)
	fmt.Println(ok)
	ok = cache.Set("qwe2", 3)
	fmt.Println(ok)
	ok = cache.Set("qwe3", 4)
	fmt.Println(ok)
	ok = cache.Set("qwe", 799)
	fmt.Println(ok)
	ok = cache.Set("qwe", 799)
	fmt.Println(ok)

	value, ok := cache.Get("qwe")
	fmt.Println(value, ok)
	value, ok = cache.Get("qwe777")
	fmt.Println(value, ok)

	cache.Clear()
	ok = cache.Set("qwe", 799)
	fmt.Println(ok)
	ok = cache.Set("qwe1", 799)
	fmt.Println(ok)
	value, ok = cache.Get("qwe1")
	fmt.Println(value, ok)
	value, ok = cache.Get("qwe777")
	fmt.Println(value, ok)
}
