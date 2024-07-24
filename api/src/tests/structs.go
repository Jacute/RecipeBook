package main

type Test struct {
	route       string
	result      string
	contentType string
	statusCode  int
}

type MultipartJson struct {
	data      string
	fieldname string
	mimetype  string
}

type MultipartFile struct {
	filepath  string
	filename  string
	fieldname string
	mimetype  string
}

type MultipartData interface {
	isMultipartData()
}

func (MultipartJson) isMultipartData() {}
func (MultipartFile) isMultipartData() {}

type MultipartTest struct {
	route       string
	result      string
	contentType string
	statusCode  int
	parts       []MultipartData
}
