package blogposts_test

import (
	"blogposts"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post1
Description: Description 1
Tags: go, programming
---
Hello world!
This is a post about Go programming.`
		secondBody = `Title: Post2
Description: Description 2
Tags: helloWorld, testing, learning, knowledge, go
---
This is another post about testing and learning Go.
It also has multiple lines of content.
And it should be read correctly.`
	)
	fs := fstest.MapFS{
		"hello-world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	got := posts[0]
	want := blogposts.Post{
		Title:       "Post1",
		Description: "Description 1",
		Tags:        []string{"go", "programming"},
		Body: `Hello world!
This is a post about Go programming.`}

	if err != nil {
		t.Fatal(err)
	}

	assertPost(t, got, want)
}

func assertPost(t *testing.T, got, want blogposts.Post) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
