package cutting_test

import (
	"net/http"
	"testing"

	"github.com/mondegor/go-webcore/mrtests/helpers"
	"github.com/stretchr/testify/suite"

	"github.com/mondegor/print-shop-back/internal/calculations/algo/section/pub/rect/cutting/controller/httpv1"
	"github.com/mondegor/print-shop-back/internal/calculations/algo/section/pub/rect/cutting/entity"
	"github.com/mondegor/print-shop-back/tests/integration"
)

type AlgoRectCuttingTestSuite struct {
	suite.Suite
	tester *integration.HttpHandlerTester
}

func (ts *AlgoRectCuttingTestSuite) SetupSuite() {
	ts.tester = integration.NewHandlerTester(ts.T())
}

func TestAlgoRectCutting(t *testing.T) {
	suite.Run(t, new(AlgoRectCuttingTestSuite))
}

func (ts *AlgoRectCuttingTestSuite) TearDownTest() {
}

func (ts *AlgoRectCuttingTestSuite) TestCalcQuantity() {
	var (
		method  = http.MethodPost
		target  = "/v1/calculations/algo/rect/cutting-quantity"
		request = httpv1.CalcCuttingQuantityRequest{
			Fragments: []string{
				"10x20",
				"18x56",
			},
			DistanceFormat: "5x3",
		}
		expectedStatusCode = http.StatusOK
		expectedResponse   = entity.QuantityResult{
			Quantity: 208,
		}
		gotResponse = entity.QuantityResult{}
	)

	statusCode, err := ts.tester.ExecRequest(
		helpers.NewHttpRequest(method, target, request),
		&gotResponse,
	)
	ts.Require().NoError(err)

	ts.Equal(expectedStatusCode, statusCode)
	ts.Equal(expectedResponse, gotResponse)
}

func (ts *AlgoRectCuttingTestSuite) TearDownSuite() {
	ts.tester.Clean()
}