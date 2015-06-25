package myAvro

import (
        "github.com/stealthly/go-avro"
)

//decode avro []byte using schema defined in encode.go
func MyDecode(coded []byte) string {
	schema, err := avro.ParseSchema(singleValueSchema)
	if err != nil {
		panic(err)
	}
	reader := avro.NewSpecificDatumReader()
	reader.SetSchema(schema)
	decoder := avro.NewBinaryDecoder(coded)
	decoded := new(Attributes)
	err = reader.Read(decoded, decoder)
        if err != nil {
                panic(err)
        }
	return decoded.Value
}
