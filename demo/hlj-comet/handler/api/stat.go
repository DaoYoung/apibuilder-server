package api

import(
	"hlj-rest/rest"
	"strconv"
	"runtime"
)

func GetStat(w rest.ResponseWriter, r *rest.Request) {
	routines := strconv.FormatInt(int64(runtime.NumGoroutine()), 10)
	stats := make(map[string]interface{})
	stats["routines"] =routines

	rest.Ok(w, stats, nil)
}
