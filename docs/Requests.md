# 🌐 Requests

This page contains documentation and information about making requests to KJP

## 📚 Required Header

In order to direct your requests to the correct portal you need to provide the ``X-Portal`` header which needs to be the
domain of the portal. e.g.

```http request
POST https://example.com/api/example
X-Portal: example.school.nz
```

If you do not provide this header you will be given an error

## 🔐 Authentication

Routes that are protected with an access key will HTTP 401 with the response
`You must provided a Authorization header to access this route` to get past this you must first authenticate using
the [Logon](Logon.md) route, and you will receive an authentication key you will need to pass this authentication key in
the
``Authorization`` header on any request that requires authentication e.g

```http request
POST https://example.com/api/example
Authorization: MY_AUTHENTICATION_TOKEN 
```

Any routes that require this will have the following line of text at the top of their documentation

*⚠️ REQUIRES AUTHENTICATION HEADER ⚠️*

## 📜 Pages

The documentation is not yet complete so not all routes have documentation listed here

- [Logon](Logon.md) - Authenticating with a portal to get an access key
