package query

import (
	"context"
	"database/sql"
	errs "errors"
	"time"

	sq "github.com/Masterminds/squirrel"
	"golang.org/x/text/language"

	"github.com/caos/zitadel/internal/domain"

	"github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/eventstore"
	"github.com/caos/zitadel/internal/query/projection"
)

type Users struct {
	SearchResponse
	Users []*User
}

type User struct {
	ID                 string
	CreationDate       time.Time
	ChangeDate         time.Time
	ResourceOwner      string
	Sequence           uint64
	State              domain.UserState
	Username           string
	LoginNames         []string
	PreferredLoginName string
	Human              *Human
	Machine            *Machine
}

type Human struct {
	FirstName         string
	LastName          string
	NickName          string
	DisplayName       string
	AvatarKey         string
	PreferredLanguage language.Tag
	Gender            domain.Gender
	Email             string
	IsEmailVerified   bool
	Phone             string
	IsPhoneVerified   bool
}

type Profile struct {
	ID                string
	CreationDate      time.Time
	ChangeDate        time.Time
	ResourceOwner     string
	Sequence          uint64
	FirstName         string
	LastName          string
	NickName          string
	DisplayName       string
	AvatarKey         string
	PreferredLanguage language.Tag
	Gender            domain.Gender
}

type Email struct {
	ID            string
	CreationDate  time.Time
	ChangeDate    time.Time
	ResourceOwner string
	Sequence      uint64
	Email         string
	IsVerified    bool
}

type Phone struct {
	ID            string
	CreationDate  time.Time
	ChangeDate    time.Time
	ResourceOwner string
	Sequence      uint64
	Phone         string
	IsVerified    bool
}

type Machine struct {
	Name        string
	Description string
}

type UserSearchQueries struct {
	SearchRequest
	Queries []SearchQuery
}

var (
	userTable = table{
		name: projection.UserTable,
	}
	UserIDCol = Column{
		name:  projection.UserIDCol,
		table: userTable,
	}
	UserCreationDateCol = Column{
		name:  projection.UserCreationDateCol,
		table: userTable,
	}
	UserChangeDateCol = Column{
		name:  projection.UserChangeDateCol,
		table: userTable,
	}
	UserResourceOwnerCol = Column{
		name:  projection.UserResourceOwnerCol,
		table: userTable,
	}
	UserStateCol = Column{
		name:  projection.UserStateCol,
		table: userTable,
	}
	UserSequenceCol = Column{
		name:  projection.UserSequenceCol,
		table: userTable,
	}
	UserUsernameCol = Column{
		name:  projection.UserUsernameCol,
		table: userTable,
	}
	UserTypeCol = Column{
		name:  projection.UserTypeCol,
		table: userTable,
	}
	userLoginNamesTable                = loginNameTable.setAlias("login_names")
	userLoginNamesUserIDCol            = LoginNameUserIDCol.setTable(userLoginNamesTable)
	userLoginNamesCol                  = LoginNameNameCol.setTable(userLoginNamesTable)
	userPreferredLoginNameTable        = loginNameTable.setAlias("preferred_login_name")
	userPreferredLoginNameUserIDCol    = LoginNameUserIDCol.setTable(userPreferredLoginNameTable)
	userPreferredLoginNameCol          = LoginNameNameCol.setTable(userPreferredLoginNameTable)
	userPreferredLoginNameIsPrimaryCol = LoginNameIsPrimaryCol.setTable(userPreferredLoginNameTable)
)

var (
	humanTable = table{
		name: projection.UserHumanTable,
	}
	// profile
	HumanUserIDCol = Column{
		name:  projection.HumanUserIDCol,
		table: humanTable,
	}
	HumanFirstNameCol = Column{
		name:  projection.HumanFirstNameCol,
		table: humanTable,
	}
	HumanLastNameCol = Column{
		name:  projection.HumanLastNameCol,
		table: humanTable,
	}
	HumanNickNameCol = Column{
		name:  projection.HumanNickNameCol,
		table: humanTable,
	}
	HumanDisplayNameCol = Column{
		name:  projection.HumanDisplayNameCol,
		table: humanTable,
	}
	HumanPreferredLanguageCol = Column{
		name:  projection.HumanPreferredLanguageCol,
		table: humanTable,
	}
	HumanGenderCol = Column{
		name:  projection.HumanGenderCol,
		table: humanTable,
	}
	HumanAvaterURLCol = Column{
		name:  projection.HumanAvaterURLCol,
		table: humanTable,
	}

	// email
	HumanEmailCol = Column{
		name:  projection.HumanEmailCol,
		table: humanTable,
	}
	HumanIsEmailVerifiedCol = Column{
		name:  projection.HumanIsEmailVerifiedCol,
		table: humanTable,
	}

	// phone
	HumanPhoneCol = Column{
		name:  projection.HumanPhoneCol,
		table: humanTable,
	}
	HumanIsPhoneVerifiedCol = Column{
		name:  projection.HumanIsPhoneVerifiedCol,
		table: humanTable,
	}
)

