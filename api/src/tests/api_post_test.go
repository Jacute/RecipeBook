package main

import (
	"RecipeBookApi/handlers"
	"RecipeBookApi/middlewares"
	"RecipeBookApi/utils"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createMultipartRequest(route string, data []MultipartData, jwt_token string) (*http.Request, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	for _, part := range data {
		partHeaders := textproto.MIMEHeader{}
		if multipartFile, ok := part.(MultipartFile); ok {
			file, err := os.Open(multipartFile.filepath)
			if err != nil {
				return nil, err
			}
			partHeaders.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, multipartFile.fieldname, multipartFile.filename))
			partHeaders.Set("Content-Type", multipartFile.mimetype)
			fileForm, err := writer.CreatePart(partHeaders)
			if err != nil {
				return nil, err
			}
			if _, err := io.Copy(fileForm, file); err != nil {
				return nil, err
			}
			file.Close()
		} else if multipartJson, ok := part.(MultipartJson); ok {
			if err := writer.WriteField(multipartJson.fieldname, multipartJson.data); err != nil {
				return nil, err
			}
		}
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", route, body)
	if err != nil {
		return nil, err
	}
	cookie := http.Cookie{
		Name:  "auth-token",
		Value: jwt_token,
	}
	req.AddCookie(&cookie)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req, nil
}

func TestCreate(t *testing.T) {
	route := "/api/recipes"
	r := gin.Default()
	r.POST(route, middlewares.AuthMiddleware(), handlers.CreateRecipe)

	token := utils.GenerateToken("admin")
	test := &MultipartTest{
		route:       route,
		result:      "Recipe created successfully",
		contentType: "text/plain; charset=utf-8",
		statusCode:  201,
		parts: []MultipartData{
			MultipartJson{data: `{
				"name": "Мартини с водкой",
				"description": "Легендарный коктейль, ставший популярным благодаря Джеймсу Бонду.",
				"ingredients": "мартини «Драй» (сухой) – 10 мл; водка – 40 мл; лимонный сок – 5 мл; оливки – 1-2 штуки; лед.",
				"steps": "Рецепт: наполнить шейкер льдом, добавить водку, интенсивно взбалтывать 10 секунд. Долить мартини, еще раз взболтать, затем перелить полученный коктейль в бокал через стрейнер (барное ситечко). На заключительном этапе добавить несколько капель лимонного сока, украсить оливками или лимонной цедрой.",
				"is_private": false
				}`,
				fieldname: "recipe",
			},
			MultipartFile{
				filepath:  "data/koktejl-s-martini-i-vodkoj.jpg",
				filename:  "test.jpg",
				fieldname: "image",
				mimetype:  "image/jpeg",
			},
		},
	}
	req, err := createMultipartRequest(test.route, test.parts, token)
	if err != nil {
		t.Error(err.Error())
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, test.statusCode, w.Code)
	assert.Equal(t, test.contentType, w.Header().Get("Content-Type"))
	assert.Equal(t, test.result, w.Body.String())
}
