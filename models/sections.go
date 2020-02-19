package models

import "time"

//Sections Model
type Sections struct {
	ID             uint32     `gorm:"id" json:"id"`
	CourseID       uint32     `gorm:"course_id" json:"courseID"`
	SequenceNumber uint32     `gorm:"sequence_number" json:"sequenceNumber"`
	Name           string     `gorm:"name" json:"name"`
	StudyMaterial  string     `gorm:"study_material" json:"studyMaterial"`
	CreatedAt      time.Time  `gorm:"created_at" json:"-"`
	UpdatedAt      time.Time  `gorm:"updated_at" json:"-"`
	DeletedAt      *time.Time `gorm:"deleted_at" json:"-"`
}

// ChnageSectionSequence Struct
type ChnageSectionSequence struct {
	SectionID       uint32 `json:"section_id"`
	CourseID        uint32 `json:"course_id"`
	CurrentLocation int    `json:"current_location"`
	ReLocation      int    `json:"re_location"`
}

type SectionInfoReq struct {
	SectionID uint32 `json:"section_id"`
	CourseID  uint32 `json:"course_id"`
	UserID    uint32 `json:"user_id"`
}

type SectionInfo struct {
	Section             Sections
	PercentageCompleted string
}
