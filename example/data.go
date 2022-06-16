package main

import (
	"math/rand"
	"strconv"
)

func getUsers() []User {
	return []User{
		{name: "Hello 1"},
		{name: "Hello 2"},
		{name: "Hello 3"},
	}
}

func getPosts(keys []string) [][]Post {
	p := make([][]Post, len(keys))
	for i := range keys {
		p[i] = make([]Post, 3)
		p[i][0] = Post{something: "made by " + keys[i] + " - " + strconv.Itoa(rand.Int())}
		p[i][1] = Post{something: "made by " + keys[i] + " - " + strconv.Itoa(rand.Int())}
		p[i][2] = Post{something: "made by " + keys[i] + " - " + strconv.Itoa(rand.Int())}
	}

	return p
}
