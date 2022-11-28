package main

import (
	"fmt"
	"log"
)

type Category struct {
	ID   int32
	Name string
	Slug string
}

type Post struct {
	ID          int32
	Categories  []Category
	Title       string
	Description string
	Slug        string
}

type Cacheable interface {
	Category | Post
}

type Cache[T Cacheable] struct {
	Data map[string]T
}

func main() {

	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":   34.09,
		"seconds": 76.12,
	}

	fmt.Printf("Generics Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))

	category := Category{
		ID:   1,
		Name: "Go Generics",
		Slug: "go-generics",
	}
	cc := New[Category]()
	cc.Set(category.Slug, category)

	post := Post{
		ID: 1,
		Categories: []Category{
			{ID: 1, Name: "Go Generics", Slug: "go-generics"},
		},
		Title:       "Generics in Golang structs",
		Description: "Here go's the text",
		Slug:        "generics-in-golang-structs",
	}

	pp := New[Post]()
	pp.Set(post.Slug, post)

	res := pp.Get(post.Slug)
	log.Println(res)

	ouai := cc.Get(category.Slug)
	log.Println(ouai)
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V

	for _, val := range m {
		s += val
	}

	return s
}

func New[T Cacheable]() *Cache[T] {
	c := Cache[T]{}
	c.Data = make(map[string]T)

	return &c
}

func (c *Cache[T]) Set(key string, value T) {
	c.Data[key] = value
}

func (c *Cache[T]) Get(key string) (v T) {
	if v, ok := c.Data[key]; ok {
		return v
	}
	return
}
