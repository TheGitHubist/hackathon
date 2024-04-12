package controller

import "time"

type session struct {
	id     int
	expiry time.Time
}

func (s session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

var Sessions = map[string]session{}
