package sessions

import (
	"errors"
)

type Store struct {
	sessions map[SessionId]*Session
}

func NewStore() *Store {
	return &Store{
		sessions: map[SessionId]*Session{},
	}
}

func (s *Store) Add(cs *Session) (sessionId SessionId) {
	sessionId = NewSessionId()
	s.sessions[sessionId] = cs
	return
}

// TODO: iterate sessions until gets 0
// TODO: account the usage
// TODO: timeout with context
func (s *Store) Iterate(sessionId SessionId) error {
	var (
		session *Session
		ok      bool
	)
	if session, ok = s.sessions[sessionId]; !ok {
		return errors.New("")
	}

	session.EvolutionManager.InitPopulation(session.Config.Population)
	// TODO: make request to runner
	session.EvolutionManager.IterateLoop()

	return nil
}
