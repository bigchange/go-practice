mockgen:
	mockgen -source testcase/test/interface.go -package mockinterfce -destination=testcase/test/mockgen/mockinterface.go
	mockgen -source testcase/db/interface.go -package mockdb -destination=testcase/db/mockgen/mockdb.go