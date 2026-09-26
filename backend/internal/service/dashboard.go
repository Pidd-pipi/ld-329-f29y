package service

import (
	"cyskillswap/internal/constants"
	"cyskillswap/internal/model"
	"cyskillswap/internal/repository"
)

func Overview() model.Overview {
	skills := repository.ListSkills()
	needs := repository.ListNeeds()
	matches := repository.ListMatches()
	appointments := repository.ListAppointments()
	orders := repository.ListSwapOrders()
	reviews := repository.ListReviews()
	messages := repository.ListMessages()
	responseTotal := 0
	for _, need := range needs {
		responseTotal += need.Responses
	}
	return model.Overview{
		Service:    constants.ServiceName,
		Categories: constants.SkillCategories,
		Metrics: map[string]int{
			"skills": len(skills), "needs": len(needs), "matches": len(matches),
			"appointments": len(appointments), "reviews": len(reviews), "unread": 3,
			"responses": responseTotal, "orders": len(orders),
		},
		Skills: skills, Needs: needs, Matches: matches,
		Appointments: appointments, SwapOrders: orders,
		Reviews: reviews, Messages: messages,
		Profile: repository.GetProfile(),
	}
}

func Skills() []model.Skill             { return repository.ListSkills() }
func Needs() []model.Need               { return repository.ListNeeds() }
func Matches() []model.Match            { return repository.ListMatches() }
func Appointments() []model.Appointment { return repository.ListAppointments() }
func Reviews() []model.Review           { return repository.ListReviews() }
func Messages() []model.Conversation    { return repository.ListMessages() }
func Profile() model.Profile            { return repository.GetProfile() }