var (
	machineTable = table{
		name: projection.UserMachineTable,
	}
	MachineUserIDCol = Column{
		name:  projection.MachineUserIDCol,
		table: machineTable,
	}
	MachineNameCol = Column{
		name:  projection.MachineNameCol,
		table: machineTable,
	}
	MachineDescriptionCol = Column{
		name:  projection.MachineDescriptionCol,
		table: machineTable,
	}
)

func (q *Queries) GetUserByID(ctx context.Context, userID string, queries ...SearchQuery) (*User, error) {
	query, scan := prepareUserQuery()
	for _, q := range queries {
		query = q.toQuery(query)
	}
	stmt, args, err := query.Where(
		sq.Eq{
			UserIDCol.identifier(): userID,
		}).ToSql()
	if err != nil {
		return nil, errors.ThrowInternal(err, "QUERY-FBg21", "Errors.Query.SQLStatment")
	}

	rows, err := q.client.QueryContext(ctx, stmt, args...)
	if err != nil {
		return nil, err
	}
	return scan(rows)
}

func (q *Queries) GetUserByLoginNameGlobal(ctx context.Context, loginName string, queries ...SearchQuery) (*User, error) {
	query, scan := prepareUserQuery()
	for _, q := range queries {
		query = q.toQuery(query)
	}
	stmt, args, err := query.Where(
		sq.Eq{
			userLoginNamesCol.identifier(): loginName,
		}).ToSql()
	if err != nil {
		return nil, errors.ThrowInternal(err, "QUERY-Dnhr2", "Errors.Query.SQLStatment")
	}

	rows, err := q.client.QueryContext(ctx, stmt, args...)
	if err != nil {
		return nil, err
	}
	return scan(rows)
}

func (q *Queries) GetHumanProfile(ctx context.Context, userID string, queries ...SearchQuery) (*Profile, error) {
	query, scan := prepareProfileQuery()
	for _, q := range queries {
		query = q.toQuery(query)
	}
	stmt, args, err := query.Where(
		sq.Eq{
			UserIDCol.identifier(): userID,
		}).ToSql()
	if err != nil {
		return nil, errors.ThrowInternal(err, "QUERY-Dgbg2", "Errors.Query.SQLStatment")
	}

	row := q.client.QueryRowContext(ctx, stmt, args...)
	return scan(row)
}

func (q *Queries) GetHumanEmail(ctx context.Context, userID string, queries ...SearchQuery) (*Email, error) {
	query, scan := prepareEmailQuery()
	for _, q := range queries {
		query = q.toQuery(query)
	}
	stmt, args, err := query.Where(
		sq.Eq{
			UserIDCol.identifier(): userID,
		}).ToSql()
	if err != nil {
		return nil, errors.ThrowInternal(err, "QUERY-BHhj3", "Errors.Query.SQLStatment")
	}

	row := q.client.QueryRowContext(ctx, stmt, args...)
	return scan(row)
}

func (q *Queries) GetHumanPhone(ctx context.Context, userID string, queries ...SearchQuery) (*Phone, error) {
	query, scan := preparePhoneQuery()
	for _, q := range queries {
		query = q.toQuery(query)
	}
	stmt, args, err := query.Where(
		sq.Eq{
			UserIDCol.identifier(): userID,
		}).ToSql()
	if err != nil {
		return nil, errors.ThrowInternal(err, "QUERY-Dg43g", "Errors.Query.SQLStatment")
	}

	row := q.client.QueryRowContext(ctx, stmt, args...)
	return scan(row)
}

