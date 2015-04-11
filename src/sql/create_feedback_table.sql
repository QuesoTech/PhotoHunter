CREATE TABLE feedback (
  id serial PRIMARY KEY,
  user_id integer REFERENCES user (id),
  datapoint_id integer REFERENCES datapoint (id),
  subject_id integer REFERENCES subject (id),
  CONSTRAINT u_constraint UNIQUE (user_id, datapoint_id, subject_id)
);
