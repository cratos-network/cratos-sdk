package types

import (
	"fmt"
)

// // QueryResResolve Queries Result Payload for a resolve query
// type QueryResValue struct {
// 	Name      string `json:"name"`
// 	Value     string `json:"value"`
// 	Namespace string `json:"namespace"`
// }

// // implement fmt.Stringer
// func (res QueryResValue) String() string {
// 	return fmt.Sprintf("%s.%s %s", res.Namespace, res.Name, res.Value)
// }

// QueryResNames Queries Result Payload for a names query
type QueryResGetAllAttributes []DemosAttribute

// implement fmt.Stringer
func (list QueryResGetAllAttributes) String() string {
	var result string
	for i, attr := range list {
		result = result + fmt.Sprintf("%s - [$s].%s.%s=%s\n", i, attr.Namespace, attr.Name, attr.Value, attr.Owner)
	}

	return result
}

type QueryResGetAllRequests []DataAccessRequest

// implement fmt.Stringer
func (list QueryResGetAllRequests) String() string {
	var result string
	for i, req := range list {
		result = result + fmt.Sprintf("%s => %s (%s)\n", i, req.From, req.DataOwner, req.Scope)
	}

	return result
}
