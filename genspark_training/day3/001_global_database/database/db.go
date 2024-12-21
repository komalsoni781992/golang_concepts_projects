package database

// avoid exported global variable most of the times
// anyone can change the state of it

type SQLDb string // allows no random string to be connection

var dB SQLDb

func InitDB(conn SQLDb) SQLDb {
	dB = conn
	return dB
}
