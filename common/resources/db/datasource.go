package db

import (
	"fmt"
)

type DataSource struct {
	host     string
	username string
	password string
	name     string
	port     int
	sslMode  string
	timezone string
}

func NewDataSource(
	host string,
	username string,
	password string,
	name string,
	port int,
	timezone string,
) *DataSource {
	sslMode := "disable"

	return &DataSource{
		host:     host,
		username: username,
		password: password,
		name:     name,
		port:     port,
		sslMode:  sslMode,
		timezone: timezone,
	}
}

func (ds *DataSource) BuildDataSourceName() string {
	dataSourceName := fmt.Sprintf("host=%s", ds.GetHost())
	dataSourceName += " " + fmt.Sprintf("user=%s", ds.GetUsername())
	dataSourceName += " " + fmt.Sprintf("password=%s", ds.GetPassword())
	dataSourceName += " " + fmt.Sprintf("dbname=%s", ds.GetName())
	dataSourceName += " " + fmt.Sprintf("port=%d", ds.GetPort())
	dataSourceName += " " + fmt.Sprintf("sslmode=%s", ds.GetSSLMode())
	dataSourceName += " " + fmt.Sprintf("TimeZone=%s", ds.GetTimezone())
	return dataSourceName
}

func (ds *DataSource) GetHost() string {
	return ds.host
}

func (ds *DataSource) GetName() string {
	return ds.name
}

func (ds *DataSource) GetPassword() string {
	return ds.password
}

func (ds *DataSource) GetPort() int {
	return ds.port
}

func (ds *DataSource) GetSSLMode() string {
	return ds.sslMode
}

func (ds *DataSource) GetTimezone() string {
	return ds.timezone
}

func (ds *DataSource) GetUsername() string {
	return ds.username
}
