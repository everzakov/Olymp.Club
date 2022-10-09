package holders

import (
	"OlympClub/pkg/database/holders"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type HolderHandler struct {
	HolderTable *holders.HolderTable
}

func (h *HolderHandler) RegisterHandler(r *mux.Router) {
	r.HandleFunc("/holders", h.GetHolders).Methods("GET", "OPTIONS")
	r.HandleFunc("/holder/{holder_id}", h.GetHolder).Methods("GET", "OPTIONS")
}

func (h *HolderHandler) GetHolder(w http.ResponseWriter, r *http.Request) {
	ans := make(map[string]interface{})
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	holderID, err := strconv.Atoi(mux.Vars(r)["holder_id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ans["error"] = "Problem with Database"
		json.NewEncoder(w).Encode(ans)
		return
	}
	holders, err := h.HolderTable.GetHolders(int32(holderID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ans["error"] = "Problem with Database"
		json.NewEncoder(w).Encode(ans)
		return
	}
	if len(holders) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		ans["error"] = "No such a holder"
	}
	w.WriteHeader(http.StatusOK)
	ans["holder"] = holders[0]
	json.NewEncoder(w).Encode(ans)
}

func (h *HolderHandler) GetHolders(w http.ResponseWriter, r *http.Request) {
	ans := make(map[string]interface{})
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	holders, err := h.HolderTable.GetAllHolders()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ans["error"] = "Problem with Database"
		json.NewEncoder(w).Encode(ans)
		return
	}
	if len(holders) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		ans["error"] = "No such a holder"
	}
	w.WriteHeader(http.StatusOK)
	ans["holders"] = holders
	json.NewEncoder(w).Encode(ans)
}
