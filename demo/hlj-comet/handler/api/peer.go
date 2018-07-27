package api

import(
	"hlj-comet/core"
	"hlj-rest/rest"
	"net/url"
	"regexp"
)

func GetPeers(w rest.ResponseWriter, r *rest.Request) {
	comet := core.GetComet()
	extra := getExtra(r.URL.Query())
	limit := r.QueryParam("limit").Default("10").Int()

	peers := comet.Pool.GetPeers(extra, limit)
	rest.Ok(w, peers, nil)
}

func getExtra(query url.Values) map[string]string {
	var extraExp = regexp.MustCompile(`^remote\[([a-zA-Z1-9_]+)\]$`)

	extra := make(map[string]string)
	for k, v := range query {
		matches := extraExp.FindStringSubmatch(k)
		if len(matches) == 2 {
			extra[matches[1]] = v[0]
		}
	}

	return extra
}