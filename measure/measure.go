package measure

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Measure struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type Handler struct {
	measure []Measure
}

func NewHandler() *Handler {
	mh := &Handler{}
	mh.measure = append(mh.measure, Measure{"1", "Килограмм"})
	mh.measure = append(mh.measure, Measure{"2", "Штука"})
	return mh
}

func (mh *Handler) GetMeasures(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mh.measure)
}
func (mh *Handler) GetMeasure(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range mh.measure {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Measure{})
}
func (mh *Handler) CreateMeasure(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var measure Measure
	_ = json.NewDecoder(r.Body).Decode(&measure)

	// Найти максимальный существующий ID
	maxID := 0
	for _, existingMeasure := range mh.measure {
		measureID, err := strconv.Atoi(existingMeasure.ID)
		if err == nil && measureID > maxID {
			maxID = measureID
		}
	}

	measure.ID = strconv.Itoa(maxID + 1)

	mh.measure = append(mh.measure, measure)

	json.NewEncoder(w).Encode(measure)
}
func (mh *Handler) UpdateMeasure(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var measure Measure
	for index, item := range mh.measure {
		if item.ID == params["id"] {
			_ = json.NewDecoder(r.Body).Decode(&measure)
			mh.measure[index].Name = measure.Name
			break
		}
	}
	json.NewEncoder(w).Encode(mh.measure)
}

func (mh *Handler) DeleteMeasure(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range mh.measure {
		if item.ID == params["id"] {
			mh.measure = append(mh.measure[:index], mh.measure[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(mh.measure)
}
