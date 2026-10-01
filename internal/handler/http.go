package handler

import (
	"encoding/json"
	"fizz-buzz/internal/db"
	"fizz-buzz/pkg/fizzbuzz"
	"log/slog"
	"net/http"
	"strconv"
)

type HTTPHandler struct {
	database *db.DB
}

func NewHTTPHandler(database *db.DB) *HTTPHandler {
	return &HTTPHandler{database: database}
}

func (h *HTTPHandler) HandleFizzBuzz(w http.ResponseWriter, r *http.Request) {
	var req fizzbuzz.Request

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	result, err := fizzbuzz.Compute(req)
	if err != nil {
		writeJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.trackRequestForStats(r, req)

	writeJSON(w, http.StatusOK, map[string]any{"data": result})
}

func (h *HTTPHandler) HandleFizzBuzzQuery(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	int1, err1 := strconv.Atoi(q.Get("int1"))
	int2, err2 := strconv.Atoi(q.Get("int2"))
	limit, err3 := strconv.Atoi(q.Get("limit"))
	str1 := q.Get("str1")
	str2 := q.Get("str2")

	if err1 != nil || err2 != nil || err3 != nil || str1 == "" || str2 == "" {
		writeJSONError(w, "Missing or invalid query parameters", http.StatusBadRequest)
		return
	}

	req := fizzbuzz.Request{
		Int1:  int1,
		Int2:  int2,
		Limit: limit,
		Str1:  str1,
		Str2:  str2,
	}

	result, err := fizzbuzz.Compute(req)
	if err != nil {
		writeJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.trackRequestForStats(r, req)

	writeJSON(w, http.StatusOK, map[string]any{"data": result})
}

func (h *HTTPHandler) HandleStats(w http.ResponseWriter, r *http.Request) {
	stat, err := h.database.GetMostFrequent(r.Context())

	if err != nil {
		slog.Error("failed to query stats", "error", err)
		writeJSONError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if stat == nil {
		writeJSONError(w, "No statistics recorded yet", http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{"data": stat})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func writeJSONError(w http.ResponseWriter, msg string, status int) {
	writeJSON(w, status, map[string]string{"error": msg})
}

func (h *HTTPHandler) trackRequestForStats(r *http.Request, req fizzbuzz.Request) {
	if h.database != nil {
		if err := h.database.TrackRequest(r.Context(), req); err != nil {
			slog.Error("failed to record stats", "error", err)
		}
	}
}
