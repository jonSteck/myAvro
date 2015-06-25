package myAvro

import (
	"github.com/stealthly/go-avro"
	"bytes"
)

var singleValueSchema = `{
	"type": "record",
	"name": "Attributes",
	"fields": [
		{"name": "value", "type": "string"}
	]
}`
//var schema, _ = avro.ParseSchema(singleValueSchema) 
//var writer = avro.NewSpecificDatumWriter()
//writer.SetSchema(schema)
//var buffer = new(bytes.Buffer)
//var encoder = avro.NewBinaryEncoder(buffer)

type Attributes struct {
	Value string
}

//Must initialize before writing
/*func EncodeInit() {
	//parse schema
	schema, err := avro.ParseSchema(singleValueSchema)
	if err != nil
		panic(err)
	}

	//init writer with schema
	writer.SetSchema(schema)
	
}*/

//encodes string to []byte using avro, maybe only initialize once?
func MyEncode(urlAttributes string) []byte {
	schema, err := avro.ParseSchema(singleValueSchema)
	if err != nil {
		panic(err)
	}

	attributes := new(Attributes)
	attributes.Value = urlAttributes

	writer := avro.NewSpecificDatumWriter()
	writer.SetSchema(schema)

	buffer := new(bytes.Buffer)
	encoder := avro.NewBinaryEncoder(buffer)

	writer.Write(attributes, encoder)
	//test
	return buffer.Bytes()
} 
