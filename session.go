package v3io

import (
	"github.com/nuclio/logger"
	"github.com/pkg/errors"
)

type Session struct {
	Sync    *SyncSession
	logger  logger.Logger
	context *Context
}

func newSession(parentLogger logger.Logger,
	context *Context,
	username string,
	password string,
	label string) (*Session, error) {

	newSyncSession, err := newSyncSession(parentLogger, context.Sync, username, password, label)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create sync session")
	}

	return &Session{
		logger:  parentLogger.GetChild("session"),
		context: context,
		Sync:    newSyncSession,
	}, nil
}

func (s *Session) NewContainer(alias string) (*Container, error) {
	return newContainer(s.logger, s, alias)
}

func (s *Session) sendRequest(request *Request) error {

	// set session
	request.session = s

	// delegate to context
	return s.context.sendRequest(request)
}
