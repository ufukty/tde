package session_post

import (
	"tde/internal/evolution"
	"tde/internal/microservices/errors/bucket"
)

type ArchiveStatus int

const (
	Accessible = ArchiveStatus(iota)
	Unauthorized
	NotFound
	Lost
)

const (
	MsgEitherDoesNotExistsOrYouDoNotHaveAccessToIt = "You don't have such archive"
)

func downloadAST(accessToken, archiveId string) (*evolution.Target, error) {
	return &evolution.Target{}, nil
}

func validateArchiveID(archiveID string) ArchiveStatus {
	return NotFound
}

func Controller(request *Request) (response *Response, errs *bucket.Bucket) {
	// var (
	// 	err error
	// )
	// {
	// 	switch validateArchiveID(request.ArchiveID) {
	// 	case NotFound:
	// 		errs.Add(detailed.New(MsgEitherDoesNotExistsOrYouDoNotHaveAccessToIt, "archive doesn't exists"))
	// 	case Unauthorized:
	// 		errs.Add(detailed.New(MsgEitherDoesNotExistsOrYouDoNotHaveAccessToIt, "archive exists but no auth"))
	// 	}

	// 	if errs.IsAny() {
	// 		return
	// 	}
	// }

	// var (
	// 	timeout time.Duration
	// 	config  *sessions.Config
	// )
	// {
	// 	timeout, err = time.ParseDuration(request.Timeout)
	// 	if err != nil {
	// 		errs.Add(detailed.AddBase(err, "Bad timeout_seconds value"))
	// 	}
	// 	config = &sessions.Config{
	// 		Runner: sessions.Runner{
	// 			Address: request.Runner.Address,
	// 			Token:   request.Runner.Token,
	// 		},
	// 		Probabilities: sessions.Probabilities{
	// 			CrossOver: request.Probabilities.CrossOver,
	// 			Mutation:  request.Probabilities.Mutation,
	// 		},
	// 		Timeout:         timeout,
	// 		Population:      request.Population,
	// 		SizeLimit:       request.SizeLimitBytes,
	// 		AllowedPackages: request.AllowedPackages,
	// 		Iterate:         request.Iterate,
	// 		TestName:        request.TestName,
	// 	}

	// 	if errs.IsAny() {
	// 		return
	// 	}
	// }

	// var target *evolution.Target
	// target, err = downloadAST("", request.ArchiveID)
	// if err != nil {
	// 	errs.Add(detailed.AddBase(err, "Could not access to archive right now."))
	// 	return
	// }

	// var (
	// 	em        = evolution.NewManager(&models.Parameters{}, &evaluation.Evaluator{})
	// 	session   = sessions.NewSession(em, config)
	// 	sessionId = sessionStore.Add(session)
	// )

	// go sessionManager.Iterate(caseId) // async
	// return &Response{
	// 	SessionId: string(sessionId),
	// 	Status:    session.Status,
	// }, nil

	return nil, nil
}
