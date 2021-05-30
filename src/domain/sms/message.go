package sms

type Message struct {
	Content     string `json:"content"`
	Destination string `json:"destination_number"`
	Format      string `json:"format"`
}

type Messages struct {
	Messages []Message `json:"messages"`
}
