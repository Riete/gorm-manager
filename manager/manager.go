package manager

import (
	"github.com/riete/gorm-manager/clause"
	"gorm.io/gorm"
)

type Manager struct {
	db    *gorm.DB
	model any
	sc    *gorm.Session
}

func (g Manager) Session() *gorm.DB {
	return g.db.Session(g.sc).Model(g.model)
}

func (g Manager) WithClauses(clauses ...clause.Clause) *gorm.DB {
	db := g.Session()
	for _, c := range clauses {
		c(db)
	}
	return db
}

func New(db *gorm.DB, model any) *Manager {
	return &Manager{db: db, model: model}
}
