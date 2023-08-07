package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/notnulldev/e-com-svelte/ent/product"
	"log"
	"path"
	"strconv"
)

const FILES_FOLDER = "files"

func (api *AppApi) GetProductById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id", 0)
	if err != nil {
		return err
	}

	all, err := api.client.Product.Query().Where(product.IDEQ(id)).Only(ctx.Context())
	if err != nil {
		return err
	}
	return ctx.JSON(all)
}

// GetMyProducts TODO
func (api *AppApi) GetMyProducts(ctx *fiber.Ctx) error {
	return api.GetAllProducts(ctx)
}

func (api *AppApi) GetAllProducts(ctx *fiber.Ctx) error {
	all, err := api.client.Product.Query().All(ctx.Context())
	if err != nil {
		return err
	}
	return ctx.JSON(all)
}

func (api *AppApi) CreateProduct(ctx *fiber.Ctx) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		return err
	}

	pName := ctx.FormValue("productName")
	pDesc := ctx.FormValue("productDescription")
	pPreviewUrl := ctx.FormValue("productPreviewUrl")

	pPriceRaw := ctx.FormValue("productPrice")
	pPrice, err := strconv.ParseFloat(pPriceRaw, 64)
	if err != nil {
		log.Println(err.Error())
		return ctx.JSON(ErrorMessage{Message: "productPrice is in invalid format"})
	}

	pStockRaw := ctx.FormValue("productStock")
	pStock, err := strconv.ParseInt(pStockRaw, 10, 32)
	if err != nil {
		log.Println(err.Error())
		return ctx.JSON(ErrorMessage{Message: "productPrice is in invalid format"})
	}

	var filePaths []string

	for _, file := range form.File["files"] {
		fName := uuid.NewString()
		err := ctx.SaveFile(file, path.Join(FILES_FOLDER, fName))
		if err != nil {
			return err
		}
		filePaths = append(filePaths, fName)
	}

	created, err := api.client.Product.Create().
		SetName(pName).
		SetPrice(pPrice).
		SetPreviewURL(pPreviewUrl).
		SetDescription(pDesc).
		SetImages(filePaths).
		SetImagesStorage("local").
		SetStock(int(pStock)).
		SetStockReserved(0).
		SetCategories([]string{}).
		Save(ctx.Context())
	if err != nil {
		return err
	}

	return ctx.JSON(created)
}
