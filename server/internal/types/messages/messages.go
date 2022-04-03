package message

import (
	"github.com/rapture-haven/virtue-of-ignorance/internal/types"
)

type Message struct {
	Event string `json:"event"`
}

type SystemInit struct {
	Event           string            `json:"event"`
	ImplicitHosting bool              `json:"implicit_hosting"`
	Locks           map[string]string `json:"locks"`
}

type SystemMessage struct {
	Event   string `json:"event"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

type MembersList struct {
	Event    string          `json:"event"`
	Memebers []*types.Member `json:"members"`
}

type Member struct {
	Event string `json:"event"`
	*types.Member
}

type MemberDisconnected struct {
	Event string `json:"event"`
	ID    string `json:"id"`
}
