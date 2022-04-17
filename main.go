package main

import "fmt"

func main() {
	c1 := NewUser("Aleks", 1)
	c2 := NewUser("George", 2)
	c3 := NewUser("Slovis", 3)
	c4 := NewUser("Bryl", 4)
	c1.CreatePost("11.04.2022", "Test post")
	fmt.Println(c1.GetPosts())
	c1.Follow(c2)
	c1.Follow(c3)
	c1.Follow(c4)
	c4.Follow(c1)
	c2.Follow(c1)
	fmt.Println(c1.GetFollowings())
	c1.UnFollow(c4)
	fmt.Println(c1.GetFollowings())
	fmt.Println(c1.GetFollowers())
	fmt.Println(c4.GetFollowings())
}
