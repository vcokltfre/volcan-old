package impl

var Interface *VolcanInterface

type VolcanInterface struct {
	Cases CaseManager
}

func SetupVolcan() error {
	Interface = &VolcanInterface{
		Cases: CaseManager{},
	}

	Error = &ErrorHandler{}

	return nil
}
