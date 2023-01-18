package loki

import (
	_ "net/http"

	"github.com/yylt/lokihub/pkg/ql"
)

type HandlerType int

const (
	IndexStats HandlerType = iota
	Series
	Label
	LabelValues
	Query
	QueryRange
	Tail
	Metrics
	Push
)

type Noder interface {
	// filter
	Match(ql ql.QueryQl) bool

	// do
	Handler(HandlerType, fn func(interface{}))

	// ready, labels is fetch succuss
	// when node is down, should return NotConnect error
	Ready() error
}
