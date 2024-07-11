package model

type Post struct {
	_id                   string
	user                  User
	firstItemName         string
	secondItemName        string
	firstItemURL          string
	secondItemURL         string
	firstItemVoteCounter  int
	secondItemVoteCounter int
	isDeleted             bool
	createDate            string
	updateDate            string
}
