CREATE TABLE timeperiods (
	id        integer PRIMARY KEY DEFAULT nextval('serial'),
	s_time    time, 
	e_time    time

);
