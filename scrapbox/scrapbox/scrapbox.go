package scrapbox

type Scrapbox struct {
	Pages []Article `json:"pages"`
}

type Article struct {
	Title string   `json:"title"`
	Lines []string `json:"lines"`
}

func NewScrapbox(articles []Article) Scrapbox {
	scrapbox := Scrapbox{}

	for _, a := range articles {
		scrapbox.Pages = append(scrapbox.Pages, a)
	}

	return scrapbox
}

func NewArticle(title string, lines []string) Article {
	article := Article{}

	article.Title = title
	article.Lines = lines

	return article
}
