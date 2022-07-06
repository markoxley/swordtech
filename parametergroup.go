package swordtech

import (
	"github.com/markoxley/daggertech"
	"github.com/markoxley/daggertech/clause"
)

// ParameterGroup is the table to hold the global parameter groupss
type ParameterGroup struct {
	daggertech.Model
	Name string `daggertech:"size:64,key:true"`
}

// GetParameters returns all the parameters for the parameter group passed as receiver
func (g *ParameterGroup) GetParameters() []*Parameter {
	result := make([]*Parameter, 0)
	if g.IsNew() {
		if m, ok := daggertech.Fetch(&Parameter{}, &daggertech.Criteria{
			Where: clause.Equal("ParameterGroupID", *(g.ID)).ToString(),
		}); ok {
			for _, model := range m {
				if p, ok := model.(*Parameter); ok {
					result = append(result, p)
				}
			}
		}
	}
	return result
}

// StandingData returns the standing data for when the table is created
func (g ParameterGroup) StandingData() []daggertech.Modeller {

	pg := []daggertech.Modeller{
		&ParameterGroup{Name: "Security"},
		&ParameterGroup{Name: "Batch"},
	}
	return pg
}
