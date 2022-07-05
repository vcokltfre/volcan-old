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
	ID        int    `json:"id,omitempty" gorm:"primaryKey;autoIncrement;uniqueIndex"`
	UserID    string `json:"user_id,omitempty" gorm:"type:varchar(255);index"`
	UserName  string `json:"user_name,omitempty" gorm:"type:varchar(255)"`
	ModID     string `json:"mod_id,omitempty" gorm:"type:varchar(255);index"`
	ModName   string `json:"mod_name,omitempty" gorm:"type:varchar(255)"`
	Type      string `json:"type,omitempty" gorm:"type:varchar(255);index"`
	Reason    string `json:"reason,omitempty" gorm:"type:text"`
	CreatedAt int64  `json:"created_at,omitempty" gorm:"index"`
	ExpiresAt int64  `json:"expires_at,omitempty" gorm:"index"`
	Notified  bool   `json:"notified,omitempty" gorm:"type:boolean"`
	MuteType  string `json:"mute_type,omitempty" gorm:"type:varchar(255);index"`
	Metadata  string `json:"metadata,omitempty" gorm:"type:text"`
}
