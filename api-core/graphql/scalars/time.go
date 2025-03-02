package scalars

import (
	"fmt"
	"io"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalTime(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		// Sadece tarih formatı: YYYY-MM-DD
		// io.WriteString(w, t.Format(`"2006-01-02"`))
		// Alternatif format (DD-MM-YYYY) için:
		io.WriteString(w, t.Format(`"02/01/2006"`))
	})
}

func UnmarshalTime(v interface{}) (time.Time, error) {
	switch v := v.(type) {
	case string:
		// YYYY-MM-DD formatı için:
		return time.Parse("2006-01-02", v)
		// DD-MM-YYYY formatı için:
		// return time.Parse("02.01.2006", v)
	default:
		return time.Time{}, fmt.Errorf("geçersiz tarih formatı: %T", v)
	}
}
