package marker

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// Represents a marker struct with recorded moments
type Marker struct {
	Moments []Moment
	lock    *sync.Mutex
}

// Represents a recorded moment in time with a name, start and end time
type Moment struct {
	Name  string
	Start time.Time
	End   time.Time
}

// Creates a new moment with the given name and the current time as the start time
func (m *Marker) Mark(name string) {
	if name == "" {
		name = fmt.Sprintf("Event %d", len(m.Moments)+1)
	}

	m.lock.Lock()
	m.Moments = append(m.Moments, Moment{
		Name:  name,
		Start: time.Now(),
	})
	m.lock.Unlock()
}

// Marks the end time of the most recent moment as the current time
func (m *Marker) Done() error {
	if len(m.Moments) == 0 {
		return fmt.Errorf("no moments marked")
	}

	m.lock.Lock()
	m.Moments[len(m.Moments)-1].End = time.Now()
	m.lock.Unlock()

	return nil
}

// Returns the duration of the moment
func (m *Moment) Elapsed() (time.Duration, error) {
	if m.End.IsZero() {
		return time.Duration(0), fmt.Errorf("no end time recorded")
	}
	return m.End.Sub(m.Start), nil
}

// Returns a formatted string representation of the marker's moments and their durations
func (m *Marker) String() (string, error) {
	result := strings.Builder{}
	for _, moment := range m.Moments {
		elapsed, err := moment.Elapsed()
		if err != nil {
			return "", err
		}

		_, err = result.WriteString(fmt.Sprintf("%s: %s\n", moment.Name, elapsed))
		if err != nil {
			return "", err
		}
	}
	return result.String(), nil
}

// Creates a new Marker instance
func New() *Marker {
	return &Marker{
		Moments: make([]Moment, 0),
		lock:    &sync.Mutex{},
	}
}
