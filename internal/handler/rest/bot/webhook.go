package bot

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/abdivasiyev/telegram-bot/internal/service/bot"
)

func (h *handler) Webhook(w http.ResponseWriter, r *http.Request) {
	var update bot.Update

	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		h.logger.Error().Err(err).Msg("error while unmarshalling telegram update")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = io.WriteString(w, "invalid signature of json")
		return
	}

	if err := h.service.Handle(r.Context(), update); err != nil {
		h.logger.Error().Err(err).Msg("could not handle telegram update")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = io.WriteString(w, "internal server error")
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = io.WriteString(w, "ok")
}
