package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" //import for side effects
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	myDatabase := OpenDataBase("./Demo.db")
	defer myDatabase.Close()
	create_tables(myDatabase)
	//addSampleStudents(myDatabase)
	//addCourses(myDatabase)
	//registerForClasses(myDatabase)
	findProbationStudents(myDatabase)
}
func OpenDataBase(dbfile string) *sql.DB {
	database, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		log.Fatal(err)
	}
	return database
}

func getMinGPA() float64 {
	fmt.Print("What is the minimum GPA for good standing:")
	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("How did we fail to read from standard in!?!?")
	}
	value = strings.TrimSpace(value)
	min_gpa, err := strconv.ParseFloat(value, 32)
	if err != nil {
		log.Fatal("oooops you typed that wrong", err)
	}
	return min_gpa
}

func findProbationStudents(database *sql.DB) {
	var firstName, lastName string
	var gpa float64
	minGpa := getMinGPA()
	selectStatement := "SELECT first_name, last_name, gpa FROM STUDENTS WHERE gpa < ?"
	resultSet, err := database.Query(selectStatement, minGpa)
	if err != nil {
		log.Fatal("Bad Query", err)
	}
	defer resultSet.Close()
	for resultSet.Next() {
		err = resultSet.Scan(&firstName, &lastName, &gpa)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s %s is on probation with a GPA of %f\n", firstName, lastName, gpa)
	}
}

func registerForClasses(database *sql.DB) {
	insertStatement := "INSERT INTO CLASS_LIST (banner_id, course_prefix, course_number,  registration_date)" +
		"VALUES(?, 'Comp', 510, DATE('now'))"
	preppedStatement, err := database.Prepare(insertStatement)
	if err != nil {
		log.Fatal("Hey prof you goofed it trying to type live", err)
	}
	for i := 1001; i <= 1008; i++ {
		preppedStatement.Exec(i)
	}
}

func addCourses(database *sql.DB) {
	var sampleData = map[string]string{
		"comp502": "Research\n(3 credits)\nPrerequisite: Consent of the department; formal application required\nOriginal research is undertaken by the graduate student in their field. This course culminates in a capstone project. For details, consult the paragraph titled “Directed or Independent Study” in the “College of Graduate Studies” section of this catalog. Offered fall and spring semesters.",
		"comp503": "Directed Study\n(1-3 credits)\nPrerequisite: Consent of the department; formal application required\nDirected study is designed for the graduate student who desires to study selected topics in a specific field. For details, consult the paragraph titled “Directed or Independent Study” in the “College of Graduate Studies” section of this catalog. Repeatable: may earn a maximum of six credits. Offered fall and spring semesters.",
		"comp510": "Topics in Programming Languages\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nThis course investigates programming language development from designer’s, user’s and implementer’s point of view. Topics include formal syntax and semantics, language system, extensible languages and control structures. There is also a survey of intralanguage features, covering ALGOL-60, ALGOL-68, Ada, Pascal, LISP, SNOBOL-4 APL, SIMULA-67, CLU, MODULA, and others. Offered periodically.",
		"comp520": "Operating Systems Principles\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nThis course examines design principles such as optimal scheduling; file systems, system integrity and security, as well as the mathematical analysis of selected aspects of operating system design. Topics include queuing theory, disk scheduling, storage management and the working set model. Design and implementation of an operating system nucleus is also studied. Offered periodically.",
		"comp525": "Design and Construction of Compilers\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nIn this course, topics will include lexical and syntactic analysis; code generation; error detection and correction; optimization techniques; models of code generators; and incremental and interactive compiling. Students will design and implement a compiler. Offered periodically.",
		"comp530": "Software Engineering\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nTopics in this course will include construction of reliable software, software tools, software testing methodologies, structured design, structured programming, software characteristics and quality and formal proofs of program correctness. Chief programmer teams and structure walk-throughs will be employed. Offered periodically.\n",
		"comp540": "Automata, Computability and Formal Languages\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nTopics in this course will include finite automata and regular languages, context- free languages, Turing machines and their variants, partial recursive functions and grammars, Church’s thesis, undecidable problems, complexity of algorithms and completeness. Offered periodically.",
		"comp545": "Analysis of Algorithms\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nThis course deals with techniques in the analysis of algorithms. Topics to be chosen from among the following: dynamic programming, search and traverse techniques, backtracking, numerical techniques, NP-hard and NP-complete problems, approximation algorithms and other topics in the analysis and design of algorithms. Offered fall semester.\n",
		"comp560": "Artificial Intelligence\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nThis course is an introduction to LISP or another AI programming language. Topics are chosen from pattern recognition, theorem proving, learning, cognitive science and vision. It also presents introduction to the basic techniques of AI such as heuristic search, semantic nets, production systems, frames, planning and other AI topics. Offered periodically.\n",
		"comp570": "Robotics\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nThis is a project-oriented course in robotics. Topics are chosen from manipulator motion and control, motion planning, legged-motion, vision, touch sensing, grasping, programming languages for robots and automated factory design. Offered periodically.",
		"comp580": "Database Systems\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nIn this course, topics will include relational, hierarchical and network data models; design theory for relational databases and query optimization; classification of data models, data languages; concurrency, integrity, privacy; modeling and measurement of access strategies; and dedicated processors, information retrieval and real time applications. Offered periodically.",
		"comp590": "Computer Architecture\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nThis course is an introduction to the internal structure of digital computers including design of gates, flip-fops, registers and memories to perform operations on numerical and other data represented in binary form; computer system analysis and design; organizational dependence on computations to be performed; and theoretical aspects of parallel and pipeline computation. Offered periodically.",
		"comp594": "Computer Networks\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nThis course provides an introduction to fundamental concepts in computer networks, including their design and implementation. Topics include network architectures and protocols, placing emphasis on protocol used in the Internet; routing; data link layer issues; multimedia networking; network security; and network management. Offered periodically.\n",
		"comp596": "Topics in Computer Science\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nIn this course, topics are chosen from program verification, formal semantics, formal language theory, concurrent programming, complexity or algorithms, programming language theory, graphics and other computer science topics. Repeatable for different topics. Offered as topics arise.",
		"comp598": " Computer Science Graduate Internship\n(3 credits)\nPrerequisite: Matriculation in the computer science master’s program; at least six credits of graduate-level course work in computer science (COMP); formal application required\nAn internship provides an opportunity to apply what has been learned in the classroom and allows the student to further professional skills. Faculty supervision allows for reflection on the internship experience and connects the applied portion of the academic study to other courses. Repeatable; may earn a maximum of six credits, however, only three credits can be used toward the degree. Graded on (P) Pass/(N) No Pass basis. Offered fall and spring semesters.\n",
	}
	insertStatement := "INSERT INTO COURSE (course_prefix, course_number, description) VALUES (?,?,?);"
	preppedStatement, err := database.Prepare(insertStatement)
	if err != nil {
		log.Fatal(err)
	}
	for course, desc := range sampleData {
		prefix := course[0:4]
		numVal := course[4:7]
		courseNum, err := strconv.Atoi(numVal)
		if err != nil {
			log.Fatal("ooops we must have mistyped", err)
		}
		preppedStatement.Exec(prefix, courseNum, desc)
	}

}

