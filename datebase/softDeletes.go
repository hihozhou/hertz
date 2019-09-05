package datebase

import "time"

type SoftDeletes struct {
	DeletedAt time.Time `json:"deleted_at"` //软删除字段
}
