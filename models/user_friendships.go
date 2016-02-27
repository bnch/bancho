package models

// UserFriendship is a friendship relationship between two users.
type UserFriendship struct {
	ID int
	// Don't ask me on what kind of drugs I was on when I decided they had to be lovers.
	// Because I have no idea.
	// Heroin, perhaps?
	Lover int
	Loved int
}