func addSampleStudents(database *sql.DB) {
	sampleNames := map[string]string{"John": "Santore", "Enping": "Li", "Michael": "Black",
		"Seikyung": "Jung", "Haleh": "Khojasteh", "Abdul": "Sattar", "Paul": "Kim", "Yiheng": "Liang"}
	statement := "INSERT INTO STUDENTS (banner_id, first_name, last_name, gpa, credits)" +
		"  VALUES (?, ?, ?, ?, ?);"
	prepped_statement, err := database.Prepare(statement)
	if err != nil {
		log.Fatal(err)
	}
	idNum := 1001
	for firstName, lastName := range sampleNames {
		randGPA := rand.Float32() + float32(rand.Intn(4))
		randCredits := rand.Intn(30)
		prepped_statement.Exec(idNum, firstName, lastName, randGPA, randCredits)
		idNum += 1
	}
}

func create_tables(database *sql.DB) {
	createStatement1 := "CREATE TABLE IF NOT EXISTS students(    " +
		"banner_id INTEGER PRIMARY KEY," +
		"first_name TEXT NOT NULL," +
		"last_name TEXT NOT NULL," +
		"gpa REAL DEFAULT 0," +
		"credits INTEGER DEFAULT 0);"
	create_course := "CREATE TABLE IF NOT EXISTS course(" +
		"    course_prefix TEXT NOT NULL," +
		"    course_number INTEGER NOT NULL," +
		"    cap INTEGER DEFAULT 20," +
		"    description TEXT," +
		"    PRIMARY KEY(course_prefix, course_number)    );"

	create_reg_statement := "CREATE TABLE IF NOT EXISTS class_list(" +
		"registration_id INTEGER PRIMARY KEY, course_prefix TEXT NOT NULL," +
		"course_number INTEGER NOT NULL," +
		"banner_id INTEGER NOT NULL," +
		"registration_date TEXT," +
		"FOREIGN KEY (banner_id) REFERENCES student (banner_id)" +
		"ON DELETE CASCADE ON UPDATE NO ACTION," +
		"FOREIGN KEY (course_prefix, course_number) REFERENCES courses (course_prefix, course_number)" +
		"ON DELETE CASCADE ON UPDATE NO ACTION" +
		");"

	database.Exec(createStatement1)
	database.Exec(create_course)
	database.Exec(create_reg_statement)
}
