package main

// Name is a struct holding a First, Last name pair for a given researcher
type Name struct {
	First string
	Last  string
}

// Researcher represents a researcher managing projects with the PhotoHunter
// researcher interface. All researchers have an ID, Name, and contact Email.
type Researcher struct {
	ID int64
	Name
	Email string
}

// NewResearcher takes a name and email and inserts it into the database, provided
// email is not already in use
func NewResearcher(name Name, email, password string) (r *Researcher, err error) {
	stmt, err := DB.Prepare("INSERT INTO researchers (fname, lname, email, pword) VALUES ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		return
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRow(name.First, name.Last, email, password).Scan(&id)
	if err != nil {
		return
	}

	r = &Researcher{
		ID:    id,
		Name:  name,
		Email: email,
	}

	return
}

// GetResearcherByID fetches a researcher from the database by its id
func GetResearcherByID(id int64) (r *Researcher, err error) {
	var fname string
	var lname string
	var email string

	err = DB.QueryRow("SELECT fname, lname, email FROM researchers WHERE id=$1", id).Scan(&fname, &lname, &email)

	name := Name{
		First: fname,
		Last:  lname,
	}

	r = &Researcher{
		ID:    id,
		Name:  name,
		Email: email,
	}

	return
}
