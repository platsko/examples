package respond

import (
	"net/http"

	"github.com/go-chi/render"

	resp "lab/go-rest-api/app/http/respond/types"
	"lab/go-rest-api/app/repo"
)

type (
	// calculateResponse describes response.
	calculateResponse struct {
		repo     *repo.DocumentRepo
		Code     int            `json:"code"`
		Client   *resp.Client   `json:"client"`
		Document *resp.Document `json:"document"`
		Display  *resp.Slip     `json:"cashierInformation,omitempty"`
		Printer  *resp.Print    `json:"printingInformation,omitempty"`
	}
)

// Render satisfies renderer interface.
func (r *calculateResponse) Render(_ http.ResponseWriter, rq *http.Request) error {
	render.Status(rq, r.Code)
	return nil
}

// CalculateResource wraps resource to respond.
func CalculateResource(rs http.ResponseWriter, rq *http.Request, rp *repo.DocumentRepo) {
	// init client repo
	c := repo.NewClient(&rp.Client)

	// init calculate response
	cr := &calculateResponse{
		repo:     rp,
		Client:   resp.NewClient(c),
		Document: resp.NewDocument(rp),
	}

	// @FIXME fake response text
	text := "Document response"
	line := resp.TextLine().SetText(text).SetWrap(resp.WordWrap)

	// create info to display and printing
	cr.Display = resp.NewSlip().AddLine(line)
	cr.Printer = resp.NewPrint().AddSlip(cr.Display)

	// set respond to response renderer
	render.Respond(rs, rq, cr)
}
