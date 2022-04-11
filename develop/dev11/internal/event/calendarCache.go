package event

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	eventID int
	mx      sync.RWMutex
	cache   map[int]*Event
}

func NewCache() *Cache {
	return &Cache{
		eventID: 1,
		cache:   make(map[int]*Event),
	}
}

func (c *Cache) CreateEventInCache(event *Event) error {
	if event.Title == "" {
		return errors.New("title is empty")
	}
	c.mx.Lock()
	defer c.mx.Unlock()
	if event.EventID == 0 {
		event.EventID = c.eventID
		c.eventID++
	}
	if _, ok := c.cache[event.EventID]; ok {
		return errors.New("event already exists")
	}
	c.cache[event.EventID] = event
	return nil
}

func (c *Cache) UpdateEventInCache(event *Event) (*Event, error) {
	if event.EventID == 0 {
		return &Event{}, errors.New("event_id is empty or equal to zero")
	}
	var eventFromMap *Event
	var ok bool
	c.mx.Lock()
	defer c.mx.Unlock()
	if eventFromMap, ok = c.cache[event.EventID]; !ok {
		return &Event{}, errors.New("event does not exist")
	} else {
		if event.Body != "" {
			eventFromMap.Body = event.Body
		}
		if event.Title != "" {
			eventFromMap.Title = event.Title
		}
		if !event.Data.IsZero() {
			eventFromMap.Data = event.Data
		}
	}
	return eventFromMap, nil
}

func (c *Cache) DeleteEventInCache(UserID int) error {
	if UserID == 0 {
		return errors.New("event_id is empty or is equal to zero")
	}
	c.mx.Lock()
	defer c.mx.Unlock()
	if _, ok := c.cache[UserID]; !ok {
		return errors.New("event does not exist")
	}
	delete(c.cache, UserID)
	return nil
}

func (c *Cache) GetEventsDayFromCache(date string) ([]Event, error) {
	if date == "" {
		return []Event{}, errors.New("date is empty")
	}
	parsedDate, err := time.Parse(LayoutIOS, date)
	if err != nil {
		return []Event{}, err
	}
	var events []Event
	c.mx.RLock()
	defer c.mx.RUnlock()
	for _, event := range c.cache {
		if event.Data.Equal(parsedDate) {
			events = append(events, *event)
		}
	}
	return events, nil
}

func (c *Cache) GetEventsWeekFromCache(date string) ([]Event, error) {
	if date == "" {
		return []Event{}, errors.New("date is empty")
	}
	parsedDate, err := time.Parse(LayoutIOS, date)
	if err != nil {
		return []Event{}, err
	}
	_, weekParsed := parsedDate.ISOWeek()
	var events []Event
	c.mx.RLock()
	defer c.mx.RUnlock()
	for _, event := range c.cache {
		if _, weekEvent := event.Data.ISOWeek(); weekParsed == weekEvent {
			events = append(events, *event)
		}
	}
	return events, nil
}

func (c *Cache) GetEventsMonthFromCache(date string) ([]Event, error) {
	if date == "" {
		return []Event{}, errors.New("date is empty")
	}
	parsedDate, err := time.Parse(LayoutIOS, date)
	if err != nil {
		return []Event{}, err
	}
	var events []Event
	c.mx.RLock()
	defer c.mx.RUnlock()
	for _, event := range c.cache {
		if event.Data.Month() == parsedDate.Month() {
			events = append(events, *event)
		}
	}
	return events, nil
}
