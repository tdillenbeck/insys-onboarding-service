package wvaultdb

import (
	"context"
	"time"

	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wlog"
	"weavelab.xyz/monorail/shared/wlib/wmetrics"
	"weavelab.xyz/monorail/shared/wlib/wvault"
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

/* Refresh attempts to extend the database credentials lease, if
   the lease cannot be extend by at least half of the requested
   extension, new credentials will be requested.
   Returns bool (able to be refreshed), error
*/

var dbForceRecreateIfLessThan = time.Second * 10

func (c *Creator) Refresh(ctx context.Context, recreate bool) (bool, error) {

	// if the token is expiring soon, force recreate
	if c.ShouldForceRecreate() {
		if recreate == false {
			wlog.InfoC(ctx, "[VaultDB] credentials expire too soon, forcing refresh")
		}
		recreate = true
	}

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

	// renew the credentials we already have
	newLeaseDuration, retryable, err := c.renew(ctx)
	if err != nil {
		if retryable {
			return false, werror.Wrap(err)
		}
		// log and continue
		wlog.WErrorC(ctx, werror.Wrap(err, "[Vault] unable to renew credentials"))
	}

	// if there's not enough time on the new lease
	if newLeaseDuration < c.requestIncrement/2 {
		// recreate the credentials by calling refresh
		// and forcing a recreate
		return c.Refresh(ctx, true)
	}

	// if we reached the end, the existing credentials
	// were able to be successfully renewed and extended
	return true, nil

}

func (c *Creator) alive(ctx context.Context) error {

	current := wvault.Secret(c)
	for {
		e := current.Expiration()
		if wvault.Clock.Now().After(e) {
			return werror.New("token is expired").Add("secretName", current.Name())
		}

		current = current.Parent()

		if current == nil {
			break
		}
	}

	return nil
}
