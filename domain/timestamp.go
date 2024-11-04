package domain

import "time"

type Timezone struct {
	TimezoneName   string
	TimezoneOffset int64
}

const dhakaTimezoneOffset = 21600

var BDTimezone = Timezone{
	TimezoneName:   "Asia/Dhaka",
	TimezoneOffset: dhakaTimezoneOffset,
}

// TimeStamp contains createdAt, updatedAt fields.
type TimeStamp struct {
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp not null;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp not null;default:current_timestamp"`
	// DeletedAt *time.Time `json:"deleted_at"` // don't need this for this application
}

// PopulateTimeStamp populates the timestamp for created_at and updated_at field.
func (t *TimeStamp) PopulateTimeStamp() {
	now := time.Now()
	t.CreatedAt = now
	t.UpdatedAt = now
}

// PopulateUpdateTimeStamp populates the timestamp for updated_at field.
func (t *TimeStamp) PopulateUpdateTimeStamp() {
	t.UpdatedAt = time.Now()
}
