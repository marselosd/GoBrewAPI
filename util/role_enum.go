package util 

type RoleENUM uint8

const (
	SWE = iota 
	QA
	PO
	UXUI
	INACTIVE
)

func (r RoleENUM) String() string {
	switch r {
	case SWE:
		return "SOFTWARE DEVELOPER"
	case QA:
		return "QUALITY ASSURANCE"
	case PO:
		return "PRODUCT OWNER"
	case UXUI:
		return "UXUI"
	case INACTIVE:
		return "INACTIVE"
	}

	return "unknown"
}