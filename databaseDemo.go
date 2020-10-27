package main


import(
"database/sql"
_ "github.com/mattn/go-sqlite3" //import for side effects
"log"
)

func main() {
	myDatabase := OpenDataBase("./Demo.db")
	defer myDatabase.Close()
	create_tables(myDatabase)
}
func OpenDataBase(dbfile string) *sql.DB{
	database, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		log.Fatal(err)
	}
	return database
}

func create_tables(database *sql.DB){
	createStatement1 := "CREATE TABLE IF NOT EXISTS students(    " +
		"banner_id INTEGER PRIMARY KEY," +
		"first_name TEXT NOT NULL," +
		"last_name TEXT NOT NULL," +
		"gpa REAL DEFAULT 0," +
		"credits INTEGER DEFAULT 0);"
	create_course :="CREATE TABLE IF NOT EXISTS course(" +
		"    course_prefix TEXT NOT NULL," +
		"    course_number INTEGER NOT NULL," +
		"    cap INTEGER DEFAULT 20," +
		"    description TEXT," +
		"    PRIMARY KEY(course_prefix, course_number)    );"

	database.Exec(createStatement1)
	database.Exec(create_course)
}