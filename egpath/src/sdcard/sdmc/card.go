// Package sdmc implements access to Secure Digital Memory Cards.
package sdmc

import (
	"errors"

	"sdcard"
)

var (
	ErrInitCMD8    = errors.New("sdmc: init CMD8")
	ErrInitACMD41  = errors.New("sdmc: init ACMD41")
	ErrStatus      = errors.New("sdmc: status") // Error flag or StateData timeout.
	ErrBadAddr     = errors.New("sdmc: bad addr")
)

type Card struct {
	host   sdcard.Host
	cap    int64
	oca    sdcard.OCR
	status sdcard.CardStatus
	rca    uint16
	ccc    uint16
}

func MakeCard(host sdcard.Host) Card {
	return Card{host: host}
}

func NewCard(host sdcard.Host) *Card {
	c := new(Card)
	*c = MakeCard(host)
	return c
}

// Cap returns card capacity as number of 512-byte blocks.
func (c *Card) Cap() int64 {
	return c.cap
}

// LastStatus returns status of the last command that returns R1
// response.
func (c *Card) LastStatus() sdcard.CardStatus {
	return c.status
}

func (c *Card) Status() (sdcard.CardStatus, error) {
	h := c.host
	c.status = h.SendCmd(sdcard.CMD13(c.rca, sdcard.Status)).R1()
	return c.status, h.Err(true)
}

const statusErrFlags = sdcard.OUT_OF_RANGE |
	sdcard.ADDRESS_ERROR |
	sdcard.BLOCK_LEN_ERROR |
	sdcard.ERASE_SEQ_ERROR |
	sdcard.ERASE_PARAM |
	sdcard.WP_VIOLATION |
	sdcard.LOCK_UNLOCK_FAILED |
	sdcard.COM_CRC_ERROR |
	sdcard.ILLEGAL_COMMAND |
	sdcard.CARD_ECC_FAILED |
	sdcard.CC_ERROR |
	sdcard.ERROR |
	sdcard.CSD_OVERWRITE |
	sdcard.WP_ERASE_SKIP |
	sdcard.CARD_ECC_DISABLED |
	sdcard.ERASE_RESET |
	sdcard.AKE_SEQ_ERROR

func (c *Card) statusCmd(cmd sdcard.Command, arg uint32) error {
	status := c.host.SendCmd(cmd, arg).R1()
	if err := c.host.Err(true); err != nil {
		c.host.SetClock(0, false)
		return err
	}
	c.status = status
	if status&statusErrFlags != 0 {
		return ErrStatus
	}
	return nil
}
