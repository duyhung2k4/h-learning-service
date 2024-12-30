package entity

import "gorm.io/gorm"

type ProcessStream struct {
	gorm.Model
	ProfileId      uint            `json:"profileId"`
	IpMergeServer  string          `json:"ipMergeServer"`
	IpStreamServer string          `json:"ipStreamServer"`
	Uuid           string          `json:"uuid"`
	Status         PROCESSS_STATUS `json:"status"`

	Profile *Profile `json:"profile" gorm:"foreignKey:ProfileId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type PROCESSS_STATUS string

const (
	PROCESS_PENDING   PROCESSS_STATUS = "PROCESS_PENDING"
	PROCESS_STREAMING PROCESSS_STATUS = "PROCESS_STREAMING"
	PROCESS_FINISHED  PROCESSS_STATUS = "PROCESS_FINISHED"
)
