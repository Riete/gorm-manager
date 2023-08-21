package manager

import (
	"gorm.io/gorm"
)

type Option func(db *gorm.DB)

func WithConditions(query any, args ...any) Option {
	return func(db *gorm.DB) {
		db.Where(query, args...)
	}
}

func WithNotConditions(query any, args ...any) Option {
	return func(db *gorm.DB) {
		db.Not(query, args...)
	}
}

func WithOrConditions(query any, args ...any) Option {
	return func(db *gorm.DB) {
		db.Or(query, args...)
	}
}

func WithSelect(query any, args ...any) Option {
	return func(db *gorm.DB) {
		db.Select(query, args...)
	}
}

func WithOmit(column ...string) Option {
	return func(db *gorm.DB) {
		db.Omit(column...)
	}
}

func WithOffset(offset int) Option {
	return func(db *gorm.DB) {
		db.Offset(offset)
	}
}

func WithLimit(limit int) Option {
	return func(db *gorm.DB) {
		db.Limit(limit)
	}
}

func WithOrder(order string) Option {
	return func(db *gorm.DB) {
		db.Order(order)
	}
}

func WithGroupByHaving(group string, having ...any) Option {
	return func(db *gorm.DB) {
		if len(having) > 0 {
			db.Group(group).Having(having[0], having[1:]...)
		} else {
			db.Group(group)
		}
	}
}

func WithDistinct(args ...any) Option {
	return func(db *gorm.DB) {
		db.Distinct(args...)
	}
}

func WithPreload(query string, conds ...any) Option {
	return func(db *gorm.DB) {
		db.Preload(query, conds...)
	}
}

func WithAttrs(attrs ...any) Option {
	return func(db *gorm.DB) {
		db.Attrs(attrs...)
	}
}

func WithAssign(attrs ...any) Option {
	return func(db *gorm.DB) {
		db.Assign(attrs...)
	}
}

func WithRawSql(sql string) Option {
	return func(db *gorm.DB) {
		db.Raw(sql)
	}
}
