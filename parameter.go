package swordtech

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/markoxley/daggertech"
	"github.com/markoxley/daggertech/clause"
)

// Parameter is the table to hold the global parameters
type Parameter struct {
	daggertech.Model
	Name             string `daggertech:"size:64,key:true"`
	Value            string `daggertech:""`
	Description      string `daggertech:""`
	ParameterGroupID string `daggertech:"type:uuid,key:true"`
}

// Int returns the int64 value of the parameter
func (p *Parameter) Int() int64 {
	if result, err := strconv.ParseInt(p.Value, 10, 64); err == nil {
		return result
	}
	return 0
}

// Float returns the float64 value of the parameter
func (p *Parameter) Float() float64 {
	if result, err := strconv.ParseFloat(p.Value, 64); err == nil {
		return result
	}
	return 0
}

// String returns the string value of the parameter
func (p *Parameter) String() string {
	return p.Value
}

// Bool returns the bool value of the paraneter
func (p *Parameter) Bool() bool {
	return strings.ToLower(p.Value) == "true"
}

// Time returns the time value of the parameter, or nil if the value cannot be converted
func (p *Parameter) Time() *time.Time {
	return stringToTime(p.Value)
}

// Set sets the value of the paraneter and automatically saves it to the database
func (p *Parameter) Set(v interface{}) {
	if boolValue, ok := v.(bool); ok {
		if boolValue {
			p.Value = "true"
		} else {
			p.Value = "false"
		}
	} else if timeValue, ok := v.(*time.Time); ok {
		p.Value = timeToString(timeValue)
	} else {
		p.Value = fmt.Sprintf("%v", v)
	}
	daggertech.Save(p)
}

// GetParameter returns the parameter specified
// If necessary, a new parameter group is created
func GetParameter(group string, name string, description string, defaultValue string) *Parameter {
	var parameterGroup *ParameterGroup
	var parameter *Parameter

	if record, ok := daggertech.First(&ParameterGroup{}, &daggertech.Criteria{
		Where: clause.Equal("Name", group).ToString(),
	}); ok {
		if parameterGroup, ok = record.(*ParameterGroup); !ok {
			parameterGroup = nil
		}
	}

	if parameterGroup == nil {
		parameterGroup = &ParameterGroup{Name: group}
		daggertech.Save(parameterGroup)
	}

	if record, ok := daggertech.First(&Parameter{}, &daggertech.Criteria{
		Where: clause.Equal("ParameterGroupID", *(parameterGroup.ID)).AndEqual("Name", name).ToString(),
	}); ok {
		if parameter, ok = record.(*Parameter); !ok {
			parameter = nil
		}
	}

	if parameter == nil {
		parameter = &Parameter{
			Name:             name,
			Description:      description,
			ParameterGroupID: *(parameterGroup.ID),
			Value:            defaultValue,
		}
		daggertech.Save(parameter)
	}
	return parameter
}

// StandingData returns the standing data for when the table is created
func (p Parameter) StandingData() []daggertech.Modeller {
	var group *ParameterGroup
	if g, ok := daggertech.First(&ParameterGroup{}, &daggertech.Criteria{
		Where: clause.Equal("Name", "Security").ToString(),
	}); ok {
		if group, ok = g.(*ParameterGroup); !ok {
			return nil
		}
	} else {
		return nil
	}
	return []daggertech.Modeller{
		&Parameter{
			Name:             "IPThreshhold",
			Value:            "60",
			Description:      "Number of minutes to check IP attempts",
			ParameterGroupID: *group.ID,
		},
		&Parameter{
			Name:             "IPAttempts",
			Value:            "6",
			Description:      "Number of attempts allowed from a unique IP address",
			ParameterGroupID: *group.ID,
		},
		&Parameter{
			Name:             "IPLockDuration",
			Value:            "60",
			Description:      "Minutes to lock IP address",
			ParameterGroupID: *group.ID,
		},
		&Parameter{
			Name:             "UserThreshhold",
			Value:            "30",
			Description:      "Number of minutes to check user attempts",
			ParameterGroupID: *group.ID,
		},
		&Parameter{
			Name:             "UserAttempts",
			Value:            "3",
			Description:      "Number of attempts allowed for a username",
			ParameterGroupID: *group.ID,
		},
		&Parameter{
			Name:             "UserLockDuration",
			Value:            "30",
			Description:      "Minutes to lock username",
			ParameterGroupID: *group.ID,
		},
	}
}
