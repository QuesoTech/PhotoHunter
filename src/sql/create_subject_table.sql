CREATE TABLE subjects (
	id        serial PRIMARY KEY,
	ds_id     integer REFERENCES dataset (id),
	target    text,
	dummies   text[]
);
