package manager

import (
	"context"

	"github.com/riete/gorm-manager/clauses"
	"gorm.io/gorm"
)

type Manager struct {
	db    *gorm.DB
	model any
	sc    *gorm.Session
}

func (g Manager) Session() *gorm.DB {
	if g.sc == nil {
		return g.db.WithContext(context.Background()).Model(g.model)
	}
	return g.db.Session(g.sc).Model(g.model)
}

func (g Manager) WithClauses(clauses ...clauses.Clause) *gorm.DB {
	db := g.Session()
	for _, clause := range clauses {
		clause(db)
	}
	return db
}

func New(db *gorm.DB, model any, sc *gorm.Session) *Manager {
	return &Manager{db: db, model: model, sc: sc}
}
