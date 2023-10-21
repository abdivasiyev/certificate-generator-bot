package bot

type (
	Response[T any] struct {
		Ok          bool   `json:"ok"`
		ErrorCode   int    `json:"error_code"`
		Result      T      `json:"result"`
		Description string `json:"description"`
	}
	DeleteWebhook struct {
		DropPendingUpdates bool `json:"drop_pending_updates"`
	}

	SetWebhook struct {
		URL         string `json:"url"`
		SecretToken string `json:"secret_token"`
	}

	CallbackQuery struct {
		ID      string   `json:"id"`
		From    *User    `json:"from"`
		Message *Message `json:"message"`
		Data    string   `json:"data"`
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
