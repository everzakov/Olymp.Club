package auth

import (
	"OlympClub/pkg/database/sessions"
	"OlympClub/pkg/database/unconfirmed"
	user_database "OlympClub/pkg/database/users"
	"OlympClub/pkg/types"
	"bytes"
	"encoding/json"
	"errors"
	"html/template"
	"net/http"
	"net/smtp"

	"github.com/gorilla/mux"
)

var (
	mailTmpl = `
	<html>
	<body>
	<div style="width: 400px; min-height: 500px; border: black 1px solid; border-radius: 10px; padding: 10px 10px 50px"> 
    <h1 style="text-align: center;">Регистрация на платформе Olymp.Club</h1>
    <p>Недавно вы регистрировались на платформе Olymp.Club</p>
    <p style="text-align: center">Чтобы <b>подтвердить</b> регистрацию,<br>нажмите на кнопку:</p>
    <a href="{{ .URLApprove }}"
       style="background-color: #0066FF; color: white; padding: 10px; border-radius: 10px; text-decoration: none; margin: auto; display: block; width: 180px; text-align: center">Подтвердить
        регистрацию</a>
    <p style="text-align: center">Или <b>перейдите</b> по ссылке:</p>
    <a href="{{ .URLApprove }}" style="word-break: break-word; text-align: center">{{ .URLApprove }}</a>
    <p style="text-align: center">Если это были не вы,<br>то <b>нажмите</b> на кнопку</p>
    <a href="{{ .URLDecline }}"
       style="background-color: #0066FF; color: white; padding: 10px; border-radius: 10px; text-decoration: none; margin: auto; display: block; width: 100px; text-align: center">Это
        был не я</a>
    <p style="text-align: center">Или <b>перейдите</b> по ссылке:</p>
    <a href="{{ .URLDecline }}" style="word-break: break-word; text-align: center">{{ .URLDecline }}</a>
	</div>
	</body>
	</html>
`
	changePasswordTmpl = `
	<div style="width: 400px; min-height: 500px; border: black 1px solid; border-radius: 10px; padding: 10px 10px 50px">
    <h1 style="text-align: center;">Смена пароля на платформе Olymp.Club</h1>
    <p>Недавно вы запросили смену пароля на платформе Olymp.Club</p>
    <p style="text-align: center">Чтобы <b>сменить пароль</b>,<br>нажмите на кнопку:</p>
    <a href="{{ .URLChange }}"
       style="background-color: #0066FF; color: white; padding: 10px; border-radius: 10px; text-decoration: none; margin: auto; display: block; width: 180px; text-align: center">Сменить
        пароль</a>
    <p style="text-align: center">Или <b>перейдите</b> по ссылке:</p>
    <a href="{{ .URLChange }}" style="word-break: break-word; text-align: center">{{ .URLChange }}</a>
    <p style="text-align: center">Если это были не вы,<br>то <b>проигнорируйте</b></p>
</div>
	`
)

type RegisterMail struct {
	URLApprove string
	URLDecline string
}

type ChangePasswordMail struct {
	URLChange string
}

type ChangePasswordRequestForm struct {
	Email string `json:"email"`
}

type ChangePassworForm struct {
	PassHash string `json:"pass_hash"`
}

type AuthForm struct {
	Email    string `json:"email"`
	PassHash string `json:"pass_hash"`
}

type RegisterForm struct {
	Email    string `json:"email"`
	PassHash string `json:"pass_hash"`
}

type AuthHandler struct {
	UnConfirmedUsersTable *unconfirmed.UnConfirmedUsersTable
	UsersModel            *user_database.UserModel
	SessionModel          *sessions.SessionModel
	MailInfo              types.MailInfo
}

