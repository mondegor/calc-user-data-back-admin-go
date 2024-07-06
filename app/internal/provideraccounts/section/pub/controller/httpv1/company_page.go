package httpv1

import (
	"net/http"

	"github.com/mondegor/go-webcore/mrpath"
	"github.com/mondegor/go-webcore/mrserver"

	"github.com/mondegor/print-shop-back/internal/provideraccounts/section/pub"
)

const (
	companyPageItemURL = "/v1/company/{rewriteName}"
)

type (
	// CompanyPage - comment struct.
	CompanyPage struct {
		parser     mrserver.RequestParserString
		sender     mrserver.ResponseSender
		useCase    pub.CompanyPageUseCase
		imgBaseURL mrpath.PathBuilder
	}
)

// NewCompanyPage - создаёт контроллер CompanyPage.
func NewCompanyPage(
	parser mrserver.RequestParserString,
	sender mrserver.ResponseSender,
	useCase pub.CompanyPageUseCase,
	imgBaseURL mrpath.PathBuilder,
) *CompanyPage {
	return &CompanyPage{
		parser:     parser,
		sender:     sender,
		useCase:    useCase,
		imgBaseURL: imgBaseURL,
	}
}

// Handlers - возвращает обработчики контроллера CompanyPage.
func (ht *CompanyPage) Handlers() []mrserver.HttpHandler {
	return []mrserver.HttpHandler{
		{Method: http.MethodGet, URL: companyPageItemURL, Func: ht.Get},
	}
}

// Get - comment method.
func (ht *CompanyPage) Get(w http.ResponseWriter, r *http.Request) error {
	item, err := ht.useCase.GetItemByRewriteName(r.Context(), ht.parser.PathParamString(r, "rewriteName"))
	if err != nil {
		return err
	}

	item.LogoURL = ht.imgBaseURL.BuildPath(item.LogoURL)

	return ht.sender.Send(w, http.StatusOK, item)
}
