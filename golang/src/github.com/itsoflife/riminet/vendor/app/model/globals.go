package model

var Rnet Riminet // the global Riminet variable

/* the common types used: possibly some to be promoted to uint64 later */
type id_t uint32
type range_t uint32
type epoch_t uint32
type energy_t uint32
type weight_t uint8
type count_t uint32 // for various countings
type label_t uint8
type scell_t uint8
type fill1_t uint8

const (
	ARR_SIZE_F      id_t     = 4096
	ARR_SIZE_S      id_t     = 4096
	INIT_REF_WEIGHT uint8 = 0x0F

	DEAD_CELL   uint8 = 0x00
	FACT_CELL   uint8 = 0x01
	SENSOR_CELL uint8 = 0x02

	MAX_SCELLS_PER_TYPE uint32 = (1 << 24) - 1

	SENSE_TYPE_NONE scell_t = 0x00
	SENSE_TYPE_1    scell_t = 0x01
	SENSE_TYPE_2    scell_t = 0x02
	SENSE_TYPE_3    scell_t = 0x03
	SENSE_TYPE_4    scell_t = 0x04

	SIMILARITY_COUNT = 5

	FILLER = 0x00  // undefined value
)

var (
	DUMMY_SCELL SCell
	DUMMY_FCELL FCell
)

func Initialize() {
	// A must (THE FIRST) function to call, to use riminet.
	DUMMY_SCELL = MakeSCell()
	DUMMY_FCELL = MakeFCell()
	Rnet = MakeRiminet()
}