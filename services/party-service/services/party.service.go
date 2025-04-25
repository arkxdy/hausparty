package services

import "hausparty/libs/db/repository"

type PartyService interface {
}

type partyService struct {
	repo repository.IPartyRepository
}

func NewPartyService(repo repository.IPartyRepository) PartyService {
	return &partyService{repo: repo}
}
