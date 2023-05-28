package model

type CandidateModel struct {
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	MiddleName      string `json:"middleName"`
	Gender          string `json:"gender"`
	City            string `json:"city"`
	District        string `json:"district"`
	DateOfBirth     string `json:"dateOfBirth"`
	Citizenship     string `json:"citizenship"`
	Education       string `json:"education"`
	Email           string `json:"email"`
	PhoneNumber     string `json:"phoneNumber"`
	Institution     string `json:"institution"`
	InstitutionCity string `json:"institutionCity"`
	Faculty         string `json:"faculty"`
	Specialization  string `json:"specialization"`
	GraduationYear  string `json:"graduationYear"`
	EducationLevel  string `json:"educationLevel"`
	WorkExperience  string `json:"workExperience"`
	InternshipField string `json:"internshipField"`
	WorkSchedule    string `json:"workSchedule"`
	ProfilePic      string `json:"profilePic"`
	VKProfile       string `json:"vkProfile"`
	TelegramProfile string `json:"telegramProfile"`
}
