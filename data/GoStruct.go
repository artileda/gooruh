package data

func SQLToStruct(st SQLTable) string {
	structString := "type " + st.Name + " struct {\n"
	for _, element := range st.Columns {
		structString += "  " + element.Name + " " + element.DataType + " \n"
	}
	structString += "}\n"
	return structString
}
