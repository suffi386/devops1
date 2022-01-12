package projection

import (
	"testing"

	"github.com/lib/pq"

	"github.com/caos/zitadel/internal/domain"
	"github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/eventstore"
	"github.com/caos/zitadel/internal/eventstore/handler"
	"github.com/caos/zitadel/internal/eventstore/repository"
	"github.com/caos/zitadel/internal/repository/usergrant"
)

func TestUserGrantProjection_reduces(t *testing.T) {
	type args struct {
		event func(t *testing.T) eventstore.Event
	}
	tests := []struct {
		name   string
		args   args
		reduce func(event eventstore.Event) (*handler.Statement, error)
		want   wantReduce
	}{
		{
			name: "reduceAdded",
			args: args{
				event: getEvent(testEvent(
					repository.EventType(usergrant.UserGrantAddedType),
					usergrant.AggregateType,
					[]byte(`{
						"userId": "user-id",
						"projectId": "project-id",
						"roleKeys": ["role"]
					}`),
				), usergrant.UserGrantAddedEventMapper),
			},
			reduce: (&UserGrantProjection{}).reduceAdded,
			want: wantReduce{
				aggregateType:    usergrant.AggregateType,
				sequence:         15,
				previousSequence: 10,
				projection:       UserGrantProjectionTable,
				executer: &testExecuter{
					executions: []execution{
						{
							expectedStmt: "INSERT INTO zitadel.projections.user_grants (id, resource_owner, creation_date, change_date, sequence, user_id, project_id, grant_id, roles, state) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
							expectedArgs: []interface{}{
								"agg-id",
								"ro-id",
								anyArg{},
								anyArg{},
								uint64(15),
								"user-id",
								"project-id",
								"",
								pq.StringArray{"role"},
								domain.UserGrantStateActive,
							},
						},
					},
				},
			},
		},
		{
			name: "reduceChanged",
			args: args{
				event: getEvent(testEvent(
					repository.EventType(usergrant.UserGrantChangedType),
					usergrant.AggregateType,
					[]byte(`{
						"roleKeys": ["role"]
					}`),
				), usergrant.UserGrantChangedEventMapper),
			},
			reduce: (&UserGrantProjection{}).reduceChanged,
			want: wantReduce{
				aggregateType:    usergrant.AggregateType,
				sequence:         15,
				previousSequence: 10,
				projection:       UserGrantProjectionTable,
				executer: &testExecuter{
					executions: []execution{
						{
							expectedStmt: "UPDATE zitadel.projections.user_grants SET (change_date, roles, sequence) = ($1, $2, $3) WHERE (id = $4)",
							expectedArgs: []interface{}{
								anyArg{},
								pq.StringArray{"role"},
								uint64(15),
								"agg-id",
							},
						},
					},
				},
			},
		},
		{
			name: "reduceCascadeChanged",
			args: args{
				event: getEvent(testEvent(
					repository.EventType(usergrant.UserGrantCascadeChangedType),
					usergrant.AggregateType,
					[]byte(`{
						"roleKeys": ["role"]
					}`),
				), usergrant.UserGrantCascadeChangedEventMapper),
			},
			reduce: (&UserGrantProjection{}).reduceChanged,
			want: wantReduce{
				aggregateType:    usergrant.AggregateType,
				sequence:         15,
				previousSequence: 10,
				projection:       UserGrantProjectionTable,
				executer: &testExecuter{
					executions: []execution{
						{
							expectedStmt: "UPDATE zitadel.projections.user_grants SET (change_date, roles, sequence) = ($1, $2, $3) WHERE (id = $4)",
							expectedArgs: []interface{}{
								anyArg{},
								pq.StringArray{"role"},
								uint64(15),
								"agg-id",
							},
						},
					},
				},
			},
		},
		{
			name: "reduceRemoved",
			args: args{
				event: getEvent(testEvent(
					repository.EventType(usergrant.UserGrantRemovedType),
					usergrant.AggregateType,
					nil,
				), usergrant.UserGrantRemovedEventMapper),
			},
			reduce: (&UserGrantProjection{}).reduceRemoved,
			want: wantReduce{
				aggregateType:    usergrant.AggregateType,
				sequence:         15,
				previousSequence: 10,
				projection:       UserGrantProjectionTable,
				executer: &testExecuter{
					executions: []execution{
						{
							expectedStmt: "DELETE FROM zitadel.projections.user_grants WHERE (id = $1)",
							expectedArgs: []interface{}{
								anyArg{},
							},
						},
					},
				},
			},
		},
		{
			name: "reduceCascadeRemoved",
			args: args{
				event: getEvent(testEvent(
					repository.EventType(usergrant.UserGrantCascadeRemovedType),
					usergrant.AggregateType,
					nil,
				), usergrant.UserGrantCascadeRemovedEventMapper),
			},
			reduce: (&UserGrantProjection{}).reduceRemoved,
			want: wantReduce{
				aggregateType:    usergrant.AggregateType,
				sequence:         15,
				previousSequence: 10,
				projection:       UserGrantProjectionTable,
				executer: &testExecuter{
					executions: []execution{
						{
							expectedStmt: "DELETE FROM zitadel.projections.user_grants WHERE (id = $1)",
							expectedArgs: []interface{}{
								anyArg{},
							},
						},
					},
				},
			},
		},
		{
			name: "reduceDeactivated",
			args: args{
				event: getEvent(testEvent(
					repository.EventType(usergrant.UserGrantDeactivatedType),
					usergrant.AggregateType,
					nil,
				), usergrant.UserGrantDeactivatedEventMapper),
			},
			reduce: (&UserGrantProjection{}).reduceDeactivated,
			want: wantReduce{
				aggregateType:    usergrant.AggregateType,
				sequence:         15,
				previousSequence: 10,
				projection:       UserGrantProjectionTable,
				executer: &testExecuter{
					executions: []execution{
						{
							expectedStmt: "UPDATE zitadel.projections.user_grants SET (change_date, state, sequence) = ($1, $2, $3) WHERE (id = $4)",
							expectedArgs: []interface{}{
								anyArg{},
								domain.UserGrantStateInactive,
								uint64(15),
								"agg-id",
							},
						},
					},
				},
			},
		},
		{
			name: "reduceReactivated",
			args: args{
				event: getEvent(testEvent(
					repository.EventType(usergrant.UserGrantReactivatedType),
					usergrant.AggregateType,
					nil,
				), usergrant.UserGrantDeactivatedEventMapper),
			},
			reduce: (&UserGrantProjection{}).reduceReactivated,
			want: wantReduce{
				aggregateType:    usergrant.AggregateType,
				sequence:         15,
				previousSequence: 10,
				projection:       UserGrantProjectionTable,
				executer: &testExecuter{
					executions: []execution{
						{
							expectedStmt: "UPDATE zitadel.projections.user_grants SET (change_date, state, sequence) = ($1, $2, $3) WHERE (id = $4)",
							expectedArgs: []interface{}{
								anyArg{},
								domain.UserGrantStateActive,
								uint64(15),
								"agg-id",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			event := baseEvent(t)
			got, err := tt.reduce(event)
			if _, ok := err.(errors.InvalidArgument); !ok {
				t.Errorf("no wrong event mapping: %v, got: %v", err, got)
			}

			event = tt.args.event(t)
			got, err = tt.reduce(event)
			assertReduce(t, got, err, tt.want)
		})
	}
}
