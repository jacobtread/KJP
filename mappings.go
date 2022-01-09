package main

type LoginResults struct {
	AccessLevel int8   `xml:"AccessLevel" json:"access_level"`
	Error       string `xml:"Error" json:"error,omitempty"`
	ErrorCode   int8   `xml:"ErrorCode" json:"error_code"`

	Success        string `xml:"Success" json:"success"`
	LoginLevel     int8   `xml:"LoginLevel" json:"login_level"`
	CurrentStudent int32  `xml:"CurrentStudent" json:"current_student"`
	Key            string `xml:"Key" json:"key"`
}

type SettingsResults struct {
	AccessLevel int8   `xml:"AccessLevel" json:"access_level"`
	Error       string `xml:"Error" json:"error,omitempty"`
	ErrorCode   int8   `xml:"ErrorCode" json:"error_code"`

	SettingsVersion        string              `xml:"SettingsVersion" json:"settings_version"`
	MiniOSVersion          string              `xml:"MiniOSVersion" json:"min_ios_version"`
	MinAndroidVersion      string              `xml:"MinAndroidVersion" json:"min_android_version"`
	StudentsAllowed        bool                `xml:"StudentsAllowed" json:"students_allowed"`
	StaffAllowed           bool                `xml:"StaffAllowed" json:"staff_allowed"`
	StudentsSavedPasswords bool                `xml:"StudentsSavedPasswords" json:"students_saved_passwords"`
	StaffSavedPasswords    bool                `xml:"StaffSavedPasswords" json:"staff_saved_passwords"`
	SchoolName             string              `xml:"SchoolName" json:"school_name"`
	LogoPath               string              `xml:"LogoPath" json:"logo_path"`
	AssessmentTypesShown   bool                `xml:"AssessmentTypesShown" json:"assessment_types_shown"`
	ShowEnrolledEntries    bool                `xml:"ShowEnrolledEntries" json:"show_enrolled_entries"`
	UserAccess             []SettingUserAccess `xml:"UserAccess>User" json:"user_access"`
	CalendarSettings       SettingsCalendar    `xml:"UserAccess>CalendarSettings" json:"calendar_settings"`
}

type SettingUserAccess struct {
	Notices         bool `xml:"Notices" json:"notices"`
	Events          bool `xml:"Events" json:"events"`
	Details         bool `xml:"Details" json:"details"`
	Timetable       bool `xml:"Timetable" json:"timetable"`
	Attendance      bool `xml:"Attendance" json:"attendance"`
	NCEA            bool `xml:"NCEA" json:"ncea"`
	Results         bool `xml:"Results" json:"results"`
	Groups          bool `xml:"Groups" json:"groups"`
	Awards          bool `xml:"Awards" json:"awards"`
	Pastoral        bool `xml:"Pastoral" json:"pastoral"`
	ReportAbsencePg bool `xml:"ReportAbsencePg" json:"report_absence_pg"`
	ReportAbsence   bool `xml:"ReportAbsence" json:"report_absence"`
}

type SettingsCalendar struct {
	Aero      string `xml:"aero" json:"aero"`
	Amazon    string `xml:"amazon" json:"amazon"`
	Auburn    string `xml:"auburn" json:"auburn"`
	Blue      string `xml:"blue" json:"blue"`
	Green     string `xml:"green" json:"green"`
	Orange    string `xml:"orange" json:"orange"`
	Purple    string `xml:"purple" json:"purple"`
	Red       string `xml:"red" json:"red"`
	Tangerine string `xml:"tangerine" json:"tangerine"`
	Teal      string `xml:"teal" json:"teal"`
	Violet    string `xml:"violet" json:"violet"`
	Black     string `xml:"black" json:"black"`
}

