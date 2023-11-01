package models

import "time"

// VisitorTracking represents the visitor_tracking table schema.
type VisitorTracking struct {
	VisitorID   string    `gorm:"primaryKey;type:varchar(36)" json:"visitor_id"`
	Timestamp   time.Time `gorm:"type:timestamptz" json:"timestamp"`
	AdblockUser bool      `gorm:"type:boolean" json:"adblock_user"`
	CampaignID  int64     `gorm:"type:bigint" json:"campaign_id"`
}

// TableName specifies the table name for the VisitorTracking struct.
func (VisitorTracking) TableName() string {
	return "visitor_tracking"
}
