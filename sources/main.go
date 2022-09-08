package main

import (
    "context"
    "fmt"
    //"html"
    "log"
    "net/http"
    "os"
    "io/ioutil"    
	"regexp"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

/*
 * Routes
 *
 * - GET /parameter/:name
 * - GET /signal/:name
 * - GET /time
 * - GET /version
 * - POST /start
 * - POST /stop
 * - PUT /parameter/:name/assign/:value
 * - PUT /signal/:name/assign/:value
 * 
 */

var routes = []route{	
    newRoute("GET",  "/parameter/([^/]+)",                  getParameter),
    newRoute("GET",  "/signal/([^/]+)",                     getSignal),
    newRoute("GET",  "/time/([^/]+)",                       getTime),
    newRoute("GET",  "/version",                            getVersion),
    newRoute("POST", "/start",                              start),
    newRoute("POST", "/stop",                               stop),
    newRoute("PUT",  "/parameter/([^/]+)/assign/([0-9]+)",  setParameter),
    newRoute("PUT",  "/signal/([^/]+)/assign/([0-9]+)",     setSignal),
}

func newRoute(method, pattern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}


func Serve(w http.ResponseWriter, r *http.Request) {
	var allow []string
	for _, route := range routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}
			ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
			route.handler(w, r.WithContext(ctx))
			return
		}
	}
	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.NotFound(w, r)
}

/*
 * Utilities
 */
type ctxKey struct{}

func getField(r *http.Request, index int) string {
	fields := r.Context().Value(ctxKey{}).([]string)
	return fields[index]
}

/*
 * Route callbacks
 */
func getParameter(w http.ResponseWriter, r *http.Request) {
	name := getField(r, 0)
	fmt.Fprintf(w, "Parameter '%s'\n", name)
}

func getSignal(w http.ResponseWriter, r *http.Request) {
	name := getField(r, 0)
	fmt.Fprintf(w, "Signal '%s'\n", name)
}

func getTime(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Time\n")
}

func getVersion(w http.ResponseWriter, r *http.Request) {
	ver, err := ioutil.ReadFile("/config/VERSION")
    check(err)
    fmt.Fprintf(w,string(ver))
}

func start(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w,"start")
}

func stop(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w,"stop")
 }

func setParameter(w http.ResponseWriter, r *http.Request) {
	name  := getField(r, 0)
    value := getField(r, 1)
	fmt.Fprintf(w, "Parameter '%s' = %s\n", name, value)
}

func setSignal(w http.ResponseWriter, r *http.Request) {
	name  := getField(r, 0)
    value := getField(r, 1)
	fmt.Fprintf(w, "Signal '%s' = %s\n", name, value)
}

/*
 * Entry-point
 */
 func main() {

    http.Handle("/", http.FileServer(http.Dir("/web")))

    router := http.HandlerFunc(Serve)

    log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), router))

}
