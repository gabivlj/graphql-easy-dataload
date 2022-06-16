package main

import (
	"log"
	"net/http"

	loader "github.com/gabivlj/grahqldl"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type Post struct {
	something string
}

func (p Post) Something() string {
	return p.something
}

type User struct {
	*loader.DataLoaderInstance[Post, string]

	name string
}

type query struct{}

func (*query) Users() []User {
	users := getUsers()
	dl := loader.NewLoader2D(func(keys []string) ([][]Post, error) {
		posts := getPosts(keys)
		return posts, nil
	})
	for i := range users {
		users[i].DataLoaderInstance = dl.LoadKey(users[i].name)
	}

	return users
}

func (u User) Name() string {
	return u.name
}

func (u User) Posts() []Post {
	v, err := u.DataLoaderInstance.Get()
	if err != nil {
		panic(err)
	}

	return v
}

func main() {
	s := `
				type Post {
					something: String!,
				}

				type User {
					posts: [Post!]!,
					name: String!
				}

                type Query {
                        users: [User!]!
                }
        `
	schema := graphql.MustParseSchema(s, &query{})
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
