package cases_api

import "github.com/vcokltfre/volcan/src/database"

var caseTypes = []string{
	database.CaseBan,
	database.CaseKick,
	database.CaseMute,
	database.CaseWarn,
	database.CaseUnban,
	database.CaseUnmute,
	database.CaseHidden,
}

type createCaseData struct {
	UserID   string `json:"user_id" validator:"required"`
	ModID    string `json:"mod_id" validator:"required"`
	Type     string `json:"type" validator:"required"`
	Reason   string `json:"reason"`
	Expires  int64  `json:"expires"`
	Notified bool   `json:"notified"`
	MuteType string `json:"mute_type"`
	Metadata string `json:"metadata"`
}

type createWarnData struct {
	UserID   string `json:"user_id" validator:"required"`
	ModID    string `json:"mod_id" validator:"required"`
	Reason   string `json:"reason"`
	Notify   bool   `json:"notify"`
	Metadata string `json:"metadata"`
}
