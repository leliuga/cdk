package types

import (
	"fmt"
)

// BytesSize returns a human readable bytes size.
func BytesSize(n uint64) string {
	switch {
	case n < 1024:
		return fmt.Sprintf("%d B", n)
	case n < 1_048_576:
		return fmt.Sprintf("%.2f KiB", float64(n)/1024)
	case n < 1_073_741_824:
		return fmt.Sprintf("%.2f MiB", float64(n)/1_048_576)
	case n < 1_099_511_627_776:
		return fmt.Sprintf("%.2f GiB", float64(n)/1_073_741_824)
	case n < 1_125_899_906_842_624:
		return fmt.Sprintf("%.2f TiB", float64(n)/1_099_511_627_776)
	case n < 1_152_921_504_606_846_976:
		return fmt.Sprintf("%.2f PiB", float64(n)/1_125_899_906_842_624)
	default:
		return fmt.Sprintf("%.2f EiB", float64(n)/1_152_921_504_606_846_976)
	}
}