func (q *Queries) SearchUsers(ctx context.Context, queries *UserSearchQueries) (*Users, error) {
	query, scan := prepareUsersQuery()
	stmt, args, err := queries.toQuery(query).ToSql()
	if err != nil {
		return nil, errors.ThrowInternal(err, "QUERY-Dgbg2", "Errors.Query.SQLStatment")
	}

	rows, err := q.client.QueryContext(ctx, stmt, args...)
	if err != nil {
		return nil, errors.ThrowInternal(err, "QUERY-AG4gs", "Errors.Internal")
	}
	users, err := scan(rows)
	if err != nil {
		return nil, err
	}
	users.LatestSequence, err = q.latestSequence(ctx, projectRolesTable)
	return users, err
}

func (q *Queries) IsUserUnique(ctx context.Context, username, email, resourceOwner string) (bool, error) {
	query, scan := prepareUserUniqueQuery()
	queries := make([]SearchQuery, 0, 3)
	if username != "" {
		usernameQuery, err := NewUserUsernameSearchQuery(username, TextEquals)
		if err != nil {
			return false, err
		}
		queries = append(queries, usernameQuery)
	}
	if email != "" {
		emailQuery, err := NewUserEmailSearchQuery(email, TextEquals)
		if err != nil {
			return false, err
		}
		queries = append(queries, emailQuery)
	}
	if resourceOwner != "" {
		resourceOwnerQuery, err := NewUserResourceOwnerSearchQuery(resourceOwner, TextEquals)
		if err != nil {
			return false, err
		}
		queries = append(queries, resourceOwnerQuery)
	}
	for _, q := range queries {
		query = q.toQuery(query)
	}
	stmt, args, err := query.ToSql()
	if err != nil {
		return false, errors.ThrowInternal(err, "QUERY-Dg43g", "Errors.Query.SQLStatment")
	}
	row := q.client.QueryRowContext(ctx, stmt, args...)
	return scan(row)
}

func (q *Queries) UserEvents(ctx context.Context, orgID, userID string, sequence uint64) ([]eventstore.Event, error) {
	query := NewUserEventSearchQuery(userID, orgID, sequence)
	return q.eventstore.Filter(ctx, query)
}

func (q *UserSearchQueries) toQuery(query sq.SelectBuilder) sq.SelectBuilder {
	query = q.SearchRequest.toQuery(query)
	for _, q := range q.Queries {
		query = q.toQuery(query)
	}
	return query
}

func (r *UserSearchQueries) AppendMyResourceOwnerQuery(orgID string) error {
	query, err := NewProjectResourceOwnerSearchQuery(orgID)
	if err != nil {
		return err
	}
	r.Queries = append(r.Queries, query)
	return nil
}

//func NewUserIDSearchQuery(value string) (SearchQuery, error) {
//	return NewTextQuery(UserIDCol, value, TextEquals)
//}

func NewUserResourceOwnerSearchQuery(value string, comparison TextComparison) (SearchQuery, error) {
	return NewTextQuery(UserResourceOwnerCol, value, comparison)
}

func NewUserUsernameSearchQuery(value string, comparison TextComparison) (SearchQuery, error) {
	return NewTextQuery(UserUsernameCol, value, comparison)
}

func NewUserFirstNameSearchQuery(value string, comparison TextComparison) (SearchQuery, error) {
	return NewTextQuery(HumanFirstNameCol, value, comparison)
}

func NewUserLastNameSearchQuery(value string, comparison TextComparison) (SearchQuery, error) {
	return NewTextQuery(HumanLastNameCol, value, comparison)
}

func NewUserNickNameSearchQuery(value string, comparison TextComparison) (SearchQuery, error) {
	return NewTextQuery(HumanNickNameCol, value, comparison)
}

func NewUserDisplayNameSearchQuery(value string, comparison TextComparison) (SearchQuery, error) {
	return NewTextQuery(HumanDisplayNameCol, value, comparison)
}

func NewUserEmailSearchQuery(value string, comparison TextComparison) (SearchQuery, error) {
	return NewTextQuery(HumanEmailCol, value, comparison)
}

func NewUserStateSearchQuery(value int32) (SearchQuery, error) {
	return NewNumberQuery(UserStateCol, value, NumberEquals)
}

func NewUserTypeSearchQuery(value int32) (SearchQuery, error) {
	return NewNumberQuery(UserStateCol, value, NumberEquals) //TODO: type
}

