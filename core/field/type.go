package titian

import ()

type Type interface {
	Name() string
	NewValue(string) (Value, error)
}

// Text represents a multi-line text field.
type Text struct{}

func (_ Text) Name() string {
	return "Text"
}

func (_ Text) NewValue(value string) (Value, error) {
	text := &TextValue{}

	if err := text.Set(value); err != nil {
		return nil, err
	}

	return text, nil
}

/*
// NumberField
type NumberField struct {
	id ftid
	unsigned bool
	integer  bool
}

func (n NumberField) Name() string {
	return "Number"
}

func (n NumberField) NewValue(value string) (string, error) {
	number, err := strconv.ParseFloat(value, 64)

	if err != nil {
		return "", err
	}

	return fmt.Sprint(number), nil
}

// FlagField
type FlagField struct{}

func (_ FlagField) Name() string {
	return "Flag"
}

func (_ FlagField) NewValue(value string) (string, error) {
	b, err := strconv.ParseBool(value)

	if err == nil {
		if b {
			return "True", nil
		}

		return "False", nil
	}

	return "", err
}

// DateField
type DateField struct {
	id
	dateonly bool
}

func (d DateField) Name() string {
	return "Date"
}

func (d DateField) NewValue() FieldValue {
	return &DateFieldValue{}
}

// ChecklistField
type ChecklistField struct {
	id
	items []string
}

func (c ChecklistField) Name() string {
	return "Checklist"
}

func (c ChecklistField) NewValue(value string) (string, error) {
	//stub
	return value, nil
}

// SelectionField
type SelectionField struct {
	id
	items []string
}

func (s SelectionField) Name() string {
	return "Selection"
}

func (s SelectionField) NewValue(value string) (string, error) {
	//stub
	return value, nil
}*/