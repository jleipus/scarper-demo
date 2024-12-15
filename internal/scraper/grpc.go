package scraper

import pb "scaper-demo/proto"

// TransformProductRPC transforms a pb.ParsedPageResponse to a Product struct.
func TransformProductRPC(resp *pb.ParsedPageResponse) Product {
	return Product{
		Name:         resp.GetName(),
		Availability: resp.GetAvailability(),
		Upc:          resp.GetUpc(),
		PriceExclTax: resp.GetPriceExclTax(),
		Tax:          resp.GetTax(),
	}
}
