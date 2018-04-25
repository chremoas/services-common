package command

import "regexp"

func IsDiscordUser(user string) bool {
	var discordUser = regexp.MustCompile(`<@\d*>`)
	return discordUser.MatchString(user)
}
