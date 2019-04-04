package main

import "fmt"

//Create author struct
// That contains the firstname and lastname and bio about the author
type author struct {
	firstName, lastName, bio string
}

//Create post struct which will have author struct
type post struct {
	title, contants string
	author
}

//Create slice of posts in website
type website struct {
	posts []post
}

//Method to get all the post from website

func (w website) getAllPosts() {
	//range over all the posts
	for _, pst := range w.posts {
		pst.contains()
	}
}

//Return full name of author
func (auth author) getFullName() string {
	return auth.firstName + " " + auth.lastName
}

//New func will create new author and return author ref

func newAuthor(firstName, lastName, bio string) author {
	return author{firstName, lastName, bio}
}

//Create new post
func newPost(title, contants string, auth author) post {
	return post{title, contants, auth}
}

//Printing post
func (pst post) contains() {
	fmt.Println("Title: ", pst.title)
	fmt.Println("Containts: ", pst.contants)
	fmt.Println("Author: ", pst.getFullName())
	fmt.Println("Bio: ", pst.bio)
	fmt.Println()
}

func main() {
	author1 := newAuthor("Ajay", "Mahar", "Go Enthusiast")
	post1 := newPost("First Post", "This is my first post.", author1)
	post2 := newPost("Second Post", "This is my second post.", author1)
	web := website{posts: []post{post1, post2}}
	fmt.Println("Containts of Website")
	web.getAllPosts()
}
