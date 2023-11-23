package forms

import "net/url"

type errs map[string][]string

func (e errs) Add(field, message string) {
	e[field] = append(e[field], message)
}

type Form struct {
	url.Values
	Errors errs
}

func New(data url.Values) *Form {
	return &Form {
		data,
		make(errs),
	}
}
