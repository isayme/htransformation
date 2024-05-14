package deleter

import (
	"net/http"

	"github.com/tomMoulard/htransformation/pkg/types"
	"github.com/tomMoulard/htransformation/pkg/utils/header"
)

func Validate(types.Rule) error {
	return nil
}

func Handle(rw http.ResponseWriter, req *http.Request, rule types.Rule) {
	if rule.SetOnResponse {
		rw.Header().Del(rule.Header)

		return
	}

	header.Delete(req, rule.Header)
}