func NewUserPreferredLoginNameSearchQuery(value string, comparison TextComparison) (SearchQuery, error) {
	return NewTextQuery(userPreferredLoginNameCol, value, comparison)
}

func prepareUserQuery() (sq.SelectBuilder, func(*sql.Rows) (*User, error)) {
	loginNamesQuery, _, err := sq.Select(
		userLoginNamesUserIDCol.identifier(),
		"ARRAY_AGG("+userLoginNamesCol.identifier()+") as login_names").
		From(userLoginNamesTable.identifier()).
		GroupBy(userLoginNamesUserIDCol.identifier()).
		ToSql()
	if err != nil {
		return sq.SelectBuilder{}, nil
	}
	preferredLoginNameQuery, preferredLoginNameArgs, err := sq.Select(
		userPreferredLoginNameUserIDCol.identifier(),
		userPreferredLoginNameCol.identifier()).
		From(userPreferredLoginNameTable.identifier()).
		Where(
			sq.Eq{
				userPreferredLoginNameIsPrimaryCol.identifier(): true,
			}).ToSql()
	if err != nil {
		return sq.SelectBuilder{}, nil
	}
	return sq.Select(
			UserIDCol.identifier(),
			UserCreationDateCol.identifier(),
			UserChangeDateCol.identifier(),
			UserResourceOwnerCol.identifier(),
			UserSequenceCol.identifier(),
			UserStateCol.identifier(),
			UserUsernameCol.identifier(),
			"login_names.login_names",
			userPreferredLoginNameCol.identifier(),
			HumanUserIDCol.identifier(),
			HumanFirstNameCol.identifier(),
			HumanLastNameCol.identifier(),
			HumanNickNameCol.identifier(),
			HumanDisplayNameCol.identifier(),
			HumanPreferredLanguageCol.identifier(),
			HumanGenderCol.identifier(),
			HumanAvaterURLCol.identifier(),
			HumanEmailCol.identifier(),
			HumanIsEmailVerifiedCol.identifier(),
			HumanPhoneCol.identifier(),
			HumanIsPhoneVerifiedCol.identifier(),
			MachineUserIDCol.identifier(),
			MachineNameCol.identifier(),
			MachineDescriptionCol.identifier(),
		).
			From(userTable.identifier()).
			LeftJoin(join(HumanUserIDCol, UserIDCol)).
			LeftJoin(join(MachineUserIDCol, UserIDCol)).
			LeftJoin("("+loginNamesQuery+") as login_names on "+userLoginNamesUserIDCol.identifier()+" = "+UserIDCol.identifier()).
			LeftJoin("("+preferredLoginNameQuery+") as preferred_login_name on "+userPreferredLoginNameUserIDCol.identifier()+" = "+UserIDCol.identifier(), preferredLoginNameArgs...).
			PlaceholderFormat(sq.Dollar),
		func(rows *sql.Rows) (*User, error) {
			u := new(User)
			preferredLoginName := sql.NullString{}

			humanID := sql.NullString{}
			firstName := sql.NullString{}
			lastName := sql.NullString{}
			nickName := sql.NullString{}
			displayName := sql.NullString{}
			preferredLanguage := sql.NullString{}
			gender := sql.NullInt32{}
			avatarKey := sql.NullString{}
			email := sql.NullString{}
			isEmailVerified := sql.NullBool{}
			phone := sql.NullString{}
			isPhoneVerified := sql.NullBool{}

			machineID := sql.NullString{}
			name := sql.NullString{}
			description := sql.NullString{}

			for rows.Next() {
				loginName := &sql.NullString{}
				isPrimary := sql.NullBool{}

				err := rows.Scan(
					&u.ID,
					&u.CreationDate,
					&u.ChangeDate,
					&u.ResourceOwner,
					&u.Sequence,
					&u.State,
					&u.Username,
					&u.LoginNames,
					&preferredLoginName,
					&humanID,
					&firstName,
					&lastName,
					&nickName,
					&displayName,
					&preferredLanguage,
					&gender,
					&avatarKey,
					&email,
					&isEmailVerified,
					&phone,
					&isPhoneVerified,
					&machineID,
					&name,
					&description,
				)
				if err != nil {
					return nil, err
				}
				if preferredLoginName.Valid {
					u.PreferredLoginName = preferredLoginName.String
				}
				if humanID.Valid {
					u.Human = &Human{
						FirstName:         firstName.String,
						LastName:          lastName.String,
						NickName:          nickName.String,
						DisplayName:       displayName.String,
						AvatarKey:         avatarKey.String,
						PreferredLanguage: language.Make(preferredLanguage.String),
						Gender:            domain.Gender(gender.Int32),
						Email:             email.String,
						IsEmailVerified:   isEmailVerified.Bool,
						Phone:             phone.String,
						IsPhoneVerified:   isPhoneVerified.Bool,
					}
				} else if machineID.Valid {
					u.Machine = &Machine{
						Name:        name.String,
						Description: description.String,
					}
				}

				if loginName.Valid {
					u.LoginNames = append(u.LoginNames, loginName.String)
				}
				if isPrimary.Valid && isPrimary.Bool {
					u.PreferredLoginName = loginName.String
				}
			}

			if err := rows.Close(); err != nil {
				return nil, errors.ThrowInternal(err, "QUERY-Dgfe2", "Errors.Query.CloseRows")
			}

			return u, nil
		}
}

