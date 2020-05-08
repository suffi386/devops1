package model

type SearchMethod int32

const (
	SEARCHMETHOD_EQUALS SearchMethod = iota
	SEARCHMETHOD_STARTS_WITH
	SEARCHMETHOD_CONTAINS
	SEARCHMETHOD_EQUALS_IGNORE_CASE
	SEARCHMETHOD_STARTS_WITH_IGNORE_CASE
	SEARCHMETHOD_CONTAINS_IGNORE_CASE
	SEARCHMETHOD_NOT_EQUALS
)
