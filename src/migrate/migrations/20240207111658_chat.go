package migrations

import (
	"context"
	"homework/domain/model/chat"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		_, err := db.NewCreateTable().
			Model((*chat.Chat)(nil)).
			ForeignKey(`(user_id) REFERENCES "user" (id) ON DELETE CASCADE`).
			ForeignKey(`(room_id) REFERENCES "room" (room_id) ON DELETE CASCADE`).
			Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		_, err := db.NewDropTable().
			Model((*chat.Chat)(nil)).
			IfExists().
			Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	})
}
