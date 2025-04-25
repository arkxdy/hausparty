package handlers

import "hauparty/services/party-service/services"

type PartyHandler struct {
	partyService services.PartyService
}

func NewPartyHandler(s services.PartyService) *PartyHandler {
	return &PartyHandler{
		partyService: s,
	}
}
