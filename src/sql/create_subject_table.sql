CREATE TABLE subjects (
	id        integer PRIMARY KEY DEFAULT nextval('serial'),
	targets   text[],
	dummies   text[]
);
