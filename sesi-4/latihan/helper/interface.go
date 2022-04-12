package helper

type EmpFunctions interface {
	GetName() string
	GetDivision() string
}

func NewEmployee(name, division string) EmpFunctions {
	return &Employee{
		Name:     name,
		Division: division,
	}
}

func (e *Employee) GetName() string {
	return e.Name
}

func (e *Employee) GetDivision() string {
	return e.Division
}
