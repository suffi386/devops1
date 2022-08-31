package query

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/zitadel/zitadel/internal/domain"
	errs "github.com/zitadel/zitadel/internal/errors"
)

func Test_ActionPrepares(t *testing.T) {
	type want struct {
		sqlExpectations sqlExpectation
		err             checkErr
	}
	tests := []struct {
		name    string
		prepare interface{}
		want    want
		object  interface{}
	}{
		{
			name:    "prepareActionsQuery no result",
			prepare: prepareActionsQuery,
			want: want{
				sqlExpectations: mockQueries(
					regexp.QuoteMeta(`SELECT projections.actions2.id,`+
						` projections.actions2.creation_date,`+
						` projections.actions2.change_date,`+
						` projections.actions2.resource_owner,`+
						` projections.actions2.sequence,`+
						` projections.actions2.action_state,`+
						` projections.actions2.name,`+
						` projections.actions2.script,`+
						` projections.actions2.timeout,`+
						` projections.actions2.allowed_to_fail,`+
						` COUNT(*) OVER ()`+
						` FROM projections.actions2`),
					nil,
					nil,
				),
			},
			object: &Actions{Actions: []*Action{}},
		},
		{
			name:    "prepareActionsQuery one result",
			prepare: prepareActionsQuery,
			want: want{
				sqlExpectations: mockQueries(
					regexp.QuoteMeta(`SELECT projections.actions2.id,`+
						` projections.actions2.creation_date,`+
						` projections.actions2.change_date,`+
						` projections.actions2.resource_owner,`+
						` projections.actions2.sequence,`+
						` projections.actions2.action_state,`+
						` projections.actions2.name,`+
						` projections.actions2.script,`+
						` projections.actions2.timeout,`+
						` projections.actions2.allowed_to_fail,`+
						` COUNT(*) OVER ()`+
						` FROM projections.actions2`),
					[]string{
						"id",
						"creation_date",
						"change_date",
						"resource_owner",
						"sequence",
						"action_state",
						"name",
						"script",
						"timeout",
						"allowed_to_fail",
						"count",
					},
					[][]driver.Value{
						{
							"id",
							testNow,
							testNow,
							"ro",
							uint64(20211109),
							domain.ActionStateActive,
							"action-name",
							"script",
							1 * time.Second,
							true,
						},
					},
				),
			},
			object: &Actions{
				SearchResponse: SearchResponse{
					Count: 1,
				},
				Actions: []*Action{
					{
						ID:            "id",
						CreationDate:  testNow,
						ChangeDate:    testNow,
						ResourceOwner: "ro",
						State:         domain.ActionStateActive,
						Sequence:      20211109,
						Name:          "action-name",
						Script:        "script",
						Timeout:       1 * time.Second,
						AllowedToFail: true,
					},
				},
			},
		},
		{
			name:    "prepareActionsQuery multiple result",
			prepare: prepareActionsQuery,
			want: want{
				sqlExpectations: mockQueries(
					regexp.QuoteMeta(`SELECT projections.actions2.id,`+
						` projections.actions2.creation_date,`+
						` projections.actions2.change_date,`+
						` projections.actions2.resource_owner,`+
						` projections.actions2.sequence,`+
						` projections.actions2.action_state,`+
						` projections.actions2.name,`+
						` projections.actions2.script,`+
						` projections.actions2.timeout,`+
						` projections.actions2.allowed_to_fail,`+
						` COUNT(*) OVER ()`+
						` FROM projections.actions2`),
					[]string{
						"id",
						"creation_date",
						"change_date",
						"resource_owner",
						"sequence",
						"action_state",
						"name",
						"script",
						"timeout",
						"allowed_to_fail",
						"count",
					},
					[][]driver.Value{
						{
							"id-1",
							testNow,
							testNow,
							"ro",
							uint64(20211109),
							domain.ActionStateActive,
							"action-name-1",
							"script",
							1 * time.Second,
							true,
						},
						{
							"id-2",
							testNow,
							testNow,
							"ro",
							uint64(20211109),
							domain.ActionStateActive,
							"action-name-2",
							"script",
							1 * time.Second,
							true,
						},
					},
				),
			},
			object: &Actions{
				SearchResponse: SearchResponse{
					Count: 2,
				},
				Actions: []*Action{
					{
						ID:            "id-1",
						CreationDate:  testNow,
						ChangeDate:    testNow,
						ResourceOwner: "ro",
						State:         domain.ActionStateActive,
						Sequence:      20211109,
						Name:          "action-name-1",
						Script:        "script",
						Timeout:       1 * time.Second,
						AllowedToFail: true,
					},
					{
						ID:            "id-2",
						CreationDate:  testNow,
						ChangeDate:    testNow,
						ResourceOwner: "ro",
						State:         domain.ActionStateActive,
						Sequence:      20211109,
						Name:          "action-name-2",
						Script:        "script",
						Timeout:       1 * time.Second,
						AllowedToFail: true,
					},
				},
			},
		},
		{
			name:    "prepareActionsQuery sql err",
			prepare: prepareActionsQuery,
			want: want{
				sqlExpectations: mockQueryErr(
					regexp.QuoteMeta(`SELECT projections.actions2.id,`+
						` projections.actions2.creation_date,`+
						` projections.actions2.change_date,`+
						` projections.actions2.resource_owner,`+
						` projections.actions2.sequence,`+
						` projections.actions2.action_state,`+
						` projections.actions2.name,`+
						` projections.actions2.script,`+
						` projections.actions2.timeout,`+
						` projections.actions2.allowed_to_fail,`+
						` COUNT(*) OVER ()`+
						` FROM projections.actions2`),
					sql.ErrConnDone,
				),
				err: func(err error) (error, bool) {
					if !errors.Is(err, sql.ErrConnDone) {
						return fmt.Errorf("err should be sql.ErrConnDone got: %w", err), false
					}
					return nil, true
				},
			},
			object: nil,
		},
		{
			name:    "prepareActionQuery no result",
			prepare: prepareActionQuery,
			want: want{
				sqlExpectations: mockQueries(
					regexp.QuoteMeta(`SELECT projections.actions2.id,`+
						` projections.actions2.creation_date,`+
						` projections.actions2.change_date,`+
						` projections.actions2.resource_owner,`+
						` projections.actions2.sequence,`+
						` projections.actions2.action_state,`+
						` projections.actions2.name,`+
						` projections.actions2.script,`+
						` projections.actions2.timeout,`+
						` projections.actions2.allowed_to_fail`+
						` FROM projections.actions2`),
					nil,
					nil,
				),
				err: func(err error) (error, bool) {
					if !errs.IsNotFound(err) {
						return fmt.Errorf("err should be zitadel.NotFoundError got: %w", err), false
					}
					return nil, true
				},
			},
			object: (*Action)(nil),
		},
		{
			name:    "prepareActionQuery found",
			prepare: prepareActionQuery,
			want: want{
				sqlExpectations: mockQuery(
					regexp.QuoteMeta(`SELECT projections.actions2.id,`+
						` projections.actions2.creation_date,`+
						` projections.actions2.change_date,`+
						` projections.actions2.resource_owner,`+
						` projections.actions2.sequence,`+
						` projections.actions2.action_state,`+
						` projections.actions2.name,`+
						` projections.actions2.script,`+
						` projections.actions2.timeout,`+
						` projections.actions2.allowed_to_fail`+
						` FROM projections.actions2`),
					[]string{
						"id",
						"creation_date",
						"change_date",
						"resource_owner",
						"sequence",
						"action_state",
						"name",
						"script",
						"timeout",
						"allowed_to_fail",
					},
					[]driver.Value{
						"id",
						testNow,
						testNow,
						"ro",
						uint64(20211109),
						domain.ActionStateActive,
						"action-name",
						"script",
						1 * time.Second,
						true,
					},
				),
			},
			object: &Action{
				ID:            "id",
				CreationDate:  testNow,
				ChangeDate:    testNow,
				ResourceOwner: "ro",
				State:         domain.ActionStateActive,
				Sequence:      20211109,
				Name:          "action-name",
				Script:        "script",
				Timeout:       1 * time.Second,
				AllowedToFail: true,
			},
		},
		{
			name:    "prepareActionQuery sql err",
			prepare: prepareActionQuery,
			want: want{
				sqlExpectations: mockQueryErr(
					regexp.QuoteMeta(`SELECT projections.actions2.id,`+
						` projections.actions2.creation_date,`+
						` projections.actions2.change_date,`+
						` projections.actions2.resource_owner,`+
						` projections.actions2.sequence,`+
						` projections.actions2.action_state,`+
						` projections.actions2.name,`+
						` projections.actions2.script,`+
						` projections.actions2.timeout,`+
						` projections.actions2.allowed_to_fail`+
						` FROM projections.actions2`),
					sql.ErrConnDone,
				),
				err: func(err error) (error, bool) {
					if !errors.Is(err, sql.ErrConnDone) {
						return fmt.Errorf("err should be sql.ErrConnDone got: %w", err), false
					}
					return nil, true
				},
			},
			object: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertPrepare(t, tt.prepare, tt.object, tt.want.sqlExpectations, tt.want.err)
		})
	}
}
