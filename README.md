# 📡 KJP

(KAMAR JSON Pass-through)

![Go](https://img.shields.io/badge/Powered%20By-Go-29BEB0?style=for-the-badge) 

![LOC](https://tokei.rs/b1/github/jacobtread/KJP)

This project is a server which maps api routes and JSON body / query parameters to KAMAR api commands and fields

This project is written in Go which allows to achieve the high level of speed and low memory usage that it provides. Which makes this an ideal
use case over KNI and KAW my previous projects which wrapped the data and decoded the XML. This is much faster and converts everything to easily parsable JSON


## ☝ How it works

Each KAMAR command is mapped to a route and will use either GET or POST depending on the method. Soon all the different
methods will be listed here. 

The parameters of the get request are also mapped to KAMAR post bodies so an example login request is the following:

```http request
POST https://example.com/api/login
Authorization: vtku
X-Portal: portal.yourschool.school.nz
Content-Type: application/json

{
    "username": {USERNAME},
    "password": {PASSWORD}
}
```

The X-Portal header is the domain that the target KAMAR portal is hosted on
and the Authorization header is the current KAMAR auth Key leaving this blank will default to the standard auth key

This request will respond with 

```json
{
    "access_level": 0,
    "error": "",
    "error_code": 0,
    "success": "YES",
    "login_level": {LOGIN_LEVEL},
    "current_student": {ID},
    "key": {ACCESS_KEY}
}
```

## 🗺️ Mapped Routes

The following is a list of the API routes and the ones marked with a checkmark have been mapped

- [x] Logon
- [x] GetSettings
- [x] GetGlobals
- [x] GetNotices
- [x] GetStudentResults
- [x] GetStudentAttendance
- [x] GetCalendar
- [x] GetEvents
- [x] GetStudentTimetable
- [x] GetStudentNCEASummary
- [x] GetStudentGroups
- [x] GetStudentAwards
- [x] GetStudentDetails
- [x] GetStudentAbsentStatus

The following routes cannot be mapped as I do not have the required access levels to map them
- [ ] GetTeacherAbsLog
- [ ] SearchStudents
- [ ] GetStudentPastoral
- [ ] GetUserDetails
- [ ] GetTeacherTimetable

### Common Error Codes

| Code | Error         |
|------|---------------|
| -2   | Key Missing   |
| -3   | Invalid Key   |
| -4   | Unknown Page  |
| -7   | Access Denied |
