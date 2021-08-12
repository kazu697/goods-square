package handler

import "github.com/gorilla/schema"

var schemaDecoder = func() *schema.Decoder {
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	return decoder
}()

var queryDecoder = func() *schema.Decoder {
	return schemaDecoder
}()
