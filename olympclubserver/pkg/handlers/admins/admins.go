package admins

import (
	"OlympClub/pkg/database/admins"
	"OlympClub/pkg/database/events"
	"OlympClub/pkg/database/holders"
	"OlympClub/pkg/database/news"
	"OlympClub/pkg/database/olympiads"
	"OlympClub/pkg/database/sessions"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type AdminHandler struct {
	HolderTable      *holders.HolderTable
	AdminModel       *admins.AdminTable
	SessionModel     *sessions.SessionModel
	BigOlympiadModel *olympiads.BigOlympiadTable
	OlympiadModel    *olympiads.OlympiadTable
	EventModel       *events.EventTable
	NewsModel        *news.NewsTable
}

func (h *AdminHandler) RegisterHandler(r *mux.Router) {
	admins := r.PathPrefix("/admin").Subrouter()
	admins.HandleFunc("/check", h.CheckAdmin).Methods("GET", "OPTIONS")
	admins.HandleFunc("/holder", h.PostHolder).Methods("POST", "OPTIONS")
	admins.HandleFunc("/big_olympiad", h.PostBigOlympiad).Methods("POST", "OPTIONS")
	admins.HandleFunc("/olympiad", h.PostOlympiad).Methods("POST", "OPTIONS")
	admins.HandleFunc("/event", h.PostEvent).Methods("POST", "OPTIONS")
	admins.HandleFunc("/news", h.PostNews).Methods("POST", "OPTIONS")

	//admins.HandleFunc("/all", h.GetAdmins).Methods("GET", "OPTIONS")
}

func (h *AdminHandler) GetAdmins(w http.ResponseWriter, r *http.Request) {
	ans := make(map[string]interface{})
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		token := r.Header.Get("Authorization")
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
		admins, err := h.AdminModel.GetAdmins()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ans["error"] = "Problem with Database"
			json.NewEncoder(w).Encode(ans)
			return
		}
		w.WriteHeader(http.StatusOK)
		ans["admins"] = admins
		json.NewEncoder(w).Encode(ans)
		return
	}
	w.WriteHeader(http.StatusOK)
	ans["ok"] = "ok"
	json.NewEncoder(w).Encode(ans)
}

func (h *AdminHandler) CheckAdmin(w http.ResponseWriter, r *http.Request) {
	ans := make(map[string]interface{})
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		token := r.Header.Get("Authorization")
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
		check, err := h.AdminModel.CheckAdmin(session[0].UserID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ans["error"] = "Problem with Database"
			json.NewEncoder(w).Encode(ans)
			return
		}
		if !check {
			w.WriteHeader(http.StatusUnauthorized)
			ans["check"] = check
			json.NewEncoder(w).Encode(ans)
			return
		}
		w.WriteHeader(http.StatusOK)
		ans["check"] = check
		json.NewEncoder(w).Encode(ans)
		return
	}
	w.WriteHeader(http.StatusOK)
	ans["ok"] = "ok"
	json.NewEncoder(w).Encode(ans)
}

