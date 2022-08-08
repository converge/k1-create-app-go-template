package health

import "net/http"

type Health struct {
}

func NewHealth() Health {
	return Health{}
}

func (handler Health) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("I'm health!"))
}
