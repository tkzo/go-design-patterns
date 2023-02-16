package main

// strategy interface

type EvictionAlgo interface {
	evict(c *Cache)
}
