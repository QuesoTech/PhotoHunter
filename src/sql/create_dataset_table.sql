CREATE TABLE dataset (
  id serial PRIMARY KEY,
  researcher_id integer REFERENCES researcher (id),
  name varchar(32)
);
