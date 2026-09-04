# URL Shortener — Development 

## Milestone 1 — URL Shortening

### Implemented

* POST `/shorten`
* Random short-code generation
* In-memory storage

### Flow

```text
HTTP Request
↓
Handler
↓
Service
↓
Storage
```

### Concepts Learned

* Go packages
* Interfaces
* Dependency injection
* `crypto/rand`
* HTTP handlers
* JSON encoding/decoding

---

## Milestone 2 — URL Redirect

### Implemented

* GET `/{shortCode}`
* URL retrieval
* HTTP redirect

### Tests

* Successful redirect
* URL not found → `404`
* Wrong HTTP method → `405`
