package query

// Options 表示一个通用的查询参数
type Options struct {
	Page     int64
	PageSize int64
	OrderBy  string
}

// NewOption 调用闭包函数设置查询参数
func NewOption(opts ...Option) Options {
	var options Options
	for _, opt := range opts {
		opt(&options)
	}
	if options.Page <= 0 {
		options.Page = 1
	}
	if options.PageSize < 1 {
		options.PageSize = 20
	}
	return options
}

type Option func(*Options)

// WithPage 设置当前页
func WithPage(page int64) Option {
	return func(qo *Options) {
		qo.Page = page
	}
}

// WithPageSize 设置每页显示数量
func WithPageSize(pageSize int64) Option {
	return func(qo *Options) {
		qo.PageSize = pageSize
	}
}

// WithOrderBy 设置排序字段，如："id desc, created_at asc"
func WithOrderBy(orderBy string) Option {
	return func(qo *Options) {
		qo.OrderBy = orderBy
	}
}
