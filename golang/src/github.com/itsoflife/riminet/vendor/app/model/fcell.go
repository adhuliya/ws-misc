package model

// FCell represents the fact cell.
// The ordering of the members is critical in saving space.
// Total size should always be 64 bytes. (for efficiency)
type FCell struct {
	// id of the FCell
	Id id_t // 4B

	// The current accumulated energy in the cell.
	Energy energy_t // 4B

	// history of activation
	History uint32 // 4B

	// Last epoch it was touched (accessed/modified).
	Epoch epoch_t // 4B

	// bit 0    : 0 dead, 1 live
	// bit 1	: 0 non-permanent, 1 permanent
	// bit 2,3  : parent 1 type,
	// bit 4,5  : parent 2 type,
	// bit 6-7  : unused
	// bit 8-11 : ref count
	// bit 12-15: unused
	// bit 16-23: valid refs; 0 invalid, 1 valid
	Misc1 uint32 // 4B

	Misc2 uint32 // 4B
	Misc3 uint32 // 4B

	// The cells that refer to this cell.
	// refs[0] and ref[1] are parent 1 and 2
	Refs [8]id_t // 32B

	// The connection strengths.
	// one quadlet bits for each connection, in order.
	// hence connection strength is 0-15 only.
	Weights uint32 // 4B  the connection strengths
}

func MakeFCell() FCell {
	fc := FCell {
		Id:			0,
		History:    0,
		Energy:     0,
        Misc1:      0,
		Misc2:		0,
		Misc3:		0,
		Epoch:      0,
	}

	for i := uint8(0); i < 8; i++ {
		fc.Refs[i] = 0
		fc.SetWeight(i, INIT_REF_WEIGHT)
	}

	return fc
}

func (fc *FCell) SetLiveness(live bool) {
	if live {
		fc.Flags1 |= 0x0010
	} else {
		fc.Flags1 &= ^uint16(0x0010)
	}
}

func (fc *FCell) IsLive() bool {
	return fc.Flags1&0x0010 != 0
}

func (fc1 *FCell) IsSimilar(fc2 *FCell) bool {
	count := CountOnes(fc1.History & fc2.History)
	if count >= SIMILARITY_COUNT {
		return true
	}
	return false
}

func (fc *FCell) SetWeight(index uint8, weight uint8) {
	// index must be between 0-7
	// weight must be between 0-15
	index = (7 - index) << 2
	mask := 0xFFFFFFFF ^ (uint32(0xF) << index)
	value := uint32(weight & 0x0F) << index
	fc.Weights = (fc.Weights & mask) | value
}

func (fc *FCell) GetWeight(index uint8) uint8 {
	// index must be between 0-7
	index = (7 - index) << 2
	return uint8((fc.Weights >> index) & 0xF)
}