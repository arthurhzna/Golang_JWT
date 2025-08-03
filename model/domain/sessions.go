package domain

import "time"

type Session struct {
	ID        int
	User_Email string
	Refresh_Token string
	Is_Revoked bool
	Created_At time.Time
	Expires_At time.Time
}
