package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type ClientName_20170120_212813 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ClientName_20170120_212813{}
	m.Created = "20170120_212813"
	migration.Register("ClientName_20170120_212813", m)
}

// Run the migrations
func (m *ClientName_20170120_212813) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE client_name(id serial primary key,name TEXT NOT NULL, ip TEXT NOT NULL)")
}

// Reverse the migrations
func (m *ClientName_20170120_212813) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE client_name")
}
