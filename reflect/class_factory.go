package main

/*
	类型工厂
*/

func Factory() {

}

func New() interface{} {
	return new(tx)
}

type tx struct {
	name string
}
