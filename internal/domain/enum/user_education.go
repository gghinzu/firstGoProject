package enum

type UserEducation string

const (
	None            UserEducation = "None"
	PrimarySchool   UserEducation = "Primary School"
	MiddleSchool    UserEducation = "Middle School"
	HighSchool      UserEducation = "High School"
	BachelorsDegree UserEducation = "Bachelor's Degree"
	MastersDegree   UserEducation = "Master's Degree"
	Doctorate       UserEducation = "Doctorate"
)
