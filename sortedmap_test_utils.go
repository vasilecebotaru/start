package sortedmap

import (
	"fmt"
	"time"
	mrand "math/rand"
)

func randStr(n int) string {
    mrand.Seed(time.Now().UTC().UnixNano())

	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_+=~[]{}|:;<>,./?"
	result := make([]byte, n)

	for i := 0; i < n; i++ {
		result[i] = chars[mrand.Intn(len(chars))]
	}
    return string(result)
}

func randRecord() Record {
	year := mrand.Intn(2129)
	if year < 1 {
		year++
	}
	mth := time.Month(mrand.Intn(12))
	if mth < 1 {
		mth++
	}
	day := mrand.Intn(28)
	if day < 1 {
		day++
	}
	return Record{
		Key: randStr(42),
		Val: time.Date(year, mth, day, 0, 0, 0, 0, time.UTC),
	}
}

func randRecords(n int) []*Record {
	records := make([]*Record, n)
	for i := range records {
		rec := randRecord()
		records[i] = &rec
	}
	return records
}

func verifyRecords(ch <-chan Record) error {
	previousRec := Record{}

	for rec := range ch {
		if previousRec.Key != "" {
			if previousRec.Val.(time.Time).After(rec.Val.(time.Time)) {
				return fmt.Errorf("%v %v",
					unsortedErr,
					fmt.Sprintf("prev: %+v, current: %+v.", previousRec, rec),
				)
			}
		}
		previousRec = rec
	}

	return nil
}

func newSortedMapFromRandRecords(n int) *SortedMap {
	records := randRecords(n)
	sm := New(nil)
	sm.BatchReplace(records...)

	return sm
}