type GlobalsResults struct {
	AccessLevel   int8   `xml:"AccessLevel" json:"access_level"`
	Error         string `xml:"Error" json:"error,omitempty"`
	ErrorCode     int8   `xml:"ErrorCode" json:"error_code"`
	NumberRecords uint16 `xml:"NumberRecords" json:"number_records"`

	PeriodDefinitions []PeriodDefinition `xml:"PeriodDefinitions>PeriodDefinition" json:"period_definitions"`
	StartTimes        []GlobalsDay       `xml:"StartTimes>Day" json:"start_times"`
}

type PeriodDefinition struct {
	Name string `xml:"PeriodName" json:"name"`
	Time string `xml:"PeriodTime" json:"time"`
}

type GlobalsDay struct {
	Index       uint16   `xml:"index,attr" json:"index"`
	PeriodTimes []string `xml:"PeriodTime" json:"times"`
}

type NoticesResult struct {
	AccessLevel   int8   `xml:"AccessLevel" json:"access_level"`
	Error         string `xml:"Error" json:"error,omitempty"`
	ErrorCode     int8   `xml:"ErrorCode" json:"error_code"`
	NumberRecords uint32 `xml:"NumberRecords" json:"number_records"`

	NoticeDate string          `xml:"NoticeDate" json:"date"`
	Meetings   []MeetingNotice `xml:"MeetingNotices>Meeting" json:"meetings"`
	Notices    []GeneralNotice `xml:"GeneralNotices>General" json:"notices"`
}

type MeetingNotice struct {
	Level     string `xml:"Level" json:"level"`
	Subject   string `xml:"Subject" json:"subject"`
	Body      string `xml:"Body" json:"body"`
	Teacher   string `xml:"Teacher" json:"teacher"`
	PlaceMeet string `xml:"PlaceMeet" json:"place"`
	DateMeet  string `xml:"DateMeet" json:"date"`
	TimeMeet  string `xml:"TimeMeet" json:"time"`
}

type GeneralNotice struct {
	Level   string `xml:"Level" json:"level"`
	Subject string `xml:"Subject" json:"subject"`
	Body    string `xml:"Body" json:"body"`
	Teacher string `xml:"Teacher" json:"teacher"`
}

type StudentResultsResults struct {
	AccessLevel   int8   `xml:"AccessLevel" json:"access_level"`
	Error         string `xml:"Error" json:"error,omitempty"`
	ErrorCode     int8   `xml:"ErrorCode" json:"error_code"`
	NumberRecords uint16 `xml:"NumberRecords" json:"number_records"`

	StudentID uint32        `xml:"StudentID" json:"student_id"`
	Results   []ResultLevel `xml:"ResultLevels>ResultLevel" json:"results"`
}

type ResultLevel struct {
	NumberResults uint32   `xml:"NumberResults" json:"number_results"`
	NCEALevel     uint8    `xml:"NCEALevel" json:"ncea_level"`
	Results       []Result `xml:"Results>Result" json:"results"`
}

type Result struct {
	Identifier      string `xml:"Number" json:"id"`
	Version         uint16 `xml:"Version" json:"version"`
	Grade           string `xml:"Grade" json:"grade"`
	Title           string `xml:"Title" json:"title"`
	SubField        string `xml:"SubField" json:"sub_field"`
	Credits         uint16 `xml:"Credits" json:"credits"`
	CreditsPassed   uint16 `xml:"CreditsPassed" json:"credits_passed"`
	ResultPublished string `xml:"ResultPublished" json:"result_published"`
}

type StudentAttendanceResults struct {
	AccessLevel int8   `xml:"AccessLevel" json:"access_level"`
	Error       string `xml:"Error" json:"error,omitempty"`
	ErrorCode   int8   `xml:"ErrorCode" json:"error_code"`

	NumberRecords uint8            `xml:"NumberRecords" json:"number_records"`
	Weeks         []AttendanceWeek `xml:"Weeks>Week" json:"weeks"`
}

type AttendanceWeek struct {
	WeekStart string   `xml:"WeekStart" json:"start"`
	Days      []string `xml:"Days>Day" json:"days"`
}

