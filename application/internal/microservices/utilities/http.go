// rq -> http.Request
// rs -> http.Response
// bq -> Request
// bs -> Response
package utilities

import (
	"tde/internal/microservices/paths"

	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
	"reflect"
	"strings"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

var (
	ErrMissingKeyInUrl = errors.New("Url has one or more missing keys")
)

func separateParams(in any) (map[string]string, map[string]any) {
	var t = reflect.TypeOf(in)
	var v = reflect.ValueOf(in)
	var fields = v.NumField()
	var url = map[string]string{}
	var body = map[string]any{}
	var tags reflect.StructTag
	for i := 0; i < fields; i++ {
		tags = t.Field(i).Tag
		if key, exists := tags.Lookup("url"); exists {
			url[key] = v.Field(i).String()
		} else if key, exists := tags.Lookup("json"); exists {
			body[key] = v.Field(i).Interface()
		}
	}
	return url, body
}

func fillUrlParamaters(muxMap map[string]string, bq any) error {
	var (
		t       = reflect.TypeOf(bq).Elem()
		v       = reflect.ValueOf(bq).Elem()
		fields  = v.NumField()
		value   string
		missing = []string{}
	)
	for i := 0; i < fields; i++ {
		if key, exists := t.Field(i).Tag.Lookup("url"); exists {
			if value, exists = muxMap[key]; exists {
				v.Field(i).Set(reflect.ValueOf(value))
			} else {
				missing = append(missing, key)
			}
		}
	}
	if len(missing) > 0 {
		return errors.Wrap(ErrMissingKeyInUrl, strings.Join(missing, ", "))
	}
	return nil
}

func NewRequest(ep paths.Endpoint, params any) (*http.Request, error) {
	var err error
	var urlParams, bodyParams = separateParams(params)
	var buffer = bytes.NewBuffer([]byte{})
	if bodyParams != nil {
		err = json.NewEncoder(buffer).Encode(bodyParams)
		if err != nil {
			return nil, errors.Wrap(err, "serializing the body")
		}
	}
	var r *http.Request
	r, err = http.NewRequest(ep.Method.String(), ep.Url(), buffer)
	if err != nil {
		return nil, errors.Wrap(err, "creating request object")
	}
	if bodyParams != nil {
		r.Header.Set("Content-Type", mime.TypeByExtension("json"))
		r.Header.Set("Content-Length", fmt.Sprintf("%d", buffer.Len()))
	}
	if urlParams != nil {
		r = mux.SetURLVars(r, urlParams)
	}
	return r, nil
}

func ParseRequest[Request any](rq *http.Request) (bq *Request, err error) {
	bq = new(Request)
	err = json.NewDecoder(rq.Body).Decode(bq)
	if err != nil && errors.Is(err, io.ErrUnexpectedEOF) {
		err = errors.Wrap(err, "parsing the request body")
		return
	}
	var vars = mux.Vars(rq)
	err = fillUrlParamaters(vars, bq)
	if err != nil {
		return bq, errors.Wrap(err, "checking url parameters")
	}
	return
}

func WriteJsonResponse(bs any, rsw http.ResponseWriter) error {
	rsw.Header().Set("Content-Type", mime.TypeByExtension("json"))
	err := json.NewEncoder(rsw).Encode(bs)
	if err != nil {
		return errors.Wrap(err, "serializing the body")
	}
	return nil
}

func ParseJsonResponse[Response any](rs *http.Response) (bs *Response, err error) {
	rs.Header.Get("Content-Type")
	bs = new(Response)
	err = json.NewDecoder(rs.Body).Decode(bs)
	if err != nil {
		return nil, errors.Wrap(err, "parsing the response body")
	}
	return
}

func Send[Request any, Response any](ep paths.Endpoint, bq Request) (*Response, error) {
	var (
		rq  *http.Request
		rs  *http.Response
		bs  = new(Response)
		err error
	)
	rq, err = NewRequest(ep, bq)
	if err != nil {
		return nil, errors.Wrap(err, "creating request")
	}
	rs, err = http.DefaultClient.Do(rq)
	if err != nil {
		return nil, errors.Wrap(err, "sending the request")
	}
	bs, err = ParseJsonResponse[Response](rs)
	if err != nil {
		return nil, errors.Wrap(err, "binding response")
	}
	return bs, nil
}
