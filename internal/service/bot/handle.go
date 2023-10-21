package bot

import "context"

func (s *service) Handle(ctx context.Context, update Update) error {
	s.logger.Debug().Any("update from telegram", update).Send()

	return nil
}
