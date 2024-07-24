package main

import (
	"RecipeBookApi/database"
	"RecipeBookApi/handlers"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var TestDB *sql.DB

func TestMain(m *testing.M) {
	database.InitDB()

	database.DB.Exec(`DROP TABLE IF EXISTS users;`)
	database.DB.Exec(`DROP TABLE IF EXISTS recipes;`)

	database.DB.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INT PRIMARY KEY AUTO_INCREMENT,
		username VARCHAR(64) UNIQUE NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		registered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`)

	database.DB.Exec(`CREATE TABLE IF NOT EXISTS recipes (
		id INT PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(255) UNIQUE NOT NULL,
		description TEXT NOT NULL,
		ingredients TEXT NOT NULL,
		steps TEXT NOT NULL,
		image_path TEXT NOT NULL,
		creator_id INTEGER NOT NULL,
		is_private BOOLEAN DEFAULT TRUE,

		FOREIGN KEY (creator_id) REFERENCES users(id)
	);`)
	database.DB.Exec("INSERT INTO users (username, email, password) VALUES ('admin', 'admin@example.com', '$2b$10$HUTcRrfAF9AeenwBy7Acj.V6WdeQZH2zgoLRHhqyrGHuiVcoVVliO');")
	database.DB.Exec("INSERT INTO recipes (name, description, ingredients, steps, image_path, creator_id, is_private) VALUES ('Борщ', 'Традиционный украинский суп с красной свеклой и мясом', '1 свекла, 2 картошки, 50 грамм соли, 300 грамм говядины, 1 морковь, 1 луковица, 2 зубчика чеснока, 1 столовая ложка томатной пасты, 1 лавровый лист, зелень по вкусу', 'Нарежьте свеклу, картошку, морковь и лук мелкими кубиками.\nОбжарьте мясо в большой кастрюле до золотистой корки.\nДобавьте нарезанные овощи, чеснок, томатную пасту и лавровый лист в кастрюлю.\nЗалейте кипящей водой, варите до готовности овощей и мяса.\nПодавайте с зеленью и сметаной.', 'images/413464-9c183c2f2550e70d5b6ad8165547800d.jpg', 1, false);")
	database.DB.Exec("INSERT INTO recipes (name, description, ingredients, steps, image_path, creator_id, is_private) VALUES ('Плов', 'Классическое блюдо восточной кухни из риса и мяса', '500 грамм баранины, 2 стакана риса, 2 моркови, 1 луковица, 1 головка чеснока, 100 грамм растительного масла, 1 столовая ложка зиры, соль и перец по вкусу', 'Нарежьте мясо и морковь ломтиками.\nОбжарьте мясо с морковью в глубокой сковороде до золотистой корки.\nДобавьте нарезанный лук и чеснок, обжаривайте до прозрачности.\nДобавьте рис и зиру, перемешайте.\nЗалейте кипящей водой, варите на медленном огне до готовности.\nПодавайте горячим.', 'images/e26b0233.jpg', 1, false);")

	exitVal := m.Run()
	database.DB.Close()
	os.RemoveAll("images")
	os.Exit(exitVal)
}

func TestGetRecipes(t *testing.T) {
	route := "/api/recipes"
	r := gin.Default()
	r.GET(route, handlers.GetRecipes)

	req, _ := http.NewRequest("GET", route, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Equal(t, `[{"id":1,"name":"Борщ","description":"Традиционный украинский суп с красной свеклой и мясом","ingredients":"1 свекла, 2 картошки, 50 грамм соли, 300 грамм говядины, 1 морковь, 1 луковица, 2 зубчика чеснока, 1 столовая ложка томатной пасты, 1 лавровый лист, зелень по вкусу","steps":"Нарежьте свеклу, картошку, морковь и лук мелкими кубиками.\nОбжарьте мясо в большой кастрюле до золотистой корки.\nДобавьте нарезанные овощи, чеснок, томатную пасту и лавровый лист в кастрюлю.\nЗалейте кипящей водой, варите до готовности овощей и мяса.\nПодавайте с зеленью и сметаной.","image_path":"images/413464-9c183c2f2550e70d5b6ad8165547800d.jpg","creator_username":"admin"},{"id":2,"name":"Плов","description":"Классическое блюдо восточной кухни из риса и мяса","ingredients":"500 грамм баранины, 2 стакана риса, 2 моркови, 1 луковица, 1 головка чеснока, 100 грамм растительного масла, 1 столовая ложка зиры, соль и перец по вкусу","steps":"Нарежьте мясо и морковь ломтиками.\nОбжарьте мясо с морковью в глубокой сковороде до золотистой корки.\nДобавьте нарезанный лук и чеснок, обжаривайте до прозрачности.\nДобавьте рис и зиру, перемешайте.\nЗалейте кипящей водой, варите на медленном огне до готовности.\nПодавайте горячим.","image_path":"images/e26b0233.jpg","creator_username":"admin"}]`, w.Body.String())
}

func TestGetRecipeByID(t *testing.T) {
	route := "/api/recipes/:id"
	r := gin.Default()
	r.GET(route, handlers.GetRecipeByID)

	testRightRoutes := []Test{
		{
			"/api/recipes/1",
			`{"id":1,"name":"Борщ","description":"Традиционный украинский суп с красной свеклой и мясом","ingredients":"1 свекла, 2 картошки, 50 грамм соли, 300 грамм говядины, 1 морковь, 1 луковица, 2 зубчика чеснока, 1 столовая ложка томатной пасты, 1 лавровый лист, зелень по вкусу","steps":"Нарежьте свеклу, картошку, морковь и лук мелкими кубиками.\nОбжарьте мясо в большой кастрюле до золотистой корки.\nДобавьте нарезанные овощи, чеснок, томатную пасту и лавровый лист в кастрюлю.\nЗалейте кипящей водой, варите до готовности овощей и мяса.\nПодавайте с зеленью и сметаной.","image_path":"images/413464-9c183c2f2550e70d5b6ad8165547800d.jpg","creator_username":"admin"}`,
			"application/json; charset=utf-8",
			200,
		},
		{
			"/api/recipes/2",
			`{"id":2,"name":"Плов","description":"Классическое блюдо восточной кухни из риса и мяса","ingredients":"500 грамм баранины, 2 стакана риса, 2 моркови, 1 луковица, 1 головка чеснока, 100 грамм растительного масла, 1 столовая ложка зиры, соль и перец по вкусу","steps":"Нарежьте мясо и морковь ломтиками.\nОбжарьте мясо с морковью в глубокой сковороде до золотистой корки.\nДобавьте нарезанный лук и чеснок, обжаривайте до прозрачности.\nДобавьте рис и зиру, перемешайте.\nЗалейте кипящей водой, варите на медленном огне до готовности.\nПодавайте горячим.","image_path":"images/e26b0233.jpg","creator_username":"admin"}`,
			"application/json; charset=utf-8",
			200,
		},
	}

	for _, test := range testRightRoutes {
		req, _ := http.NewRequest("GET", test.route, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		ok1 := assert.Equal(t, test.statusCode, w.Code)
		ok2 := assert.Equal(t, test.contentType, w.Header().Get("Content-Type"))
		ok3 := assert.Equal(t, test.result, w.Body.String())
		if !ok1 {
			t.Errorf("Route: %s, Status Code: %d; Want: %d", test.route, w.Code, test.statusCode)
		}
		if !ok2 {
			t.Errorf("Route: %s, Content-Type: %s; Want: %s", test.route, w.Header().Get("Content-Type"), test.contentType)
		}
		if !ok3 {
			t.Errorf("Route: %s, Response: %s; Want: %s", test.route, w.Body.String(), test.result)
		}
	}

	testFailRoutes := []Test{
		{
			"/api/recipes/+-,,.dsa123",
			"Bad ID",
			"text/plain; charset=utf-8",
			400,
		},
		{
			"/api/recipes/555555",
			"Recipe not found",
			"text/plain; charset=utf-8",
			404,
		},
	}

	for _, test := range testFailRoutes {
		req, _ := http.NewRequest("GET", test.route, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		ok1 := assert.Equal(t, test.statusCode, w.Code)
		ok2 := assert.Equal(t, "text/plain; charset=utf-8", w.Header().Get("Content-Type"))
		ok3 := assert.Equal(t, test.result, w.Body.String())
		if !ok1 {
			t.Errorf("Route: %s, Status Code: %d; Want: %d", test.route, w.Code, test.statusCode)
		}
		if !ok2 {
			t.Errorf("Route: %s, Content-Type: %s; Want: %s", test.route, w.Header().Get("Content-Type"), test.contentType)
		}
		if !ok3 {
			t.Errorf("Route: %s, Response: %s; Want: %s", test.route, w.Body.String(), test.result)
		}
	}
}
