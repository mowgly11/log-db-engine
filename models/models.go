package models

type IndexEntry struct {
	SegmentName string
	Offset int
}

const SEGEMENT_SIZE_LIMIT_KB int32 = 1