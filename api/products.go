package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/notnulldev/e-com-svelte/ent/product"
	"log"
	"os"
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

func (api *AppApi) DeleteProduct(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		ctx.Response().SetStatusCode(400)
		err := ctx.JSON(ErrorMessage{
			Message: fmt.Sprintf("product with ID %d not found", id),
		})
		if err != nil {
			return err
		}
	}

	productToDelete, err := api.client.Product.Query().Where(product.IDEQ(id)).Only(ctx.Context())
	if err != nil {
		return err
	}

	url := productToDelete.PreviewURL
	_ = deleteFile(url)
	for _, f := range productToDelete.Images {
		_ = deleteFile(f)
	}

	deleted, err := api.client.Product.Delete().Where(product.IDEQ(id)).Exec(ctx.Context())
	if err != nil {
		return err
	}

	if deleted != 1 {
		err := ctx.JSON(ErrorMessage{
			Message: fmt.Sprintf("Failed to delete product with id [%d]", id),
		})
		if err != nil {
			return err
		}
	}

	return ctx.SendStatus(204)
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

func (api *AppApi) UpdateProduct(ctx *fiber.Ctx) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		return err
	}

	pIdRaw := ctx.FormValue("productId")
	pId, err := strconv.ParseInt(pIdRaw, 10, 64)
	if err != nil {
		log.Println(err.Error())
		return ctx.JSON(ErrorMessage{Message: "productId is in invalid format"})
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
		// save only new files: todo: optimize
		_ = ctx.SaveFile(file, path.Join(FILES_FOLDER, fName))
		filePaths = append(filePaths, fName)
	}

	created, err := api.client.Product.Update().
		Where(product.IDEQ(int(pId))).
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

func deleteFile(file string) error {
	return os.Remove(fmt.Sprintf("./files/%s", file))
}
