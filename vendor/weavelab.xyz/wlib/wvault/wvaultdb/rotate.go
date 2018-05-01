package wvaultdb

import (
	"context"
	"time"

	"weavelab.xyz/wlib/werror"
	"weavelab.xyz/wlib/wlog"
	"weavelab.xyz/wlib/wmetrics"
	"weavelab.xyz/wlib/wvault"
)

const (
	metricAction = "wvaultDBAction"
)

func init() {
	wmetrics.SetLabels(metricAction, "action")
}

type refreshableCredentials interface {
	UpdateCredentials(string, string) error
}

func (c *Creator) setDB(db refreshableCredentials) {
	c.Lock()
	defer c.Unlock()

	c.db = db
}

/*
	wmetrics.Incr(1, metricAction, "renewed")
	wmetrics.Incr(1, metricAction, "recreated")
	wmetrics.Incr(1, metricAction, "revoked")

		err := r.db.UpdateCredentials(creds.Username, creds.Password)
*/

func (c *Creator) Refresh(ctx context.Context, recreate bool) (bool, error) {

	// if we were forced to recreate, recreate the database credentials
	// and respond indicating that we did not refresh existing credentials
	if recreate {
		err := c.createCredentials(ctx)
		if err != nil {
			return false, werror.Wrap(err)
		}

		// tell the database that the credentials changed
		err = c.db.UpdateCredentials(c.Username(), c.Password())
		if err != nil {
			return false, werror.Wrap(err, "unable to update credentials")
		}

		return false, nil
	}

	newLeaseDuration, retryable, err := c.renew(ctx)
	if err != nil && retryable {
		return false, werror.Wrap(err)
	} else if err != nil {
		// log and continue
		wlog.WErrorC(ctx, werror.Wrap(err, "[Vault]"))
	}

	// if there's enough time left on the lease that we
	// don't care to renew
	if newLeaseDuration >= c.requestIncrement/2 {
		return false, nil
	}

	// recreate the credentials
	err = c.createCredentials(ctx)
	if err != nil {
		return false, werror.Wrap(err, "unable to recreate credentials")
	}

	// tell the database that the credentials changed
	err = c.db.UpdateCredentials(c.Username(), c.Password())
	if err != nil {
		return false, werror.Wrap(err, "unable to update credentials")
	}

	return true, nil
}

func (c *Creator) alive(ctx context.Context) error {

	current := wvault.Secret(c)
	for {
		e := current.Expiration()
		if time.Now().After(e) {
			return werror.New("token is expired").Add("secretName", current.Name())
		}

		current = current.Parent()

		if current == nil {
			break
		}
	}

	return nil
}
