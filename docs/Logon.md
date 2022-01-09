# 🔑 Logon Request

The login request authenticates you with the KAMAR portal and returns an access token, student ID, and access level of
the logged in account

### 🛣️ Route

**POST** ```/api/login```

### 📁 Parameters

| Key      | Type   | Description                 |
|----------|--------|-----------------------------|
| username | string | The username of the account |
| password | string | The password of the account |

### Example

The following is an example request t

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

### ✔️ Example Response

```json
{
  // Present on every request indicates the access level of the current key
  "access_level": 0,
  "error_code": 0,
  // Always YES on success
  "success": "YES",
  // The login level of the account, Permissions are listed in GetSettings
  "login_level": 0,
  // The current student ID
  "current_student": 0,
  // The authentication key
  "key": "🔑"
}
```

### ❓ Incorrect Username/Password response

This response will occur if you send the incorrect username/password

```json
{
  // Present on every request indicates the access level of the current key
  "access_level": 0,
  // The error message
  "error": "Incorrect Username or Password entered",
  "error_code": 0
}
```

### 🐢 Rate Limited Error response

This response will occur if you send the incorrect password for an account too many times

```json
{
  // Present on every request indicates the access level of the current key
  "access_level": 0,
  // The error message
  "error": "Too many failed login attempts, try again in {TIME} minutes",
  "error_code": 0
}
```