func prepareProfileQuery() (sq.SelectBuilder, func(*sql.Row) (*Profile, error)) {
	return sq.Select(
			UserIDCol.identifier(),
			UserCreationDateCol.identifier(),
			UserChangeDateCol.identifier(),
			UserResourceOwnerCol.identifier(),
			UserSequenceCol.identifier(),
			HumanUserIDCol.identifier(),
			HumanFirstNameCol.identifier(),
			HumanLastNameCol.identifier(),
			HumanNickNameCol.identifier(),
			HumanDisplayNameCol.identifier(),
			HumanPreferredLanguageCol.identifier(),
			HumanGenderCol.identifier(),
			HumanAvaterURLCol.identifier()).
			From(projectRolesTable.identifier()).
			LeftJoin(join(HumanUserIDCol, UserIDCol)).
			PlaceholderFormat(sq.Dollar),
		func(row *sql.Row) (*Profile, error) {
			p := new(Profile)

			humanID := sql.NullString{}
			firstName := sql.NullString{}
			lastName := sql.NullString{}
			nickName := sql.NullString{}
			displayName := sql.NullString{}
			preferredLanguage := sql.NullString{}
			gender := sql.NullInt32{}
			avatarKey := sql.NullString{}
			err := row.Scan(
				&p.ID,
				&p.CreationDate,
				&p.ChangeDate,
				&p.ResourceOwner,
				&p.Sequence,
				&humanID,
				&firstName,
				&lastName,
				&nickName,
				&displayName,
				&preferredLanguage,
				&gender,
				&avatarKey,
			)
			if err != nil {
				if errs.Is(err, sql.ErrNoRows) {
					return nil, errors.ThrowNotFound(err, "QUERY-HNhb3", "Errors.User.NotFound")
				}
				return nil, errors.ThrowInternal(err, "QUERY-Rfheq", "Errors.Internal")
			}
			if !humanID.Valid {
				return nil, errors.ThrowPreconditionFailed(nil, "QUERY-WLTce", "Errors.User.NotHuman")
			}

			p.FirstName = firstName.String
			p.LastName = lastName.String
			p.NickName = nickName.String
			p.DisplayName = displayName.String
			p.AvatarKey = avatarKey.String
			p.PreferredLanguage = language.Make(preferredLanguage.String)
			p.Gender = domain.Gender(gender.Int32)

			return p, nil
		}
}

func prepareEmailQuery() (sq.SelectBuilder, func(*sql.Row) (*Email, error)) {
	return sq.Select(
			UserIDCol.identifier(),
			UserCreationDateCol.identifier(),
			UserChangeDateCol.identifier(),
			UserResourceOwnerCol.identifier(),
			UserSequenceCol.identifier(),
			HumanUserIDCol.identifier(),
			HumanEmailCol.identifier(),
			HumanIsEmailVerifiedCol.identifier()).
			From(projectRolesTable.identifier()).
			LeftJoin(join(HumanUserIDCol, UserIDCol)).
			PlaceholderFormat(sq.Dollar),
		func(row *sql.Row) (*Email, error) {
			e := new(Email)

			humanID := sql.NullString{}
			email := sql.NullString{}
			isEmailVerified := sql.NullBool{}

			err := row.Scan(
				&e.ID,
				&e.CreationDate,
				&e.ChangeDate,
				&e.ResourceOwner,
				&e.Sequence,
				&humanID,
				&email,
				&isEmailVerified,
			)
			if err != nil {
				if errs.Is(err, sql.ErrNoRows) {
					return nil, errors.ThrowNotFound(err, "QUERY-Hms2s", "Errors.User.NotFound")
				}
				return nil, errors.ThrowInternal(err, "QUERY-Nu42d", "Errors.Internal")
			}
			if !humanID.Valid {
				return nil, errors.ThrowPreconditionFailed(nil, "QUERY-pt7HY", "Errors.User.NotHuman")
			}

			e.Email = email.String
			e.IsVerified = isEmailVerified.Bool

			return e, nil
		}
}

