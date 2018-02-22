package simplehttp

import (
	"github.com/ljun20160606/simplehttp/simplehttputil"
	"io"
	"net/http"
	"strings"
)

type Request struct {
	method       string
	url          *strings.Builder
	header       http.Header
	querys       [][2]string
	forms        map[string][]string
	body         io.Reader
	jsonData     interface{}
	cookies      []*http.Cookie
	charset      simplehttputil.Charset
	clearCookies bool
	retry        int
	Client       Client
}

func NewRequest(client Client) *Request {
	return &Request{header: http.Header{}, Client: client}
}

func (r *Request) SetMethod(method string) *Request {
	r.method = method
	return r
}

func (r *Request) Get() *Request {
	r.method = http.MethodGet
	return r
}

func (r *Request) Post() *Request {
	r.method = http.MethodPost
	return r
}

func (r *Request) Patch() *Request {
	r.method = http.MethodPatch
	return r
}

func (r *Request) Connect() *Request {
	r.method = http.MethodConnect
	return r
}

func (r *Request) Delete() *Request {
	r.method = http.MethodDelete
	return r
}

func (r *Request) MethodHead() *Request {
	r.method = http.MethodHead
	return r
}

func (r *Request) Options() *Request {
	r.method = http.MethodOptions
	return r
}

func (r *Request) SetCookies(c []*http.Cookie) *Request {
	r.cookies = c
	return r
}

func (r *Request) SetStringCookies(cookie string) *Request {
	return r.SetCookies(ReadCookies([]string{cookie}, ""))
}

func (r *Request) AddCookie(ck *http.Cookie) *Request {
	if r.cookies == nil {
		r.cookies = []*http.Cookie{ck}
		return r
	}
	r.cookies = append(r.cookies, ck)
	return r
}

func (r *Request) SetUrl(url string) *Request {
	builder := strings.Builder{}
	builder.WriteString(url)
	r.url = &builder
	return r
}

func (r *Request) ClearCookies() *Request {
	r.clearCookies = true
	return r
}

func (r *Request) Query(k, v string) *Request {
	if r.querys == nil {
		r.querys = [][2]string{}
	}
	r.querys = append(r.querys, [2]string{k, v})
	return r
}

func (r *Request) QueryArray(k string, vs []string) *Request {
	if r.querys == nil {
		r.querys = [][2]string{}
	}
	for _, v := range vs {
		r.querys = append(r.querys, [2]string{k, v})
	}
	return r
}

func (r *Request) SetQuerys(querys [][2]string) *Request {
	r.querys = querys
	return r
}

func (r *Request) Form(k string, v string) *Request {
	if r.forms == nil {
		r.forms = make(map[string][]string)
		r.Head(ContentType, "application/x-www-form-urlencoded")
	}
	r.forms[k] = append(r.forms[k], v)
	return r
}

func (r *Request) FormForce(k string, v string) *Request {
	if r.forms == nil {
		r.forms = make(map[string][]string)
		r.Head(ContentType, "application/x-www-form-urlencoded")
	}
	r.forms[k] = []string{v}
	return r
}

func (r *Request) FormArray(k string, v []string) *Request {
	if r.forms == nil {
		r.forms = make(map[string][]string)
		r.Head(ContentType, "application/x-www-form-urlencoded")
	}
	r.forms[k] = v
	return r
}

func (r *Request) SetForms(forms map[string][]string) *Request {
	if r.forms == nil {
		r.Head(ContentType, "application/x-www-form-urlencoded")
	}
	if forms == nil {
		r.forms = make(map[string][]string)
	} else {
		r.forms = forms
	}
	return r
}

func (r *Request) SetJSON(data interface{}) *Request {
	r.jsonData = data
	return r
}

func (r *Request) SetBody(body io.Reader) *Request {
	r.body = body
	return r
}

func (r *Request) SetRetry(retry int) *Request {
	r.retry = retry
	return r
}

func (r *Request) RefererInHeader(refer string) *Request {
	return r.Head(Referer, refer)
}

func (r *Request) OriginInHeader(origin string) *Request {
	return r.Head(Origin, origin)
}

func (r *Request) Head(k, v string) *Request {
	r.header.Set(k, v)
	return r
}

func (r *Request) GB18030() *Request {
	r.charset = simplehttputil.GB18030
	return r
}

func (r *Request) UTF8() *Request {
	r.charset = simplehttputil.UTF8
	return r
}

func (r *Request) Send() (resp *Response) {
	return r.Client.Send(r)
}
