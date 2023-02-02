package handlers

import "github.com/gorilla/mux"

// интерфейс, чтобы показать, как работать с интерфейсами
type Handler interface {
	RegisterHandler(r *mux.Router)
}
