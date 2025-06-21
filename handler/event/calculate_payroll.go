package event

import (
	"attendance/model/event"
	"context"
	"encoding/json"
	"fmt"
)

// CalculatePayroll ...
func (h *Handler) CalculatePayroll(ctx context.Context, msg []byte) error {
	raw := &event.MessageFormat{}
	err := json.Unmarshal(msg, raw)
	if err != nil {
		return fmt.Errorf("failed to unmarshal raw message: %w", err)
	}

	data, _ := json.Marshal(raw.Data)
	updateData := event.CalculatePayrollJob{}
	err = json.Unmarshal(data, &updateData)
	if err != nil {
		return fmt.Errorf("failed to unmarshal data message: %w", err)
	}

	return h.Bootstrap.Service.Payroll.CalculatePayroll(ctx, updateData)
}
