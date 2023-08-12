package prototype

type Db struct {
  DbName string
}

var _ Connection = &Db{}

func (d *Db) ConnectionHealthCheck() bool{
  if d.DbName == "" {
    return false
  }
  return true
}

func (d *Db) ConnectionInfo() string {
  if (d.DbName == "") {
    return "sorry, the connection is lost"
  }
  return "DB connection"
}

func (d *Db) Clone() Connection {
  // create a deep copy 
  return &Db{
    DbName: d.DbName,
  }
}

func (d *Db) AbortConnection(){
  d.DbName = ""
}

func NewDb() *Db{
  return &Db{
    DbName: "postgres",
  }
}