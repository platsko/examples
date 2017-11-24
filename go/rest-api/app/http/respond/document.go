package respond

import (
	"net/http"

	"github.com/go-chi/render"

	resp "lab/go-rest-api/app/http/respond/types"
	"lab/go-rest-api/app/repo"
)

type (
	// documentResponse describes response.
	documentResponse struct {
		repo     *repo.DocumentRepo
		Code     int            `json:"code"`
		Client   *resp.Client   `json:"client"`
		Document *resp.Document `json:"document"`
		Display  *resp.Slip     `json:"cashierInformation,omitempty"`
		Printer  *resp.Print    `json:"printingInformation,omitempty"`
	}
)

// Render satisfies renderer interface.
func (r *documentResponse) Render(_ http.ResponseWriter, rq *http.Request) error {
	render.Status(rq, r.Code)
	return nil
}

// DocumentResource wraps resource to respond.
func DocumentResource(rs http.ResponseWriter, rq *http.Request, rp *repo.DocumentRepo) {
	// init client repo
	cr := repo.NewClient(&rp.Client)

	// init document response
	dr := &documentResponse{
		repo:     rp,
		Client:   resp.NewClient(cr),
		Document: resp.NewDocument(rp),
	}

	// @TODO text to respond
	text := "Document response"
	line := resp.TextLine().SetText(text).SetWrap(resp.WordWrap)

	// init a slip info to display and printing
	dr.Display = resp.NewSlip().AddLine(line)
	dr.Printer = resp.NewPrint().AddSlip(dr.Display)

	// set respond to response renderer
	render.Respond(rs, rq, dr)
}
