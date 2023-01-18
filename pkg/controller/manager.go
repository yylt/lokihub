package controller

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/yylt/lokihub/pkg/loki"
)

type Manager struct {
	mux *loki.Mux
}

func New(mux *loki.Mux, routes *route.RouterGroup) *Manager {
	m := &Manager{
		mux: mux,
	}
	m.probe(routes)
	return m
}

func (m *Manager) probe(routes *route.RouterGroup) {
	routes.GET("/loki/api/v1/index/stats", m.IndexStatsHandler)
	routes.GET("/loki/api/v1/label/:name/values", m.LabelValueHandler)
	routes.GET("/loki/api/v1/labels", m.LabelsHandler)
	routes.GET("/loki/api/v1/query_range", m.LogQueryRangeHandler)
	// ... TODO
}

// loki/api/v1/query
func (m *Manager) LogQueryHandler(c context.Context, ctx *app.RequestContext) {

}

// loki/api/v1/query_range
func (m *Manager) LogQueryRangeHandler(c context.Context, ctx *app.RequestContext) {

}

// loki/api/v1/labels
func (m *Manager) LabelsHandler(c context.Context, ctx *app.RequestContext) {

}

// loki/api/v1/label/<name>/values
func (m *Manager) LabelValueHandler(c context.Context, ctx *app.RequestContext) {

}

// loki/api/v1/series
func (m *Manager) SeriesHandler(c context.Context, ctx *app.RequestContext) {

}

// loki/api/v1/tail
func (m *Manager) TailHandler(c context.Context, ctx *app.RequestContext) {

}

// loki/api/v1/metrics
func (m *Manager) MetricsHandler(c context.Context, ctx *app.RequestContext) {

}

// loki/api/v1/push
func (m *Manager) PushHandler(w http.ResponseWriter, r *http.Request) {

}

// loki/api/v1/ready
func (m *Manager) ReadyHandler(w http.ResponseWriter, r *http.Request) {

}

// loki/api/v1/index/stats
func (m *Manager) IndexStatsHandler(c context.Context, ctx *app.RequestContext) {

}
