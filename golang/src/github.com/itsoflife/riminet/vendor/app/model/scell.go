package model

// SCell represents the sensor cell.
// There is no parent child relation between SCells.
// The ordering of the members is critical in saving space.
// Total size should always be 64 bytes. (for efficiency)
type SCell struct {
	// MSB is the Sensor Cell type (0-255).
	// The sense type of sensor cell.
	Id id_t //4B

	// The range of values.
	Lower uint32 //4B lower limit
	Upper uint32 //4B upper limit

	// The current accumulated energy in the cell.
	Energy energy_t //4B

	// history of activation (the last <=32 epocs only)
	History uint32  //4B

	// Misc info
	// Bit 0    : 0 dead, 1 live
	// Bit 1    : permanent cell
	// Bit 2-3  : unused
	// Bit 4-7  : current reference count (0-8)
	// Bit 8-15 : valid refs : 0 invalid, 1 valid
	// Bit 16-31: unused
	Misc1 uint32 //4B

	// Last epoch it was touched (accessed/modified).
	Epoch epoch_t //4B

	// The cells that refer to this cell.
	// Sensor cells have no parents.
	Refs [8]id_t //32B

	// The connection strengths.
	// one quadlet bits for each connection, in order.
	// hence connection strength is 0-15 only.
	Weights uint32 //4B
}

func MakeSCell() SCell {
	sc := SCell{
		Id:         0,
		Misc1:      0,
		Lower:      0,
		Upper:      0,
		Energy:     0,
		Epoch:      0,
	}

	for i := uint8(0); i < 8; i++ {
		sc.Refs[i] = 0
		sc.SetWeight(i, INIT_REF_WEIGHT)
	}

	return sc
}

// IsLive returns true if the cell is live.
func (sc *SCell) IsLive() bool {
	return (sc.Misc1 & 0x01) != 0
}

func (sc *SCell) SetLiveness(live bool) {
	if live {
		sc.Misc1 |= 0x01
	} else {
		sc.Misc1 &= ^uint8(0x01)
	}
}

func (sc *SCell) SetWeight(index uint8, weight uint8) {
	// index must be between 0-7
	// weight must be between 0-15
	index = (7 - index) << 2
	mask := 0xFFFFFFFF ^ (uint32(0xF) << index)
	value := uint32(weight & 0x0F) << index
	sc.Weights = (sc.Weights & mask) | value
}

func (sc *SCell) GetWeight(index uint8) uint8 {
	// index must be between 0-7
	index = 7 - index
	index = index << 2 // index * 4
	return uint8((sc.Weights >> index) & 0xF)
}

// check if value is in the range of the sensor cell
func (sc *SCell) IsInRange(value uint32) bool {
	return value >= sc.Lower && value <= sc.Upper
}

// split the range of the sensor cell
func (sc *SCell) Split() {

}
