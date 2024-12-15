package parser_test

import (
	"scaper-demo/internal/parser"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseHTMLContent(t *testing.T) {
	t.Run("sample page", func(t *testing.T) {
		t.Parallel()

		result, err := parser.ParseHTMLContent(strings.NewReader(sampleHTML))
		require.NoError(t, err)

		assert.Equal(t, "A Light in the Attic", result.Name)
		assert.Equal(t, "In stock (22 available)", result.Availability)
		assert.Equal(t, "a897fe39b1053632", result.Upc)
		assert.Equal(t, "£51.77", result.PriceExclTax)
		assert.Equal(t, "£0.00", result.Tax)
	})

	t.Run("empty page", func(t *testing.T) {
		t.Parallel()

		_, err := parser.ParseHTMLContent(strings.NewReader(""))
		require.EqualError(t, err, "no body found in the HTML document")
	})
}

const sampleHTML = `<div class="content">
    <div id="promotions">
    </div>
    <div id="content_inner">
        <article class="product_page"><!-- Start of product page -->
            <div class="row">
                <div class="col-sm-6">
                    <div id="product_gallery" class="carousel">
                        <div class="thumbnail">
                            <div class="carousel-inner">
                                <div class="item active">
                                    <img src="../../media/cache/fe/72/fe72f0532301ec28892ae79a629a293c.jpg"
                                        alt="A Light in the Attic" />
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="col-sm-6 product_main">
                    <h1>A Light in the Attic</h1>
                    <p class="price_color">£51.77</p>
                    <p class="instock availability">
                        <i class="icon-ok"></i>

                        In stock (22 available)

                    </p>
                    <p class="star-rating Three">
                        <i class="icon-star"></i>
                        <i class="icon-star"></i>
                        <i class="icon-star"></i>
                        <i class="icon-star"></i>
                        <i class="icon-star"></i>
                    </p>
                    <hr />
                    <div class="alert alert-warning" role="alert"><strong>Warning!</strong> This is a demo
                        website for web scraping purposes. Prices and ratings here were randomly assigned
                        and have no real meaning.</div>
                </div><!-- /col-sm-6 -->
            </div><!-- /row -->
            <div id="product_description" class="sub-header">
                <h2>Product Description</h2>
            </div>
            <p>It's hard to imagine a world without A Light in the Attic. This now-classic collection of
                poetry and drawings from Shel Silverstein celebrates its 20th anniversary with this special
                edition. Silverstein's humorous and creative verse can amuse the dowdiest of readers.
                Lemon-faced adults and fidgety kids sit still and read these rhythmic words and laugh and
                smile and love th It's hard to imagine a world without A Light in the Attic. This
                now-classic collection of poetry and drawings from Shel Silverstein celebrates its 20th
                anniversary with this special edition. Silverstein's humorous and creative verse can amuse
                the dowdiest of readers. Lemon-faced adults and fidgety kids sit still and read these
                rhythmic words and laugh and smile and love that Silverstein. Need proof of his genius?
                RockabyeRockabye baby, in the treetopDon't you know a treetopIs no safe place to rock?And
                who put you up there,And your cradle, too?Baby, I think someone down here'sGot it in for
                you. Shel, you never sounded so good. ...more</p>



            <div class="sub-header">
                <h2>Product Information</h2>
            </div>
            <table class="table table-striped">
                <tr>
                    <th>UPC</th>
                    <td>a897fe39b1053632</td>
                </tr>
                <tr>
                    <th>Product Type</th>
                    <td>Books</td>
                </tr>
                <tr>
                    <th>Price (excl. tax)</th>
                    <td>£51.77</td>
                </tr>
                <tr>
                    <th>Price (incl. tax)</th>
                    <td>£51.77</td>
                </tr>
                <tr>
                    <th>Tax</th>
                    <td>£0.00</td>
                </tr>
                <tr>
                    <th>Availability</th>
                    <td>In stock (22 available)</td>
                </tr>
                <tr>
                    <th>Number of reviews</th>
                    <td>0</td>
                </tr>
            </table>
            <section>
                <div id="reviews" class="sub-header">
                </div>
            </section>
        </article><!-- End of product page -->
    </div>
</div>`
