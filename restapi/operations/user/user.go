package user

import (
	"github.com/eure/si2018-server-side/entities"
	"github.com/eure/si2018-server-side/models"
	"github.com/eure/si2018-server-side/repositories"
	si "github.com/eure/si2018-server-side/restapi/summerintern"
	"github.com/go-openapi/runtime/middleware"
)

func GetUsers(p si.GetUsersParams) middleware.Responder {
	usertokenHandler := repositories.NewUserTokenRepository()
	userlikeHandler := repositories.NewUserLikeRepository()
	userHandler := repositories.NewUserRepository()
	//ngur := repositories.NewGetUserRepository()
	// find userid
	token := p.Token
	limit := int(p.Limit)
	offset := int(p.Offset)

	// validate limit, offset
	if limit < offset {
		return GetProfileBadRequestErr()
	}

	usertoken, err := usertokenHandler.GetByToken(token)
	if err != nil {
		return GetUserRespUnauthErr()
	}
	// Is There a collect user token?
	if usertoken == nil {
		return GetProfileBadRequestErr()
	}
	// find userlike
	userlike, err := userlikeHandler.FindLikeAll(usertoken.UserID)
	if err != nil {
		return GetUserRespInternalErr()
	}
	// find user
	userprofile, err := userHandler.GetByUserID(usertoken.UserID)
	if err != nil {
		return GetUserRespInternalErr()
	}
	oppositeGenger := userprofile.GetOppositeGender()
	ent, err := userHandler.FindWithCondition(limit, offset, oppositeGenger, userlike)
	if err != nil {
		return GetUserRespInternalErr()
	}

	var ud []*models.User
	for i := 0; i < len(ent); i++ {
		us := ent[i].Build()
		ud = append(ud, &us)
	}

	return si.NewGetUsersOK().WithPayload(ud)

}

func GetProfileByUserID(p si.GetProfileByUserIDParams) middleware.Responder {
	userHandler := repositories.NewUserRepository()
	usertokenHandler := repositories.NewUserTokenRepository()
	token := p.Token
	myuserid := p.UserID

	userprofile, err := userHandler.GetByUserID(myuserid)
	if err != nil {
		return GetProfileInternalErr()
	}
	/*
		userprofile, err := userHandler.GetByUserIDWithImage(myuserid)
		fmt.Println(userprofile)
	*/
	// Is There a exist UserProfile?
	if userprofile == nil {
		return GetProfileNotFoundErr()
	}
	usertoken, err := usertokenHandler.GetByToken(token)
	if err != nil {
		return GetProfileRespUnauthErr()

	}
	// Is There a collect user token?
	if usertoken == nil {
		return GetProfileBadRequestErr()
	}
	// Is token and userid is match?
	if userprofile.ID != usertoken.UserID {
		return GetProfileBadRequestErr()
	}
	sEnt := userprofile.Build()

	return si.NewGetProfileByUserIDOK().WithPayload(&sEnt)
}

func PutProfile(p si.PutProfileParams) middleware.Responder {
	nur := repositories.NewUserRepository()
	nutr := repositories.NewUserTokenRepository()
	userID := p.UserID
	putParams := p.Params
	//Find user
	user, err := nur.GetByUserID(userID)
	if err != nil {
		return PutProfileInternalErr()
	}
	// Is there a User through token
	usertoken, err := nutr.GetByToken(putParams.Token)
	if err != nil {
		return PutProfileInternalErr()
	}
	// Is threre token Authorized
	if usertoken == nil {
		return PutProfileRespUnauthErr()
	}

	if usertoken.UserID != userID {
		return PutProfileForbiddenErr()
	}
	binduser(putParams, user)
	err = nur.Update(user)
	if err != nil {
		return PutProfileInternalErr()
	}
	// Is update User profile
	respuser, err := nur.GetByUserID(userID)
	if err != nil {
		return PutProfileInternalErr()
	}
	updateuser := respuser.Build()
	return si.NewPutProfileOK().WithPayload(&updateuser)
}

