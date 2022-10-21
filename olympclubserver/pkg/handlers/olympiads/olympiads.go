package olympiads

import (
	"OlympClub/pkg/database/news"
	"OlympClub/pkg/database/olympiads"
	database "OlympClub/pkg/database/olympiads"
	"OlympClub/pkg/database/sessions"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type OlympiadHandler struct {
	OlympiadModel     *database.OlympiadTable
	BigOlympiadModel  *database.BigOlympiadTable
	SessionModel      *sessions.SessionModel
	OlympiadUserModel *database.OlympiadUserTable
	NewsModel         *news.NewsTable
}

func (h *OlympiadHandler) RegisterHandler(r *mux.Router) {
	r.HandleFunc("/olympiad", h.CreateOlympiad).Methods("POST")
	r.HandleFunc("/big_olympiads", h.GetBigOlympiads).Methods("GET", "OPTIONS")
	r.HandleFunc("/olympiads", h.GetOlympiads).Methods("GET", "OPTIONS")
	r.HandleFunc("/olympiad/{big_olympiad_id}", h.GetBigOlympiad).Methods("GET", "OPTIONS")
	r.HandleFunc("/olympiad/{big_olympiad_id}/olympiads", h.GetOlympiadsByBigOlympiadId)
	r.HandleFunc("/olympiad/{big_olympiad_id}/{olympiad_id}", h.GetOlympiadById).Methods("GET", "OPTIONS")
	r.HandleFunc("/olympiad/{big_olympiad_id}/{olympiad_id}/news", h.GetOlympiadNewsById).Methods("GET", "OPTIONS")

	r.HandleFunc("/olympiads/my", h.GetUserOlympiads).Methods("GET", "OPTIONS")
	r.HandleFunc("/olympiad/{big_olympiad_id}/{olympiad_id}/add", h.AddOlympiadToUser).Methods("GET", "OPTIONS")
}

func (h *OlympiadHandler) CreateOlympiad(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.Header().Set("Content-Type", "application/json")
	ans := make(map[string]interface{})
	name := r.Form.Get("name")
	subject := r.Form.Get("subject")
	level := r.Form.Get("level")
	img := r.Form.Get("img")
	short := r.Form.Get("short")
	status := r.Form.Get("status")
	grade := r.Form.Get("grade")
	bigOlympiadID, err := strconv.Atoi(r.Form.Get("big_olympiad_id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ans["error"] = "Problem with Database"
		json.NewEncoder(w).Encode(ans)
		return
	}
	holderID, err := strconv.Atoi(r.Form.Get("holder_id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ans["error"] = "Problem with Database"
		json.NewEncoder(w).Encode(ans)
		return
	}
	website := r.Form.Get("website")
	olympiad := database.Olympiad{
		Name:          name,
		Subject:       subject,
		Level:         level,
		Img:           img,
		Short:         short,
		Status:        status,
		BigOlympiadID: int32(bigOlympiadID),
		Grade:         grade,
		HolderID:      int32(holderID),
		Website:       website,
	}
	olympiad, err = h.OlympiadModel.CreateOlympiad(olympiad)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ans["error"] = "Problem with Database"
		json.NewEncoder(w).Encode(ans)
		return
	}
	ans["olympiad_id"] = olympiad.ID
	json.NewEncoder(w).Encode(ans)
	w.WriteHeader(http.StatusOK)
}

func (h *OlympiadHandler) GetOlympiads(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	subject := r.URL.Query().Get("subject")
	level := r.URL.Query().Get("level")
	grade := r.URL.Query().Get("grade")
	filter := database.NewOlympiadFilter()
	filter.Subject = subject
	filter.Level = level
	filter.Grade = grade
	olympiads, err := h.OlympiadModel.GetOlympiads(filter)
	ans := make(map[string]interface{})
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ans["error"] = "Problem with Database"
		json.NewEncoder(w).Encode(ans)
		return
	}
	// fmt.Println(olympiads)
	w.WriteHeader(http.StatusOK)
	ans["olympiads"] = olympiads
	json.NewEncoder(w).Encode(ans)
}

func (h *OlympiadHandler) GetBigOlympiads(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	olympiads, err := h.BigOlympiadModel.GetBigOlympiads(olympiads.NewBigOlympiadFilter())
	ans := make(map[string]interface{})
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ans["error"] = "Problem with Database"
		json.NewEncoder(w).Encode(ans)
		return
	}
	// fmt.Println(olympiads)
	w.WriteHeader(http.StatusOK)
	ans["olympiads"] = olympiads
	json.NewEncoder(w).Encode(ans)
}

