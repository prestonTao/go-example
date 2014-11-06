func main(){
	mm := mongo.NewMongoDBConn()
	mm.Connect("mongodb://redis:redis@localhost:27017/rrest")
	defer mm.Stop()
	mm.SetDBName("rrest")
	mm.SetTableName("machinea")
	//mm.Delete(bson.M{"_id": bson.ObjectIdHex("52735a28ad289175e9000001")})
	//mm.Delete(bson.M{"usage":"eeeee"})
	//mm.DeleteID(bson.ObjectIdHex("527359b4ad289175c3000001"))
	fmt.Printf("end\n")
}