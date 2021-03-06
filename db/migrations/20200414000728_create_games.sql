-- migrate:up
CREATE TABLE games (
  id serial NOT NULL,
  state int NOT NULL DEFAULT 0,
  phase int NOT NULL DEFAULT 0,
  phase_count int NOT NULL DEFAULT 0,
  channel_id varchar(255),
  created_at TIMESTAMP NOT NULL DEFAULT now(),

  primary key (id)
);

-- migrate:down
DROP TABLE games;