type CalendarResults struct {
	AccessLevel int8   `xml:"AccessLevel" json:"access_level"`
	Error       string `xml:"Error" json:"error,omitempty"`
	ErrorCode   int8   `xml:"ErrorCode" json:"error_code"`

	NumberRecords uint16        `xml:"NumberRecords" json:"number_records"`
	Days          []CalendarDay `xml:"Days>Day" json:"days"`
}

type EventsResults struct {
	AccessLevel int8   `xml:"AccessLevel" json:"access_level"`
	Error       string `xml:"Error" json:"error,omitempty"`
	ErrorCode   int8   `xml:"ErrorCode" json:"error_code"`

	NumberRecords uint16  `xml:"NumberRecords" json:"number_records"`
	Events        []Event `xml:"Events>Event" json:"events"`
}

type Event struct {
	Title        string `xml:"Title" json:"title"`
	Location     string `xml:"Location" json:"location"`
	Details      string `xml:"Details" json:"details"`
	Priority     int8   `xml:"Priority" json:"priority"`
	Public       bool   `xml:"Public" json:"public"`
	Student      bool   `xml:"Student" json:"student"`
	CG1          bool   `xml:"CG1" json:"cg_1"`
	CG2          bool   `xml:"CG2" json:"cg_2"`
	Staff        bool   `xml:"Staff" json:"staff"`
	Color        string `xml:"Color" json:"color"`
	ColorLabel   string `xml:"ColorLabel" json:"color_label"`
	DateTimeInfo string `xml:"DateTimeInfo,omitempty" json:"date_time,omitempty"`
	TimeStart    string `xml:"DateTimeStart,omitempty" json:"time_start,omitempty"`
	TimeFinish   string `xml:"DateTimeFinish,omitempty" json:"time_finish,omitempty"`
	DateStart    string `xml:"Start" json:"date_start"`
	DateFinish   string `xml:"Finish" json:"date_finish"`
}

type CalendarDay struct {
	Date     string `xml:"Date" json:"date"`
	Status   string `xml:"Status" json:"status"`
	DayTT    uint16 `xml:"DayTT,omitempty" json:"day_tt,omitempty"`
	Term     uint16 `xml:"Term,omitempty" json:"term,omitempty"`
	TermA    uint16 `xml:"TermA,omitempty" json:"term_a,omitempty"`
	Week     uint16 `xml:"Week,omitempty" json:"week,omitempty"`
	WeekA    uint16 `xml:"WeekA,omitempty" json:"week_a,omitempty"`
	WeekYear uint16 `xml:"WeekYear,omitempty" json:"week_year,omitempty"`
	TermYear uint16 `xml:"TermYear,omitempty" json:"term_year,omitempty"`
}

type StudentTimetableResults struct {
	AccessLevel   int8   `xml:"AccessLevel" json:"access_level"`
	Error         string `xml:"Error" json:"error,omitempty"`
	ErrorCode     int8   `xml:"ErrorCode" json:"error_code"`
	NumberRecords uint16 `xml:"NumberRecords" json:"number_records"`

	TTGrid string `xml:"TTGrid" json:"tt_grid"`

	Student TimetableStudent `xml:"Students>Student" json:"student"`
}

type TimetableStudent struct {
	ID    int32  `xml:"IDNumber" json:"id"`
	Level int8   `xml:"Level" json:"level"`
	Tutor string `xml:"Tutor" json:"tutor"`

	TimetableData TimetableWeeks `xml:"TimetableData" json:"data"`
}

type TimetableWeeks struct {
	Weeks []TimetableWeek `xml:",any" json:"weeks"`
}

type TimetableWeek struct {
	Days []string `xml:",any" json:"days"`
}

type StudentNCEASummaryResults struct {
	AccessLevel   int8   `xml:"AccessLevel" json:"access_level"`
	Error         string `xml:"Error" json:"error,omitempty"`
	ErrorCode     int8   `xml:"ErrorCode" json:"error_code"`
	NumberRecords uint16 `xml:"NumberRecords" json:"number_records"`

	Students []NCEAStudent `xml:"Students>Student" json:"students"`
}

