package tapi

type Type string

const (
	STRING        Type = "string"
	STRINGARRAY        = "string[]"
	INT32              = "int32"
	INT32ARRAY         = "int32[]"
	INT64              = "int64"
	INT64ARRAY         = "int64[]"
	FLOAT              = "float"
	FLOATARRAY         = "float[]"
	BOOL               = "bool"
	BOOLARRAY          = "bool[]"
	GEOPOINT           = "geopoint"
	GEOPOINTARRAY      = "geopoint[]"
	OBJECT             = "object" // object is comparable to a go struct
	OBJECTARRAY        = "object[]"
	STRINGPTR          = "string*" // special type that can be string or []string
)
