# 📡 KJP

(KAMAR JSON Pass-through)

This project is a server which maps api routes and JSON body / query parameters to KAMAR api commands and fields

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
- [ ] GetNotices
- [ ] GetStudentResults
- [ ] GetStudentAttendance
- [ ] GetCalendar
- [ ] GetEvents
- [ ] GetStudentTimetable
- [ ] GetStudentNCEASumarry
- [ ] GetStudentGroups
- [ ] GetStudentAwards
- [ ] GetTeacherAbsLog
- [ ] SearchStudents
- [ ] GetStudentDetails
- [ ] GetStudentPastoral
- [ ] GetUserDetails
- [ ] GetStudentAbsentStatus
- [ ] GetTeacherTimetable

### Common Error Codes

| Code | Error         |
|------|---------------|
| -2   | Key Missing   |
| -3   | Invalid Key   |
| -4   | Unknown Page  |
| -7   | Access Denied |
