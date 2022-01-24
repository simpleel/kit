package query

// Options 表示一个通用的查询参数
type Options struct {
	Start   int64
	Limit   int64
	OrderBy string
}

// NewOption 调用闭包函数设置查询参数
func NewOption(opts ...Option) Options {
	var options Options
	for _, opt := range opts {
		opt(&options)
	}
	if options.Start <= 0 {
		options.Start = 1
	}
	if options.Limit < 1 {
		options.Limit = 20
	}
	return options
}

type Option func(*Options)

// WithStart 设置当前页
func WithStart(start int64) Option {
	return func(qo *Options) {
		qo.Start = start
	}
}

// WithLimit 设置每页显示数量
func WithLimit(limit int64) Option {
	return func(qo *Options) {
		qo.Limit = limit
	}
}

// WithOrderBy 设置排序字段，如："id desc, created_at asc"
func WithOrderBy(orderBy string) Option {
	return func(qo *Options) {
		qo.OrderBy = orderBy
	}
}