func (h *AuthHandler) RegisterHandler(r *mux.Router) {
	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/", h.GetAuthForm).Methods("GET", "OPTIONS")
	auth.HandleFunc("/post", h.PostAuthForm).Methods("POST", "OPTIONS")

	register := r.PathPrefix("/register").Subrouter()
	register.HandleFunc("/", h.GetRegisterForm).Methods("GET", "OPTIONS")
	register.HandleFunc("/post", h.PostRegisterForm).Methods("POST", "OPTIONS")
	register.HandleFunc("/verify", h.VerifyUser).Methods("GET", "OPTIONS")
	register.HandleFunc("/decline", h.DeclineUser).Methods("GET", "OPTIONS")

	change := r.PathPrefix("/password").Subrouter()
	change.HandleFunc("/request", h.GetChangePasswordRequestForm).Methods("GET", "OPTIONS")
	change.HandleFunc("/request/post", h.PostChangePasswordRequestForm).Methods("POST", "OPTIONS")
	change.HandleFunc("/change", h.GetChangePasswordRequestForm).Methods("GET", "OPTIONS")
	change.HandleFunc("/change/post", h.PostChangePasswordForm).Methods("POST", "OPTIONS")
}

func (h *AuthHandler) GetAuthForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	ans := make(map[string]interface{})
	w.WriteHeader(http.StatusOK)
	ans["ok"] = "ok"
	json.NewEncoder(w).Encode(ans)
}

func (h *AuthHandler) GetRegisterForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	ans := make(map[string]interface{})
	w.WriteHeader(http.StatusOK)
	ans["ok"] = "ok"
	json.NewEncoder(w).Encode(ans)
}

func (h *AuthHandler) GetChangePasswordRequestForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	ans := make(map[string]interface{})
	w.WriteHeader(http.StatusOK)
	ans["ok"] = "ok"
	json.NewEncoder(w).Encode(ans)
}

func (h *AuthHandler) GetChangePasswordForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	ans := make(map[string]interface{})
	w.WriteHeader(http.StatusOK)
	ans["ok"] = "ok"
	json.NewEncoder(w).Encode(ans)
}

func (h *AuthHandler) PostAuthForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		var f AuthForm
		err := json.NewDecoder(r.Body).Decode(&f)
		if err != nil {
			ans := make(map[string]interface{})
			w.WriteHeader(http.StatusBadRequest)
			ans["error"] = "Problem with request"
			json.NewEncoder(w).Encode(ans)
			return
		}
		users, err := h.UsersModel.GetUsersByEmailAndPassword(f.Email, f.PassHash)
		if errors.Is(err, user_database.ErrUserDoesntExists) || len(users) == 0 {
			ans := make(map[string]interface{})
			w.WriteHeader(http.StatusBadRequest)
			ans["error"] = "User doesn't exist"
			json.NewEncoder(w).Encode(ans)
			return
		}
		if err != nil {
			ans := make(map[string]interface{})
			w.WriteHeader(http.StatusInternalServerError)
			ans["error"] = "Problem with Database"
			json.NewEncoder(w).Encode(ans)
			return
		}
		session, err := h.SessionModel.CreateSession(users[0].ID)
		if err != nil {
			ans := make(map[string]interface{})
			w.WriteHeader(http.StatusInternalServerError)
			ans["error"] = "Problem with Database"
			json.NewEncoder(w).Encode(ans)
			return
		}
		w.WriteHeader(http.StatusOK)
		ans := make(map[string]interface{})
		ans["token"] = session.Token
		json.NewEncoder(w).Encode(ans)
		return
	}
	ans := make(map[string]interface{})
	w.WriteHeader(http.StatusOK)
	ans["ok"] = "ok"
	json.NewEncoder(w).Encode(ans)
}

