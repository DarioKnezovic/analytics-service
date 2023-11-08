package models

type ModalCTRTracking struct {
	Session         string `json:"session"`
	Timestamp       string `json:"timestamp"`
	InteractionType string `json:"interaction_type"`
	ObjectID        string `json:"object_id"`
	AdditionalData  string `json:"additional_data"`
	CampaignID      int    `json:"campaign_id"`
}

func (ModalCTRTracking) TableName() string {
	return "modal_ctr_tracking"
}
