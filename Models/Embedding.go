package Models

import (
	"github.com/lib/pq"
	"time"
)

type TUEmbed struct {
	EmbedID         uint            `gorm:"primaryKey;AUTO_INCREMENT;uniqueIndex" json:"id"`
	UserUID         string          `gorm:"type:varchar" json:"user_uid"`
	EmbeddingVector pq.Float64Array `gorm:"type:float[];not null" json:"embedding_vector"`
	ModelVersion    string          `gorm:"type:varchar(50)" json:"model_version"`
	CreatedAt       time.Time       `json:"created_at"`
}
