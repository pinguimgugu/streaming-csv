package resource

import (
	"encoding/csv"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/pinguimgugu/streaming-csv/domain/contract"
)

// StreamingGetUser struct
type StreamingGetUser struct {
	UserRepo contract.UserRepository
}

// Handle method
func (a *StreamingGetUser) Handler(c echo.Context) error {
	res := c.Response()
	w := csv.NewWriter(res)
	header := res.Header()
	header.Set(echo.HeaderContentType, echo.MIMEOctetStream)
	header.Set(echo.HeaderContentDisposition, "attachment; filename="+"users.csv")
	header.Set("Content-Transfer-Encoding", "binary")
	header.Set("Expires", "0")
	w.Write([]string{"ID", "NAME", "AGE"})
	res.WriteHeader(http.StatusOK)

	for record := range a.UserRepo.StreamingAllUsers() {
		if err := w.Write([]string{strconv.Itoa(record.ID), record.Name, strconv.Itoa(record.Age)}); err != nil {
			return nil
		}
		w.Flush()
	}

	return nil
}