type NCEAStudent struct {
	CreditsTotal    CreditSummary  `xml:"CreditsTotal" json:"credits_total"`
	CreditsInternal CreditSummary  `xml:"CreditsInternal" json:"credits_internal"`
	CreditsExternal CreditSummary  `xml:"CreditsExternal" json:"credits_external"`
	YearTotals      []YearSummary  `xml:"YearTotals>YearTotal" json:"year_totals"`
	LevelTotals     []LevelSummary `xml:"LevelTotals>LevelTotal" json:"level_totals"`
	NCEASummary     NCEASummary    `xml:"NCEA" json:"ncea"`
}

type CreditSummary struct {
	NotAchieved uint32 `xml:"NotAchieved" json:"not_achieved"`
	Achieved    uint32 `xml:"Achieved" json:"achieved"`
	Merit       uint32 `xml:"Merit" json:"merit"`
	Excellence  uint32 `xml:"Excellence" json:"excellence"`
	Total       uint32 `xml:"Total" json:"total"`
	Attempted   uint32 `xml:"Attempted" json:"attempted"`
}

type YearSummary struct {
	Year        string `xml:"Year" json:"year"`
	NotAchieved uint32 `xml:"NotAchieved" json:"not_achieved"`
	Achieved    uint32 `xml:"Achieved" json:"achieved"`
	Merit       uint32 `xml:"Merit" json:"merit"`
	Excellence  uint32 `xml:"Excellence" json:"excellence"`
	Total       uint32 `xml:"Total" json:"total"`
	Attempted   uint32 `xml:"Attempted" json:"attempted"`
}

type LevelSummary struct {
	Level       uint16 `xml:"Level" json:"level"`
	NotAchieved uint32 `xml:"NotAchieved" json:"not_achieved"`
	Achieved    uint32 `xml:"Achieved" json:"achieved"`
	Merit       uint32 `xml:"Merit" json:"merit"`
	Excellence  uint32 `xml:"Excellence" json:"excellence"`
	Total       uint32 `xml:"Total" json:"total"`
	Attempted   uint32 `xml:"Attempted" json:"attempted"`
}

type NCEASummary struct {
	L1           string `xml:"L1NCEA" json:"level_1"`
	L2           string `xml:"L2NCEA" json:"level_2"`
	L3           string `xml:"L3NCEA" json:"level_3"`
	UELiteracy   string `xml:"NCEAUELIT" json:"ue_literacy"`
	Numeracy     string `xml:"NCEANUM" json:"numeracy"`
	Lvl1Literacy string `xml:"NCEAL1LIT" json:"l1_literacy"`
}

type StudentGroupsResults struct {
	AccessLevel   int8               `xml:"AccessLevel" json:"access_level"`
	Error         string             `xml:"Error" json:"error,omitempty"`
	ErrorCode     int8               `xml:"ErrorCode" json:"error_code"`
	NumberRecords uint16             `xml:"NumberRecords" json:"number_records"`
	Years         []StudentGroupYear `xml:"Years>Year" json:"years"`
}

type StudentGroupYear struct {
	Grid         string  `xml:"Grid" json:"grid"`
	NumberGroups uint16  `xml:"NumberGroups" json:"number_groups"`
	Groups       []Group `xml:"Groups>Group" json:"groups"`
}

type Group struct {
	Name    string `xml:"Name" json:"name"`
	Teacher string `xml:"Teacher" json:"teacher"`
}

type StudentAwardsResults struct {
	AccessLevel   int8        `xml:"AccessLevel" json:"access_level"`
	Error         string      `xml:"Error" json:"error,omitempty"`
	ErrorCode     int8        `xml:"ErrorCode" json:"error_code"`
	NumberRecords uint16      `xml:"NumberRecords" json:"number_records"`
	NumberYears   uint16      `xml:"NumberYears" json:"number_years"`
	Years         []AwardYear `xml:"Years>Year" json:"years"`
}

