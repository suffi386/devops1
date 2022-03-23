package projection

import (
	"context"

	"github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/eventstore"
	"github.com/caos/zitadel/internal/eventstore/handler"
	"github.com/caos/zitadel/internal/eventstore/handler/crdb"
	"github.com/caos/zitadel/internal/repository/instance"
)

const (
	IAMProjectionTable = "projections.iam"

	IAMColumnID              = "id"
	IAMColumnChangeDate      = "change_date"
	IAMColumnGlobalOrgID     = "global_org_id"
	IAMColumnProjectID       = "iam_project_id"
	IAMColumnSequence        = "sequence"
	IAMColumnSetUpStarted    = "setup_started"
	IAMColumnSetUpDone       = "setup_done"
	IAMColumnDefaultLanguage = "default_language"
)

type IAMProjection struct {
	crdb.StatementHandler
}

func NewIAMProjection(ctx context.Context, config crdb.StatementHandlerConfig) *IAMProjection {
	p := new(IAMProjection)
	config.ProjectionName = IAMProjectionTable
	config.Reducers = p.reducers()
	config.InitCheck = crdb.NewTableCheck(
		crdb.NewTable([]*crdb.Column{
			crdb.NewColumn(IAMColumnID, crdb.ColumnTypeText),
			crdb.NewColumn(IAMColumnChangeDate, crdb.ColumnTypeTimestamp),
			crdb.NewColumn(IAMColumnGlobalOrgID, crdb.ColumnTypeText, crdb.Default("")),
			crdb.NewColumn(IAMColumnProjectID, crdb.ColumnTypeText, crdb.Default("")),
			crdb.NewColumn(IAMColumnSequence, crdb.ColumnTypeInt64),
			crdb.NewColumn(IAMColumnSetUpStarted, crdb.ColumnTypeInt64, crdb.Default(0)),
			crdb.NewColumn(IAMColumnSetUpDone, crdb.ColumnTypeInt64, crdb.Default(0)),
			crdb.NewColumn(IAMColumnDefaultLanguage, crdb.ColumnTypeText, crdb.Default("")),
		},
			crdb.NewPrimaryKey(IAMColumnID),
		),
	)
	p.StatementHandler = crdb.NewStatementHandler(ctx, config)
	return p
}

func (p *IAMProjection) reducers() []handler.AggregateReducer {
	return []handler.AggregateReducer{
		{
			Aggregate: instance.AggregateType,
			EventRedusers: []handler.EventReducer{
				{
					Event:  instance.GlobalOrgSetEventType,
					Reduce: p.reduceGlobalOrgSet,
				},
				{
					Event:  instance.ProjectSetEventType,
					Reduce: p.reduceIAMProjectSet,
				},
				{
					Event:  instance.DefaultLanguageSetEventType,
					Reduce: p.reduceDefaultLanguageSet,
				},
				{
					Event:  instance.SetupStartedEventType,
					Reduce: p.reduceSetupEvent,
				},
				{
					Event:  instance.SetupDoneEventType,
					Reduce: p.reduceSetupEvent,
				},
			},
		},
	}
}

func (p *IAMProjection) reduceGlobalOrgSet(event eventstore.Event) (*handler.Statement, error) {
	e, ok := event.(*instance.GlobalOrgSetEvent)
	if !ok {
		return nil, errors.ThrowInvalidArgumentf(nil, "HANDL-2n9f2", "reduce.wrong.event.type %s", iam.GlobalOrgSetEventType)
	}
	return crdb.NewUpsertStatement(
		e,
		[]handler.Column{
			handler.NewCol(IAMColumnID, e.Aggregate().InstanceID),
			handler.NewCol(IAMColumnChangeDate, e.CreationDate()),
			handler.NewCol(IAMColumnSequence, e.Sequence()),
			handler.NewCol(IAMColumnGlobalOrgID, e.OrgID),
		},
	), nil
}

func (p *IAMProjection) reduceIAMProjectSet(event eventstore.Event) (*handler.Statement, error) {
	e, ok := event.(*instance.ProjectSetEvent)
	if !ok {
		return nil, errors.ThrowInvalidArgumentf(nil, "HANDL-30o0e", "reduce.wrong.event.type %s", iam.ProjectSetEventType)
	}
	return crdb.NewUpsertStatement(
		e,
		[]handler.Column{
			handler.NewCol(IAMColumnID, e.Aggregate().InstanceID),
			handler.NewCol(IAMColumnChangeDate, e.CreationDate()),
			handler.NewCol(IAMColumnSequence, e.Sequence()),
			handler.NewCol(IAMColumnProjectID, e.ProjectID),
		},
	), nil
}

func (p *IAMProjection) reduceDefaultLanguageSet(event eventstore.Event) (*handler.Statement, error) {
	e, ok := event.(*instance.DefaultLanguageSetEvent)
	if !ok {
		return nil, errors.ThrowInvalidArgumentf(nil, "HANDL-30o0e", "reduce.wrong.event.type %s", iam.DefaultLanguageSetEventType)
	}
	return crdb.NewUpsertStatement(
		e,
		[]handler.Column{
			handler.NewCol(IAMColumnID, e.Aggregate().InstanceID),
			handler.NewCol(IAMColumnChangeDate, e.CreationDate()),
			handler.NewCol(IAMColumnSequence, e.Sequence()),
			handler.NewCol(IAMColumnDefaultLanguage, e.Language.String()),
		},
	), nil
}

func (p *IAMProjection) reduceSetupEvent(event eventstore.Event) (*handler.Statement, error) {
	e, ok := event.(*instance.SetupStepEvent)
	if !ok {
		return nil, errors.ThrowInvalidArgumentf(nil, "HANDL-d9nfw", "reduce.wrong.event.type %v", []eventstore.EventType{iam.SetupDoneEventType, iam.SetupStartedEventType})
	}
	columns := []handler.Column{
		handler.NewCol(IAMColumnID, e.Aggregate().InstanceID),
		handler.NewCol(IAMColumnChangeDate, e.CreationDate()),
		handler.NewCol(IAMColumnSequence, e.Sequence()),
	}
	if e.EventType == instance.SetupStartedEventType {
		columns = append(columns, handler.NewCol(IAMColumnSetUpStarted, e.Step))
	} else {
		columns = append(columns, handler.NewCol(IAMColumnSetUpDone, e.Step))
	}
	return crdb.NewUpsertStatement(
		e,
		columns,
	), nil
}
