CREATE TABLE timeperiods (
	id        integer PRIMARY KEY DEFAULT nextval('serial'),
	ds_id     integer REFERENCES dataset (id),
	s_time    time, 
	e_time    time

);
