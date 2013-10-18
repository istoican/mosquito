package mosquito


import (
	"regexp"
)

type Route struct {
	method		Method
	regexp		*regexp.Regexp
	params		[]string
	handler		Handler
}

func (r *Route) Match(method Method, path string) (params map[string]string, found bool) {
	if r.method != method {
		return nil, false
	}

	matches := r.regexp.FindAllStringSubmatch(path, -1)

	if matches == nil {
		return nil, false
	}

	names := r.regexp.SubexpNames()
	for i := 1; i < len(names); i++ {
		params[names[i]] = matches[0][i]
	}
	return params, true
}

func NewRoute(path string, method Method, handler Handler) *Route {
	route := &Route{}

	route.method = method

	re := regexp.MustCompile(":[^/#?()]+")

	regexpString := re.ReplaceAllStringFunc(path, func(m string) string {
	    route.params = append(route.params, m[1:len(m)])
		return "(?P<" + m[1:len(m)] + ">[^/#?]+)"
	})

	route.regexp = regexp.MustCompile(regexpString)

	route.handler = handler

	return route
}
