CREATE TABLE feedback (
  id integer PRIMARY KEY,
  user_id integer REFERENCES user (id),
  datapoint_id integer REFERENCES datapoint (id),
  dataset_id integer REFERENCES dataset (id),
  CONSTRAINT u_constraint UNIQUE (user_id, datapoint_id, dataset_id)
);
