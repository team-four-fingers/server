package timezone

import "time"

var (
	KST = mustLoadKST()
)

func MustLoadLocation(name string) *time.Location {
	l, err := time.LoadLocation(name)
	if err != nil {
		panic(err)
	}
	return l
}

func mustLoadKST() *time.Location {
	return MustLoadLocation("Asia/Seoul")
}
