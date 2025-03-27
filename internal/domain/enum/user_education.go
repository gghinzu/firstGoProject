package enum

type UserEducation int

const (
	None            UserEducation = 1
	PrimarySchool   UserEducation = 2
	MiddleSchool    UserEducation = 3
	HighSchool      UserEducation = 4
	BachelorsDegree UserEducation = 5
	MastersDegree   UserEducation = 6
	Doctorate       UserEducation = 7
)

func (e UserEducation) String() string {
	switch e {
	case None:
		return "None"
	case PrimarySchool:
		return "Primary School"
	case MiddleSchool:
		return "Middle School"
	case HighSchool:
		return "High School"
	case BachelorsDegree:
		return "Bachelor's Degree"
	case MastersDegree:
		return "Master's Degree"
	case Doctorate:
		return "Doctorate"
	default:
		return "Unknown"
	}
}
