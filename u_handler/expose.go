package u_handler

const (
	LimitKey  string = "limit"
	OffsetKey string = "offset"
	SortByKey string = "sort_by"
	OrderKey  string = "order"
)

type LogConfig interface {
	StatusConfig() map[int]bool
}
