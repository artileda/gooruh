package data

type SQLTable struct {
	Name    string
	Columns []SQLCol
}

func (st SQLTable) New(name string) {
	st.Name = name
}

func (st SQLTable) GenerateQueries() string {
	queries := "CREATE TABLE `" + st.Name + "` ( \n"
	for _, column := range st.Columns {
		queries += column.GenerateQueries() + "\n"
	}
	queries += ")"
	return queries
}

type SQLCol struct {
	Name     string
	DataType string
	Length   string
	NotNull  bool
	AutoInc  bool
	Primary  bool
	Unique   bool
	Foreign  bool
}

func (s SQLCol) New() {
	s.Name = ""
	s.DataType = "string"
	s.Length = "100"
	s.NotNull = false
	s.AutoInc = false
	s.Primary = false
	s.Unique = false
	s.Foreign = false
}

func GoToSQLTypes(types string) string {
	switch types {
	case "string":
		return "TEXT"
	case "uint":
		return "INTERGER"
	case "int":
		return "INTEGER"
	case "bool":
		return "BOOLEAN"
	case "date":
		return "DATE"
	default:
		return "UNKTYPE"
	}
}

func (s SQLCol) GenerateQueries() string {
	queries := "   "

	queries += ("`" + s.Name + "` ")
	queries += (GoToSQLTypes(s.DataType) + "(" + s.Length + ") ")

	if s.NotNull {
		queries += " NOT NULL "
	}

	if s.AutoInc {
		queries += " AUTO INCREMENT "

	}

	if s.Primary {
		queries += " PRIMARY KEY "
	}

	if s.Unique {
		queries += " UNIQUE "
	}

	if s.Foreign {

	}

	return queries
}

func ParseParameterToSQL(input []string) SQLCol {
	dup := SQLCol{}
	dup.New()

	dup.Name = input[0]
	dup.DataType = input[1]
	dup.Length = input[2]

	for _, element := range input {
		switch element {
		case "primary":
			dup.Primary = true
		case "notnull":
			dup.NotNull = true
		case "autoinc":
			dup.AutoInc = true
		case "unique":
			dup.Unique = true
		case "foreign":
			dup.Foreign = true
		}
	}

	return dup
}
