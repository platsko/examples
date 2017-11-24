package respond

import (
	"net/http"

	"github.com/go-chi/render"

	resp "lab/go-rest-api/app/http/respond/types"
	"lab/go-rest-api/app/repo"
)

type (
	// confirmResponse describes response.
	confirmResponse struct {
		repo    *repo.DocumentRepo
		Code    int         `json:"code"`
		Display *resp.Slip  `json:"cashierInformation,omitempty"`
		Printer *resp.Print `json:"printingInformation,omitempty"`
	}
)

// Render satisfies renderer interface.
func (r *confirmResponse) Render(_ http.ResponseWriter, rq *http.Request) error {
	render.Status(rq, r.Code)
	return nil
}

// ConfirmResponse wraps resource to respond.
func ConfirmResponse(rs http.ResponseWriter, rq *http.Request, rp *repo.DocumentRepo) {
	// init confirm response
	r := &confirmResponse{repo: rp}

	// @FIXME fake response text
	text := "Confirm response"
	line := resp.TextLine().SetText(text).SetWrap(resp.WordWrap)

	// create info to display and printing
	r.Display = resp.NewSlip().AddLine(line)
	r.Printer = resp.NewPrint().AddSlip(r.Display)

	// set respond to response renderer
	render.Respond(rs, rq, r)
}
