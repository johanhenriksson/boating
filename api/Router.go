package api

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/gorilla/sessions"

    "github.com/johanhenriksson/edetalj-backend/views"
)

type RoutedService interface {
    Path()      string
    Routes()    Routes
}

type RouteArgs struct {
    Vars        map[string]string
    Writer      http.ResponseWriter
    Request     *http.Request
    Session     *sessions.Session
}

type RouteHandlerFunc func(RouteArgs)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    Handler     RouteHandlerFunc
}

type Routes []Route

type Router struct {
    router      *mux.Router
    session     *sessions.CookieStore
    services    []RoutedService
}

func (r *Router) Mux() *mux.Router {
    return r.router
}

func (r *Router) Files(prefix string, dir string) {
    r.Mux().PathPrefix(prefix).Handler(http.FileServer(http.Dir(dir)))
}

func (router *Router) Register(srv RoutedService) {
    for _, route := range srv.Routes() {
        path := fmt.Sprintf("%s%s", srv.Path(), route.Pattern);
        router.Mux().
            Methods(route.Method).
            Path(path).
            Name(route.Name).
            HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                router.Route(route, w, r)
            })
    }
}

func (router *Router) RegisterView(view *views.View) {
    router.Mux().
        Methods("GET").
        Path(view.Pattern).
        HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            view := views.UserView { }
            view.Render(views.ViewContext{
                Writer: w,
                Vars: make(map[string]interface{}),
            })

        })
}

func (router *Router) Route(route Route, w http.ResponseWriter, r *http.Request) {
    session, _ := router.session.Get(r, "session")
    params := RouteArgs {
        Request:    r,
        Writer:     w,
        Vars:       mux.Vars(r),
        Session:    session,
    }
    route.Handler(params)
}

func NewRouter() *Router {
    router := mux.NewRouter()

    return &Router {
        router: router,
        session: sessions.NewCookieStore([]byte("secret password")),
    }
}
