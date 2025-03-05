package public_routes

import (
	"context"
	"net/http"

	"github.com/carsonkrueger/main/internal/types"
	"github.com/carsonkrueger/main/templates/layouts"
	"github.com/carsonkrueger/main/templates/pages"
	"github.com/carsonkrueger/main/tools"
	"github.com/go-chi/chi/v5"
)

type Home struct {
	types.WithAppContext
}

func (r *Home) Path() string {
	return "/"
}

func (hw *Home) PublicRoute(r chi.Router) {
	r.Get("/", hw.redirect_home)
	r.Get("/home", hw.home)
	r.Get("/signup", hw.signup)
}

func (hw *Home) redirect_home(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, "/home", http.StatusMovedPermanently)
}

func (hw *Home) home(res http.ResponseWriter, req *http.Request) {
	hw.AppCtx.Lgr.Info("Logging /home route")
	home := layouts.Main(pages.Home())
	err := home.Render(context.Background(), res)
	if err != nil {
		tools.RequestHttpError(hw.AppCtx.Lgr, res, 500, err)
	}
}

func (hw *Home) signup(res http.ResponseWriter, req *http.Request) {
	hw.AppCtx.Lgr.Info("Logging /signup route")
	home := layouts.Main(pages.HomeSignup())
	err := home.Render(context.Background(), res)
	if err != nil {
		tools.RequestHttpError(hw.AppCtx.Lgr, res, 500, err)
	}
}
