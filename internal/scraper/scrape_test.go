package scraper

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtractProductURLs(t *testing.T) {
	reader := strings.NewReader(sampleHTML)

	// Extract the product URLs
	urls, err := extractProductURLs(reader)
	require.NoError(t, err)

	// Check the extracted URLs
	require.Len(t, urls, 3)
	assert.Equal(t, "https://books.toscrape.com/catalogue/../../a-light-in-the-attic_1000/index.html", urls[0])
	assert.Equal(t, "https://books.toscrape.com/catalogue/../../tipping-the-velvet_999/index.html", urls[1])
	assert.Equal(t, "https://books.toscrape.com/catalogue/../../soumission_998/index.html", urls[2])
}

const sampleHTML = `<div class="row">
    <aside class="sidebar col-sm-4 col-md-3">
        <div id="promotions_left">
        </div>
    </aside>

    <div class="col-sm-8 col-md-9">
        <div class="page-header action">
            <h1>Books</h1>
        </div>
        <div id="messages">
        </div>
        <div id="promotions">
        </div>
        <form method="get" class="form-horizontal">
            <div style="display:none">
            </div>
            <strong>1000</strong> results - showing <strong>1</strong> to <strong>20</strong>.
        </form>

        <section>
            <div class="alert alert-warning" role="alert"><strong>Warning!</strong> This is a demo website for web
                scraping purposes. Prices and ratings here were randomly assigned and have no real meaning.</div>
            <div>
                <ol class="row">
                    <li class="col-xs-6 col-sm-4 col-md-3 col-lg-3">
                        <article class="product_pod">
                            <div class="image_container">
                                <a href="../../a-light-in-the-attic_1000/index.html"><img
                                        src="../../../media/cache/2c/da/2cdad67c44b002e7ead0cc35693c0e8b.jpg"
                                        alt="A Light in the Attic" class="thumbnail"></a>
                            </div>
                            <p class="star-rating Three">
                                <i class="icon-star"></i>
                                <i class="icon-star"></i>
                                <i class="icon-star"></i>
                                <i class="icon-star"></i>
                                <i class="icon-star"></i>
                            </p>
                            <h3><a href="../../a-light-in-the-attic_1000/index.html" title="A Light in the Attic">A
                                    Light in the ...</a></h3>
                            <div class="product_price">
                                <p class="price_color">£51.77</p>
                                <p class="instock availability">
                                    <i class="icon-ok"></i>

                                    In stock

                                </p>
                                <form>
                                    <button type="submit" class="btn btn-primary btn-block"
                                        data-loading-text="Adding...">Add to basket</button>
                                </form>
                            </div>
                        </article>
                    </li>
                    <li class="col-xs-6 col-sm-4 col-md-3 col-lg-3">
                        <article class="product_pod">
                            <div class="image_container">
                                <a href="../../tipping-the-velvet_999/index.html"><img
                                        src="../../../media/cache/26/0c/260c6ae16bce31c8f8c95daddd9f4a1c.jpg"
                                        alt="Tipping the Velvet" class="thumbnail"></a>
                            </div>
                            <p class="star-rating One">
                                <i class="icon-star"></i>
                                <i class="icon-star"></i>
                                <i class="icon-star"></i>
                                <i class="icon-star"></i>
                                <i class="icon-star"></i>
                            </p>
                            <h3><a href="../../tipping-the-velvet_999/index.html" title="Tipping the Velvet">Tipping the
                                    Velvet</a></h3>
                            <div class="product_price">
                                <p class="price_color">£53.74</p>
                                <p class="instock availability">
                                    <i class="icon-ok"></i>
                                    In stock
                                </p>
                                <form>
                                    <button type="submit" class="btn btn-primary btn-block"
                                        data-loading-text="Adding...">Add to basket</button>
                                </form>
                            </div>
                        </article>
                    </li>
                    <li class="col-xs-6 col-sm-4 col-md-3 col-lg-3">
                        <article class="product_pod">
                            <div class="image_container">
                                <a href="../../soumission_998/index.html"><img
                                        src="../../../media/cache/3e/ef/3eef99c9d9adef34639f510662022830.jpg"
                                        alt="Soumission" class="thumbnail"></a>
                            </div>
                            <p class="star-rating One">
                                <i class="icon-star"></i>
                                <i class="icon-star"></i>
                                <i class="icon-star"></i>
                                <i class="icon-star"></i>
                                <i class="icon-star"></i>
                            </p>
                            <h3><a href="../../soumission_998/index.html" title="Soumission">Soumission</a></h3>
                            <div class="product_price">
                                <p class="price_color">£50.10</p>
                                <p class="instock availability">
                                    <i class="icon-ok"></i>

                                    In stock

                                </p>
                                <form>
                                    <button type="submit" class="btn btn-primary btn-block"
                                        data-loading-text="Adding...">Add to basket</button>
                                </form>
                            </div>
                        </article>
                    </li>
                </ol>
                <div>
                    <ul class="pager">
                        <li class="current">
                            Page 1 of 50
                        </li>
                        <li class="next"><a href="page-2.html">next</a></li>
                    </ul>
                </div>
            </div>
        </section>
    </div>
</div><!-- /row -->`
