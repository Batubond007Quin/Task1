package protobuf

import "context"

type Server struct {
}

func (s *Server) ChangeText(context context.Context, message *Message) (*Message, error) {
	body := message.Body
	result := "[CHANGED]" + body

	resultMessage := Message{}
	resultMessage.Body = result

	return &resultMessage, nil
}
