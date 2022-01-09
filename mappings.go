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
	AccessLevel int8   `xml:"AccessLevel" json:"access_level"`
	Error       string `xml:"Error" json:"error,omitempty"`
	ErrorCode   int8   `xml:"ErrorCode" json:"error_code"`

	NumberRecords     uint16             `xml:"NumberRecords" json:"number_records"`
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
	AccessLevel int8   `xml:"AccessLevel" json:"access_level"`
	Error       string `xml:"Error" json:"error,omitempty"`
	ErrorCode   int8   `xml:"ErrorCode" json:"error_code"`

	NoticeDate    string `xml:"NoticeDate" json:"date"`
	NumberRecords uint32 `xml:"NumberRecords" json:"number_records"`

	Meetings []MeetingNotice `xml:"MeetingNotices>Meeting" json:"meetings"`
	Notices  []GeneralNotice `xml:"GeneralNotices>General" json:"notices"`
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
	AccessLevel int8   `xml:"AccessLevel" json:"access_level"`
	Error       string `xml:"Error" json:"error,omitempty"`
	ErrorCode   int8   `xml:"ErrorCode" json:"error_code"`

	NumberRecords uint16 `xml:"NumberRecords" json:"number_records"`
	StudentID     uint32 `xml:"StudentID" json:"student_id"`

	Results []ResultLevel `xml:"ResultLevels>ResultLevel" json:"results"`
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
		Response:   func() interface{} { return &StudentAttendanceResults{} },
		Parameters: map[string]ParameterMapping{
			"grid": {
				Name:     "Grid",
				Required: true,
			},
		},
	},
}
