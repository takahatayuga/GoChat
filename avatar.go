package main

import (
	"errors"
)

//ErrNoAvatarはAvatarインスタンスがアバターのURLを返すことができない場合に発生するエラー
var ErrorNoAvatarURL = errors.New("chat: アバターのURLを取得できません")

//Avatarはユーザのプロフィール画像を表す型
type Avatar interface {
	//GetAvatarURLは指定されたクライアントのアバターのURLを返す
	//問題が発生した場合にはエラーを返す
	//特にURLを取得できなかった場合はErrorNoAvatarURLを返す
	GetAvatarURL(c *client) (string, error)
}

type AuthAvatar struct{}

var UseAuthAvatar AuthAvatar

func (_ AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrorNoAvatarURL
}
