package events

import (
	event_database "OlympClub/pkg/database/events"
	"OlympClub/pkg/database/news"
	session_database "OlympClub/pkg/database/sessions"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type EventHandler struct {
	EventModel     *event_database.EventTable
	SessionModel   *session_database.SessionModel
	EventUserModel *event_database.EventUserTable
	NewsModel      *news.NewsTable
}

func (h *EventHandler) RegisterHandler(r *mux.Router) {
	r.HandleFunc("/event", h.CreateEvent).Methods("POST")
	r.HandleFunc("/events", h.GetEvents).Methods("GET", "OPTIONS")
	r.HandleFunc("/events/my", h.GetUserEvents).Methods("GET", "OPTIONS")
	r.HandleFunc("/event/{event_id}", h.GetEventById).Methods("GET", "OPTIONS")
	r.HandleFunc("/event/{event_id}/news", h.GetEventNewsById).Methods("GET", "OPTIONS")
	r.HandleFunc("/event/{event_id}/add", h.AddEventToUser).Methods("GET", "OPTIONS")

}

func (h *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.Header().Set("Content-Type", "application/json")
	ans := make(map[string]interface{})
	name := r.Form.Get("name")
	description := r.Form.Get("description")
	short := r.Form.Get("short")
	img := r.Form.Get("img")
	status := r.Form.Get("status")
	holderID, err := strconv.Atoi(r.Form.Get("holder_id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ans["error"] = "Problem with Database"
		json.NewEncoder(w).Encode(ans)
		return
	}
	website := r.Form.Get("website")

	event := event_database.Event{
		Name:        name,
		Description: description,
		Short:       short,
		Img:         img,
		Status:      status,
		HolderID:    int32(holderID),
		Website:     website,
	}
	event, err = h.EventModel.CreateEvent(event)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ans["error"] = "Problem with Database"
		json.NewEncoder(w).Encode(ans)
		return
	}
	ans["event_id"] = event.ID
	json.NewEncoder(w).Encode(ans)
	w.WriteHeader(http.StatusOK)
}

func (h *EventHandler) GetEvents(w http.ResponseWriter, r *http.Request) {
	ans := make(map[string]interface{})
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()
	eventStr := r.Form.Get("id")
	eventId, err := strconv.Atoi(eventStr)
	eventShort := r.Form.Get("short")
	if err != nil {
		eventId = -1
		eventShort = eventStr
	}
	holderStr := r.Form.Get("holder_id")
	holderId, err := strconv.Atoi(holderStr)
	if err != nil {
		holderId = -1
	}

	filter := event_database.EventFilter{
		ID:       int32(eventId),
		Short:    eventShort,
		HolderID: int32(holderId),
	}
	events, err := h.EventModel.GetEvents(filter)
	fmt.Println(events)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ans["error"] = "Problem with Database"
		json.NewEncoder(w).Encode(ans)
		return
	}
	fmt.Println(events)
	w.WriteHeader(http.StatusOK)
	ans["events"] = events
	json.NewEncoder(w).Encode(ans)
}

func (h *EventHandler) GetEventById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	ans := make(map[string]interface{})
	eventStr := mux.Vars(r)["event_id"]
	filter := event_database.NewEventFilter()
	eventID, err := strconv.Atoi(eventStr)
	if err != nil {
		filter.Short = eventStr
	} else {
		filter.ID = int32(eventID)
	}
	events, err := h.EventModel.GetEvents(filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ans["error"] = "Problem with Database"
		json.NewEncoder(w).Encode(ans)
		return
	}
	if len(events) == 0 {
		w.WriteHeader(http.StatusNotFound)
		ans["error"] = "Olympiad not found"
		json.NewEncoder(w).Encode(ans)
		return
	}

	w.WriteHeader(http.StatusOK)
	ans["event"] = events[0]
	fmt.Println(ans)
	json.NewEncoder(w).Encode(ans)
}

func (h *EventHandler) GetEventNewsById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	ans := make(map[string]interface{})
	eventStr := mux.Vars(r)["event_id"]
	filter := event_database.NewEventFilter()
	eventID, err := strconv.Atoi(eventStr)
	if err != nil {
		filter.Short = eventStr
	} else {
		filter.ID = int32(eventID)
	}
	events, err := h.EventModel.GetEvents(filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ans["error"] = "Problem with Database"
		json.NewEncoder(w).Encode(ans)
		return
	}
	if len(events) == 0 {
		w.WriteHeader(http.StatusNotFound)
		ans["error"] = "Olympiad not found"
		json.NewEncoder(w).Encode(ans)
		return
	}
	event := events[0]
	newsFiler := news.NewsFilter{
		ID:    -1,
		Table: "Events",
		Key:   event.ID,
	}
	news, err := h.NewsModel.GetNews(newsFiler)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ans["error"] = "Problem with Database"
		json.NewEncoder(w).Encode(ans)
		return
	}
	w.WriteHeader(http.StatusOK)
	ans["news"] = news
	fmt.Println(ans)
	json.NewEncoder(w).Encode(ans)
}

func (h *EventHandler) GetUserEvents(w http.ResponseWriter, r *http.Request) {
	ans := make(map[string]interface{})
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		token := r.Header.Get("Authorization")
		token = token[7:]
		session, err := h.SessionModel.GetSessions(token)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ans["error"] = "Problem with Database"
			json.NewEncoder(w).Encode(ans)
			return
		}
		if len(session) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			ans["error"] = "No such user"
			json.NewEncoder(w).Encode(ans)
			return
		}
		events, err := h.EventUserModel.GetEvents(session[0].UserID)
		fmt.Println(events)
		fmt.Println(events)
		w.WriteHeader(http.StatusOK)
		ans["events"] = events
		json.NewEncoder(w).Encode(ans)
		return
	}
	w.WriteHeader(http.StatusOK)
	ans["ok"] = "ok"
	json.NewEncoder(w).Encode(ans)
}

func (h *EventHandler) AddEventToUser(w http.ResponseWriter, r *http.Request) {
	ans := make(map[string]interface{})
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		token := r.Header.Get("Authorization")
		if len(token) < 7 {
			w.WriteHeader(http.StatusBadRequest)
			ans["error"] = "Wrong format of token"
			json.NewEncoder(w).Encode(ans)
			return
		}
		token = token[7:]
		session, err := h.SessionModel.GetSessions(token)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ans["error"] = "Problem with Database"
			json.NewEncoder(w).Encode(ans)
			return
		}
		if len(session) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			ans["error"] = "No such user"
			json.NewEncoder(w).Encode(ans)
			return
		}
		eventStr := mux.Vars(r)["event_id"]
		eventID, err := strconv.Atoi(eventStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			ans["error"] = "Wrong format of event"
			json.NewEncoder(w).Encode(ans)
			return
		}
		err = h.EventUserModel.CreateConnection(session[0].UserID, int32(eventID))
		if errors.Is(err, event_database.ErrConnectionExists) {
			w.WriteHeader(http.StatusBadRequest)
			ans["error"] = "Already is subscribed"
			json.NewEncoder(w).Encode(ans)
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ans["error"] = "Problem with Database"
			json.NewEncoder(w).Encode(ans)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	ans["ok"] = "ok"
	json.NewEncoder(w).Encode(ans)
}
