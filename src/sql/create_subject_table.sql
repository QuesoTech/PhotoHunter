CREATE TABLE subjects (
	id        integer PRIMARY KEY DEFAULT nextval('serial'),
	ds_id     integer REFERENCES dataset (id),
	targets   text[],
	dummies   text[]
);
