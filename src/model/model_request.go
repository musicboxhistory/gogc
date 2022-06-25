package model

import (
	"net/url"
)

type Request struct {
	Params map[string]string
	Query url.Values
}
