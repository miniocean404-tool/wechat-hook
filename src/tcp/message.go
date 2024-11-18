package tcp

type Message struct {
	Content            string  `json:"content"`
	CreateTime         int64   `json:"createTime"`
	DisplayFullContent string  `json:"displayFullContent"`
	FromUser           string  `json:"fromUser"`
	MsgID              float64 `json:"msgId"`
	MsgSequence        int64   `json:"msgSequence"`
	PID                int64   `json:"pid"`
	Signature          string  `json:"signature"`
	ToUser             string  `json:"toUser"`
	Type               int64   `json:"type"`
}