type AwardYear struct {
	Grid         string  `xml:"Grid" json:"grid"`
	NumberAwards uint16  `xml:"NumberAwards" json:"number_awards"`
	Awards       []Award `xml:"Awards>Award"`
}

type Award struct {
	Title   string `xml:"Title" json:"title"`
	Teacher string `xml:"Teacher" json:"teacher,omitempty"`
	Details string `xml:"Details" json:"details,omitempty"`
}

type StudentDetailsResults struct {
	AccessLevel   int8    `xml:"AccessLevel" json:"access_level"`
	Error         string  `xml:"Error" json:"error,omitempty"`
	ErrorCode     int8    `xml:"ErrorCode" json:"error_code"`
	NumberRecords uint16  `xml:"NumberRecords" json:"number_records"`
	Student       Student `xml:"Students>Student" json:"student"`
}

type Student struct {
	KAMARVersion         string `xml:"KAMAR_Version" json:"kamar_version"`
	StudentID            int32  `xml:"StudentID" json:"student_id"`
	FirstName            string `xml:"FirstName" json:"first_name"`
	ForeNames            string `xml:"ForeNames" json:"fore_names"`
	LastName             string `xml:"LastName" json:"last_name"`
	FirstNameLegal       string `xml:"FirstNameLegal" json:"first_name_legal"`
	ForeNamesLegal       string `xml:"ForeNamesLegal" json:"fore_names_legal"`
	LastNameLegal        string `xml:"LastNameLegal" json:"last_name_legal"`
	Gender               string `xml:"Gender" json:"gender"`
	Ethnicity            string `xml:"Ethnicity" json:"ethnicity"`
	DateBirth            string `xml:"DateBirth" json:"date_birth"`
	Age                  string `xml:"Age" json:"age"`
	NSN                  string `xml:"NSN" json:"nsn"`
	DoctorName           string `xml:"DoctorName" json:"doctor_name"`
	DoctorPhone          string `xml:"DoctorPhone" json:"doctor_phone"`
	DoctorAddress        string `xml:"DoctorAddress" json:"doctor_address"`
	DentistName          string `xml:"DentistName" json:"dentist_name"`
	DentistPhone         string `xml:"DentistPhone" json:"dentist_phone"`
	DentistAddress       string `xml:"DentistAddress" json:"dentist_address"`
	AllowedPanadol       string `xml:"AllowedPanadol" json:"allowed_panadol"`
	AllowedIbuprofen     string `xml:"AllowedIbuprofen" json:"allowed_ibuprofen"`
	HealthFlag           string `xml:"HealthFlag" json:"health_flag"`
	Medical              string `xml:"Medical" json:"medical"`
	Reactions            string `xml:"Reactions" json:"reactions"`
	Vaccinations         string `xml:"Vaccinations" json:"vaccinations"`
	SpecialCircumstances string `xml:"SpecialCircumstances" json:"special_circumstances"`
	HomePhone            string `xml:"HomePhone" json:"home_phone"`
	HomeAddress          string `xml:"HomeAddress" json:"home_address"`
	ParentTitle          string `xml:"ParentTitle" json:"parent_title"`
	ParentSalutation     string `xml:"ParentSalutation" json:"parent_salutation"`
	ParentEmail          string `xml:"ParentEmail" json:"parent_email"`
	MotherRelation       string `xml:"MotherRelation" json:"mother_relation"`
	MotherName           string `xml:"MotherName" json:"mother_name"`
	MotherStatus         string `xml:"MotherStatus" json:"mother_status"`
	MotherEmail          string `xml:"MotherEmail" json:"mother_email"`
	MotherPhoneHome      string `xml:"MotherPhoneHome" json:"mother_phone_home"`
	MotherPhoneCell      string `xml:"MotherPhoneCell" json:"mother_phone_cell"`
	MotherPhoneWork      string `xml:"MotherPhoneWork" json:"mother_phone_work"`
	MotherPhoneExtn      string `xml:"MotherPhoneExtn" json:"mother_phone_extn"`
	MotherOccupation     string `xml:"MotherOccupation" json:"mother_occupation"`
	MotherWorkAddress    string `xml:"MotherWorkAddress" json:"mother_work_address"`
	MotherNotes          string `xml:"MotherNotes" json:"mother_notes"`
	FatherRelation       string `xml:"FatherRelation" json:"father_relation"`
	FatherName           string `xml:"FatherName" json:"father_name"`
	FatherStatus         string `xml:"FatherStatus" json:"father_status"`
	FatherEmail          string `xml:"FatherEmail" json:"father_email"`
	FatherPhoneHome      string `xml:"FatherPhoneHome" json:"father_phone_home"`
	FatherPhoneCell      string `xml:"FatherPhoneCell" json:"father_phone_cell"`
	FatherPhoneWork      string `xml:"FatherPhoneWork" json:"father_phone_work"`
	FatherPhoneExtn      string `xml:"FatherPhoneExtn" json:"father_phone_extn"`
	FatherOccupation     string `xml:"FatherOccupation" json:"father_occupation"`
	FatherWorkAddress    string `xml:"FatherWorkAddress" json:"father_work_address"`
	FatherNotes          string `xml:"FatherNotes" json:"father_notes"`
	StudentMobile        string `xml:"StudentMobile" json:"student_mobile"`
	StudentEmail         string `xml:"StudentEmail" json:"student_email"`
	StudentSchoolEmail   string `xml:"StudentSchoolEmail" json:"student_school_email"`
	GeneralNotes         string `xml:"GeneralNotes" json:"general_notes"`
	HealthNotes          string `xml:"HealthNotes" json:"health_notes"`
	EmergencyName        string `xml:"EmergencyName" json:"emergency_name"`
	EmergencyPhoneHome   string `xml:"EmergencyPhoneHome" json:"emergency_phone_home"`
	EmergencyPhoneCell   string `xml:"EmergencyPhoneCell" json:"emergency_phone_cell"`
	EmergencyPhoneWork   string `xml:"EmergencyPhoneWork" json:"emergency_phone_work"`
	EmergencyPhoneExtn   string `xml:"EmergencyPhoneExtn" json:"emergency_phone_extn"`
	EmergencyNotes       string `xml:"EmergencyNotes" json:"emergency_notes"`
}

