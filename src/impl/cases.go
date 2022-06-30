package impl

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/vcokltfre/volcan/src/config"
	"github.com/vcokltfre/volcan/src/core"
	"github.com/vcokltfre/volcan/src/database"
)

type CaseManager struct {}

func resolveMember(userID string) (*discordgo.Member, error) {
	member, err := core.Session.GuildMember(config.Config.GetPrimaryGuild(), userID)
	if err != nil {
		return nil, err
	}

	if member == nil {
		return nil, fmt.Errorf("Unable to find member %s", userID)
	}

	return member, nil
}

func resolveCaseContext(userID, modID string) (*discordgo.Member, *discordgo.Member, error) {
	member, err := resolveMember(userID)
	if err != nil {
		return nil, nil, err
	}

	mod, err := resolveMember(modID)
	if err != nil {
		return nil, nil, err
	}

	return member, mod, nil
}

func notifyMember(member *discordgo.Member, reason string) {
	channel, err := core.Session.UserChannelCreate(member.User.ID)
	if err != nil {
		Error.Dispatch(err)
	}

	_, err = core.Session.ChannelMessageSend(channel.ID, reason)
	if err != nil {
		Error.Dispatch(err)
	}
}

func (m *CaseManager) CreateCase(userID, userName, modID, modName, typ, reason string, expires int64) (*database.Case, error) {
	db_case := &database.Case{
		UserID:    userID,
		UserName:  userName,
		ModID:     modID,
		ModName:   modName,
		Type:      typ,
		Reason:    reason,
		CreatedAt: time.Now().Unix(),
		ExpiresAt: expires,
	}

	err := database.DB.Create(db_case).Error
	if err != nil {
		return nil, err
	}

	return db_case, nil
}

func (m *CaseManager) WarnUser(userID string, modID string, reason string, notify bool) (*database.Case, error) {
	member, mod, err := resolveCaseContext(userID, modID)
	if err != nil {
		return nil, err
	}

	if notify {
		notifyMember(member, reason)
	}

	return m.CreateCase(userID, member.User.Username, modID, mod.User.Username, database.CaseWarn, reason, 0)
}
