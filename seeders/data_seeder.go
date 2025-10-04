package seeders

import (
	"ecommerce-backend/models"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// RunSeeder inserts dummy data (users, categories, products, example order)
func RunSeeder(db *gorm.DB) {
	// -----------------------
	// Users (with bcrypt)
	// -----------------------
	users := []models.User{
		{Name: "Admin", Email: "admin@example.com", Role: "admin"},
		{Name: "Alice", Email: "alice@example.com", Role: "customer"},
		{Name: "Bob", Email: "bob@example.com", Role: "customer"},
	}

	for _, u := range users {
		var existing models.User
		if err := db.Where("email = ?", u.Email).First(&existing).Error; err == gorm.ErrRecordNotFound {
			hashed, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost) // default pwd
			u.Password = string(hashed)
			db.Create(&u)
		}
	}

	// -----------------------
	// Categories
	// -----------------------
	cats := []models.Category{
		{Name: "Electronics"},
		{Name: "Fashion"},
		{Name: "Books"},
		{Name: "Groceries"},
	}

	for _, c := range cats {
		db.FirstOrCreate(&c, models.Category{Name: c.Name})
	}

	// -----------------------
	// Products
	// -----------------------
	// Note: assume category IDs 1..n exist; if not, GORM will set accordingly
	products := []models.Product{
	{Name: "Laptop Gaming XYZ", Price: 15000000, Stock: 8, CategoryID: 1},
	{Name: "Smartphone Ultra", Price: 9000000, Stock: 15, CategoryID: 1},
	{Name: "T-Shirt Casual", Price: 120000, Stock: 100, CategoryID: 2},
	{Name: "Novel: The Adventure", Price: 90000, Stock: 40, CategoryID: 3},
	{Name: "Rice 5kg", Price: 120000, Stock: 50, CategoryID: 4},
	}

	for _, p := range products {
		db.FirstOrCreate(&p, models.Product{Name: p.Name})
	}

	// -----------------------
	// Example Order (for Alice)
	// -----------------------
	// only create a sample order if none exist for alice@example.com
	var alice models.User
	if err := db.Where("email = ?", "alice@example.com").First(&alice).Error; err == nil {
		var count int64
		db.Model(&models.Order{}).Where("user_id = ?", alice.ID).Count(&count)
		if count == 0 {
			// fetch two products
			var prod1, prod2 models.Product
			db.Where("name = ?", "Laptop Gaming XYZ").First(&prod1)
			db.Where("name = ?", "T-Shirt Casual").First(&prod2)

			items := []models.OrderItem{
				{ProductID: prod1.ID, Quantity: 1, Subtotal: prod1.Price * 1},
				{ProductID: prod2.ID, Quantity: 2, Subtotal: prod2.Price * 2},
			}
			total := items[0].Subtotal + items[1].Subtotal

			order := models.Order{
				UserID: alice.ID,
				Total:  total,
				Status: "pending",
				Items:  items,
			}
			db.Create(&order)

			// reduce product stock
			prod1.Stock -= 1
			prod2.Stock -= 2
			db.Save(&prod1)
			db.Save(&prod2)
			fmt.Println("✅ Sample order created for alice@example.com")
		}
	}

	fmt.Println("✅ Seeder finished.")
}

func Seed(db *gorm.DB) {
	// Admin user
	hashed, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	admin := models.User{Name: "Admin", Email: "admin@mail.com", Password: string(hashed), Role: "admin"}
	db.FirstOrCreate(&admin, models.User{Email: admin.Email})

	// Customer user
	hashed2, _ := bcrypt.GenerateFromPassword([]byte("budi123"), bcrypt.DefaultCost)
	user := models.User{Name: "Budi", Email: "budi@mail.com", Password: string(hashed2), Role: "customer"}
	db.FirstOrCreate(&user, models.User{Email: user.Email})

	// Categories
	books := models.Category{Name: "Books"}
	electronics := models.Category{Name: "Electronics"}
	groceries := models.Category{Name: "Groceries"}
	db.FirstOrCreate(&books, books)
	db.FirstOrCreate(&electronics, electronics)
	db.FirstOrCreate(&groceries, groceries)

	// Products
	db.FirstOrCreate(&models.Product{Name: "Laptop ASUS", Price: 15000000, Stock: 10, CategoryID: electronics.ID})
	db.FirstOrCreate(&models.Product{Name: "Smartphone Samsung", Price: 5000000, Stock: 15, CategoryID: electronics.ID})
	db.FirstOrCreate(&models.Product{Name: "Novel The Adventure", Price: 90000, Stock: 40, CategoryID: books.ID})
	db.FirstOrCreate(&models.Product{Name: "Rice 5kg", Price: 120000, Stock: 50, CategoryID: groceries.ID})
}


func Seed1(db *gorm.DB) {
	// seed categories
	categories := []models.Category{
		{Name: "Electronics"},
		{Name: "Fashion"},
	}
	for _, c := range categories {
		db.FirstOrCreate(&c, models.Category{Name: c.Name})
	}

	// seed products
	products := []models.Product{
		{Name: "Smartphone Ultra", Price: 9000000, Stock: 10, CategoryID: 1},
		{Name: "T-Shirt Casual", Price: 120000, Stock: 100, CategoryID: 2},
	}
	for _, p := range products {
		db.FirstOrCreate(&p, models.Product{Name: p.Name})
	}

	// seed admin user
	admin := models.User{Name: "Admin", Email: "admin@mail.com", Password: "hashed_password", Role: "admin"}
	db.FirstOrCreate(&admin, models.User{Email: admin.Email})
}
