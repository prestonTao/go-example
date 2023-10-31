package main

import (
	// "../mgo"
	"fmt"
)

type MongoDBConn struct {
	session   *mgo.Session
	dbName    string
	tableName string
}

func NewMongoDBConn() *MongoDBConn {
	return &MongoDBConn{}
}

func (m *MongoDBConn) Connect(url string) *mgo.Session {
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	m.session = session
	session.SetMode(mgo.Monotonic, true)
	return m.session
}

func (m *MongoDBConn) Stop() {
	m.session.Close()
}

func (m *MongoDBConn) SetDBName(dbName string) {
	m.dbName = dbName
}

func (m *MongoDBConn) SetTableName(tableName string) {
	m.tableName = tableName
}

/*only find One
 */
func (m *MongoDBConn) FindOne(query interface{}, result interface{}) interface{} {
	Collection := m.session.DB(m.dbName).C(m.tableName)
	err := Collection.Find(query).One(&result)
	if err != nil {
		panic(err)
		return nil
	}
	return result
}

/*find all
 */
func (m *MongoDBConn) FindAll(result interface{}) interface{} {
	Collection := m.session.DB(m.dbName).C(m.tableName)
	err := Collection.Find(nil).All(result)
	if err != nil {
		panic(err)
		fmt.Printf("Log Info======:\n", err)
		return nil
	}
	return result
}

/*find one or more from condition
 */
func (m *MongoDBConn) Find(query interface{}, result interface{}) interface{} {
	Collection := m.session.DB(m.dbName).C(m.tableName)
	err := Collection.Find(query).All(result)
	if err != nil {
		panic(err)
		fmt.Printf("Log Info======:\n", err)
		return nil
	}
	return result
}

/*insert data*/
func (m *MongoDBConn) Insert(data interface{}) {
	Collection := m.session.DB(m.dbName).C(m.tableName)
	err := Collection.Insert(data)
	if err != nil {
		panic(err)
		fmt.Printf("Log Info======:\n", err)
	}
}

func (m *MongoDBConn) Update(selector interface{}, change interface{}) {
	Collection := m.session.DB(m.dbName).C(m.tableName)
	err := Collection.Update(selector, change)
	if err != nil {
		panic(err)
		fmt.Printf("Log Info======:\n", err)
	}
}

func (m *MongoDBConn) Delete(data interface{}) {
	Collection := m.session.DB(m.dbName).C(m.tableName)
	err := Collection.Remove(data)
	if err != nil {
		panic(err)
		fmt.Printf("Log Info======:\n", err)
	}
}

func (m *MongoDBConn) DeleteID(id interface{}) {
	Collection := m.session.DB(m.dbName).C(m.tableName)
	err := Collection.RemoveId(id)
	if err != nil {
		panic(err)
		fmt.Printf("Log Info======:\n", err)
	}
}
