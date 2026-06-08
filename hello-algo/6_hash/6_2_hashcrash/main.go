package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Pair struct {
	key int
	value string
}

type HashMapChaining struct {
	_size int
	_capacity int
	_loadThres float64
	_extendRatio int
	_buckets [][]Pair	
}

func newHashMapChaining() *HashMapChaining {
	buckets := make([][]Pair, 4)
	for i := 0; i < 4; i++ {
		buckets[i] = make([]Pair, 0)
	}
	
	return &HashMapChaining {
		_size: 0,
		_capacity: 4,
		_loadThres: 2.0 / 3.0,
		_extendRatio: 2,
		_buckets: buckets,
	}
}

func (s *HashMapChaining) hasFunc(key int) int {
	return key % s._capacity
}

func (s *HashMapChaining) loadFactor() float64 {
	return float64(s._size) / float64(s._capacity)	
}

func (s *HashMapChaining) get(key int) string {
	index := s.hasFunc(key)
	bucket := s._buckets[index]
	for _, p := range bucket {
		if p.key == key {
			return p.value
		}	
	}
	return ""
}

func (s *HashMapChaining) put(key int, val string) {
	if s.loadFactor() > s._loadThres {
		s.extend()	
	}
	
	index := s.hasFunc(key)
	for i:= range s._buckets[index] {
		if s._buckets[index][i].key == key {
			return
		}
	}
	
	p := Pair {
		key:key,
		value:val,
	}
	
	s._buckets[index] = append(s._buckets[index], p)
	s._size += 1
}

func (s *HashMapChaining) remove(key int) {
	index := s.hasFunc(key)
	for i, p := range s._buckets[index] {
		if p.key == key {
			s._buckets[index] = append(s._buckets[index][:i], s._buckets[index][i+1:]...)
			s._size -= 1
			break
		}
	}
}

func (s *HashMapChaining) extend() {
	buckets := make([][]Pair, len(s._buckets))
	for i:= 0; i < len(s._buckets); i++ {
		buckets[i] = make([]Pair, len(s._buckets[i]))
		copy(buckets[i], s._buckets[i])
	}
	
	s._capacity *= s._extendRatio
	s._buckets = make([][]Pair, s._capacity)
	for i:= 0; i < s._capacity; i++ {
		s._buckets[i] = make([]Pair, 0)
	}
	
	s._size = 0
	for _, bucket := range buckets {
		for _, p := range bucket {
			s.put(p.key, p.value)
		}
	}
}

func (s *HashMapChaining) logself() {

	fmt.Printf("\n\n",)
	var builder strings.Builder
	for i, bucket := range s._buckets {
		builder.WriteString("[")
		builder.WriteString("<" + strconv.Itoa(i) + ">:")
		for _, p := range bucket {
			builder.WriteString(strconv.Itoa(p.key) + " -> " + p.value + " ")
		}
		builder.WriteString("]\n")
		fmt.Printf(builder.String())
		builder.Reset()
	}
	fmt.Printf("\n\n")
	
}

func main() {
	hmap := newHashMapChaining()
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