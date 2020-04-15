-- migrate:up
CREATE TABLE roles (
  id serial NOT NULL,
  name varchar(128),
  description TEXT,

  created_at TIMESTAMP NOT NULL DEFAULT now(),

  PRIMARY KEY (id)
);

-- migrate:down
DROP TABLE roles;

