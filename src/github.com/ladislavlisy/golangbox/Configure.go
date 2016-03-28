package golangbox

var EMPTY_PENDING_NAMES = []uint32{}

func PendingArticleNames1(code1 uint32) []uint32 {
	return []uint32{code1}
}

func ConfigureContractTermArticles() []*Article {
	configCategory := CATEGORY_TERMS

	articleArray := []*Article{
		NewArticle(ARTICLE_CONTRACT_EMPL_TERM, configCategory,
			EMPTY_PENDING_NAMES),
		NewArticle(ARTICLE_POSITION_EMPL_TERM, configCategory,
			PendingArticleNames1(
				ARTICLE_CONTRACT_EMPL_TERM)),
	}
	return articleArray
}

func ConfigureArticles() []*Article {
	return ConfigureContractTermArticles()
}