func preparePhoneQuery() (sq.SelectBuilder, func(*sql.Row) (*Phone, error)) {
	return sq.Select(
			UserIDCol.identifier(),
			UserCreationDateCol.identifier(),
			UserChangeDateCol.identifier(),
			UserResourceOwnerCol.identifier(),
			UserSequenceCol.identifier(),
			HumanUserIDCol.identifier(),
			HumanPhoneCol.identifier(),
			HumanIsPhoneVerifiedCol.identifier()).
			From(projectRolesTable.identifier()).
			LeftJoin(join(HumanUserIDCol, UserIDCol)).
			PlaceholderFormat(sq.Dollar),
		func(row *sql.Row) (*Phone, error) {
			e := new(Phone)

			humanID := sql.NullString{}
			email := sql.NullString{}
			isPhoneVerified := sql.NullBool{}

			err := row.Scan(
				&e.ID,
				&e.CreationDate,
				&e.ChangeDate,
				&e.ResourceOwner,
				&e.Sequence,
				&humanID,
				&email,
				&isPhoneVerified,
			)
			if err != nil {
				if errs.Is(err, sql.ErrNoRows) {
					return nil, errors.ThrowNotFound(err, "QUERY-DAvb3", "Errors.User.NotFound")
				}
				return nil, errors.ThrowInternal(err, "QUERY-Bmf2h", "Errors.Internal")
			}
			if !humanID.Valid {
				return nil, errors.ThrowPreconditionFailed(nil, "QUERY-hliQl", "Errors.User.NotHuman")
			}

			e.Phone = email.String
			e.IsVerified = isPhoneVerified.Bool

			return e, nil
		}
}

func prepareUserUniqueQuery() (sq.SelectBuilder, func(*sql.Row) (bool, error)) {
	return sq.Select(
			UserIDCol.identifier(),
			UserStateCol.identifier(),
			UserUsernameCol.identifier(),
			HumanUserIDCol.identifier(),
			HumanEmailCol.identifier(),
			HumanIsEmailVerifiedCol.identifier()).
			From(projectRolesTable.identifier()).
			LeftJoin(join(HumanUserIDCol, UserIDCol)).
			PlaceholderFormat(sq.Dollar),
		func(row *sql.Row) (bool, error) {
			userID := sql.NullString{}
			state := sql.NullInt32{}
			username := sql.NullString{}
			humanID := sql.NullString{}
			email := sql.NullString{}
			isEmailVerified := sql.NullBool{}

			err := row.Scan(
				&userID,
				&state,
				&username,
				&humanID,
				&email,
				&isEmailVerified,
			)
			if err != nil {
				if errs.Is(err, sql.ErrNoRows) {
					return false, errors.ThrowNotFound(err, "QUERY-Rbnaq", "Errors.User.NotFound")
				}
				return false, errors.ThrowInternal(err, "QUERY-Cxces", "Errors.Internal")
			}
			return userID.Valid, nil
		}
}

