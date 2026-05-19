package controller

import (
	"strconv"

	"backend/model"

	"github.com/gofiber/fiber/v2"
)

var beasiswaList = []model.Beasiswa{
	{
		ID: 1,
		Nama: "Beasiswa Prestasi/Akademik",
		Syarat: "IPK minimal 3.5",
		Deadline: "30 Juni 2026",
	},
	{
		ID: 2,
		Nama: "Beasiswa Anak Pos Indonesia",
		Syarat: "Anak pegawai Pos Indonesia",
		Deadline: "15 Juli 2026",
	},
	{
		ID: 3,
		Nama: "Beasiswa Non-Akademik",
		Syarat: "Prestasi olahraga/seni",
		Deadline: "20 Juli 2026",
	},
}

func GetBeasiswa(c *fiber.Ctx) error {
	return c.JSON(beasiswaList)
}

func GetDetailBeasiswa(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	for _, item := range beasiswaList {
		if item.ID == id {
			return c.JSON(item)
		}
	}

	return c.Status(404).JSON(fiber.Map{
		"message": "Beasiswa tidak ditemukan",
	})
}

func AddBeasiswa(c *fiber.Ctx) error {

	var baru model.Beasiswa

	if err := c.BodyParser(&baru); err != nil {
		return err
	}

	beasiswaList = append(beasiswaList, baru)

	return c.JSON(fiber.Map{
		"message": "Beasiswa berhasil ditambahkan",
		"data": baru,
	})
}

func UpdateBeasiswa(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	var update model.Beasiswa

	c.BodyParser(&update)

	for i, item := range beasiswaList {

		if item.ID == id {

			beasiswaList[i] = update

			return c.JSON(fiber.Map{
				"message": "Beasiswa berhasil diupdate",
				"data": update,
			})
		}
	}

	return c.Status(404).JSON(fiber.Map{
		"message": "Beasiswa tidak ditemukan",
	})
}

func DeleteBeasiswa(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	for i, item := range beasiswaList {

		if item.ID == id {

			beasiswaList = append(beasiswaList[:i], beasiswaList[i+1:]...)

			return c.JSON(fiber.Map{
				"message": "Beasiswa berhasil dihapus",
			})
		}
	}

	return c.Status(404).JSON(fiber.Map{
		"message": "Beasiswa tidak ditemukan",
	})
}