package data

func SQLToStruct(st PSQLTable) string {
	structString := "type " + st.Name + " struct {\n"
	for _, element := range st.Columns {
		var datatype = "*" + element.DataType

		if element.NotNull {
			datatype = element.DataType
		}
		structString += "  " + element.Name + " " + datatype + " \n"
	}
	structString += "}\n"
	return structString
}
