package main

type set struct {
	items map[interface{}]int64
}

func newSet() *set {
	return &set{
		items: make(map[interface{}]int64),
	}
}

func (s *set) add(item interface{}) bool {
	if _, exists := s.items[item]; !exists {
		s.items[item] = 1
		return true
	}

	s.items[item]++
	return false
}

func (s *set) contains(item interface{}) bool {
	_, alreadySeen := s.items[item]
	return alreadySeen
}

func (s *set) count(item interface{}) int64 {
	if count, alreadySeen := s.items[item]; alreadySeen {
		return count
	}

	return 0
}
