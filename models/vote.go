package models

type ticketType int
type resultType int
type statusType int

const (
	agree ticketType = iota
	refusal
	abstain
)

const (
	pass resultType = iota
	not
	pending
	cancel
)

const (
	wait statusType = iota
	begin
	end
)

type VoteBase struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	GroupID string `json:"groupid"`
}

type VoteStatistics struct {
	VoteBase
	Result  resultType `json:"result"`
	Status  statusType `json:"status"`
	Tickets []Ballot   `json:"tickets"`
}

type Ballot struct {
	ID     string     `json:"id"`
	Type   ticketType `json:"type"`
	BoxID  string     `json:"from"`
	VoteID string     `json:"voteid"`
}
