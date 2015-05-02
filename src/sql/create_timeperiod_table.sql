CREATE TABLE timeperiods (
	id        serial PRIMARY KEY,
	ds_id     integer REFERENCES dataset (id),
	s_time    time, 
	e_time    time

);
