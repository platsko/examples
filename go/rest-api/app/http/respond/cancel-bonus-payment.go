package respond

import (
	"net/http"

	"github.com/go-chi/render"

	resp "lab/go-rest-api/app/http/respond/types"
	"lab/go-rest-api/app/repo"
)

type (
	// cancelBonusPaymentResponse describes response.
	cancelBonusPaymentResponse struct {
		repo     *repo.DocumentRepo
		Code     int            `json:"code"`
		Document *resp.Document `json:"document"`
		Display  *resp.Slip     `json:"cashierInformation,omitempty"`
		Printer  *resp.Print    `json:"printingInformation,omitempty"`
	}
)

// Render satisfies renderer interface.
func (r *cancelBonusPaymentResponse) Render(_ http.ResponseWriter, rq *http.Request) error {
	render.Status(rq, r.Code)
	return nil
}

// CancelBonusPaymentResponse wraps resource to respond.
func CancelBonusPaymentResponse(rs http.ResponseWriter, rq *http.Request, rp *repo.DocumentRepo) {
	// init calculate response
	r := &cancelBonusPaymentResponse{
		repo:     rp,
		Document: resp.NewDocument(rp),
	}

	// @FIXME fake response text
	text := "Cancel bonus payment response"
	line := resp.TextLine().SetText(text).SetWrap(resp.WordWrap)

	// create info to display and printing
	r.Display = resp.NewSlip().AddLine(line)
	r.Printer = resp.NewPrint().AddSlip(r.Display)

	// set respond to response renderer
	render.Respond(rs, rq, r)
}