func prepareUsersQuery() (sq.SelectBuilder, func(*sql.Rows) (*Users, error)) {
	loginNamesQuery, _, err := sq.Select(
		userLoginNamesUserIDCol.identifier(),
		"ARRAY_AGG("+userLoginNamesCol.identifier()+") as login_names").
		From(userLoginNamesTable.identifier()).
		GroupBy(userLoginNamesUserIDCol.identifier()).
		ToSql()
	if err != nil {
		return sq.SelectBuilder{}, nil
	}
	preferredLoginNameQuery, preferredLoginNameArgs, err := sq.Select(
		userPreferredLoginNameUserIDCol.identifier(),
		userPreferredLoginNameCol.identifier()).
		From(userPreferredLoginNameTable.identifier()).
		Where(
			sq.Eq{
				userPreferredLoginNameIsPrimaryCol.identifier(): true,
			}).ToSql()
	if err != nil {
		return sq.SelectBuilder{}, nil
	}
	return sq.Select(
			UserIDCol.identifier(),
			UserCreationDateCol.identifier(),
			UserChangeDateCol.identifier(),
			UserResourceOwnerCol.identifier(),
			UserSequenceCol.identifier(),
			UserStateCol.identifier(),
			UserUsernameCol.identifier(),
			"login_names.login_names",
			userPreferredLoginNameCol.identifier(),
			HumanUserIDCol.identifier(),
			HumanFirstNameCol.identifier(),
			HumanLastNameCol.identifier(),
			HumanNickNameCol.identifier(),
			HumanDisplayNameCol.identifier(),
			HumanPreferredLanguageCol.identifier(),
			HumanGenderCol.identifier(),
			HumanAvaterURLCol.identifier(),
			HumanEmailCol.identifier(),
			HumanIsEmailVerifiedCol.identifier(),
			HumanPhoneCol.identifier(),
			HumanIsPhoneVerifiedCol.identifier(),
			MachineUserIDCol.identifier(),
			MachineNameCol.identifier(),
			MachineDescriptionCol.identifier(),
			countColumn.identifier()).
			From(projectRolesTable.identifier()).
			LeftJoin(join(HumanUserIDCol, UserIDCol)).
			LeftJoin(join(MachineUserIDCol, UserIDCol)).
			LeftJoin("("+loginNamesQuery+") as login_names on "+userLoginNamesUserIDCol.identifier()+" = "+UserIDCol.identifier()).
			LeftJoin("("+preferredLoginNameQuery+") as preferred_login_name on "+userPreferredLoginNameUserIDCol.identifier()+" = "+UserIDCol.identifier(), preferredLoginNameArgs...).
			PlaceholderFormat(sq.Dollar),
		func(rows *sql.Rows) (*Users, error) {
			users := make([]*User, 0)
			var count uint64
			for rows.Next() {
				u := new(User)
				preferredLoginName := sql.NullString{}

				humanID := sql.NullString{}
				firstName := sql.NullString{}
				lastName := sql.NullString{}
				nickName := sql.NullString{}
				displayName := sql.NullString{}
				preferredLanguage := sql.NullString{}
				gender := sql.NullInt32{}
				avatarKey := sql.NullString{}
				email := sql.NullString{}
				isEmailVerified := sql.NullBool{}
				phone := sql.NullString{}
				isPhoneVerified := sql.NullBool{}

				machineID := sql.NullString{}
				name := sql.NullString{}
				description := sql.NullString{}

				err := rows.Scan(
					&u.ID,
					&u.CreationDate,
					&u.ChangeDate,
					&u.ResourceOwner,
					&u.Sequence,
					&u.State,
					&u.Username,
					&u.LoginNames,
					&preferredLoginName,
					&humanID,
					&firstName,
					&lastName,
					&nickName,
					&displayName,
					&preferredLanguage,
					&gender,
					&avatarKey,
					&email,
					&isEmailVerified,
					&phone,
					&isPhoneVerified,
					&machineID,
					&name,
					&description,
					&count,
				)
				if err != nil {
					return nil, err
				}

				if preferredLoginName.Valid {
					u.PreferredLoginName = preferredLoginName.String
				}

				if humanID.Valid {
					u.Human = &Human{
						FirstName:         firstName.String,
						LastName:          lastName.String,
						NickName:          nickName.String,
						DisplayName:       displayName.String,
						AvatarKey:         avatarKey.String,
						PreferredLanguage: language.Make(preferredLanguage.String),
						Gender:            domain.Gender(gender.Int32),
						Email:             email.String,
						IsEmailVerified:   isEmailVerified.Bool,
						Phone:             phone.String,
						IsPhoneVerified:   isPhoneVerified.Bool,
					}
				} else if machineID.Valid {
					u.Machine = &Machine{
						Name:        name.String,
						Description: description.String,
					}
				}

				users = append(users, u)
			}

			if err := rows.Close(); err != nil {
				return nil, errors.ThrowInternal(err, "QUERY-frhbd", "Errors.Query.CloseRows")
			}

			return &Users{
				Users: users,
				SearchResponse: SearchResponse{
					Count: count,
				},
			}, nil
		}
}
