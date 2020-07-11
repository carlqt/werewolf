-- migrate:up
CREATE TABLE notifications (
  id serial NOT NULL,
  game_id int,
  message text,
  notify_at TIMESTAMP,
  sent boolean NOT NULL DEFAULT false,
  created_at TIMESTAMP NOT NULL DEFAULT now(),

  FOREIGN KEY (game_id) REFERENCES games(id),
  PRIMARY KEY (id)
);

-- migrate:down
DROP TABLE notifications;