package main

import (
	"fmt"
)

type Pair struct {
	key int
	val string	
}

type ArrayHashMap struct {
	_buckets []*Pair
}

func newArrayHashMap() *ArrayHashMap {
	buckets := make([]*Pair, 100)
	return &ArrayHashMap {
		_buckets:buckets,
	}
}

func (s *ArrayHashMap) hasFunc(key int) int {
	index := key % 100
	return index	
}

func (s *ArrayHashMap) get(key int) string {
	index := s.hasFunc(key)
	pair := s._buckets[index]
	if pair == nil {
		return "Not Found"
	}
	
	return pair.val
}

func (s *ArrayHashMap) put(key int, val string) {
	pair := &Pair{key:key, val:val}
	index := s.hasFunc(key)
	s._buckets[index] = pair
}

func (s *ArrayHashMap) remove(key int) {
	index := s.hasFunc(key)
	s._buckets[index] = nil
}

func (s *ArrayHashMap) pairSet() []*Pair {
	var pairs []*Pair
	for _, pair := range s._buckets {
		if pair != nil {
			pairs = append(pairs, pair)
		}
	}
	return pairs
}

func (s *ArrayHashMap) keySet() []int {
	var keys []int
	for _, pair := range s._buckets {
		if pair != nil {
			keys = append(keys, pair.key)
		}
	}
	return keys
}

func (s *ArrayHashMap) valueSet() []string {
	var values []string
	for _, pair := range s._buckets {
		if pair != nil {
			values = append(values, pair.val)
		}
	}
	return values
}

func (s *ArrayHashMap) logself() {
	for _, pair := range s._buckets {
		if pair != nil {
			fmt.Println(pair.key, "->", pair.val)
		}
	}
}

func main() {
	hmap := newArrayHashMap()
	//hmap[12836] = "xiaoha"
	//hmap[15937] = "xiaoluo"
	//hmap[16750] = "xiaosuan"
	//hmap[13276] = "xiaofa"
	//hmap[10583] = "xiaoya"
	
	hmap.put(12836, "xiaoha")
	hmap.put(15937, "xiaoluo")
	hmap.put(16750, "xiaosuan")
	hmap.put(13276, "xiaofa")
	hmap.put(10583, "xiaoya")
	
	name := hmap.get(15937)
	fmt.Printf("name:%s\n", name)
	
	hmap.remove(10583)
	
	hmap.logself()
	
	//name := hmap[15937]
}