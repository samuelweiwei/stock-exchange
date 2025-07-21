package captcha

import (
	"crypto/rand"
	"io"
)

type Driver interface {
	GenerateIdCaptcha(target string) (id, captcha string)
}

func NewDriver(length int) Driver {
	return &DigitalCaptchaDriver{length: length}
}

type DigitalCaptchaDriver struct {
	length int
}

var idLen = 20
var idChars = []byte(TxtNumbers + TxtAlphabet)

const (
	//TxtNumbers chacters for numbers.
	TxtNumbers = "012346789"
	//TxtAlphabet characters for alphabet.
	TxtAlphabet = "ABCDEFGHJKMNOQRSTUVXYZabcdefghjkmnoqrstuvxyz"
)

func (d *DigitalCaptchaDriver) GenerateIdCaptcha(target string) (id, captcha string) {
	id = d.RandomId()
	captcha = d.randomDigits(d.length)
	return
}

func (d *DigitalCaptchaDriver) RandomId() string {
	b := d.randomBytesMod(idLen, byte(len(idChars)))
	for i, c := range b {
		b[i] = idChars[c]
	}
	return string(b)
}
func (d *DigitalCaptchaDriver) randomDigits(length int) (res string) {
	digits := d.randomBytesMod(length, 10)
	res = d.parseDigitsToString(digits)
	return
}

// randomBytes returns a byte slice of the given length read from CSPRNG.
func (d *DigitalCaptchaDriver) randomBytes(length int) (b []byte) {
	b = make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		panic("captcha: error reading random source: " + err.Error())
	}
	return
}

// randomBytesMod returns a byte slice of the given length, where each byte is
// a random number modulo mod.
func (d *DigitalCaptchaDriver) randomBytesMod(length int, mod byte) (b []byte) {
	if length == 0 {
		return nil
	}
	if mod == 0 {
		panic("captcha: bad mod argument for randomBytesMod")
	}
	maxrb := 255 - byte(256%int(mod))
	b = make([]byte, length)
	i := 0
	for {
		r := d.randomBytes(length + (length / 4))
		for _, c := range r {
			if c > maxrb {
				// Skip this number to avoid modulo bias.
				continue
			}
			b[i] = c % mod
			i++
			if i == length {
				return
			}
		}
	}
}

func (d *DigitalCaptchaDriver) parseDigitsToString(bytes []byte) string {
	stringB := make([]byte, len(bytes))
	for idx, by := range bytes {
		stringB[idx] = by + '0'
	}
	return string(stringB)
}