type StudentAbsentStatus struct {
	AccessLevel   int8   `xml:"AccessLevel" json:"access_level"`
	Error         string `xml:"Error" json:"error,omitempty"`
	ErrorCode     int8   `xml:"ErrorCode" json:"error_code"`
	NumberRecords uint16 `xml:"NumberRecords" json:"number_records"`

	Student AbsenceStudent `xml:"Students>Student" json:"student"`
}

type AbsenceStudent struct {
	HJustified   int32 `xml:"HalfDaysJ" json:"hd_justified"`
	HUnjustified int32 `xml:"HalfDaysU" json:"hd_unjustified"`
	HOverseas    int32 `xml:"HalfDaysO" json:"hd_overseas"`
	HTotal       int32 `xml:"HalfDaysT" json:"hd_total"`
	HOpen        int32 `xml:"HalfDaysOpen" json:"hd_open"`

	FJustified   int32 `xml:"FullDaysJ" json:"justified"`
	FUnjustified int32 `xml:"FullDaysU" json:"unjustified"`
	FOverseas    int32 `xml:"FullDaysO" json:"overseas"`
	FTotal       int32 `xml:"FullDaysT" json:"total"`
	FOpen        int32 `xml:"FullDaysOpen" json:"open"`

	PercentJustified   int16 `xml:"PctgeJ" json:"percent_justified"`
	PercentUnjustified int16 `xml:"PctgeU" json:"percent_unjustified"`
	PercentOverseas    int16 `xml:"PctgeO" json:"percent_overseas"`
	PercentTotal       int16 `xml:"PctgeT" json:"percent_total"`
	PercentPresent     int16 `xml:"PctgeP" json:"percent_present"`
}

