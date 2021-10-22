package period2018

import (
	. "github.com/shopspring/decimal"
	"github.com/mzdyhrave/payrollgo-legalios/internal/providers"
	"github.com/mzdyhrave/payrollgo-legalios/internal/props"
	"github.com/mzdyhrave/payrollgo-legalios/internal/types"
)

type providerSocial2018 struct {
	providers.ProviderBase
}

func NewProviderSocial2018() providers.IProviderSocial {
	return providerSocial2018{
		ProviderBase: providers.ProviderBase{Version: types.GetVersionId(2018)},
	}
}

func (b providerSocial2018) GetProps(period types.IPeriod) props.IPropsSocial {
	return props.NewPropsSocial(b.Version,
		b.MaxAnnualsBasis(period),
		b.FactorEmployer(period),
		b.FactorEmployerHigher(period),
		b.FactorEmployee(period),
		b.FactorEmployeeGarant(period),
		b.FactorEmployeeReduce(period),
		b.MarginIncomeEmp(period),
		b.MarginIncomeAgr(period))
}

func (b providerSocial2018) MaxAnnualsBasis(period types.IPeriod) int32 {
	return SOCIAL_MAX_ANNUALS_BASIS
}

func (b providerSocial2018) FactorEmployer(period types.IPeriod) Decimal {
	return NewFromFloat(SOCIAL_FACTOR_EMPLOYER)
}

func (b providerSocial2018) FactorEmployerHigher(period types.IPeriod) Decimal {
	return NewFromFloat(SOCIAL_FACTOR_EMPLOYER_HIGHER)
}

func (b providerSocial2018) FactorEmployee(period types.IPeriod) Decimal {
	return NewFromFloat(SOCIAL_FACTOR_EMPLOYEE)
}

func (b providerSocial2018) FactorEmployeeGarant(period types.IPeriod) Decimal {
	return NewFromFloat(SOCIAL_FACTOR_EMPLOYEE_GARANT)
}

func (b providerSocial2018) FactorEmployeeReduce(period types.IPeriod) Decimal {
	return NewFromFloat(SOCIAL_FACTOR_EMPLOYEE_REDUCE)
}

func (b providerSocial2018) MarginIncomeEmp(period types.IPeriod) int32 {
	return SOCIAL_MARGIN_INCOME_EMP
}

func (b providerSocial2018) MarginIncomeAgr(period types.IPeriod) int32 {
	return SOCIAL_MARGIN_INCOME_AGR
}
