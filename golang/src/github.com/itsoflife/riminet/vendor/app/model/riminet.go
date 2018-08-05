package model

// Riminet is a global structure that holds the network and its state info.
type Riminet struct {
	// Riminet id, for future use
	Id id_t `json:"id"`

	// stores the current epoch
	Epoch epoch_t

	// all fCells
	FCells    []FCell
	MaxFcId   id_t          // stores max id used
	FreeFcIds map[id_t]fill1_t // unused ids <= maxFcId

	SCells    map[scell_t][]SCell  // all sCells categorized by type
	MaxScId   map[scell_t]id_t         // stores max id used
	FreeScIds map[scell_t]map[id_t]fill1_t // unused ids <= maxScId
}


func MakeRiminet() Riminet {
	rn := Riminet{
		Id:        0,
		Epoch:     0,
		FCells:    make([]FCell, ARR_SIZE_F),
		MaxFcId:   0,
		FreeFcIds: make(map[id_t]fill1_t),

		SCells:    make(map[scell_t][]SCell),
		MaxScId:   make(map[scell_t]id_t),
		FreeScIds: make(map[scell_t]map[id_t]fill1_t),
	}

	return rn
}

// GetEpoch returns the current epoch without incrementing.
func (r *Riminet) GetEpoch() epoch_t {
	return r.Epoch
}

// NextEpoch increments the current epoch and returns the value.
func (r *Riminet) NextEpoch() epoch_t {
	r.Epoch += 1
	return r.Epoch
}

// GetFreeFcId returns an id which is not is use.
// It returns an id from the pool of available ids or
// if pool is empty, return (maxFcId+1) and increment it.
func (r *Riminet) GetFreeFcId() id_t {
	var id id_t = 0
	if len(r.FreeFcIds) > 0 {
		for k, _ := range r.FreeFcIds {
			id = k
			break
		}
		delete(r.FreeFcIds, id)
		return id
	} else {
		r.MaxFcId += 1

		if r.MaxFcId > id_t(cap(r.FCells)) {
			// Error State
		}

		if r.MaxFcId == id_t(cap(r.FCells)) {
			// Array is Full.
			// Add a dummy dead cell to increase size.
			DUMMY_FCELL.SetLiveness(false) // just to be sure
			r.FCells = append(r.FCells, DUMMY_FCELL)
		}

		return r.MaxFcId
	}
}

// ReturnFreeFcId adds a free id to the pool of available ids.
func (r *Riminet) ReturnFreeFcId(id id_t) {
	r.FreeFcIds[id] = FILLER
}

// GetFreeScId returns an id which is not is use.
// It returns an id from the pool of available ids or
// if pool is empty, return (maxScId+1) and increment it.
func (r *Riminet) GetFreeScId(scell_type scell_t) id_t {
	freeIds := r.FreeScIds[scell_type]
	maxId  := r.MaxScId[scell_type]
	var id id_t = 0
	if len(freeIds) > 0 {
		for k, _ := range freeIds {
			id = k
			break
		}
		delete(freeIds, id)
		return id
	} else {
		maxId += 1
		r.MaxScId[scell_type] = maxId

		scells := r.SCells[scell_type]

		if maxId > id_t(cap(scells)) {
			// Error State
		}

		if maxId == id_t(cap(scells)) {
			// Array is Full.
			// Add a dummy dead cell to increase size.
			DUMMY_SCELL.SetLiveness(false) // just to be sure
			scells = append(scells, DUMMY_SCELL)
			r.SCells[scell_type] = scells
		}

		return maxId
	}
}

// ReturnFreeScId adds a free id to the pool of available ids.
func (r *Riminet) AddFreeScId(id id_t) {
	r.FreeScIds[scell_t(id >> 24)][id] = fill1_t(FILLER)
}
