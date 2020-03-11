package erroraggregator

import (
	"errors"
	"testing"
)

func TestErrorAggregator(t *testing.T) {
	errAgg := New()

	if errAgg.GotErrors() == true {
		t.Errorf("error aggregator emits errors without errors")
		return
	}

	errAgg.Append(errors.New("Xablau"))

	if errAgg.GotErrors() != true {
		t.Errorf("error aggregator dont count errors")
		return
	}

	if errAgg.GetErrorMessages().Error() != "Xablau" {
		t.Errorf("error on generated error message %s", errAgg.GetErrorMessages())
		return
	}
}
