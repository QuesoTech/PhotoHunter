CREATE TABLE dataset (
  id integer PRIMARY KEY,
  researcher_id integer REFERENCES researcher (id),
  name varchar(32)
);
