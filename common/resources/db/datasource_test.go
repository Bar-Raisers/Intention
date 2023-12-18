package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildDataSourceName(t *testing.T) {
	// Given
	host := "host"
	username := "user"
	password := "password"
	name := "name"
	port := 5432
	timezone := "timezone"

	dataSource := NewDataSource(
		host,
		username,
		password,
		name,
		port,
		timezone,
	)

	// When
	dataSourceName := dataSource.BuildDataSourceName()

	// Then
	expectedDataSourceName := "host=host user=user password=password dbname=name port=5432 sslmode=disable TimeZone=timezone"
	assert.Equal(t, expectedDataSourceName, dataSourceName)
}

func TestGetHost(t *testing.T) {
	// Given
	expectedHost := "test_host"
	dataSource := &DataSource{
		host: expectedHost,
	}

	// When
	host := dataSource.GetHost()

	// Then
	assert.Equal(t, expectedHost, host)
}

func TestGetName(t *testing.T) {
	// Given
	expectedName := "test_name"
	dataSource := &DataSource{
		name: expectedName,
	}

	// When
	name := dataSource.GetName()

	// Then
	assert.Equal(t, expectedName, name)
}

func TestGetPassword(t *testing.T) {
	// Given
	expectedPassword := "test_password"
	dataSource := &DataSource{
		password: expectedPassword,
	}

	// When
	password := dataSource.GetPassword()

	// Then
	assert.Equal(t, expectedPassword, password)
}

func TestGetPort(t *testing.T) {
	// Given
	expectedPort := 1000
	dataSource := &DataSource{
		port: expectedPort,
	}

	// When
	port := dataSource.GetPort()

	// Then
	assert.Equal(t, expectedPort, port)
}

func TestGetSSLMode(t *testing.T) {
	// Given
	expectedSSLMode := "test_sslmode"
	dataSource := &DataSource{
		sslMode: expectedSSLMode,
	}

	// When
	sslMode := dataSource.GetSSLMode()

	// Then
	assert.Equal(t, expectedSSLMode, sslMode)
}

func TestGetTimezone(t *testing.T) {
	// Given
	expectedTimezone := "test_timezone"
	dataSource := &DataSource{
		timezone: expectedTimezone,
	}

	// When
	timezone := dataSource.GetTimezone()

	// Then
	assert.Equal(t, expectedTimezone, timezone)
}

func TestGetUsername(t *testing.T) {
	// Given
	expectedUsername := "test_username"
	dataSource := &DataSource{
		username: expectedUsername,
	}

	// When
	username := dataSource.GetUsername()

	// Then
	assert.Equal(t, expectedUsername, username)
}

func TestNewDataSource(t *testing.T) {
	// Given
	expectedHost := "test_host"
	expectedUsername := "test_username"
	expectedPassword := "test_password"
	expectedName := "test_name"
	expectedPort := 1000
	expectedTimezone := "test_timezone"
	expectedSSLMode := "disable"

	// When
	dataSource := NewDataSource(
		expectedHost,
		expectedUsername,
		expectedPassword,
		expectedName,
		expectedPort,
		expectedTimezone,
	)

	// Then
	assert.Equal(t, expectedHost, dataSource.host)
	assert.Equal(t, expectedUsername, dataSource.username)
	assert.Equal(t, expectedPassword, dataSource.password)
	assert.Equal(t, expectedName, dataSource.name)
	assert.Equal(t, expectedPort, dataSource.port)
	assert.Equal(t, expectedTimezone, dataSource.timezone)
	assert.Equal(t, expectedSSLMode, dataSource.sslMode)
}
