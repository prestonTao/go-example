func main() {

	//var re  []Machine
	mm := mongo.NewMongoDBConn()
	mm.Connect("mongodb://redis:redis@localhost:27017/rrest")
	defer mm.Stop()
	mm.SetDBName("rrest")
	mm.SetTableName("machinea")
	mm.Insert(&Qe{
		Id_: bson.NewObjectId(),
		Usage:"eeeeea",
	})

	//tt := mm.FindAll(&re)
	//fmt.Printf("%v\n",tt)
	fmt.Printf("end\n")


}
