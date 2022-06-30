package database

const (
	CaseWarn   = "warn"
	CaseMute   = "mute"
	CaseUnmute = "unmute"
	CaseKick   = "kick"
	CaseBan    = "ban"
	CaseUnban  = "unban"
	CaseHidden = "hidden"
)

type Case struct {
	ID        int    `json:"id" gorm:"primaryKey;autoIncrement;uniqueIndex"`
	UserID    string `json:"user_id" gorm:"type:varchar(255);index"`
	UserName  string `json:"user_name" gorm:"type:varchar(255)"`
	ModID     string `json:"mod_id" gorm:"type:varchar(255);index"`
	ModName   string `json:"mod_name" gorm:"type:varchar(255)"`
	Type      string `json:"type" gorm:"type:varchar(255);index"`
	Reason    string `json:"reason" gorm:"type:text"`
	CreatedAt int64  `json:"created_at" gorm:"index"`
	ExpiresAt int64  `json:"expires_at" gorm:"index"`
}
