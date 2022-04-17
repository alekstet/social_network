package main

import (
	"fmt"
)

type User struct {
	Name           string
	Id             int
	ListFollowers  map[int]User
	ListFollowings map[int]User
	Posts          []Post
	Stories        []Story
	Activity       []string
}

type Post struct {
	Id       int
	Date     string
	Image    string
	Likes    int
	Comments []Comment
}

type Story struct {
	Id    int
	Date  string
	Image string
}

type Comment struct {
	Date string
	Name string
	Text string
}

func NewUser(name string, id int) *User {
	return &User{
		Name:           name,
		Id:             id,
		ListFollowers:  make(map[int]User),
		ListFollowings: make(map[int]User),
		Posts:          []Post{},
		Stories:        []Story{},
		Activity:       []string{},
	}
}

type IUser interface {
	NewContent(interface{})

	Notification(string)
	NotifyAboutContent(interface{})

	Follow(user *User)
	UnFollow(user *User)
	RemoveFollower(int)
	BeRemoved(int)

	GetFollowers() map[int]string
	GetFollowings() map[int]string
	GetPosts() map[int]Post
	GetActivity() map[int]string

	CreatePost(string, string)
	CreateStory(string, string)
}

func (u *User) NewContent(content interface{}) {
	for _, i := range u.ListFollowers {
		i.NotifyAboutContent(content)
	}
}

func (u *User) CreatePost(date, image string) {
	post := Post{Date: date, Image: image}
	u.Posts = append(u.Posts, post)
	u.NewContent(post)
}

func (u *User) CreateStory(date, image string) {
	story := Story{Date: date, Image: image}
	u.Stories = append(u.Stories, story)
	u.NewContent(story)
}

func (u *User) NotifyAboutContent(content interface{}) {
	switch cont := content.(type) {
	case Post:
		fmt.Printf("Id %v, see new post: %s, %s\n", u.Id, cont.Image, cont.Date)
	case Story:
		fmt.Printf("Id %v, see new story: %s, %s\n", u.Id, cont.Image, cont.Date)
	}
}

func (u *User) Follow(user *User) {
	u.ListFollowings[user.Id] = *user
	user.ListFollowers[u.Id] = *u
	msg := fmt.Sprintf("%s start followed you", u.Name)
	user.Notification(msg)
}

func (u *User) UnFollow(user *User) {
	delete(u.ListFollowings, user.Id)
	delete(user.ListFollowers, u.Id)
	msg := fmt.Sprintf("%s finish followed you", u.Name)
	user.Notification(msg)
}

func (u *User) Notification(msg string) {
	u.Activity = append(u.Activity, msg)
}

func (u *User) BeRemoved(id int) {
	delete(u.ListFollowings, id)
}

func (u *User) RemoveFollower(id int) {
	removedFollower := u.ListFollowings[id]
	removedFollower.BeRemoved(u.Id)
	removedFollower.Notification("User deleted you from subscriptions:(")
	delete(u.ListFollowers, id)
}

func (u *User) GetFollowers() map[int]string {
	followers := make(map[int]string)
	for _, i := range u.ListFollowers {
		followers[i.Id] = i.Name
	}
	return followers
}

func (u *User) GetFollowings() map[int]string {
	followings := make(map[int]string)
	for _, i := range u.ListFollowings {
		followings[i.Id] = i.Name
	}
	return followings
}

func (u *User) GetPosts() map[int]Post {
	posts := make(map[int]Post)
	for _, i := range u.Posts {
		posts[i.Id] = i
	}
	return posts
}

func (u *User) GetActivity() map[int]string {
	activity := make(map[int]string)
	for _, i := range u.ListFollowings {
		activity[i.Id] = i.Name
	}
	return activity
}
