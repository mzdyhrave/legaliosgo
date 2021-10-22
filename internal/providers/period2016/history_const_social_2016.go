package period2016

import (
	"github.com/mzdyhrave/payrollgo-legalios/internal/providers/period2015"
)

const (
	SOCIAL_VERSION_CODE int16 = 2016
	SOCIAL_MAX_ANNUALS_BASIS = 1296288
	SOCIAL_FACTOR_EMPLOYER float64 = period2015.SOCIAL_FACTOR_EMPLOYER
	SOCIAL_FACTOR_EMPLOYER_HIGHER float64 = period2015.SOCIAL_FACTOR_EMPLOYER_HIGHER
	SOCIAL_FACTOR_EMPLOYEE float64 = period2015.SOCIAL_FACTOR_EMPLOYEE
	SOCIAL_FACTOR_EMPLOYEE_REDUCE float64 = 0
	SOCIAL_FACTOR_EMPLOYEE_GARANT float64 = 0
	SOCIAL_MARGIN_INCOME_EMP = period2015.SOCIAL_MARGIN_INCOME_EMP
	SOCIAL_MARGIN_INCOME_AGR = period2015.SOCIAL_MARGIN_INCOME_AGR
)

