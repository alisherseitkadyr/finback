package handler

import (
	"encoding/json"
	"finback/internal/content/usecase"
	"net/http"
	"strings"
)

type HTTPHandler struct {
	service *usecase.Service
}

func NewHTTPHandler(service *usecase.Service) *HTTPHandler {
	return &HTTPHandler{service: service}
}

func (h *HTTPHandler) ListTopics(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{"items": h.service.ListTopics()})
}

func (h *HTTPHandler) GetTopic(w http.ResponseWriter, r *http.Request) {
	topicID := strings.TrimPrefix(r.URL.Path, "/topics/")
	item, err := h.service.GetTopic(topicID)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func (h *HTTPHandler) GetLesson(w http.ResponseWriter, r *http.Request) {
	lessonID := strings.TrimPrefix(r.URL.Path, "/lessons/")
	lesson, err := h.service.GetLesson(lessonID)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"lesson": lesson})
}

func (h *HTTPHandler) SubmitAssessment(w http.ResponseWriter, r *http.Request) {
	var req usecase.SubmitAssessmentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	result, err := h.service.SubmitAssessment(req)
	if err != nil {
		status := http.StatusBadRequest
		if err.Error() == "user not found" {
			status = http.StatusNotFound
		}
		writeError(w, status, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, result)
}

func (h *HTTPHandler) GetRecommendations(w http.ResponseWriter, r *http.Request) {
	result, err := h.service.GetRecommendations(r.URL.Query().Get("user_id"))
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, result)
}

func (h *HTTPHandler) CompleteLesson(w http.ResponseWriter, r *http.Request) {
	var req usecase.CompleteLessonRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	result, err := h.service.CompleteLesson(req)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, result)
}

func (h *HTTPHandler) GetProgress(w http.ResponseWriter, r *http.Request) {
	result, err := h.service.GetProgress(r.URL.Query().Get("user_id"))
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, result)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
