package entities

import (
	"github.com/eure/si2018-server-side/models"
	"github.com/go-openapi/strfmt"
)

type MatchUserResponse struct {
	User
	MatchedAt strfmt.DateTime
}

func (res MatchUserResponse) Build() models.MatchUserResponse {
	return models.MatchUserResponse{
		MatchedAt:      res.MatchedAt,
		ID:             res.ID,
		Gender:         res.Gender,
		Birthday:       res.Birthday,
		Nickname:       res.Nickname,
		ImageURI:       res.ImageURI,
		Tweet:          res.Tweet,
		Introduction:   res.Introduction,
		ResidenceState: res.ResidenceState,
		HomeState:      res.HomeState,
		Education:      res.Education,
		Job:            res.Job,
		AnnualIncome:   res.AnnualIncome,
		Height:         res.Height,
		BodyBuild:      res.BodyBuild,
		MaritalStatus:  res.MaritalStatus,
		Child:          res.Child,
		WhenMarry:      res.WhenMarry,
		WantChild:      res.WantChild,
		Smoking:        res.Smoking,
		Drinking:       res.Drinking,
		Holiday:        res.Holiday,
		HowToMeet:      res.HowToMeet,
		CostOfDate:     res.CostOfDate,
		NthChild:       res.NthChild,
		Housework:      res.Housework,
		CreatedAt:      res.CreatedAt,
		UpdatedAt:      res.UpdatedAt,
	}
}

func (res *MatchUserResponse) ApplyUser(u User) {
	res.ID = u.ID
	res.Gender = u.Gender
	res.Birthday = u.Birthday
	res.Nickname = u.Nickname
	res.ImageURI = u.ImageURI
	res.Tweet = u.Tweet
	res.Introduction = u.Introduction
	res.ResidenceState = u.ResidenceState
	res.HomeState = u.HomeState
	res.Education = u.Education
	res.Job = u.Job
	res.AnnualIncome = u.AnnualIncome
	res.Height = u.Height
	res.BodyBuild = u.BodyBuild
	res.MaritalStatus = u.MaritalStatus
	res.Child = u.Child
	res.WhenMarry = u.WhenMarry
	res.WantChild = u.WantChild
	res.Smoking = u.Smoking
	res.Drinking = u.Drinking
	res.Holiday = u.Holiday
	res.HowToMeet = u.HowToMeet
	res.CostOfDate = u.CostOfDate
	res.NthChild = u.NthChild
	res.Housework = u.Housework
	res.CreatedAt = u.CreatedAt
	res.UpdatedAt = u.UpdatedAt
}

type MatchUserResponses []MatchUserResponse

func (responses *MatchUserResponses) Build() []*models.MatchUserResponse {
	var sResponses []*models.MatchUserResponse

	for _, response := range *responses {
		sResponse := response.Build()
		sResponses = append(sResponses, &sResponse)
	}
	return sResponses
}
