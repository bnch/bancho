package packethandler

// BotName is the name of the bot that handles automated stuff.
// BanchoBot automatically gets pink name. Other bots can't get pink name.
const BotName = "BanchoBot"

// BotID is the user ID of the bot. Set it to an ID that doesn't exist in your db (for instance 1).
// Do not set it to 2 because if so the user won't be able to contact the bot.
const BotID = 1

// SendMessage sends a private message to an user. (user being their token)
func SendMessage(user string, message string) {
	sess := Sessions[user]
	if sess == nil {
		return
	}
	sess.Push(ChatMessage{
		From:    BotName,
		To:      sess.User.Name,
		Content: message,
		UserID:  BotID,
	}.ToPacketNoIgnore())
}
