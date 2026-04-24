package model

type Relation string

const (
	OrHigher Relation = ">="
	Equal    Relation = "=="
	OrUnder  Relation = "<="
	NoEqual  Relation = "!="
)

type Uwasa struct {
	Name     string
	Target   string
	Relation Relation
}
