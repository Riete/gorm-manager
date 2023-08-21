package manager

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
)

type GormManager[T any] struct {
	model *T
	db    *gorm.DB
}

func (g *GormManager[T]) Session(opts ...Option) *gorm.DB {
	return g.SessionContext(context.Background(), opts...)
}

func (g *GormManager[T]) SessionContext(ctx context.Context, opts ...Option) *gorm.DB {
	db := g.db.WithContext(ctx).Order("")
	for _, opt := range opts {
		opt(db)
	}
	return db
}

func (g *GormManager[T]) First(opts ...Option) (*T, error) {
	return g.FirstContext(context.Background(), opts...)
}

func (g *GormManager[T]) FirstContext(ctx context.Context, opts ...Option) (*T, error) {
	model := new(T)
	return model, g.SessionContext(ctx, opts...).First(model).Error
}

func (g *GormManager[T]) Create(model *T) (*T, error) {
	return g.CreateContext(context.Background(), model)
}

func (g *GormManager[T]) CreateContext(ctx context.Context, model *T) (*T, error) {
	return model, g.SessionContext(ctx).Create(model).Error
}

func (g *GormManager[T]) BatchCreate(models ...*T) ([]*T, error) {
	return g.BatchCreateContext(context.Background(), models...)
}

func (g *GormManager[T]) BatchCreateContext(ctx context.Context, models ...*T) ([]*T, error) {
	return models, g.SessionContext(ctx).Create(models).Error
}

func (g *GormManager[T]) FirstOrCreate(model *T, opts ...Option) (*T, error) {
	return g.FirstOrCreateContext(context.Background(), model, opts...)
}

func (g *GormManager[T]) FirstOrCreateContext(ctx context.Context, model *T, opts ...Option) (*T, error) {
	return model, g.SessionContext(ctx, opts...).FirstOrCreate(model).Error
}

func (g *GormManager[T]) Update(column string, value any, opts ...Option) (int64, error) {
	return g.UpdateContext(context.Background(), column, value, opts...)
}

func (g *GormManager[T]) UpdateContext(ctx context.Context, column string, value any, opts ...Option) (int64, error) {
	model := new(T)
	result := g.SessionContext(ctx, opts...).Model(model).Update(column, value)
	return result.RowsAffected, result.Error
}

func (g *GormManager[T]) Updates(update T, opts ...Option) (int64, error) {
	return g.UpdatesContext(context.Background(), update, opts...)
}

func (g *GormManager[T]) UpdatesContext(ctx context.Context, update T, opts ...Option) (int64, error) {
	model := new(T)
	result := g.SessionContext(ctx, opts...).Model(model).Updates(update)
	return result.RowsAffected, result.Error
}

func (g *GormManager[T]) UpdateColumn(column string, value any, opts ...Option) (int64, error) {
	return g.UpdateColumnContext(context.Background(), column, value, opts...)
}

func (g *GormManager[T]) UpdateColumnContext(ctx context.Context, column string, value any, opts ...Option) (int64, error) {
	model := new(T)
	result := g.SessionContext(ctx, opts...).Model(model).UpdateColumn(column, value)
	return result.RowsAffected, result.Error
}

func (g *GormManager[T]) UpdateColumns(update T, opts ...Option) (int64, error) {
	return g.UpdateColumnsContext(context.Background(), update, opts...)
}

func (g *GormManager[T]) UpdateColumnsContext(ctx context.Context, update T, opts ...Option) (int64, error) {
	model := new(T)
	result := g.SessionContext(ctx, opts...).Model(model).UpdateColumns(update)
	return result.RowsAffected, result.Error
}

func (g *GormManager[T]) Delete(opts ...Option) error {
	return g.DeleteContext(context.Background(), opts...)
}

func (g *GormManager[T]) DeleteContext(ctx context.Context, opts ...Option) error {
	model := new(T)
	return g.SessionContext(ctx, opts...).Delete(model).Error
}

func (g *GormManager[T]) Find(opts ...Option) ([]T, error) {
	return g.FindContext(context.Background(), opts...)
}