func (h *AuthHandler) PostRegisterForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		var f RegisterForm
		err := json.NewDecoder(r.Body).Decode(&f)
		if err != nil {
			ans := make(map[string]interface{})
			w.WriteHeader(http.StatusBadRequest)
			ans["error"] = "Problem with request"
			json.NewEncoder(w).Encode(ans)
			return
		}

		// Use the r.PostForm.Get() method to retrieve the relevant data fields
		// from the r.PostForm map.
		check, err := h.UnConfirmedUsersTable.IsUserConfirmed(f.Email)
		if err != nil {
			ans := make(map[string]interface{})
			w.WriteHeader(http.StatusInternalServerError)
			ans["error"] = "Problem with Database"
			json.NewEncoder(w).Encode(ans)
			return
		}
		if check {
			ans := make(map[string]interface{})
			w.WriteHeader(http.StatusBadRequest)
			ans["error"] = "User exists"
			json.NewEncoder(w).Encode(ans)
			return
		}
		user, err := h.UnConfirmedUsersTable.CreateUnconfirmedUser(f.Email, f.PassHash)
		if err != nil {
			ans := make(map[string]interface{})
			w.WriteHeader(http.StatusInternalServerError)
			ans["error"] = "Problem with Database"
			json.NewEncoder(w).Encode(ans)
			return
		}
		auth := smtp.PlainAuth("", h.MailInfo.User, h.MailInfo.Password, h.MailInfo.Host)
		to := []string{
			f.Email,
		}

		mail := RegisterMail{
			URLApprove: h.MailInfo.FrontURL + "/register/verify?token1=" + user.Token1 + "&token2=" + user.Token2,
			URLDecline: h.MailInfo.FrontURL + "/register/decline?token1=" + user.Token1 + "&token2=" + user.Token2,
		}
		tmpl := template.New("mail")
		if tmpl, err = tmpl.Parse(mailTmpl); err != nil {
			// fmt.Println(err)
		}
		var buf bytes.Buffer
		tmpl.Execute(&buf, mail)

		subject := "Subject: Test email from Go!\n"
		mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

		msg := []byte(subject + mime + buf.String())
		_ = smtp.SendMail(h.MailInfo.Addr, auth, h.MailInfo.From, to, msg)

		// fmt.Println("email: ", f.Email, f.PassHash)
	}

	ans := make(map[string]interface{})
	w.WriteHeader(http.StatusOK)
	ans["ok"] = "ok"
	json.NewEncoder(w).Encode(ans)
}

func (h *AuthHandler) VerifyUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()
	token1 := r.URL.Query().Get("token1")
	token2 := r.URL.Query().Get("token2")
	// fmt.Println(token1, token2)
	if r.Method == "GET" {
		uc_users, err := h.UnConfirmedUsersTable.GetUsersByTokens(token1, token2)
		if err != nil {
			ans := make(map[string]interface{})
			w.WriteHeader(http.StatusInternalServerError)
			ans["error"] = "Problem with Database"
			json.NewEncoder(w).Encode(ans)
			return
		}
		if len(uc_users) == 0 {
			ans := make(map[string]interface{})
			w.WriteHeader(http.StatusBadRequest)
			ans["error"] = "User doesn't exist"
			json.NewEncoder(w).Encode(ans)
			return
		}
		check, err := h.UnConfirmedUsersTable.IsUserConfirmed(uc_users[0].Email)
		if check {
			ans := make(map[string]interface{})
			w.WriteHeader(http.StatusBadRequest)
			ans["error"] = "User is already confirmed"
			json.NewEncoder(w).Encode(ans)
			return
		}
		if err != nil {
			ans := make(map[string]interface{})
			w.WriteHeader(http.StatusInternalServerError)
			ans["error"] = "Problem with Database"
			json.NewEncoder(w).Encode(ans)
			return
		}
		_, err = h.UnConfirmedUsersTable.ConfirmUser(token1, token2)
		// fmt.Println(err)
		if err != nil {
			ans := make(map[string]interface{})
			w.WriteHeader(http.StatusBadRequest)
			ans["error"] = "Problem with Database"
			json.NewEncoder(w).Encode(ans)
			return
		}
	}
	ans := make(map[string]interface{})
	w.WriteHeader(http.StatusOK)
	ans["ok"] = "ok"
	json.NewEncoder(w).Encode(ans)
}

