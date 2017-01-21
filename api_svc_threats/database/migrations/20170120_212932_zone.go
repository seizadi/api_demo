package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Zone_20170120_212932 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Zone_20170120_212932{}
	m.Created = "20170120_212932"
	migration.Register("Zone_20170120_212932", m)
}

// Run the migrations
func (m *Zone_20170120_212932) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE zone(id serial primary key,fqdn TEXT NOT NULL)")
}

// Reverse the migrations
func (m *Zone_20170120_212932) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE zone")
}
