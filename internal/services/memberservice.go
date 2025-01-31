package services

import (
	"github.com/nyae44/GoLibrary/internal/models"
	"github.com/nyae44/GoLibrary/internal/repositories"
)

type memberService struct {
	memberRepo repositories.MemberRepository
}

func NewMemberService(memberRepo repositories.MemberRepository) *memberService {
	return &memberService{memberRepo: memberRepo}
}

func (s *memberService) CreateMember(member *models.Member) error {
	err := s.memberRepo.Create(member)
	if err != nil {
		return err
	}
	return nil
}

func (s *memberService) UpdateMember(member *models.Member) error {
	err, _ := s.memberRepo.Update(member)
	if err != nil {
		return err
	}
	return nil
}

func (s *memberService) DeleteMember(id uint) error {
	err := s.memberRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
func (s *memberService) GetMember(id uint) (*models.Member, error) {
	return s.memberRepo.GetById(id)
}

func (s *memberService) GetMemberByName(name string) (*models.Member, error) {
	return s.memberRepo.GetByName(name)
}

func (s *memberService) ListMembers() ([]*models.Member, error) {
	return s.memberRepo.GetAllMembers()
}
