package serveo

import (
	"bytes"
	"context"
	"os/exec"
)

func RunServeo(ctx context.Context) error {

	cmd := exec.CommandContext(ctx, "", "")

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
