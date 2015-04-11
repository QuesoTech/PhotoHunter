CREATE TABLE datapoint (
  id serial PRIMARY KEY,
  dataset_id integer REFERENCES dataset (id),
  image_url varchar(256)
);
