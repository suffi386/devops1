package domain

type User interface {
	GetUsername() string
	GetState() UserState
}

type UserState int32

const (
	UserStateUnspecified UserState = iota
	UserStateActive
	UserStateInactive
	UserStateDeleted
	UserStateLocked
	UserStateSuspend
	UserStateInitial

	userStateCount
)

func (f UserState) Valid() bool {
	return f >= 0 && f < userStateCount
}

func (s UserState) Exists() bool {
	return s != UserStateUnspecified && s != UserStateDeleted
}

func (s UserState) NotDisabled() bool {
	return s == UserStateActive || s == UserStateInitial
}

func (s UserState) Active() bool {
	return s.Valid() && s.Exists() && s.NotDisabled()
}

type UserType int32

const (
	UserTypeUnspecified UserType = iota
	UserTypeHuman
	UserTypeMachine
	userTypeCount
)

func (f UserType) Valid() bool {
	return f >= 0 && f < userTypeCount
}

type UserAuthMethodType int32

const (
	UserAuthMethodTypeUnspecified UserAuthMethodType = iota
	UserAuthMethodTypeOTP
	UserAuthMethodTypeU2F
	UserAuthMethodTypePasswordless
	userAuthMethodTypeCount
)

func (f UserAuthMethodType) Valid() bool {
	return f >= 0 && f < userAuthMethodTypeCount
}

type PersonalAccessTokenState int32

const (
	PersonalAccessTokenStateUnspecified PersonalAccessTokenState = iota
	PersonalAccessTokenStateActive
	PersonalAccessTokenStateRemoved

	personalAccessTokenStateCount
)

func (f PersonalAccessTokenState) Valid() bool {
	return f >= 0 && f < personalAccessTokenStateCount
}
