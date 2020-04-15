-- migrate:up
CREATE TABLE players (
  id serial NOT NULL,
  game_id int,
  role_id int,
  name varchar(128),
  uid varchar(255),
  state int NOT NULL DEFAULT 0,
  created_at TIMESTAMP NOT NULL DEFAULT now(),

  FOREIGN KEY (role_id) REFERENCES roles(id),
  FOREIGN KEY (game_id) REFERENCES games(id),
  PRIMARY KEY (id)
);

-- migrate:down
DROP TABLE players;
