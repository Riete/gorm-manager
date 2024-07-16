package clause

import (
	"gorm.io/gorm"
)

type Clause func(db *gorm.DB)

func Where(query any, args ...any) Clause {
	return func(db *gorm.DB) {
		db.Where(query, args...)
	}
}

func Not(query any, args ...any) Clause {
	return func(db *gorm.DB) {
		db.Not(query, args...)
	}
}

func Or(query any, args ...any) Clause {
	return func(db *gorm.DB) {
		db.Or(query, args...)
	}
}

func Select(query any, args ...any) Clause {
	return func(db *gorm.DB) {
		db.Select(query, args...)
	}
}

func Omit(column ...string) Clause {
	return func(db *gorm.DB) {
		db.Omit(column...)
	}
}

func Offset(offset int) Clause {
	return func(db *gorm.DB) {
		db.Offset(offset)
	}
}

func Limit(limit int) Clause {
	return func(db *gorm.DB) {
		db.Limit(limit)
	}
}

func Order(order string) Clause {
	return func(db *gorm.DB) {
		db.Order(order)
	}
}

func Group(group string) Clause {
	return func(db *gorm.DB) {
		db.Group(group)
	}
}

func Having(query any, args ...any) Clause {
	return func(db *gorm.DB) {
		db.Having(query, args...)
	}
}

func Distinct(args ...any) Clause {
	return func(db *gorm.DB) {
		db.Distinct(args...)
	}
}

func Preload(query string, conds ...any) Clause {
	return func(db *gorm.DB) {
		db.Preload(query, conds...)
	}
}

func Attrs(attrs ...any) Clause {
	return func(db *gorm.DB) {
		db.Attrs(attrs...)
	}
}

func Assign(attrs ...any) Clause {
	return func(db *gorm.DB) {
		db.Assign(attrs...)
	}
}

func RawSql(sql string) Clause {
	return func(db *gorm.DB) {
		db.Raw(sql)
	}
}
