package controller

import (
	"context"

	"backend/config"
	"backend/model"

	"github.com/gofiber/fiber/v2"
)

func DaftarBeasiswa(c *fiber.Ctx) error {

	collection :=
		config.MongoClient.
			Database("ulbi_beasiswa").
			Collection("pendaftaran")

	var data model.Pendaftaran

	if err := c.BodyParser(&data); err != nil {

		return c.Status(400).JSON(fiber.Map{
			"message": "Data tidak valid",
		})
	}

	data.Status = "Pending"

	_, err := collection.InsertOne(
		context.Background(),
		data,
	)

	if err != nil {

		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Pendaftaran berhasil",
	})
}

func CekStatus(c *fiber.Ctx) error {

	npm := c.Params("npm")

	collection :=
		config.MongoClient.
			Database("ulbi_beasiswa").
			Collection("pendaftaran")

	var result model.Pendaftaran

	err := collection.FindOne(
		context.Background(),
		fiber.Map{
			"npm": npm,
		},
	).Decode(&result)

	if err != nil {

		return c.Status(404).JSON(fiber.Map{
			"message": "Data tidak ditemukan",
		})
	}

	return c.JSON(fiber.Map{
		"nama": result.NamaMahasiswa,
		"npm": result.NPM,
		"status": result.Status,
		"beasiswa": result.Beasiswa,
	})
}

func UpdateStatus(c *fiber.Ctx) error {

	npm := c.Params("npm")

	status := c.Query("status")

	collection :=
		config.MongoClient.
			Database("ulbi_beasiswa").
			Collection("pendaftaran")

	filter := fiber.Map{
		"npm": npm,
	}

	update := fiber.Map{
		"$set": fiber.Map{
			"status": status,
		},
	}

	_, err := collection.UpdateOne(
		context.Background(),
		filter,
		update,
	)

	if err != nil {

		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Status berhasil diupdate",
	})
}

func DeletePendaftaran(c *fiber.Ctx) error {

	npm := c.Params("npm")

	collection :=
		config.MongoClient.
			Database("ulbi_beasiswa").
			Collection("pendaftaran")

	_, err := collection.DeleteOne(
		context.Background(),
		fiber.Map{
			"npm": npm,
		},
	)

	if err != nil {

		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil dihapus",
	})
}