func (g *GormManager[T]) FindContext(ctx context.Context, opts ...Option) ([]T, error) {
	var models []T
	return models, g.SessionContext(ctx, opts...).Find(&models).Error
}

func (g *GormManager[T]) Count(opts ...Option) (int64, error) {
	return g.CountContext(context.Background(), opts...)
}

func (g *GormManager[T]) CountContext(ctx context.Context, opts ...Option) (int64, error) {
	var count int64
	model := new(T)
	return count, g.SessionContext(ctx, opts...).Model(model).Count(&count).Error
}

func (g *GormManager[T]) Transaction(f func(tx *gorm.DB) error) error {
	return g.TransactionContext(context.Background(), f)
}

func (g *GormManager[T]) TransactionContext(ctx context.Context, f func(tx *gorm.DB) error) error {
	return g.SessionContext(ctx).Transaction(f)
}

func (g *GormManager[T]) Scan(dst any, opts ...Option) error {
	return g.ScanContext(context.Background(), dst, opts...)
}

func (g *GormManager[T]) ScanContext(ctx context.Context, dst any, opts ...Option) error {
	return g.SessionContext(ctx, opts...).Scan(dst).Error
}

func (g *GormManager[T]) Rows(opts ...Option) (*sql.Rows, error) {
	return g.RowsContext(context.Background(), opts...)
}

func (g *GormManager[T]) RowsContext(ctx context.Context, opts ...Option) (*sql.Rows, error) {
	return g.SessionContext(ctx, opts...).Rows()
}

func (g *GormManager[T]) Exec(sql string, args ...any) error {
	return g.ExecContext(context.Background(), sql, args...)
}

func (g *GormManager[T]) ExecContext(ctx context.Context, sql string, args ...any) error {
	return g.SessionContext(ctx).Exec(sql, args...).Error
}

func (g *GormManager[T]) AssociationFind(column string, dst any, opts ...Option) error {
	return g.AssociationFindContext(context.Background(), column, dst, opts...)
}

func (g *GormManager[T]) AssociationFindContext(ctx context.Context, column string, dst any, opts ...Option) error {
	return g.SessionContext(ctx, opts...).Association(column).Find(dst)
}

func (g *GormManager[T]) AssociationAppend(column string, dst []any, opts ...Option) error {
	return g.AssociationAppendContext(context.Background(), column, dst, opts...)
}

func (g *GormManager[T]) AssociationAppendContext(ctx context.Context, column string, dst []any, opts ...Option) error {
	return g.SessionContext(ctx, opts...).Association(column).Append(dst...)
}

func (g *GormManager[T]) AssociationReplace(column string, dst []any, opts ...Option) error {
	return g.AssociationReplaceContext(context.Background(), column, dst, opts...)
}

func (g *GormManager[T]) AssociationReplaceContext(ctx context.Context, column string, dst []any, opts ...Option) error {
	return g.SessionContext(ctx, opts...).Association(column).Replace(dst...)
}

func (g *GormManager[T]) AssociationDelete(column string, dst []any, opts ...Option) error {
	return g.AssociationDeleteContext(context.Background(), column, dst, opts...)
}

func (g *GormManager[T]) AssociationDeleteContext(ctx context.Context, column string, dst []any, opts ...Option) error {
	return g.SessionContext(ctx, opts...).Association(column).Delete(dst...)
}

func (g *GormManager[T]) AssociationClear(column string, opts ...Option) error {
	return g.AssociationClearContext(context.Background(), column, opts...)
}

func (g *GormManager[T]) AssociationClearContext(ctx context.Context, column string, opts ...Option) error {
	return g.SessionContext(ctx, opts...).Association(column).Clear()
}

func (g *GormManager[T]) AssociationCount(column string, opts ...Option) int64 {
	return g.AssociationCountContext(context.Background(), column, opts...)
}

func (g *GormManager[T]) AssociationCountContext(ctx context.Context, column string, opts ...Option) int64 {
	return g.SessionContext(ctx, opts...).Association(column).Count()
}

func NewGormManager[T any](m *T, db *gorm.DB) GormManager[T] {
	return GormManager[T]{model: m, db: db}
}
