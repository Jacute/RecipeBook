package main

import (
	"RecipeBookApi/database"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AddTestRecipes() {
	database.DB.Exec(`CREATE TABLE IF NOT EXISTS recipes (
		id INT PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(255) UNIQUE NOT NULL,
		description TEXT NOT NULL,
		ingredients TEXT NOT NULL,
		steps TEXT NOT NULL,
		image_path TEXT NOT NULL,
		creator_id INTEGER,
		is_private BOOLEAN DEFAULT TRUE
	);`)
	database.DB.Exec("INSERT INTO recipes (name, description, ingredients, steps, image_path, creator_id, is_private) VALUES ('Борщ', 'Традиционный украинский суп с красной свеклой и мясом', '1 свекла, 2 картошки, 50 грамм соли, 300 грамм говядины, 1 морковь, 1 луковица, 2 зубчика чеснока, 1 столовая ложка томатной пасты, 1 лавровый лист, зелень по вкусу', 'Нарежьте свеклу, картошку, морковь и лук мелкими кубиками.\nОбжарьте мясо в большой кастрюле до золотистой корки.\nДобавьте нарезанные овощи, чеснок, томатную пасту и лавровый лист в кастрюлю.\nЗалейте кипящей водой, варите до готовности овощей и мяса.\nПодавайте с зеленью и сметаной.', 'images/413464-9c183c2f2550e70d5b6ad8165547800d.jpg', null, false);")
	database.DB.Exec("INSERT INTO recipes (name, description, ingredients, steps, image_path, creator_id, is_private) VALUES ('Плов', 'Классическое блюдо восточной кухни из риса и мяса', '500 грамм баранины, 2 стакана риса, 2 моркови, 1 луковица, 1 головка чеснока, 100 грамм растительного масла, 1 столовая ложка зиры, соль и перец по вкусу', 'Нарежьте мясо и морковь ломтиками.\nОбжарьте мясо с морковью в глубокой сковороде до золотистой корки.\nДобавьте нарезанный лук и чеснок, обжаривайте до прозрачности.\nДобавьте рис и зиру, перемешайте.\nЗалейте кипящей водой, варите на медленном огне до готовности.\nПодавайте горячим.', 'images/e26b0233.jpg', null, false);")
}

func TestGetRecipes(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	database.InitDB()
	defer database.DB.Close()
	AddTestRecipes()

	req, _ := http.NewRequest("GET", "/api/recipes", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Equal(t, `[{"id":1,"name":"Борщ","description":"Традиционный украинский суп с красной свеклой и мясом","ingredients":"1 свекла, 2 картошки, 50 грамм соли, 300 грамм говядины, 1 морковь, 1 луковица, 2 зубчика чеснока, 1 столовая ложка томатной пасты, 1 лавровый лист, зелень по вкусу","steps":"Нарежьте свеклу, картошку, морковь и лук мелкими кубиками.\nОбжарьте мясо в большой кастрюле до золотистой корки.\nДобавьте нарезанные овощи, чеснок, томатную пасту и лавровый лист в кастрюлю.\nЗалейте кипящей водой, варите до готовности овощей и мяса.\nПодавайте с зеленью и сметаной.","image_path":"images/413464-9c183c2f2550e70d5b6ad8165547800d.jpg"},{"id":2,"name":"Плов","description":"Классическое блюдо восточной кухни из риса и мяса","ingredients":"500 грамм баранины, 2 стакана риса, 2 моркови, 1 луковица, 1 головка чеснока, 100 грамм растительного масла, 1 столовая ложка зиры, соль и перец по вкусу","steps":"Нарежьте мясо и морковь ломтиками.\nОбжарьте мясо с морковью в глубокой сковороде до золотистой корки.\nДобавьте нарезанный лук и чеснок, обжаривайте до прозрачности.\nДобавьте рис и зиру, перемешайте.\nЗалейте кипящей водой, варите на медленном огне до готовности.\nПодавайте горячим.","image_path":"images/e26b0233.jpg"}]`, w.Body.String())
}

func TestGetRecipeByID(t *assert.TestingT) {
	// TODO: Implement test for GetRecipeByID
}
