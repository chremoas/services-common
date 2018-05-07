package command

import "regexp"

func IsDiscordUser(user string) bool {
	var discordUser = regexp.MustCompile(`<@\d*>`)
	return discordUser.MatchString(user)
}

func ExtractUserId(user string) string {
	return user[2 : len(user)-1]
}
