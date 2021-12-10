package entity

type UserType int

//go:generate stringer -type=UserType -linecomment
const (
	UserTypePublisher UserType = iota + 1 //publisher
	UserTypeWechat                        //wechat
	UserTypeLogin                         //login
	UserTypeTourist                       //touristt

)