type ContentMapping struct {
	Method     string
	Command    string
	RequireKey bool
	Parameters map[string]ParameterMapping
	Response   func() interface{}
}

type ParameterMapping struct {
	Name     string
	Required bool
}

var Mappings = map[string]ContentMapping{
	"login": {
		Method:  "POST",
		Command: "Logon",
		Parameters: map[string]ParameterMapping{
			"username": {
				Name:     "Username",
				Required: true,
			},
			"password": {
				Name:     "Password",
				Required: true,
			},
		},
		Response: func() interface{} { return &LoginResults{} },
	},
	"settings": {
		Method:   "GET",
		Command:  "GetSettings",
		Response: func() interface{} { return &SettingsResults{} },
	},
	"globals": {
		Method:   "GET",
		Command:  "GetGlobals",
		Response: func() interface{} { return &GlobalsResults{} },
	},
	"notices": {
		Method:  "GET",
		Command: "GetNotices",
		Parameters: map[string]ParameterMapping{
			"date": {
				Name:     "Date",
				Required: true,
			},
		},
		Response: func() interface{} { return &NoticesResult{} },
	},
	"student_results": {
		Method:     "GET",
		RequireKey: true,
		Command:    "GetStudentResults",
		Response:   func() interface{} { return &StudentResultsResults{} },
	},
	"student_attendance": {
		Method:     "GET",
		RequireKey: true,
		Command:    "GetStudentAttendance",
		Parameters: map[string]ParameterMapping{
			"grid": {
				Name:     "Grid",
				Required: true,
			},
		},
		Response: func() interface{} { return &StudentAttendanceResults{} },
	},
	"calendar": {
		Method:  "GET",
		Command: "GetCalendar",
		Parameters: map[string]ParameterMapping{
			"year": {
				Name:     "Year",
				Required: true,
			},
		},
		Response: func() interface{} { return &CalendarResults{} },
	},
	"events": {
		Method:  "GET",
		Command: "GetEvents",
		Parameters: map[string]ParameterMapping{
			"start": {
				Name:     "DateStart",
				Required: true,
			},
			"end": {
				Name:     "DateFinish",
				Required: false,
			},
		},
		Response: func() interface{} { return &EventsResults{} },
	},
	"student_timetable": {
		Method:  "GET",
		Command: "GetStudentTimetable",
		Parameters: map[string]ParameterMapping{
			"grid": {
				Name:     "Grid",
				Required: true,
			},
			"id": {
				Name:     "StudentID",
				Required: false,
			},
		},
		Response: func() interface{} { return &StudentTimetableResults{} },
	},
	"student_ncea_summary": {
		Method:     "GET",
		RequireKey: true,
		Command:    "GetStudentNCEASummary",
		Response:   func() interface{} { return &StudentNCEASummaryResults{} },
	},
	"student_groups": {
		Method:     "GET",
		RequireKey: true,
		Command:    "GetStudentGroups",
		Response:   func() interface{} { return &StudentGroupsResults{} },
	},
	"student_awards": {
		Method:     "GET",
		RequireKey: true,
		Command:    "GetStudentAwards",
		Response:   func() interface{} { return &StudentAwardsResults{} },
	},
	"student_details": {
		Method:     "GET",
		RequireKey: true,
		Command:    "GetStudentDetails",
		Response:   func() interface{} { return &StudentDetailsResults{} },
	},
	"student_absence_stats": {
		Method:     "GET",
		RequireKey: true,
		Command:    "GetStudentAbsenceStats",
		Parameters: map[string]ParameterMapping{
			"grid": {
				Name:     "Grid",
				Required: true,
			},
		},
		Response: func() interface{} { return &StudentAbsentStatus{} },
	},
}
