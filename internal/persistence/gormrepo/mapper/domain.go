package mapper

import "github.com/Arsfiqball/codec/internal/value/domain"

type DomainDAO struct {
	ID       string `gorm:"column:id,primaryKey"`
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}

func NewDomainDAO(ent domain.Entity) DomainDAO {
	return DomainDAO{
		ID:       ent.ID(),
		Name:     ent.Name(),
		Email:    ent.Email(),
		Password: ent.Password(),
	}
}

func (DomainDAO) TableName() string {
	return "domains"
}

func (dao DomainDAO) ToEntity() domain.Entity {
	return domain.NewEntityWithData(
		dao.ID,
		dao.Name,
		dao.Email,
		dao.Password,
	)
}

type DomainStatDAO struct {
	Name  string `gorm:"column:name"`
	Count int    `gorm:"column:count"`
}

func (DomainStatDAO) TableName() string {
	return "domains"
}

func (dao DomainStatDAO) ToEntity() domain.Stat {
	return domain.NewStat(
		dao.Name,
		dao.Count,
	)
}
