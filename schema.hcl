table "games" {
  schema = schema.public
  column "id" {
    null = false
    type = serial
  }
  column "state" {
    null    = false
    type    = integer
    default = 0
  }
  column "phase" {
    null    = false
    type    = integer
    default = 0
  }
  column "phase_count" {
    null    = false
    type    = integer
    default = 0
  }
  column "channel_id" {
    null = true
    type = character_varying(255)
  }
  column "created_at" {
    null    = false
    type    = timestamp
    default = sql("now()")
  }
  primary_key {
    columns = [column.id]
  }
}
table "notifications" {
  schema = schema.public
  column "id" {
    null = false
    type = serial
  }
  column "game_id" {
    null = true
    type = integer
  }
  column "message" {
    null = true
    type = text
  }
  column "notify_at" {
    null = true
    type = timestamp
  }
  column "sent" {
    null    = false
    type    = boolean
    default = false
  }
  column "created_at" {
    null    = false
    type    = timestamp
    default = sql("now()")
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "notifications_game_id_fkey" {
    columns     = [column.game_id]
    ref_columns = [table.games.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
}
table "players" {
  schema = schema.public
  column "id" {
    null = false
    type = serial
  }
  column "game_id" {
    null = true
    type = integer
  }
  column "role_id" {
    null = true
    type = integer
  }
  column "name" {
    null = true
    type = character_varying(128)
  }
  column "uid" {
    null = true
    type = character_varying(255)
  }
  column "state" {
    null    = false
    type    = integer
    default = 0
  }
  column "created_at" {
    null    = false
    type    = timestamp
    default = sql("now()")
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "players_game_id_fkey" {
    columns     = [column.game_id]
    ref_columns = [table.games.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "players_role_id_fkey" {
    columns     = [column.role_id]
    ref_columns = [table.roles.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
}
table "roles" {
  schema = schema.public
  column "id" {
    null = false
    type = serial
  }
  column "name" {
    null = true
    type = character_varying(128)
  }
  column "description" {
    null = true
    type = text
  }
  column "created_at" {
    null    = false
    type    = timestamp
    default = sql("now()")
  }
  primary_key {
    columns = [column.id]
  }
}
table "schema_migrations" {
  schema = schema.public
  column "version" {
    null = false
    type = character_varying(255)
  }
  primary_key {
    columns = [column.version]
  }
}
schema "public" {
}