func binduser(user si.PutProfileBody, ent *entities.User) {
	ent.AnnualIncome = user.AnnualIncome
	ent.BodyBuild = user.BodyBuild
	ent.Child = user.Child
	ent.CostOfDate = user.CostOfDate
	ent.Drinking = user.Drinking
	ent.Education = user.Education
	ent.Height = user.Height
	ent.Holiday = user.Holiday
	ent.HomeState = user.HomeState
	ent.Housework = user.Housework
	ent.HowToMeet = user.HowToMeet
	ent.ImageURI = user.ImageURI
	ent.Introduction = user.Introduction
	ent.Job = user.Job
	ent.MaritalStatus = user.MaritalStatus
	ent.Nickname = user.Nickname
	ent.NthChild = user.NthChild
	ent.ResidenceState = user.ResidenceState
	ent.Smoking = user.Smoking
	ent.Tweet = user.Tweet
	ent.WantChild = user.WantChild
	ent.WhenMarry = user.WhenMarry

}

// return 400 Bad Request
func GetUserRespBadReqestErr() middleware.Responder {
	return si.NewGetUsersBadRequest().WithPayload(
		&si.GetUsersBadRequestBody{
			Code:    "400",
			Message: "Bad Request",
		})
}

// return 401 Token Is Invalid
func GetUserRespUnauthErr() middleware.Responder {
	return si.NewGetUsersUnauthorized().WithPayload(
		&si.GetUsersUnauthorizedBody{
			Code:    "401",
			Message: "Token Is Invalid",
		})
}

// return 50 Internal Server Error
func GetUserRespInternalErr() middleware.Responder {
	return si.NewGetUsersInternalServerError().WithPayload(
		&si.GetUsersInternalServerErrorBody{
			Code:    "500",
			Message: "Internal Server Error",
		})
}

// return 400 Bad Request
func PutProfileBadRequestErr() middleware.Responder {
	return si.NewPutProfileBadRequest().WithPayload(
		&si.PutProfileBadRequestBody{
			Code:    "400",
			Message: "Bad Request",
		})
}

// return 401 Token Is Invalid
func PutProfileRespUnauthErr() middleware.Responder {
	return si.NewPutProfileUnauthorized().WithPayload(
		&si.PutProfileUnauthorizedBody{
			Code:    "401",
			Message: "Token Is Invalid",
		})
}

// return 403 Forbidden
func PutProfileForbiddenErr() middleware.Responder {
	return si.NewPutProfileForbidden().WithPayload(
		&si.PutProfileForbiddenBody{
			Code:    "403",
			Message: "Forbidden",
		})
}

// return 500 Internal Server Error
func PutProfileInternalErr() middleware.Responder {
	return si.NewPutProfileInternalServerError().WithPayload(
		&si.PutProfileInternalServerErrorBody{
			Code:    "500",
			Message: "Internal Server Error",
		})
}

// return 400 bad Request
func GetProfileBadRequestErr() middleware.Responder {
	return si.NewGetProfileByUserIDBadRequest().WithPayload(
		&si.GetProfileByUserIDBadRequestBody{
			Code:    "400",
			Message: "Bad Request",
		})
}

// return 401 Token Is Invalid
func GetProfileRespUnauthErr() middleware.Responder {
	return si.NewGetProfileByUserIDUnauthorized().WithPayload(
		&si.GetProfileByUserIDUnauthorizedBody{
			Code:    "401",
			Message: "Token Is Invalid",
		})
}

// return 404 User Not Found
func GetProfileNotFoundErr() middleware.Responder {
	return si.NewGetProfileByUserIDNotFound().WithPayload(
		&si.GetProfileByUserIDNotFoundBody{
			Code:    "404",
			Message: "User Not Found",
		})
}

// return 500 Internal Server Error
func GetProfileInternalErr() middleware.Responder {
	return si.NewGetProfileByUserIDInternalServerError().WithPayload(
		&si.GetProfileByUserIDInternalServerErrorBody{
			Code:    "500",
			Message: "Internal Server Error",
		})
}
