package lbot

import (
	"errors"
)

const (
	EventMessage   = eventString("138311609000106303")
	EventOperation = eventString("138311609100106403")

	TextMessage     = 1
	ImageMessage    = 2
	VideoMessage    = 3
	AudioMessage    = 4
	LocationMessage = 7
	StickerMessage  = 8
	ContactMessage  = 10

	OpAddedFriend = 4
	OpBlocked     = 8

	ToTypeUser = 1

	DefaultToChannel = 1383378250
)

var (
	ErrUserExceed = errors.New("User exceed limition")
)
