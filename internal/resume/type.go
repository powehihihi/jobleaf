package resume

type Resume struct {
	Contact    *Contact
	Summary    string
	Education  []*Education
	Experience []*Experience
	Projects   []*Project
	Skills     *Skills
}

type Contact struct {
	Name     string
	Mobile   string
	Mail     string
	Github   string
	Linkedin string
	Telegram string
	Website  string
}

type Education struct {
	Name       string
	URL        string
	Degree     string
	Location   string
	Duration   string
	GPA        string
	Coursework string
	Thesis     *Thesis
}

type Thesis struct {
	Title string
	URL   string
}

type Experience struct {
	Company *Company
	Roles   []*Role
}

type Company struct {
	Name     string
	URL      string
	Location string
}

type Role struct {
	Team     string
	Position string
	Duration string
	Tags     []string
	Summary  []string
}

type Project struct {
	Name     string
	URL      string
	Duration string
	Tags     []string
	Summary  []string
}

type Skills struct {
	Languages []string
	Libraries []string
	Tools     []string
}
