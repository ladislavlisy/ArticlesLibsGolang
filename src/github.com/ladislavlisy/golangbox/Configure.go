package golangbox

var EMPTY_PENDING_NAMES = []uint32{}

func PendingArticleNames1(code1 uint32) []uint32 {
	return []uint32{code1}
}

func PendingArticleNames2(code1 uint32, code2 uint32) []uint32 {
	return []uint32{code1, code2}
}

func PendingArticleNames3(code1 uint32, code2 uint32, code3 uint32) []uint32 {
	return []uint32{code1, code2, code3 }
}

func PendingArticleNames4(code1 uint32, code2 uint32, code3 uint32, code4 uint32) []uint32 {
	return []uint32{code1, code2, code3, code4}
}

func PendingArticleNames5(code1 uint32, code2 uint32, code3 uint32, code4 uint32, code5 uint32) []uint32 {
	return []uint32{code1, code2, code3, code4, code5}
}

func PendingArticleNames9(code1 uint32, code2 uint32, code3 uint32, code4 uint32, code5 uint32, code6 uint32, code7 uint32, code8 uint32, code9 uint32) []uint32 {
	return []uint32{code1, code2, code3, code4, code5, code6, code7, code8, code9}
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

func ConfigurePositionTimeArticles() []*Article {
	configCategory := CATEGORY_TIMES

	articleArray := []*Article{
		NewArticle(ARTICLE_SCHEDULE_WORK, configCategory,
		PendingArticleNames1(
			ARTICLE_POSITION_EMPL_TERM)),
		NewArticle(ARTICLE_TIMESHEET_SCHEDULE, configCategory,
		PendingArticleNames1(
			ARTICLE_SCHEDULE_WORK)),
		NewArticle(ARTICLE_TIMESHEET_WORKING, configCategory,
		PendingArticleNames2(
			ARTICLE_TIMESHEET_SCHEDULE,
			ARTICLE_POSITION_EMPL_TERM)),
		NewArticle(ARTICLE_TIMESHEET_ABSENCE, configCategory,
		PendingArticleNames1(
			ARTICLE_TIMESHEET_WORKING)),
		NewArticle(ARTICLE_TIMEHOURS_WORKING, configCategory,
		PendingArticleNames1(
			ARTICLE_TIMESHEET_WORKING)),
		NewArticle(ARTICLE_TIMEHOURS_ABSENCE, configCategory,
		PendingArticleNames1(
			ARTICLE_TIMESHEET_ABSENCE)),
	}
	return articleArray
}

func ConfigureGrossIncomeArticles() []*Article {
	configCategory := CATEGORY_AMOUNT

	articleArray := []*Article{
		NewArticle(ARTICLE_SALARY_BASE, configCategory,
		PendingArticleNames2(
			ARTICLE_TIMEHOURS_WORKING,
			ARTICLE_TIMEHOURS_ABSENCE)),
	}
	return articleArray
}

func ConfigureTotalIncomeArticles() []*Article {
	configCategory := CATEGORY_FINAL

	articleArray := []*Article{
		NewArticle(ARTICLE_INCOME_GROSS, configCategory,
		EMPTY_PENDING_NAMES),
		NewArticle(ARTICLE_INCOME_NETTO, configCategory,
		PendingArticleNames9(
			ARTICLE_INCOME_GROSS,
			ARTICLE_TAXING_ADVANCES_TOTAL,
			ARTICLE_TAXING_BONUS_CHILD,
			ARTICLE_TAXING_WITHHOLD_GENERAL,
			ARTICLE_HEALTH_EMPLOYEE_GENERAL,
			ARTICLE_HEALTH_EMPLOYEE_MANDATORY,
			ARTICLE_SOCIAL_EMPLOYEE_GENERAL,
			ARTICLE_SOCIAL_EMPLOYEE_PENSION,
			ARTICLE_GARANT_EMPLOYEE_PENSION)),
	}
	return articleArray
}

func ConfigureNettoDeductsArticles() []*Article {
	configCategory := CATEGORY_NETTO

	articleArray := []*Article{
		NewArticle(ARTICLE_TAXING_ADVANCES_TOTAL, configCategory,
		PendingArticleNames2(
			ARTICLE_TAXING_ADVANCES_GENERAL,
			ARTICLE_TAXING_ADVANCES_SOLIDARY)),
		NewArticle(ARTICLE_TAXING_ADVANCES_GENERAL, configCategory,
		PendingArticleNames1(
			ARTICLE_TAXING_ADVANCES_BASIS_GENERAL)),
		NewArticle(ARTICLE_TAXING_ADVANCES_SOLIDARY, configCategory,
		PendingArticleNames1(
			ARTICLE_TAXING_ADVANCES_BASIS_SOLIDARY)),
		NewArticle(ARTICLE_TAXING_WITHHOLD_GENERAL, configCategory,
		PendingArticleNames1(
			ARTICLE_TAXING_WITHHOLD_BASIS_GENERAL)),
		NewArticle(ARTICLE_HEALTH_EMPLOYEE_GENERAL, configCategory,
		PendingArticleNames1(
			ARTICLE_HEALTH_BASIS_GENERAL)),
		NewArticle(ARTICLE_HEALTH_EMPLOYEE_MANDATORY, configCategory,
		PendingArticleNames1(
			ARTICLE_HEALTH_BASIS_MANDATORY)),
		NewArticle(ARTICLE_SOCIAL_EMPLOYEE_GENERAL, configCategory,
		PendingArticleNames1(
			ARTICLE_SOCIAL_BASIS_GENERAL)),
		NewArticle(ARTICLE_SOCIAL_EMPLOYEE_PENSION, configCategory,
		PendingArticleNames1(
			ARTICLE_SOCIAL_BASIS_PENSION)),
		NewArticle(ARTICLE_GARANT_EMPLOYEE_PENSION, configCategory,
		PendingArticleNames1(
			ARTICLE_GARANT_BASIS_PENSION)),
	}
	return articleArray
}

func ConfigureBasisHealthArticles() []*Article {
	configCategory := CATEGORY_NETTO

	articleArray := []*Article{
		NewArticle(ARTICLE_HEALTH_INCOME_SUBJECT, configCategory,
		EMPTY_PENDING_NAMES),
		NewArticle(ARTICLE_HEALTH_INCOME_PARTICIP, configCategory,
		PendingArticleNames1(
			ARTICLE_HEALTH_INCOME_SUBJECT)),
		NewArticle(ARTICLE_HEALTH_BASIS_GENERAL, configCategory,
		PendingArticleNames1(
			ARTICLE_HEALTH_INCOME_PARTICIP)),
		NewArticle(ARTICLE_HEALTH_BASIS_MANDATORY, configCategory,
		PendingArticleNames2(
			ARTICLE_HEALTH_BASIS_GENERAL,
			ARTICLE_HEALTH_INCOME_PARTICIP)),
		NewArticle(ARTICLE_HEALTH_BASIS_LEGALCAP, configCategory,
		PendingArticleNames2(
			ARTICLE_HEALTH_BASIS_GENERAL,
			ARTICLE_HEALTH_INCOME_PARTICIP)),
	}
	return articleArray
}

func ConfigureBasisSocialArticles() []*Article {
	configCategory := CATEGORY_NETTO

	articleArray := []*Article{
		NewArticle(ARTICLE_SOCIAL_INCOME_SUBJECT, configCategory,
		EMPTY_PENDING_NAMES),
		NewArticle(ARTICLE_SOCIAL_INCOME_PARTICIP, configCategory,
		PendingArticleNames1(
			ARTICLE_SOCIAL_INCOME_SUBJECT)),
		NewArticle(ARTICLE_SOCIAL_BASIS_GENERAL, configCategory,
		PendingArticleNames1(
			ARTICLE_SOCIAL_INCOME_PARTICIP)),
		NewArticle(ARTICLE_SOCIAL_BASIS_PENSION, configCategory,
		PendingArticleNames1(
			ARTICLE_SOCIAL_INCOME_PARTICIP)),
		NewArticle(ARTICLE_SOCIAL_BASIS_LEGALCAP, configCategory,
		PendingArticleNames1(
			ARTICLE_SOCIAL_INCOME_PARTICIP)),
	}
	return articleArray
}

func ConfigureBasisGarantArticles() []*Article {
	configCategory := CATEGORY_NETTO

	articleArray := []*Article{
		NewArticle(ARTICLE_GARANT_INCOME_SUBJECT, configCategory,
		EMPTY_PENDING_NAMES),
		NewArticle(ARTICLE_GARANT_INCOME_PARTICIP, configCategory,
		PendingArticleNames1(
			ARTICLE_GARANT_INCOME_SUBJECT)),
		NewArticle(ARTICLE_GARANT_BASIS_PENSION, configCategory,
		PendingArticleNames1(
			ARTICLE_GARANT_INCOME_PARTICIP)),
		NewArticle(ARTICLE_GARANT_BASIS_LEGALCAP, configCategory,
		PendingArticleNames1(
			ARTICLE_GARANT_INCOME_PARTICIP)),
	}
	return articleArray
}

func ConfigureBasisTaxingArticles() []*Article {
	configCategory := CATEGORY_NETTO

	articleArray := []*Article{
		NewArticle(ARTICLE_TAXING_INCOME_SUBJECT, configCategory,
		EMPTY_PENDING_NAMES),
		NewArticle(ARTICLE_TAXING_INCOME_HEALTH, configCategory,
		EMPTY_PENDING_NAMES),
		NewArticle(ARTICLE_TAXING_INCOME_SOCIAL, configCategory,
		EMPTY_PENDING_NAMES),
	}
	return articleArray
}

func ConfigureBasisAdvancesArticles() []*Article {
	configCategory := CATEGORY_NETTO

	articleArray := []*Article{
		NewArticle(ARTICLE_TAXING_ADVANCES_INCOME, configCategory,
		PendingArticleNames1 (
			ARTICLE_TAXING_INCOME_SUBJECT)),
		NewArticle(ARTICLE_TAXING_ADVANCES_HEALTH, configCategory,
		PendingArticleNames1 (
			ARTICLE_TAXING_INCOME_HEALTH)),
		NewArticle(ARTICLE_TAXING_ADVANCES_SOCIAL, configCategory,
		PendingArticleNames1 (
			ARTICLE_TAXING_INCOME_SOCIAL)),
		NewArticle(ARTICLE_TAXING_ADVANCES_BASIS_GENERAL, configCategory,
		PendingArticleNames3 (
			ARTICLE_TAXING_ADVANCES_INCOME,
			ARTICLE_TAXING_ADVANCES_HEALTH,
			ARTICLE_TAXING_ADVANCES_SOCIAL)),
		NewArticle(ARTICLE_TAXING_ADVANCES_BASIS_SOLIDARY, configCategory,
		PendingArticleNames1 (
			ARTICLE_TAXING_ADVANCES_BASIS_GENERAL)),
	}
	return articleArray
}

func ConfigureBasisWithholdArticles() []*Article {
	configCategory := CATEGORY_NETTO

	articleArray := []*Article{
		NewArticle(ARTICLE_TAXING_WITHHOLD_INCOME, configCategory,
		PendingArticleNames1(
			ARTICLE_TAXING_INCOME_SUBJECT)),
		NewArticle(ARTICLE_TAXING_WITHHOLD_HEALTH, configCategory,
		PendingArticleNames1(
			ARTICLE_TAXING_INCOME_HEALTH)),
		NewArticle(ARTICLE_TAXING_WITHHOLD_SOCIAL, configCategory,
		PendingArticleNames1(
			ARTICLE_TAXING_INCOME_SOCIAL)),
		NewArticle(ARTICLE_TAXING_WITHHOLD_BASIS_GENERAL, configCategory,
		PendingArticleNames3(
			ARTICLE_TAXING_WITHHOLD_INCOME,
			ARTICLE_TAXING_WITHHOLD_HEALTH,
			ARTICLE_TAXING_WITHHOLD_SOCIAL)),
	}
	return articleArray
}

func ConfigureAllowanceTaxisArticles() []*Article {
	configCategory := CATEGORY_NETTO

	articleArray := []*Article{
		NewArticle(ARTICLE_TAXING_ALLOWANCE_PAYER, configCategory,
		PendingArticleNames1(
			ARTICLE_TAXING_ADVANCES_INCOME)),
		NewArticle(ARTICLE_TAXING_ALLOWANCE_DISABILITY, configCategory,
		PendingArticleNames1(
			ARTICLE_TAXING_ADVANCES_INCOME)),
		NewArticle(ARTICLE_TAXING_ALLOWANCE_STUDYING, configCategory,
		PendingArticleNames1(
			ARTICLE_TAXING_ADVANCES_INCOME)),
		NewArticle(ARTICLE_TAXING_ALLOWANCE_CHILD, configCategory,
		PendingArticleNames1(
			ARTICLE_TAXING_ADVANCES_INCOME)),
	}
	return articleArray
}

func ConfigureRebateTaxisArticles() []*Article {
	configCategory := CATEGORY_NETTO

	articleArray := []*Article{
		NewArticle(ARTICLE_TAXING_REBATE_PAYER, configCategory,
		PendingArticleNames4(
			ARTICLE_TAXING_ALLOWANCE_PAYER,
			ARTICLE_TAXING_ADVANCES_TOTAL,
			ARTICLE_TAXING_ALLOWANCE_DISABILITY,
			ARTICLE_TAXING_ALLOWANCE_STUDYING)),
		NewArticle(ARTICLE_TAXING_REBATE_CHILD, configCategory,
		PendingArticleNames3(
			ARTICLE_TAXING_ALLOWANCE_CHILD,
			ARTICLE_TAXING_ADVANCES_TOTAL,
			ARTICLE_TAXING_REBATE_PAYER)),
		NewArticle(ARTICLE_TAXING_BONUS_CHILD, configCategory,
		PendingArticleNames3(
			ARTICLE_TAXING_ADVANCES_TOTAL,
			ARTICLE_TAXING_REBATE_PAYER,
			ARTICLE_TAXING_REBATE_CHILD)),
	}
	return articleArray
}

func ConfigureArticles() []*Article {
	return ConfigureContractTermArticles()
}

