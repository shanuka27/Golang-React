package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/shanuka27/go-fiber-postgres/models"
	"github.com/shanuka27/go-fiber-postgres/storage"
	"gorm.io/gorm"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Item struct {
	Name         string `json:"name"`
	UnitPrice    string `json:"unit_price"`
	ItemCategory string `json:"item_category"`
}

type Invoice struct {
	Name        string `json:"name"`
	MobileNo    string `json:"mobile_no"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	BillingType string `json:"billing_type"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateItem(context *fiber.Ctx) error {
	item := Item{}

	err := context.BodyParser(&item)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Create(&item).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create item"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "item has been added"})
	return nil
}

func (r *Repository) DeleteItem(context *fiber.Ctx) error {
	itemModel := models.Items{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := r.DB.Delete(itemModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete item",
		})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "item delete successfully",
	})
	return nil
}

func (r *Repository) GetItems(context *fiber.Ctx) error {
	itemModels := &[]models.Items{}

	err := r.DB.Find(itemModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get items"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "items fetched successfully",
		"data":    itemModels,
	})
	return nil
}

func (r *Repository) GetItemByID(context *fiber.Ctx) error {

	id := context.Params("id")
	itemModel := &models.Items{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	fmt.Println("the ID is", id)

	err := r.DB.Where("id = ?", id).First(itemModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the book"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "item id fetched successfully",
		"data":    itemModel,
	})
	return nil
}

func (r *Repository) UpdateItem(context *fiber.Ctx) error {
	item := Item{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := context.BodyParser(&item)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Model(&models.Items{}).Where("id = ?", id).Updates(item).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not update item"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "item has been updated"})
	return nil
}

func (r *Repository) CreateInvoice(context *fiber.Ctx) error {
	invoice := Invoice{}

	err := context.BodyParser(&invoice)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Create(&invoice).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create invoice"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "invoice has been added"})
	return nil
}

func (r *Repository) DeleteInvoice(context *fiber.Ctx) error {
	invoiceModel := models.Invoices{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := r.DB.Delete(invoiceModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete invoice",
		})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "invoice delete successfully",
	})
	return nil
}

func (r *Repository) GetInvoices(context *fiber.Ctx) error {
	invoiceModels := &[]models.Invoices{}

	err := r.DB.Find(invoiceModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get invoices"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "invoices fetched successfully",
		"data":    invoiceModels,
	})
	return nil
}

func (r *Repository) UpdateInvoice(context *fiber.Ctx) error {
	invoice := Invoice{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := context.BodyParser(&invoice)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Model(&models.Invoices{}).Where("id = ?", id).Updates(invoice).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not update invoice"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "invoice has been updated"})
	return nil
}

func (r *Repository) GetInvoicesByID(context *fiber.Ctx) error {
	id := context.Params("id")
	invoiceModel := &models.Invoices{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	fmt.Println("the ID is", id)

	err := r.DB.Where("id = ?", id).First(invoiceModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the invoice"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "invoice id fetched successfully",
		"data":    invoiceModel,
	})
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_items", r.CreateItem)
	api.Delete("delete_item/:id", r.DeleteItem)
	api.Get("/get_item/:id", r.GetItemByID)
	api.Get("/item", r.GetItems)
	api.Put("/update_item/:id", r.UpdateItem)

	api.Post("/create_invoice", r.CreateInvoice)
	api.Delete("/delete_invoice/:id", r.DeleteInvoice)
	api.Get("/get_invoice/:id", r.GetInvoicesByID)
	api.Get("/invoice", r.GetInvoices)
	api.Put("/update_invoice/:id", r.UpdateInvoice)

}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("could not load the database")
	}

	err = models.MigrateItems(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	err = models.MigrateInvoices(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowMethods: "GET,POST,DELETE,PUT",
		AllowHeaders: "Content-Type, Authorization",
	}))

	r.SetupRoutes(app)
	app.Listen(":8081")

}