func (h *AuthHandler) DeclineUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		r.ParseForm()
		token1 := r.URL.Query().Get("token1")
		token2 := r.URL.Query().Get("token2")
		err := h.UnConfirmedUsersTable.DeleteUser(token1, token2)
		if errors.Is(err, unconfirmed.ErrUserDoesntExists) {
			ans := make(map[string]interface{})
			w.WriteHeader(http.StatusBadRequest)
			ans["error"] = "No such user"
			json.NewEncoder(w).Encode(ans)
			return
		}
		if errors.Is(err, unconfirmed.ErrUserIsAlreadyConfirmed) {
			ans := make(map[string]interface{})
			w.WriteHeader(http.StatusBadRequest)
			ans["error"] = "User is already confirmed"
			json.NewEncoder(w).Encode(ans)
			return
		}
		if err != nil {
			ans := make(map[string]interface{})
			w.WriteHeader(http.StatusInternalServerError)
			ans["error"] = "Problem with Database"
			json.NewEncoder(w).Encode(ans)
			return
		}
	}
	ans := make(map[string]interface{})
	w.WriteHeader(http.StatusOK)
	ans["ok"] = "ok"
	json.NewEncoder(w).Encode(ans)
}

func (h *AuthHandler) PostChangePasswordRequestForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		var f ChangePasswordRequestForm
		err := json.NewDecoder(r.Body).Decode(&f)
		email := f.Email
		users, err := h.UsersModel.GetUsersByEmail(email)
		// fmt.Println(err)
		if errors.Is(err, unconfirmed.ErrUserDoesntExists) {
			ans := make(map[string]interface{})
			w.WriteHeader(http.StatusInternalServerError)
			ans["error"] = "No such user"
			json.NewEncoder(w).Encode(ans)
			return
		}
		if err != nil {
			ans := make(map[string]interface{})
			w.WriteHeader(http.StatusInternalServerError)
			ans["error"] = "Problem with Database"
			json.NewEncoder(w).Encode(ans)
			return
		}
		auth := smtp.PlainAuth("", h.MailInfo.User, h.MailInfo.Password, h.MailInfo.Host)
		to := []string{
			email,
		}

		mail := ChangePasswordMail{
			URLChange: h.MailInfo.FrontURL + "/password/change?token1=" + users[0].Token1 + "&token2=" + users[0].Token2,
		}
		tmpl := template.New("mail")
		if tmpl, err = tmpl.Parse(changePasswordTmpl); err != nil {
			// fmt.Println(err)
		}
		var buf bytes.Buffer
		tmpl.Execute(&buf, mail)

		subject := "Subject: Test email from Go!\n"
		mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

		msg := []byte(subject + mime + buf.String())
		_ = smtp.SendMail(h.MailInfo.Addr, auth, h.MailInfo.From, to, msg)
	}
	ans := make(map[string]interface{})
	w.WriteHeader(http.StatusOK)
	ans["ok"] = "ok"
	json.NewEncoder(w).Encode(ans)
}

func (h *AuthHandler) PostChangePasswordForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		var f ChangePassworForm
		err := json.NewDecoder(r.Body).Decode(&f)
		if err != nil {
			ans := make(map[string]interface{})
			w.WriteHeader(http.StatusBadRequest)
			ans["error"] = "Wrong format of password"
			json.NewEncoder(w).Encode(ans)
			return
		}
		new_password := f.PassHash
		r.ParseForm()
		token1 := r.URL.Query().Get("token1")
		token2 := r.URL.Query().Get("token2")
		users, err := h.UsersModel.GetUsersByTokens(token1, token2)
		if errors.Is(err, user_database.ErrUserDoesntExists) || len(users) == 0 {
			ans := make(map[string]interface{})
			w.WriteHeader(http.StatusBadRequest)
			ans["error"] = "Wrong tokens"
			json.NewEncoder(w).Encode(ans)
			return
		}
		if err != nil {
			ans := make(map[string]interface{})
			w.WriteHeader(http.StatusInternalServerError)
			ans["error"] = "Problem with Database"
			json.NewEncoder(w).Encode(ans)
			return
		}
		err = h.UsersModel.UpdatePassword(token1, token2, new_password)
		if err != nil {
			ans := make(map[string]interface{})
			w.WriteHeader(http.StatusInternalServerError)
			ans["error"] = "Problem with Database"
			json.NewEncoder(w).Encode(ans)
			return
		}
	}
	ans := make(map[string]interface{})
	w.WriteHeader(http.StatusOK)
	ans["ok"] = "ok"
	json.NewEncoder(w).Encode(ans)
}