func (h *OlympiadHandler) GetBigOlympiad(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	ans := make(map[string]interface{})
	bigOlympiad := mux.Vars(r)["big_olympiad_id"]
	filter := database.NewBigOlympiadFilter()
	bigOlympiadID, err := strconv.Atoi(bigOlympiad)
	if err != nil {
		filter.Short = bigOlympiad
	} else {
		filter.ID = int32(bigOlympiadID)
	}
	bigOlympiads, err := h.BigOlympiadModel.GetBigOlympiads(filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ans["error"] = "Problem with Database"
		json.NewEncoder(w).Encode(ans)
		return
	}
	if len(bigOlympiads) == 0 {
		w.WriteHeader(http.StatusNotFound)
		ans["error"] = "Olympiad not found"
		json.NewEncoder(w).Encode(ans)
		return
	}

	w.WriteHeader(http.StatusOK)
	ans["big_olympiad"] = bigOlympiads[0]
	// fmt.Println(ans)
	json.NewEncoder(w).Encode(ans)
}

func (h *OlympiadHandler) GetOlympiadById(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	ans := make(map[string]interface{})

	olympiad := mux.Vars(r)["olympiad_id"]
	olympiadID, err := strconv.Atoi(olympiad)
	olympiadShort := ""
	if err != nil {
		olympiadShort = olympiad
		olympiadID = 0
	}
	bigOlympiad := mux.Vars(r)["big_olympiad_id"]
	bigOlympiadID, err := strconv.Atoi(bigOlympiad)
	if err != nil {
		bigOlympiadFilter := database.NewBigOlympiadFilter()
		bigOlympiadFilter.Short = bigOlympiad
		olympiads, err := h.BigOlympiadModel.GetBigOlympiads(bigOlympiadFilter)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ans["error"] = "Problem with Database"
			json.NewEncoder(w).Encode(ans)
			return
		}
		if len(olympiads) != 0 {
			bigOlympiadID = int(olympiads[0].ID)
		} else {
			bigOlympiadID = 0
		}
	}
	filter := database.NewOlympiadFilter()
	filter.OlympiadID = int32(olympiadID)
	filter.BigOlympiadID = int32(bigOlympiadID)
	filter.OlympiadShort = olympiadShort
	olympiads, err := h.OlympiadModel.GetOlympiads(filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ans["error"] = "Problem with Database"
		json.NewEncoder(w).Encode(ans)
		return
	}
	// fmt.Println(olympiads)
	if len(olympiads) != 0 {
		w.WriteHeader(http.StatusOK)
		ans["olympiad"] = olympiads[0]
		json.NewEncoder(w).Encode(ans)
	} else {
		w.WriteHeader(http.StatusNotFound)
		ans["error"] = "Olympiad not found"
		json.NewEncoder(w).Encode(ans)
	}
}

