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

// Хэндлер чтобы работать с олимпиадами
type OlympiadHandler struct {
	OlympiadModel     *database.OlympiadTable
	BigOlympiadModel  *database.BigOlympiadTable
	SessionModel      *sessions.SessionModel
	OlympiadUserModel *database.OlympiadUserTable
	NewsModel         *news.NewsTable
}

func (h *OlympiadHandler) RegisterHandler(r *mux.Router) {
	// регистрируем хэнлеры
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

// Создаём олимпиаду
func (h *OlympiadHandler) CreateOlympiad(w http.ResponseWriter, r *http.Request) {
	// Парсим форму
	r.ParseForm()
	// Выставляем "Content-Type
	w.Header().Set("Content-Type", "application/json")
	ans := make(map[string]interface{})

	// Получаем нужные данные из формы (input[name="name"])
	name := r.Form.Get("name")
	subject := r.Form.Get("subject")
	level := r.Form.Get("level")
	img := r.Form.Get("img")
	short := r.Form.Get("short")
	status := r.Form.Get("status")
	grade := r.Form.Get("grade")

	// Переводим строку в число
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

	// Создаём новую олимпиаду
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
	// кодируем ответ
	json.NewEncoder(w).Encode(ans)
	// выставляем ответ
	w.WriteHeader(http.StatusOK)
}

// Получаем олимпиады
func (h *OlympiadHandler) GetOlympiads(w http.ResponseWriter, r *http.Request) {
	// парсим форму
	r.ParseForm()
	// получаем значения из query blahblah.com/?subject=123
	subject := r.URL.Query().Get("subject")
	level := r.URL.Query().Get("level")
	grade := r.URL.Query().Get("grade")

	// Создаём новый фильтр
	filter := database.NewOlympiadFilter()
	filter.Subject = subject
	filter.Level = level
	filter.Grade = grade

	// Получаем олимпиады с помощью фильтра
	olympiads, err := h.OlympiadModel.GetOlympiads(filter)

	ans := make(map[string]interface{})

	// Выставляем некоторые переменные
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
	// выставляем флаг и кодируем ответ
	w.WriteHeader(http.StatusOK)
	ans["olympiads"] = olympiads
	json.NewEncoder(w).Encode(ans)
}

// Получение больших олимпиад
func (h *OlympiadHandler) GetBigOlympiads(w http.ResponseWriter, r *http.Request) {
	// Парсим форму
	r.ParseForm()
	olympiads, err := h.BigOlympiadModel.GetBigOlympiads(olympiads.NewBigOlympiadFilter())
	ans := make(map[string]interface{})

	// Выставляем переменные
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
	// выставляем флаг и кодируем ответ
	w.WriteHeader(http.StatusOK)
	ans["olympiads"] = olympiads
	json.NewEncoder(w).Encode(ans)
}

// Получение информации о большой информации
func (h *OlympiadHandler) GetBigOlympiad(w http.ResponseWriter, r *http.Request) {
	// Выставляем нужные переменные
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	ans := make(map[string]interface{})

	// Берём переменную из роутера /olympiad/{big_olympiad_id}
	bigOlympiad := mux.Vars(r)["big_olympiad_id"]
	// фильтр
	filter := database.NewBigOlympiadFilter()
	bigOlympiadID, err := strconv.Atoi(bigOlympiad)
	// костыль чтобы можно было по номеру и по короткому имени находить олимпиаду
	if err != nil {
		filter.Short = bigOlympiad
	} else {
		filter.ID = int32(bigOlympiadID)
	}
	// Ищем олимпиады
	bigOlympiads, err := h.BigOlympiadModel.GetBigOlympiads(filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ans["error"] = "Problem with Database"
		json.NewEncoder(w).Encode(ans)
		return
	}
	if len(bigOlympiads) == 0 {
		// выставляем флаг и кодируем ответ
		w.WriteHeader(http.StatusNotFound)
		ans["error"] = "Olympiad not found"
		json.NewEncoder(w).Encode(ans)
		return
	}
	// выставляем флаг и кодируем ответ
	w.WriteHeader(http.StatusOK)
	ans["big_olympiad"] = bigOlympiads[0]
	json.NewEncoder(w).Encode(ans)
}

// Получение информации об олимпиаде
func (h *OlympiadHandler) GetOlympiadById(w http.ResponseWriter, r *http.Request) {
	// парсим форму
	r.ParseForm()
	// выставляем переменные
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	ans := make(map[string]interface{})

	// Берём переменную из роутера
	olympiad := mux.Vars(r)["olympiad_id"]
	olympiadID, err := strconv.Atoi(olympiad)
	olympiadShort := ""
	if err != nil {
		olympiadShort = olympiad
		olympiadID = 0
	}

	// Берём переменную из роутера
	bigOlympiad := mux.Vars(r)["big_olympiad_id"]
	bigOlympiadID, err := strconv.Atoi(bigOlympiad)
	if err != nil {
		// если указан только короткое имя
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
	// делаем фильтр
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

	// выставляем флаг и кодируем ответ
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

// получаем информацию о новостях об олимпиаде
func (h *OlympiadHandler) GetOlympiadNewsById(w http.ResponseWriter, r *http.Request) {
	// парсим
	r.ParseForm()
	// выставляем переменные
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	ans := make(map[string]interface{})

	// Берём переменную из роутера
	olympiad := mux.Vars(r)["olympiad_id"]
	olympiadID, err := strconv.Atoi(olympiad)
	olympiadShort := ""
	if err != nil {
		olympiadShort = olympiad
		olympiadID = 0
	}
	// Берём переменную из роутера
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
	// делаем фильтр
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
	if len(olympiads) != 0 {
		olympiad := olympiads[0]

		// фильтер для новостей
		newsFilter := news.NewsFilter{
			ID:    -1,
			Table: "Olympiads",
			Key:   olympiad.ID,
		}
		news, err := h.NewsModel.GetNews(newsFilter)
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

// получить олимпиады с помощью большой олимпиады
func (h *OlympiadHandler) GetOlympiadsByBigOlympiadId(w http.ResponseWriter, r *http.Request) {
	// парсим
	r.ParseForm()
	// выставляем переменные
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	ans := make(map[string]interface{})

	// Берём переменную из роутера
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
	// делаем фильтр
	filter := database.NewOlympiadFilter()
	filter.BigOlympiadID = int32(bigOlympiadID)
	olympiads, err := h.OlympiadModel.GetOlympiads(filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ans["error"] = "Problem with Database"
		json.NewEncoder(w).Encode(ans)
		return
	}
	// выставляем флаг и кодируем ответ
	w.WriteHeader(http.StatusOK)
	ans["olympiads"] = olympiads
	json.NewEncoder(w).Encode(ans)
}

// получить олимпиады пользователя
func (h *OlympiadHandler) GetUserOlympiads(w http.ResponseWriter, r *http.Request) {
	// выставляем нужные переменные
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")
	ans := make(map[string]interface{})

	if r.Method == "GET" {
		// берём токен пользователя
		token := r.Header.Get("Authorization")
		if len(token) < 7 {
			w.WriteHeader(http.StatusBadRequest)
			ans["error"] = "Wrong format of token"
			json.NewEncoder(w).Encode(ans)
			return
		}
		// получаем сессию
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
		// получаем олимпиаду с помощью user id
		olympiads, err := h.OlympiadUserModel.GetOlympiads(session[0].UserID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ans["error"] = "Problem with Database"
			json.NewEncoder(w).Encode(ans)
			return
		}
		w.WriteHeader(http.StatusOK)
		ans["olympiads"] = olympiads
		json.NewEncoder(w).Encode(ans)
		return
	}
	w.WriteHeader(http.StatusOK)
	ans["ok"] = "ok"
	json.NewEncoder(w).Encode(ans)
}

// добавить олимпиаду пользователю
func (h *OlympiadHandler) AddOlympiadToUser(w http.ResponseWriter, r *http.Request) {
	// выставляем переменные
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")
	ans := make(map[string]interface{})
	if r.Method == "GET" {
		// получаем токен
		token := r.Header.Get("Authorization")
		if len(token) < 7 {
			w.WriteHeader(http.StatusBadRequest)
			ans["error"] = "Wrong format of token"
			json.NewEncoder(w).Encode(ans)
			return
		}
		// получаем сессию
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
		// получаем олимпиаду из переменной роутера
		olympiadStr := mux.Vars(r)["olympiad_id"]
		olympiadID, err := strconv.Atoi(olympiadStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			ans["error"] = "Wrong format of olympiad"
			json.NewEncoder(w).Encode(ans)
			return
		}
		// добавить связь между пользователем и олимпиадой
		err = h.OlympiadUserModel.CreateConnection(session[0].UserID, int32(olympiadID))
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
