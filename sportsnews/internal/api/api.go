package api

import (
	"context"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/think-free/sportsnews/lib/logging"
	"github.com/think-free/sportsnews/sportsnews/internal/cliparams"
	"github.com/think-free/sportsnews/sportsnews/internal/database"
)

type Api struct {
	cp *cliparams.ClientParameters
	db database.Database

	mux *mux.Router
}

func New(ctx context.Context, cp *cliparams.ClientParameters, db database.Database) *Api {
	a := &Api{
		cp:  cp,
		db:  db,
		mux: mux.NewRouter(),
	}

	a.mux.HandleFunc("/status", a.getStatus)
	a.mux.HandleFunc("/provider/realise/v1/teams/{team}/news", a.getNews)
	a.mux.HandleFunc("/provider/realise/v1/teams/{team}/news/{id}", a.getNewsID)

	return a
}

func (a *Api) Run() error {
	return http.ListenAndServe(a.cp.ListenAddress, a.mux)
}

func (a *Api) getStatus(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(JSONStatusOK))
}

func (a *Api) getNews(w http.ResponseWriter, r *http.Request) {
	ctx := a.getContextLog(r.Context(), "getNews", r.Host)

	vars := mux.Vars(r)
	team := vars["team"]
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		logging.L(ctx).Errorf("error parsing page: %s", err.Error())
		page = -1
	}

	logging.L(ctx).Debugf("getting news for team '%s'", team)

	news, err := a.db.GetNews(ctx, team, page)
	a.writeResponse(ctx, w, news, err)
}

func (a *Api) getNewsID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	defer w.Header().Set("Content-Type", "application/json")

	ctx := a.getContextLog(r.Context(), "getNewsID", r.Host)

	vars := mux.Vars(r)
	team := vars["team"]
	id := vars["id"]

	logging.L(ctx).Debugf("getting news '%s' for team '%s'", id, team)

	news, err := a.db.GetNewsByID(ctx, team, id)
	a.writeResponse(ctx, w, news, err)
}

func (a *Api) writeResponse(ctx context.Context, w http.ResponseWriter, data interface{}, err error) {
	status := StatusOK
	if err != nil {
		logging.L(ctx).Errorf("error getting data: %s", err.Error())
		status = StatusError
	}

	_, err = w.Write(NewResponse(ctx, status, data))
	if err != nil {
		logging.L(ctx).Errorf("error writing response: %s", err.Error())
	}
}

func (a *Api) getContextLog(ctx context.Context, method, remote string) context.Context {
	ctx = logging.SetTag(ctx, "api_remote", remote)
	ctx = logging.SetTag(ctx, "api_call", method)
	ctx = logging.SetTag(ctx, "api_uuid", uuid.NewString())
	return ctx
}