func (h *OlympiadHandler) GetOlympiadNewsById(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	ans := make(map[string]interface{})

	olympiad := mux.Vars(r)["olympiad_id"]
	olympiadID, err := strconv.Atoi(olympiad)
	olympiadShort := ""
	if err != nil {
		olympiadShort = olympiad
		olympiadID = 0
	}
	bigOlympiad := mux.Vars(r)["big_olympiad_id"]
	bigOlympiadID, err := strconv.Atoi(bigOlympiad)
	if err != nil {
		bigOlympiadFilter := database.NewBigOlympiadFilter()
		bigOlympiadFilter.Short = bigOlympiad
		olympiads, err := h.BigOlympiadModel.GetBigOlympiads(bigOlympiadFilter)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ans["error"] = "Problem with Database"
			json.NewEncoder(w).Encode(ans)
			return
		}
		if len(olympiads) != 0 {
			bigOlympiadID = int(olympiads[0].ID)
		} else {
			bigOlympiadID = 0
		}
	}
	filter := database.NewOlympiadFilter()
	filter.OlympiadID = int32(olympiadID)
	filter.BigOlympiadID = int32(bigOlympiadID)
	filter.OlympiadShort = olympiadShort
	olympiads, err := h.OlympiadModel.GetOlympiads(filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ans["error"] = "Problem with Database"
		json.NewEncoder(w).Encode(ans)
		return
	}
	// fmt.Println(olympiads)
	if len(olympiads) != 0 {
		olympiad := olympiads[0]

		newsFiler := news.NewsFilter{
			ID:    -1,
			Table: "Olympiads",
			Key:   olympiad.ID,
		}
		news, err := h.NewsModel.GetNews(newsFiler)
		// fmt.Println(err, news)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ans["error"] = "Problem with Database"
			json.NewEncoder(w).Encode(ans)
			return
		}
		ans["news"] = news
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ans)
	} else {
		w.WriteHeader(http.StatusNotFound)
		ans["error"] = "Olympiad not found"
		json.NewEncoder(w).Encode(ans)
	}
}

func (h *OlympiadHandler) GetOlympiadsByBigOlympiadId(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	ans := make(map[string]interface{})

	bigOlympiadShort := mux.Vars(r)["big_olympiad_id"]
	bigOlympiadID, err := strconv.Atoi(bigOlympiadShort)
	if err != nil {
		bigOlympiadFilter := database.NewBigOlympiadFilter()
		bigOlympiadFilter.Short = bigOlympiadShort
		olympiads, err := h.BigOlympiadModel.GetBigOlympiads(bigOlympiadFilter)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ans["error"] = "Problem with Database"
			json.NewEncoder(w).Encode(ans)
			return
		}
		if len(olympiads) != 0 {
			bigOlympiadID = int(olympiads[0].ID)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			ans["error"] = "Problem with Database"
			json.NewEncoder(w).Encode(ans)
			return
		}
	}
	filter := database.NewOlympiadFilter()
	filter.BigOlympiadID = int32(bigOlympiadID)
	olympiads, err := h.OlympiadModel.GetOlympiads(filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ans["error"] = "Problem with Database"
		json.NewEncoder(w).Encode(ans)
		return
	}
	// fmt.Println(olympiads)
	w.WriteHeader(http.StatusOK)
	ans["olympiads"] = olympiads
	// fmt.Println(ans)
	json.NewEncoder(w).Encode(ans)
}

func (h *OlympiadHandler) GetUserOlympiads(w http.ResponseWriter, r *http.Request) {
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
		olympiads, err := h.OlympiadUserModel.GetOlympiads(session[0].UserID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ans["error"] = "Problem with Database"
			json.NewEncoder(w).Encode(ans)
			return
		}
		// fmt.Println(olympiads)
		w.WriteHeader(http.StatusOK)
		ans["olympiads"] = olympiads
		json.NewEncoder(w).Encode(ans)
		return
	}
	w.WriteHeader(http.StatusOK)
	ans["ok"] = "ok"
	json.NewEncoder(w).Encode(ans)
}

func (h *OlympiadHandler) AddOlympiadToUser(w http.ResponseWriter, r *http.Request) {
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
		// fmt.Println(session)
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
		eventStr := mux.Vars(r)["olympiad_id"]
		eventID, err := strconv.Atoi(eventStr)
		// fmt.Println(eventStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			ans["error"] = "Wrong format of olympiad"
			json.NewEncoder(w).Encode(ans)
			return
		}
		err = h.OlympiadUserModel.CreateConnection(session[0].UserID, int32(eventID))
		// fmt.Println(err)
		if errors.Is(err, olympiads.ErrConnectionExists) {
			w.WriteHeader(http.StatusBadRequest)
			ans["error"] = "Is already subscribed"
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
