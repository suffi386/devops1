package org

import (
	"github.com/caos/zitadel/internal/eventstore/v2"
	"github.com/caos/zitadel/internal/v2/repository/policy/login"
)

var (
	LoginPolicyAddedEventType   = orgEventTypePrefix + login.LoginPolicyAddedEventType
	LoginPolicyChangedEventType = orgEventTypePrefix + login.LoginPolicyChangedEventType
)

type LoginPolicyReadModel struct{ login.LoginPolicyReadModel }

func (rm *LoginPolicyReadModel) AppendEvents(events ...eventstore.EventReader) {
	for _, event := range events {
		switch e := event.(type) {
		case *LoginPolicyAddedEvent:
			rm.ReadModel.AppendEvents(&e.LoginPolicyAddedEvent)
		case *LoginPolicyChangedEvent:
			rm.ReadModel.AppendEvents(&e.LoginPolicyChangedEvent)
		case *login.LoginPolicyAddedEvent, *login.LoginPolicyChangedEvent:
			rm.ReadModel.AppendEvents(e)
		}
	}
}

type LoginPolicyAddedEvent struct {
	login.LoginPolicyAddedEvent
}

type LoginPolicyChangedEvent struct {
	login.LoginPolicyChangedEvent
}
