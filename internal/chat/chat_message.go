package chat

type ChatMessage interface {
	Content() []byte
}

type WSChatMessage struct {
	typ int
	content []byte
}

func (wscm WSChatMessage) Content() []byte {
	return wscm.content
}
