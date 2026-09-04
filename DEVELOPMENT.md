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

---

## Milestone 3 — URL Validation

### Implemented

* URL validation before shortening
* Only `http` and `https` schemes allowed
* URL host cannot be empty
* Invalid URLs return an error
* Invalid URL requests return `400 Bad Request`

### Validation Flow

```text
POST /shorten
↓
Handler
↓
Service
↓
validateURL()
↓
Invalid → Error → 400 Bad Request
Valid → Generate short code → Store URL
```

### Tests

* Valid URL
* Invalid random text
* Missing scheme
* Unsupported scheme
* Missing host
* Invalid URL returns `400 Bad Request`

### Concepts Learned

* Go `net/url`
* URL parsing
* Input validation
* Error propagation
* HTTP `400 Bad Request`
* HTTP `405 Method Not Allowed`
* `httptest`
* Service-layer testing
* Handler testing
