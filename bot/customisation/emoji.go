package customisation

import (
	"fmt"
	"github.com/rxdn/gdl/objects"
	"github.com/rxdn/gdl/objects/guild/emoji"
)

type CustomEmoji struct {
	Name     string
	Id       uint64
	Animated bool
}

func NewCustomEmoji(name string, id uint64, animated bool) CustomEmoji {
	return CustomEmoji{
		Name: name,
		Id:   id,
	}
}

func (e CustomEmoji) String() string {
	if e.Animated {
		return fmt.Sprintf("<a:%s:%d>", e.Name, e.Id)
	} else {
		return fmt.Sprintf("<:%s:%d>", e.Name, e.Id)
	}
}

func (e CustomEmoji) BuildEmoji() *emoji.Emoji {
	return &emoji.Emoji{
		Id:       objects.NewNullableSnowflake(e.Id),
		Name:     e.Name,
		Animated: e.Animated,
	}
}

var (
	EmojiId         = NewCustomEmoji("id", 1358547901617541401, false)
	EmojiOpen       = NewCustomEmoji("open", 1358547911478218802, false)
	EmojiOpenTime   = NewCustomEmoji("opentime", 1358547930797314150, false)
	EmojiClose      = NewCustomEmoji("close", 1358547873259720856, false)
	EmojiCloseTime  = NewCustomEmoji("closetime", 1358547882936107048, false)
	EmojiReason     = NewCustomEmoji("reason", 1358547992428285982, false)
	EmojiSubject    = NewCustomEmoji("subject", 1358548001718796418, false)
	EmojiTranscript = NewCustomEmoji("transcript", 1358548021322846305, false)
	EmojiClaim      = NewCustomEmoji("claim", 1358547860140064878, false)
	EmojiPanel      = NewCustomEmoji("panel", 1358547945884090378, false)
	EmojiRating     = NewCustomEmoji("rating", 1358547967988072671, false)
	EmojiStaff      = NewCustomEmoji("staff", 1358547978826158213, false)
	EmojiThread     = NewCustomEmoji("thread", 1358548012024201237, false)
	EmojiBulletLine = NewCustomEmoji("bulletline", 1358547835553054904, false)
	EmojiPatreon    = NewCustomEmoji("patreon", 1358547956617314354, false)
	EmojiDiscord    = NewCustomEmoji("discord", 1358547892503318578, false)
	//EmojiTime       = NewCustomEmoji("time", 974006684622159952, false)
)

// PrefixWithEmoji Useful for whitelabel bots
func PrefixWithEmoji(s string, emoji CustomEmoji, includeEmoji bool) string {
	if includeEmoji {
		return fmt.Sprintf("%s %s", emoji, s)
	} else {
		return s
	}
}
