CREATE TABLE researchers (
	id        serial PRIMARY KEY,
	fname     varchar(20) NOT NULL,
	lname     varchar(25) NOT NULL,
	email     varchar(50) UNIQUE NOT NULL		 

);
