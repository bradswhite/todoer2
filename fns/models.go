package fns

import (
  "time"
  "github.com/google/uuid"
)

type Todo struct {
  ID        uuid.UUID
  Username  string
  Text      string
  Completed bool
  Timestamp time.Time
}

/*type Todo struct {
  ID        uuid.UUID   `db:"id`
  Username  string      `db:"username"`
  Text      string      `db:"text"`
  Completed bool        `db:"completed"`
  Timestamp time.Time   `db:"timestamp"`
}*/

