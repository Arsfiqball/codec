package gormrepo

import (
	"context"

	"github.com/Arsfiqball/codec/internal/persistence/gormrepo/mapper"
	"github.com/Arsfiqball/codec/internal/value/domain"

	"github.com/Arsfiqball/talkback"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type Domain struct {
	tracer trace.Tracer
	db     *gorm.DB
}

func NewDomain(
	tracer trace.Tracer,
	db *gorm.DB,
) *Domain {
	return &Domain{
		tracer: tracer,
		db:     db,
	}
}

func (d *Domain) Create(ctx context.Context, ent domain.Entity) (domain.Entity, error) {
	ctx, span := d.tracer.Start(ctx, "github.com/Arsfiqball/codec/internal/persistence/gormrepo/domain/Create")
	defer span.End()

	row := mapper.NewDomainDAO(ent)

	if err := d.db.WithContext(ctx).Create(&row).Error; err != nil {
		return domain.Entity{}, err
	}

	return row.ToEntity(), nil
}

func (d *Domain) Update(ctx context.Context, ent domain.Entity) (domain.Entity, error) {
	ctx, span := d.tracer.Start(ctx, "github.com/Arsfiqball/codec/internal/persistence/gormrepo/domain/Update")
	defer span.End()

	row := mapper.NewDomainDAO(ent)

	if err := d.db.WithContext(ctx).Save(&row).Error; err != nil {
		return domain.Entity{}, err
	}

	return row.ToEntity(), nil
}

func (d *Domain) Delete(ctx context.Context, ent domain.Entity) (domain.Entity, error) {
	ctx, span := d.tracer.Start(ctx, "github.com/Arsfiqball/codec/internal/persistence/gormrepo/domain/Delete")
	defer span.End()

	row := mapper.NewDomainDAO(ent)

	if err := d.db.WithContext(ctx).Delete(&row).Error; err != nil {
		return domain.Entity{}, err
	}

	return row.ToEntity(), nil
}

func (d *Domain) GetOne(ctx context.Context, q domain.Query) (domain.Entity, error) {
	ctx, span := d.tracer.Start(ctx, "github.com/Arsfiqball/codec/internal/persistence/gormrepo/domain/GetOne")
	defer span.End()

	q.Limit = 1

	ents, err := d.GetList(ctx, q)
	if err != nil {
		return domain.Entity{}, err
	}

	if len(ents) == 0 {
		return domain.Entity{}, domain.ErrNotFound
	}

	return ents[0], nil
}

func (d *Domain) GetList(ctx context.Context, q domain.Query) ([]domain.Entity, error) {
	ctx, span := d.tracer.Start(ctx, "github.com/Arsfiqball/codec/internal/persistence/gormrepo/domain/GetList")
	defer span.End()

	var (
		rows []mapper.DomainDAO
		ents []domain.Entity
	)

	plan, err := talkback.ToSqlPlan(q.Query, domainTranslations, domainPreloadable)
	if err != nil {
		return nil, err
	}

	tx := d.db.WithContext(ctx).Model(&mapper.DomainDAO{})

	if plan.Where != "" {
		tx = tx.Where(plan.Where, plan.WhereArgs...)
	}

	if q.Search != "" {
		tx = tx.Where("name LIKE ?", "%"+q.Search+"%")
	}

	tx = tx.Select("*")

	if plan.Order != "" {
		tx = tx.Order(plan.Order)
	}

	if plan.Limit > 0 {
		tx = tx.Limit(plan.Limit)
	} else {
		tx = tx.Limit(100)
	}

	if plan.Offset > 0 {
		tx = tx.Offset(plan.Offset)
	} else {
		tx = tx.Offset(0)
	}

	for _, preload := range plan.Preload {
		tx = tx.Preload(preload)
	}

	if err := tx.Find(&rows).Error; err != nil {
		return nil, err
	}
	for _, row := range rows {
		ents = append(ents, row.ToEntity())
	}

	return ents, nil
}

func (d *Domain) GetStat(ctx context.Context, query domain.Query) ([]domain.Stat, error) {
	ctx, span := d.tracer.Start(ctx, "github.com/Arsfiqball/codec/internal/persistence/gormrepo/domain/GetStat")
	defer span.End()

	var (
		ents []domain.Stat
		rows []struct {
			Name  string `gorm:"column:name"`
			Count int    `gorm:"column:count"`
		}
	)

	plan, err := talkback.ToSqlPlan(query.Query, domainTranslations, domainPreloadable)
	if err != nil {
		return nil, err
	}

	tx := d.db.WithContext(ctx).Model(&mapper.DomainDAO{})

	if plan.Where != "" {
		tx = tx.Where(plan.Where, plan.WhereArgs...)
	}

	if query.Search != "" {
		tx = tx.Where("name LIKE ?", "%"+query.Search+"%")
	}

	if plan.Group != "" {
		tx = tx.Group(plan.Group)
	}

	if plan.Select != "" {
		tx = tx.Select(plan.Select)
	} else {
		return nil, domain.ErrMissingAccumulator
	}

	if plan.Order != "" {
		tx = tx.Order(plan.Order)
	}

	if plan.Limit > 0 {
		tx = tx.Limit(plan.Limit)
	} else {
		tx = tx.Limit(100)
	}

	if plan.Offset > 0 {
		tx = tx.Offset(plan.Offset)
	} else {
		tx = tx.Offset(0)
	}

	if err := tx.Find(&rows).Error; err != nil {
		return nil, err
	}

	for _, row := range rows {
		ents = append(ents, domain.NewStat(row.Name, row.Count))
	}

	return ents, nil
}
