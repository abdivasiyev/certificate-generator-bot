package bot

const (
	PrivateChat    ChatType = "private"
	GroupChat      ChatType = "group"
	SuperGroupChat ChatType = "supergroup"
	ChannelChat    ChatType = "channel"
)

type (
	SendPhoto struct {
		ChatID           string    `json:"chat_id"`
		Photo            string    `json:"photo"`
		Caption          string    `json:"caption,omitempty"`
		ParseMode        string    `json:"parse_mode,omitempty"`
		ReplyToMessageID int64     `json:"reply_to_message_id,omitempty"`
		ReplyMarkup      *Keyboard `json:"reply_markup,omitempty"` // create with generic for *KeyboardMarkup types
	}

	EditReplyMarkup struct {
		ChatID      int64     `json:"chat_id"`
		MessageID   int64     `json:"message_id"`
		ReplyMarkup *Keyboard `json:"reply_markup,omitempty"` // create with generic for *KeyboardMarkup types
	}

	EditMessage struct {
		ChatID      int64     `json:"chat_id"`
		MessageID   int64     `json:"message_id"`
		Text        string    `json:"text"`
		ParseMode   string    `json:"parse_mode,omitempty"`
		ReplyMarkup *Keyboard `json:"reply_markup,omitempty"` // create with generic for *KeyboardMarkup types
	}

	InlineKeyboardButton struct {
		Text         string `json:"text"`
		URL          string `json:"url,omitempty"`
		CallbackData string `json:"callback_data,omitempty"`
	}

	Keyboard struct {
		InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard,omitempty"`
	}

	SendMessage struct {
		ChatID                   int64     `json:"chat_id"`
		Text                     string    `json:"text"`
		ParseMode                string    `json:"parse_mode,omitempty"`
		ReplyToMessageID         int64     `json:"reply_to_message_id,omitempty"`
		AllowSendingWithoutReply bool      `json:"allow_sending_without_reply,omitempty"`
		ReplyMarkup              *Keyboard `json:"reply_markup,omitempty"` // create with generic for *KeyboardMarkup types
	}

	DeleteWebhook struct {
		DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`
	}

	SetWebhook struct {
		URL         string `json:"url"`
		SecretToken string `json:"secret_token,omitempty"`
	}

	Response[T any] struct {
		Ok          bool   `json:"ok"`
		ErrorCode   int    `json:"error_code"`
		Result      T      `json:"result"`
		Description string `json:"description"`
	}

	CallbackQuery struct {
		ID      string   `json:"id"`
		From    *User    `json:"from"`
		Message *Message `json:"message"`
		Data    string   `json:"data"`
	}

	ChatType string

	Chat struct {
		ID        int64    `json:"id"`
		Type      ChatType `json:"type"`
		Username  string   `json:"username"`
		FirstName string   `json:"first_name"`
		LastName  string   `json:"last_name"`
	}

	User struct {
		ID        int64  `json:"id"`
		IsBot     bool   `json:"is_bot"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Username  string `json:"username"`
	}

	Message struct {
		MessageID int64  `json:"message_id"`
		From      *User  `json:"from"`
		Chat      Chat   `json:"chat"`
		Date      int64  `json:"date"`
		Text      string `json:"text"`
	}

	Update struct {
		UpdateID      int            `json:"update_id"`
		Message       *Message       `json:"message"`
		EditedMessage *Message       `json:"edited_message"`
		CallbackQuery *CallbackQuery `json:"callback_query"`
	}
)
