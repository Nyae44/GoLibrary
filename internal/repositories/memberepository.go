package repositories

import (
	"github.com/nyae44/GoLibrary/internal/models"
	"gorm.io/gorm"
)

type MemberRepository interface {
	Create(member *models.Member) error
	Update(member *models.Member) (error, error)
	Delete(id uint) error
	GetById(id uint) (*models.Member, error)
	GetByName(name string) (*models.Member, error)
	GetAllMembers() ([]*models.Member, error)
}

type memberRepository struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) MemberRepository {
	return &memberRepository{db: db}
}

func (r *memberRepository) Create(member *models.Member) error {
	return r.db.Create(member).Error
}

func (r *memberRepository) Update(member *models.Member) (error, error) {
	return r.db.Save(member).Error, nil
}

func (r *memberRepository) Delete(id uint) error {
	return r.db.Delete(&models.Member{}, id).Error
}
func (r *memberRepository) GetById(id uint) (*models.Member, error) {
	var member models.Member
	err := r.db.First(&member, id).Error
	return &member, err
}

func (r *memberRepository) GetByName(name string) (*models.Member, error) {
	var member models.Member
	err := r.db.Where("name = ?", name).First(&member).Error
	return &member, err
}

func (r *memberRepository) GetAllMembers() ([]*models.Member, error) {
	var members []*models.Member
	err := r.db.Find(&members).Error
	return members, err

}
