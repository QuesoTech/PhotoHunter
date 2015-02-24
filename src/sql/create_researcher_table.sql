CREATE TABLE researchers (
	id        integer PRIMARY KEY DEFAULT nextval('serial'),
	fname     varchar(20) NOT NULL,
	lname     varchar(25) NOT NULL,
	email     varchar(50) NOT NULL		 

);
