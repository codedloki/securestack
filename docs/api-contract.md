# SecureStack – API Contract

## Overview

SecureStack is a framework for discovering, managing, and securely accessing
open-source productivity tools through a unified interface.

The platform focuses on:
- Reducing repeated authentication across tools
- Controlled access to tools
- Secure, isolated execution environments
- Minimal setup for end users

This document defines the current and planned API surface.

---

## API Principles

- JSON over HTTP
- Stateless requests
- Authentication via JWT
- Versioning handled through releases
- Backend acts as a control plane, not the tool itself

---

## Authentication Model (v0.2.0)

**Status:** Designed / partially implemented

- JWT-based authentication
- Tokens issued on successful login
- Tokens passed via `Authorization` header
- No external identity providers

### Authorization Header

Authorization: Bearer <jwt_token>

---

## Standard Error Format

All error responses follow this structure:

```json
{
  "error": {
    "code": "ERROR_CODE",
    "message": "Human readable message"
  }
}
```

---

## Public Endpoints (No Authentication)

### Health Check

**GET** `/health`

Used to verify server availability.

Response:
```json
{
  "status": "ok"
}
```

---

## Authentication Endpoints

### Login

**POST** `/auth/login`

Authenticates a user and issues a JWT.

Request:
```json
{
  "username": "string",
  "password": "string"
}
```

Response:
```json
{
  "token": "jwt_token",
  "expires_in": 3600
}
```

Possible Errors:
- INVALID_CREDENTIALS
- USER_DISABLED

---

### Logout (Planned)

**POST** `/auth/logout`

Token invalidation / blacklist (not implemented yet).

---

## Tool Registry Endpoints

### List Tools

**GET** `/tools`

Returns all tools accessible to the authenticated user.

Authentication:
- Required

Response:
```json
[
  {
    "id": "string",
    "name": "OnlyOffice",
    "category": "productivity",
    "deployment": "container",
    "maintained": true
  }
]
```

---

### Get Tool Details

**GET** `/tools/{tool_id}`

Returns detailed metadata for a tool.

Authentication:
- Required

Response:
```json
{
  "id": "string",
  "name": "OnlyOffice",
  "description": "Open-source office suite",
  "category": "productivity",
  "deployment": "container",
  "security_model": "isolated",
  "maintained": true
}
```

---

## Tool Session Management (Planned)

### Start Tool Session

**POST** `/tools/{tool_id}/start`

### Stop Tool Session

**POST** `/tools/{tool_id}/stop`

---

## Role-Based Access Control (Future)

RBAC is not part of v0.2.0.

---

## Version Notes

- v0.1.0 – Tool registry and config validation
- v0.2.0 – Authentication introduced (partial)
