package handlers

import (
	"net/http"
	"wb_l2/develop/dev11/internal/converter"
	"wb_l2/develop/dev11/internal/event"
)

type Handler struct {
	Cache *event.Cache
}

func (h *Handler) CreateEvent(w http.ResponseWriter, req *http.Request) {
	evt, err := event.UnmarshalRequest(req)
	if err != nil {
		converter.JsonResponse(w, http.StatusNotFound, err.Error(), true)
		return
	}
	err = h.Cache.CreateEventInCache(&evt)
	if err != nil {
		converter.JsonResponse(w, http.StatusNotFound, err.Error(), true)
		return
	}
	converter.JsonResponse(w, http.StatusOK, evt, false)
}

func (h *Handler) UpdateEvent(w http.ResponseWriter, req *http.Request) {
	evt, err := event.UnmarshalRequest(req)
	if err != nil {
		converter.JsonResponse(w, http.StatusNotFound, err.Error(), true)
		return
	}
	updatedEvent, err := h.Cache.UpdateEventInCache(&evt)
	if err != nil {
		converter.JsonResponse(w, http.StatusNotFound, err.Error(), true)
		return
	}
	converter.JsonResponse(w, http.StatusOK, updatedEvent, false)
}

func (h *Handler) DeleteEvent(w http.ResponseWriter, req *http.Request) {
	evt, err := event.UnmarshalRequest(req)
	if err != nil {
		converter.JsonResponse(w, http.StatusNotFound, err.Error(), true)
		return
	}
	err = h.Cache.DeleteEventInCache(evt.EventID)
	if err != nil {
		converter.JsonResponse(w, http.StatusNotFound, err.Error(), true)
		return
	}
	converter.JsonResponse(w, http.StatusOK, "event is deleted", false)
}

func (h *Handler) GetEventsForDay(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	date := query.Get("date")
	events, err := h.Cache.GetEventsDayFromCache(date)
	if err != nil {
		converter.JsonResponse(w, http.StatusNotFound, err.Error(), true)
		return
	}
	converter.JsonResponse(w, http.StatusOK, events, false)
}

func (h *Handler) GetEventsForWeek(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	date := query.Get("date")
	events, err := h.Cache.GetEventsWeekFromCache(date)
	if err != nil {
		converter.JsonResponse(w, http.StatusNotFound, err.Error(), true)
		return
	}
	converter.JsonResponse(w, http.StatusOK, events, false)
}

func (h *Handler) GetEventsForMonth(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	date := query.Get("date")
	events, err := h.Cache.GetEventsMonthFromCache(date)
	if err != nil {
		converter.JsonResponse(w, http.StatusNotFound, err.Error(), true)
		return
	}
	converter.JsonResponse(w, http.StatusOK, events, false)
}
