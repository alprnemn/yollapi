package common

import "time"

type userKey string

const (
	QueryTimeoutDuration         = time.Second * 5
	MaxBytes                     = 1 << 20
	UserCtx              userKey = "user"
	UserExpTime                  = time.Minute
)
