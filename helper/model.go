package helper

import (
	"belajar-golang-restfull-api/model/domain"
	"belajar-golang-restfull-api/model/web"
)

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, web.CategoryResponse(category))
	}

	return categoryResponses
}
