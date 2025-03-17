package drinking

const (
	PENDING = "pending"
	MAKING  = "making"
	SERVED  = "served"
)

const (
	COFFEE = "coffee"
	JUICE  = "juice"
	TEA    = "tea"
	NONE   = "none"
)

type Drinking struct {
	status string
}
