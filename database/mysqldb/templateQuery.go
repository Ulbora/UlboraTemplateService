package mysqldb

// ContentQuery is a content select query
const (
	InsertTemplateQuery               = "INSERT INTO template (name, application, active, client_id) VALUES (?, ?, ?, ?) "
	UpdateTemplateQuery               = "UPDATE template set active = ? where id = ? and client_id = ? "
	UpdateClearNonActiveTemplateQuery = "UPDATE template set active = ? where id != ? and client_id = ? "
	TemplateGetActiveQuery            = "select * from template WHERE active = 1 and application = ? and client_id = ?"
	TemplateGetByClientQuery          = "select * from template WHERE application = ? and client_id = ? order by name"
	TemplateDeleteQuery               = "DELETE from template WHERE  id = ? and client_id = ?"
	ConnectionTestQuery               = "SELECT count(*) from template"
)