func (h *AdminHandler) PostHolder(w http.ResponseWriter, r *http.Request) {
	ans := make(map[string]interface{})
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Type")
	w.Header().Set("Content-Type", "multipart/form-data")
	if r.Method == "POST" {
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
		check, err := h.AdminModel.CheckAdmin(session[0].UserID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ans["error"] = "Problem with Database"
			json.NewEncoder(w).Encode(ans)
			return
		}
		if !check {
			w.WriteHeader(http.StatusUnauthorized)
			ans["error"] = "No an admin"
			json.NewEncoder(w).Encode(ans)
			return
		}
		r.ParseMultipartForm(10 << 20) //10 MB
		file, image, err := r.FormFile("holder-logo")
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer file.Close()

		image.Filename = time.Now().String() + image.Filename
		dst, err := os.Create("static/img/" + image.Filename)
		if err != nil {
			log.Println("error creating file", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()
		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		name := r.PostFormValue("holder-name")
		holder := holders.Holder{
			Name: name,
			Logo: image.Filename,
		}
		// fmt.Println(holder)
		h.HolderTable.InsertHolder(holder)
	}
	w.WriteHeader(http.StatusOK)
	ans["ok"] = "ok"
	json.NewEncoder(w).Encode(ans)
}

func (h *AdminHandler) PostBigOlympiad(w http.ResponseWriter, r *http.Request) {
	ans := make(map[string]interface{})
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Type")
	w.Header().Set("Content-Type", "multipart/form-data")
	if r.Method == "POST" {
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
		check, err := h.AdminModel.CheckAdmin(session[0].UserID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ans["error"] = "Problem with Database"
			json.NewEncoder(w).Encode(ans)
			return
		}
		if !check {
			w.WriteHeader(http.StatusUnauthorized)
			ans["error"] = "No an admin"
			json.NewEncoder(w).Encode(ans)
			return
		}
		r.ParseMultipartForm(10 << 20) //10 MB
		file, image, err := r.FormFile("big_olympiad_logo")
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer file.Close()

		image.Filename = time.Now().String() + image.Filename
		dst, err := os.Create("static/img/" + image.Filename)
		if err != nil {
			log.Println("error creating file", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()
		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		name := r.PostFormValue("big_olympiad_name")
		short := r.PostFormValue("short")
		description := r.PostFormValue("description")
		status := r.PostFormValue("status")
		bigOlympiad := olympiads.BigOlympiad{
			Name:        name,
			Short:       short,
			Description: description,
			Status:      status,
			Logo:        image.Filename,
		}
		//fmt.Println(bigOlympiad)
		_, err = h.BigOlympiadModel.CreateBigOlympiad(bigOlympiad)
		if errors.Is(err, olympiads.ErrBigOlympiadIsAlreadyExisted) {
			w.WriteHeader(http.StatusInternalServerError)
			ans["error"] = "Big Olympiad Is Already Existed"
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

func (h *AdminHandler) PostOlympiad(w http.ResponseWriter, r *http.Request) {
	ans := make(map[string]interface{})
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Type")
	w.Header().Set("Content-Type", "multipart/form-data")
	if r.Method == "POST" {
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
		check, err := h.AdminModel.CheckAdmin(session[0].UserID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ans["error"] = "Problem with Database"
			json.NewEncoder(w).Encode(ans)
			return
		}
		if !check {
			w.WriteHeader(http.StatusUnauthorized)
			ans["error"] = "No an admin"
			json.NewEncoder(w).Encode(ans)
			return
		}
		r.ParseMultipartForm(10 << 20) //10 MB
		file, image, err := r.FormFile("olympiad_logo")
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer file.Close()

		image.Filename = time.Now().String() + image.Filename
		dst, err := os.Create("static/img/" + image.Filename)
		if err != nil {
			log.Println("error creating file", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()
		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		name := r.PostFormValue("name")
		subject := r.PostFormValue("subject")
		status := r.PostFormValue("status")
		level := r.PostFormValue("level")
		grade := r.PostFormValue("grade")
		holder := r.PostFormValue("holder")
		holderId, err := strconv.Atoi(holder)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			ans["error"] = "Wrong format of holder"
			json.NewEncoder(w).Encode(ans)
			return
		}
		bigOlympiad := r.PostFormValue("big_olympiad")
		bigOlympiadId, err := strconv.Atoi(bigOlympiad)
		// fmt.Println("ok")

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			ans["error"] = "Wrong format of big olympiad"
			json.NewEncoder(w).Encode(ans)
			return
		}
		short := r.PostFormValue("short")
		website := r.PostFormValue("website")

		olympiad := olympiads.Olympiad{
			Name:          name,
			Subject:       subject,
			Status:        status,
			Level:         level,
			Grade:         grade,
			BigOlympiadID: int32(bigOlympiadId),
			Short:         short,
			Website:       website,
			HolderID:      int32(holderId),
			Img:           image.Filename,
		}
		// fmt.Println(bigOlympiad)
		_, err = h.OlympiadModel.CreateOlympiad(olympiad)
		if errors.Is(err, olympiads.ErrOlympiadIsAlreadyExisted) {
			w.WriteHeader(http.StatusInternalServerError)
			ans["error"] = "Olympiad Is Already Existed"
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

func (h *AdminHandler) PostEvent(w http.ResponseWriter, r *http.Request) {
	ans := make(map[string]interface{})
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Type")
	w.Header().Set("Content-Type", "multipart/form-data")
	if r.Method == "POST" {
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
		check, err := h.AdminModel.CheckAdmin(session[0].UserID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ans["error"] = "Problem with Database"
			json.NewEncoder(w).Encode(ans)
			return
		}
		if !check {
			w.WriteHeader(http.StatusUnauthorized)
			ans["error"] = "No an admin"
			json.NewEncoder(w).Encode(ans)
			return
		}
		r.ParseMultipartForm(10 << 20) //10 MB
		file, image, err := r.FormFile("event_logo")
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer file.Close()

		image.Filename = time.Now().String() + image.Filename
		dst, err := os.Create("static/img/" + image.Filename)
		if err != nil {
			log.Println("error creating file", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()
		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		name := r.PostFormValue("name")
		status := r.PostFormValue("status")
		holder := r.PostFormValue("holder")
		holderId, err := strconv.Atoi(holder)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			ans["error"] = "Wrong format of holder"
			json.NewEncoder(w).Encode(ans)
			return
		}
		// fmt.Println("ok")

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			ans["error"] = "Wrong format of big olympiad"
			json.NewEncoder(w).Encode(ans)
			return
		}
		short := r.PostFormValue("short")
		website := r.PostFormValue("website")
		description := r.PostFormValue("description")

		event := events.Event{
			Name:        name,
			Status:      status,
			Short:       short,
			Website:     website,
			HolderID:    int32(holderId),
			Description: description,
			Img:         image.Filename,
		}
		// fmt.Println(event)
		_, err = h.EventModel.CreateEvent(event)
		if errors.Is(err, olympiads.ErrOlympiadIsAlreadyExisted) {
			w.WriteHeader(http.StatusInternalServerError)
			ans["error"] = "Olympiad Is Already Existed"
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

func (h *AdminHandler) PostNews(w http.ResponseWriter, r *http.Request) {
	ans := make(map[string]interface{})
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Type")
	w.Header().Set("Content-Type", "multipart/form-data")
	if r.Method == "POST" {
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
		check, err := h.AdminModel.CheckAdmin(session[0].UserID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ans["error"] = "Problem with Database"
			json.NewEncoder(w).Encode(ans)
			return
		}
		if !check {
			w.WriteHeader(http.StatusUnauthorized)
			ans["error"] = "No an admin"
			json.NewEncoder(w).Encode(ans)
			return
		}
		r.ParseMultipartForm(10 << 20) //10 MB
		title := r.PostFormValue("title")
		description := r.PostFormValue("description")
		Table := r.PostFormValue("table")
		keyStr := r.PostFormValue("key")
		key, err := strconv.Atoi(keyStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			ans["error"] = "Wrong format of key"
			json.NewEncoder(w).Encode(ans)
			return
		}

		news := news.News{
			Title:       title,
			Table:       Table,
			Key:         int32(key),
			Description: description,
		}
		// fmt.Println(news)
		err = h.NewsModel.InsertNews(news)
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
