package db

func SearchProfile(Id string) error {
	/*ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("User")
	var profile models.User

	objId, _ := primitive.ObjectIDFromHex(Id)

	condition := bson.M{
		"_id": objId,
	}

	err := col.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""
	if err != nil {
		fmt.Println("An error ocurred while seraching the user " + err.Error() + profile.Email + objId.Hex())
		return profile, err
	}*/

	return nil
}
