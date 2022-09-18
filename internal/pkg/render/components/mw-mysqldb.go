package components

type MwMySQLDb struct {
	Instance string `json:"instance"`

	DBName string `json:"dbname"`
}

func NewMwMySQLDb(instance, dbname string) Component {
	return &MwMySQLDb{
		instance,
		dbname,
	}
}

func (m *MwMySQLDb) GetName() string {
	return m.DBName
}

func (m *MwMySQLDb) GetType() string {
	return "k-mw-mysqldb"
}
