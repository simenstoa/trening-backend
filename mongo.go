package main

import (
	"gopkg.in/mgo.v2"
)

type Session struct {
	session *mgo.Session
}

func NewSession() (*Session, error) {
	session, err := mgo.Dial("host.docker.internal:27017")
	if err != nil {
		return nil, err
	}
	return &Session{session}, err
}

func (s *Session) Copy() *Session {
	return &Session{s.session.Copy()}
}

func (s *Session) GetCollection(db string, col string) *mgo.Collection {
	return s.session.DB(db).C(col)
}

func (s *Session) Close() {
	if s.session != nil {
		s.session.Close()
	}
}
