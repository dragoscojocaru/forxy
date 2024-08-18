package response

import "net/http"

type ChannelMessage struct {
	idx      int
	response http.Response
}

func NewChannelMessage(idx int, response http.Response) *ChannelMessage {
	tmp := new(ChannelMessage)

	tmp.idx = idx
	tmp.response = response

	return tmp
}

func GetIdx(ChannelMessage *ChannelMessage) int {
	return ChannelMessage.idx
}

func GetResponse(ChannelMessage *ChannelMessage) http.Response {
	return ChannelMessage.response
}
