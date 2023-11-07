package model

import "time"

type Document struct {
	Id        string
	Published bool
	WebsiteId string
	CreatedAt *time.Time
}
