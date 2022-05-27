package models

type DataBaseObjectType string

const (
	DATABASE                       DataBaseObjectType = "DATABASE"
	TABLE                          DataBaseObjectType = "TABLE"
	USER                           DataBaseObjectType = "USER"
	MACRO                          DataBaseObjectType = "MACRO"
	VIEW                           DataBaseObjectType = "VIEW"
	STANDARD_FUNCTION              DataBaseObjectType = "STANDARD_FUNCTION"
	TRIGGER                        DataBaseObjectType = "TRIGGER"
	AUTHORIZATION                  DataBaseObjectType = "AUTHORIZATION"
	HASH_INDEX                     DataBaseObjectType = "HASH_INDEX"
	JAR                            DataBaseObjectType = "JAR"
	JOIN_INDEX                     DataBaseObjectType = "JOIN_INDEX"
	STORED_PROCEDURE               DataBaseObjectType = "STORED_PROCEDURE"
	AGGREGATE_FUNCTION             DataBaseObjectType = "AGGREGATE_FUNCTION"
	COMBINED_AGGREGATE_FUNCTIONS   DataBaseObjectType = "COMBINED_AGGREGATE_FUNCTIONS"
	EXTERNAL_PROCEDURE             DataBaseObjectType = "EXTERNAL_PROCEDURE"
	GLOP_SET                       DataBaseObjectType = "GLOP_SET"
	INSTANCE_OR_CONSTRUCTOR_METHOD DataBaseObjectType = "INSTANCE_OR_CONSTRUCTOR_METHOD"
	JOURNAL                        DataBaseObjectType = "JOURNAL"
	NO_PI_TABLE                    DataBaseObjectType = "NO_PI_TABLE"
	ORDERED_ANALYTICAL_FUNCTION    DataBaseObjectType = "ORDERED_ANALYTICAL_FUNCTION"
	QUEUE_TABLE                    DataBaseObjectType = "QUEUE_TABLE"
	TABLE_FUNCTION                 DataBaseObjectType = "TABLE_FUNCTION"
	TABLE_OPERATOR                 DataBaseObjectType = "TABLE_OPERATOR"
	USER_DEFINED_METHOD            DataBaseObjectType = "USER_DEFINED_METHOD"
	USER_DEFINED_DATA_TYPE         DataBaseObjectType = "USER_DEFINED_DATA_TYPE"
	SERVER_OBJECT                  DataBaseObjectType = "SERVER_OBJECT"
	USER_INSTALLED_FILE            DataBaseObjectType = "USER_INSTALLED_FILE"
)
