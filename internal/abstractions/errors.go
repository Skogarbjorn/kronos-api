package abstractions

import "net/http"

type ErrorWriter func(w http.ResponseWriter, err error)
