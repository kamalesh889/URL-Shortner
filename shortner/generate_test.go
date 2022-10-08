package shortner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratorshortUrl(t *testing.T) {
	longurl := "https://www.makemytrip.com/flight/search?itinerary=IXB-BBI-30/09/2022&tripType=O&paxType=A-1_C-0_I-0&intl=false&cabinClass=E&ccde=IN&lang=eng"
	generateurl := GenerateshorlUrl(longurl)

	assert.Equal(t, generateurl, "HSVSfwZjGAM")

}
