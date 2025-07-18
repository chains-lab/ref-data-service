package cli

import (
	"context"

	"github.com/chains-lab/ref-data-svc/internal/config"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

func Start(ctx context.Context, cfg config.Config, log *logrus.Logger, app *app.App) error {
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error { return api.Run(ctx, cfg, log, app) })

	return eg.Wait()
}
