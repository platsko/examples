package respond

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"

	resp "lab/go-rest-api/app/http/respond/types"
	"lab/go-rest-api/app/repo"
)

type (
	// clientResponse describes response.
	clientResponse struct {
		repo    *repo.ClientRepo
		Code    int         `json:"code"`
		Display *resp.Slip  `json:"cashierInformation,omitempty"`
		Printer *resp.Print `json:"printingInformation,omitempty"`
	}
)

// Render satisfies renderer interface.
func (r *clientResponse) Render(_ http.ResponseWriter, rq *http.Request) error {
	render.Status(rq, r.Code)
	return nil
}

// ClientResource wraps resource to respond.
func ClientResource(rs http.ResponseWriter, rq *http.Request, rp *repo.ClientRepo) {
	// init confirm response
	r := &clientResponse{repo: rp}

	// @FIXME fake response text
	balance := rp.Balance()
	text := fmt.Sprintf("Account %s has %d bonuses.", rp.Phone, balance)
	line := resp.TextLine().SetText(text).SetWrap(resp.WordWrap)

	// create info to display and printing
	r.Display = resp.NewSlip().AddLine(line)
	r.Printer = resp.NewPrint().AddSlip(r.Display)

	// set respond to response renderer
	render.Respond(rs, rq, r)
}
