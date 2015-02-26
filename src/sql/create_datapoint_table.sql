CREATE TABLE datapoint (
  id integer DEFAULT nextval('serial'),
  dataset_id integer REFERENCES dataset (id),
  image_url varchar(256)
);
