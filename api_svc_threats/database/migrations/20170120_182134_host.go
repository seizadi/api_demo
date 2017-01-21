package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Host_20170120_182134 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Host_20170120_182134{}
	m.Created = "20170120_182134"
	migration.Register("Host_20170120_182134", m)
}

// Run the migrations
func (m *Host_20170120_182134) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE host(id serial primary key,fqdn TEXT NOT NULL)")
}

// Reverse the migrations
func (m *Host_20170120_182134) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE host")
}
