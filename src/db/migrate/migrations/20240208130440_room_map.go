package migrations

import (
	"context"
	"homework/domain/model/room"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		_, err := db.NewCreateTable().
			Model((*room.RoomMap)(nil)).
			ForeignKey(`(user_id) REFERENCES "user" (id) ON DELETE CASCADE`).
			ForeignKey(`(room_id) REFERENCES "room" (id) ON DELETE CASCADE`).
			Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		_, err := db.NewDropTable().
			Model((*room.RoomMap)(nil)).
			IfExists().
			Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	})
}
