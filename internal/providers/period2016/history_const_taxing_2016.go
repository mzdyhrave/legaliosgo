package period2016

import (
	"github.com/mzdyhrave/payrollgo-legalios/internal/providers/period2015"
)

const (
	TAXING_VERSION_CODE int16 = 2016
	TAXING_ALLOWANCE_PAYER = period2015.TAXING_ALLOWANCE_PAYER
	TAXING_ALLOWANCE_DISAB_1ST = period2015.TAXING_ALLOWANCE_DISAB_1ST
	TAXING_ALLOWANCE_DISAB_2ND = period2015.TAXING_ALLOWANCE_DISAB_2ND
	TAXING_ALLOWANCE_DISAB_3RD = period2015.TAXING_ALLOWANCE_DISAB_3RD
	TAXING_ALLOWANCE_STUDY = period2015.TAXING_ALLOWANCE_STUDY
	TAXING_ALLOWANCE_CHILD_1ST = period2015.TAXING_ALLOWANCE_CHILD_1ST
	TAXING_ALLOWANCE_CHILD_2ND = period2015.TAXING_ALLOWANCE_CHILD_2ND
	TAXING_ALLOWANCE_CHILD_3RD = period2015.TAXING_ALLOWANCE_CHILD_3RD
	TAXING_FACTOR_ADVANCES float64 = period2015.TAXING_FACTOR_ADVANCES
	TAXING_FACTOR_WITHHOLD float64 = period2015.TAXING_FACTOR_WITHHOLD
	TAXING_FACTOR_SOLITARY float64 = period2015.TAXING_FACTOR_SOLITARY
	TAXING_MIN_AMOUNT_OF_TAXBONUS = period2015.TAXING_MIN_AMOUNT_OF_TAXBONUS
	TAXING_MAX_AMOUNT_OF_TAXBONUS = period2015.TAXING_MAX_AMOUNT_OF_TAXBONUS
	TAXING_MARGIN_INCOME_OF_TAXBONUS = (SALARY_MIN_MONTHLY_WAGE / 2)
	TAXING_MARGIN_INCOME_OF_ROUNDING = period2015.TAXING_MARGIN_INCOME_OF_ROUNDING
	TAXING_MARGIN_INCOME_OF_WITHHOLD = period2015.TAXING_MARGIN_INCOME_OF_WITHHOLD
	TAXING_MARGIN_INCOME_OF_SOLITARY = (4 * 27006)
	TAXING_MARGIN_INCOME_OF_WHT_EMP = period2015.TAXING_MARGIN_INCOME_OF_WHT_EMP
	TAXING_MARGIN_INCOME_OF_WHT_AGR = period2015.TAXING_MARGIN_INCOME_OF_WHT_AGR

	VAR05_TAXING_ALLOWANCE_CHILD_2ND = 1417
	VAR05_TAXING_ALLOWANCE_CHILD_3RD = 1717
)

