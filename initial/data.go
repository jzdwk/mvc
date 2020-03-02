package initial

//init db
var InitialData = []string{

	//Create admin
	`	REPLACE INTO  user  ( id, name, email, phone, deleted, admin, last_login,create_time, update_time, password ) VALUES (1,'admin','admin@gmail.com','11111111111',0,1,now(),now(),now(),'efeb5259dc1ed1b9315fc0df4ded4a0e45fa66045a3e68a8ed88ffcd7c92956f30f91c1e8dd9ba91d6301655a3c54238b8fe');`,
}
