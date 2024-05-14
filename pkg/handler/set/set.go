package set

import (
	"net/http"

	"github.com/tomMoulard/htransformation/pkg/types"
	"github.com/tomMoulard/htransformation/pkg/utils/header"
)

func Validate(rule types.Rule) error {
	if rule.Header == "" {
		return types.ErrMissingRequiredFields
	}

	return nil
}

func Handle(rw http.ResponseWriter, req *http.Request, rule types.Rule) {
	if rule.SetOnResponse {
		rw.Header().Set(rule.Header, rule.Value)

		return
	}

	header.Set(req, rule.Header, rule.Value)
}
