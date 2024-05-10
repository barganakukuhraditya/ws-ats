package controller

import (
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	cek "github.com/barganakukuhraditya/be_ats/module"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetAllJadwal(c *fiber.Ctx) error {
    ps, err := cek.GetAllJadwal()
    if err != nil {
        return err
    }
    
    return c.JSON(ps)
}
