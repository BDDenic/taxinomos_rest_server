package categories

/*
{
    "data": {
        "type": "categories",
        "id": "1105",
        "attributes": {
            "website-category-id": 1105,
            "maincategory": "Automotive",
            "subcategory": "Wholesale trade of motor vehicle parts and accessories",
            "lang": "en",
            "description": "This is a test description"
        },
        "relationships": {
            "measurements": {
                "meta": {
                    "total": 1
                },
                "links": {
                    "self": "https://classify-rest.labs.nic.at/api/v1/categories/1105/relationships/measurements",
                    "related": "https://classify-rest.labs.nic.at/api/v1/categories/1105/measurements"
                }
            }
        },
        "links": {
            "self": "https://classify-rest.labs.nic.at/api/v1/categories/1105"
        }
    }
}
*/

type CategoryData struct {
	Type          string                `json:"type"`
	Id            string                `json:"id"`
	Attributes    CategoryAttributes    `json:"attributes"`
	Relationships CategoryRelationships `json:"relationships"`
	Links         CategoryLinks         `json:"links"`
}

type CategoryAttributes struct {
	WebsiteCategoryId int    `json:"website-category-id"`
	MainCategory      string `json:"maincategory"`
	SubCategory       string `json:"subcategory"`
	Language          string `json:"lang"`
	Description       string `json:"description"`
}

type CategoryRelationships struct {
	Measurements struct {
		Total int `json:"total"`
	} `json:"measurements"`
}

type CategoryLinks struct {
	Self string `json:"self"`
}

type Category struct {
	Data CategoryData `json:"data"`
}

type Categories []Categories
