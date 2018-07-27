package rest

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var (
	// ErrJsonPayloadEmpty is returned when the JSON payload is empty.
	ErrJsonPayloadEmpty = errors.New("JSON payload is empty")
)

// Request inherits from http.Request, and provides additional methods.
type Request struct {
	*http.Request

	// Map of parameters that have been matched in the URL Path.
	PathParams map[string]string

	// Environment used by middlewares to communicate.
	Env map[string]interface{}

	// a copy of body after called DecodeJsonPayload
	BodyCopy []byte
}

// PathParam provides a convenient access to the PathParams map.
func (r *Request) PathParam(name string) *Value {
	return &Value{r.PathParams[name]}
}

// EnvParam provides a convenient access to the Env map.
// But only support Env of string value
func (r *Request) EnvParam(name string) *Value {
	str, _ := r.Env[name].(string)
	return &Value{str}
}

// QueryParam provides a convenient access to Query Values.
func (r *Request) QueryParam(name string) *Value {
	return &Value{r.URL.Query().Get(name)}
}

// QueryParam provides a convenient access to Query Values.
func (r *Request) SetQueryParam(name string, value string) {
	values := r.URL.Query()
	values.Set(name, value)
	r.URL.RawQuery = values.Encode()
}

// PostParam provides a convenient access to Post Values.
func (r *Request) PostParam(name string) *Value {
	return &Value{r.PostFormValue(name)}
}

// PostParam provides a convenient access to Post Values.
func (r *Request) SetPostParam(name string, value string) {
	r.PostFormValue("")
	r.PostForm.Set(name, value)
}

// FormParam provides a convenient access to FromValue.
func (r *Request) FormParam(name string) *Value {
	return &Value{r.FormValue(name)}
}

// FormParam provides a convenient access to FromValue.
func (r *Request) SetFormParam(name string, value string){
	r.FormValue("")
	r.Form.Set(name, value)
}

// DecodeJsonPayload reads the request body and decodes the JSON using json.Unmarshal.
func (r *Request) DecodeJsonPayload(v interface{}) error {
	if r.BodyCopy == nil {
		var err error
		r.BodyCopy, err = ioutil.ReadAll(r.Body)
		r.Body.Close()
		if err != nil {
			return err
		}
	}

	if len(r.BodyCopy) == 0 {
		return ErrJsonPayloadEmpty
	}
	return json.Unmarshal(r.BodyCopy, v)
}

// SetJsonPayload sets the request body
func (r *Request) SetJsonPayload(k string, v interface{}) error {
	var err error

	tmp := make(map[string]interface{})
	r.DecodeJsonPayload(&tmp)
	tmp[k] = v
	r.BodyCopy, err = json.Marshal(tmp)
	return err
}

// BaseUrl returns a new URL object with the Host and Scheme taken from the request.
// (without the trailing slash in the host)
func (r *Request) BaseUrl() *url.URL {
	scheme := r.URL.Scheme
	if scheme == "" {
		scheme = "http"
	}

	// HTTP sometimes gives the default scheme as HTTP even when used with TLS
	// Check if TLS is not nil and given back https scheme
	if scheme == "http" && r.TLS != nil {
		scheme = "https"
	}

	host := r.Host
	if len(host) > 0 && host[len(host)-1] == '/' {
		host = host[:len(host)-1]
	}

	return &url.URL{
		Scheme: scheme,
		Host:   host,
	}
}

// UrlFor returns the URL object from UriBase with the Path set to path, and the query
// string built with queryParams.
func (r *Request) UrlFor(path string, queryParams map[string][]string) *url.URL {
	baseUrl := r.BaseUrl()
	baseUrl.Path = path
	if queryParams != nil {
		query := url.Values{}
		for k, v := range queryParams {
			for _, vv := range v {
				query.Add(k, vv)
			}
		}
		baseUrl.RawQuery = query.Encode()
	}
	return baseUrl
}

// CorsInfo contains the CORS request info derived from a rest.Request.
type CorsInfo struct {
	IsCors      bool
	IsPreflight bool
	Origin      string
	OriginUrl   *url.URL

	// The header value is converted to uppercase to avoid common mistakes.
	AccessControlRequestMethod string

	// The header values are normalized with http.CanonicalHeaderKey.
	AccessControlRequestHeaders []string
}

// GetCorsInfo derives CorsInfo from Request.
func (r *Request) GetCorsInfo() *CorsInfo {

	origin := r.Header.Get("Origin")

	var originUrl *url.URL
	var isCors bool

	if origin == "" {
		isCors = false
	} else if origin == "null" {
		isCors = true
	} else {
		var err error
		originUrl, err = url.ParseRequestURI(origin)
		isCors = err == nil && r.Host != originUrl.Host
	}

	reqMethod := r.Header.Get("Access-Control-Request-Method")

	reqHeaders := []string{}
	rawReqHeaders := r.Header[http.CanonicalHeaderKey("Access-Control-Request-Headers")]
	for _, rawReqHeader := range rawReqHeaders {
		if len(rawReqHeader) == 0 {
			continue
		}
		// net/http does not handle comma delimited headers for us
		for _, reqHeader := range strings.Split(rawReqHeader, ",") {
			reqHeaders = append(reqHeaders, http.CanonicalHeaderKey(strings.TrimSpace(reqHeader)))
		}
	}

	isPreflight := isCors && r.Method == "OPTIONS" && reqMethod != ""

	return &CorsInfo{
		IsCors:                      isCors,
		IsPreflight:                 isPreflight,
		Origin:                      origin,
		OriginUrl:                   originUrl,
		AccessControlRequestMethod:  strings.ToUpper(reqMethod),
		AccessControlRequestHeaders: reqHeaders,
	}
}
