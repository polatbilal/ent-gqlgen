package scalars

import (
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/shopspring/decimal"
)

// Decimal wraps NullDecimal for GraphQL - NULL safe!
type Decimal = decimal.NullDecimal

// MarshalDecimal serializes NullDecimal to GraphQL
func MarshalDecimal(d Decimal) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		if !d.Valid {
			w.Write([]byte("null"))
			return
		}
		io.WriteString(w, fmt.Sprintf(`"%s"`, d.Decimal.String()))
	})
}

// UnmarshalDecimal deserializes GraphQL input to NullDecimal
func UnmarshalDecimal(v interface{}) (Decimal, error) {
	if v == nil {
		return decimal.NullDecimal{Valid: false}, nil
	}

	var dec decimal.Decimal
	var err error

	switch v := v.(type) {
	case string:
		dec, err = decimal.NewFromString(v)
	case int:
		dec = decimal.NewFromInt(int64(v))
	case int64:
		dec = decimal.NewFromInt(v)
	case float64:
		dec = decimal.NewFromFloat(v)
	default:
		return decimal.NullDecimal{}, fmt.Errorf("decimal formatı geçersiz: %T", v)
	}

	if err != nil {
		return decimal.NullDecimal{}, err
	}

	return decimal.NullDecimal{Decimal: dec, Valid: true}, nil
}
