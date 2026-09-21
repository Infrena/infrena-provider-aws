package awsprov

import "testing"

// TestTheDeclaredCeilingIsWellUnderTheAccountQuota.
//
// Cloud Control allows about 137 simultaneous requests, and one infrena
// operation is at most one simultaneous request: crud.go sends a single call
// and await() polls one at a time with backoff between, so N concurrent
// operations peak at N requests and usually fewer.
//
// The bound asserted here is the REASONING, not the number. A future edit that
// raises the ceiling past a quarter of the quota is either wrong or has new
// information, and either way should have to change this line and say why.
func TestTheDeclaredCeilingIsWellUnderTheAccountQuota(t *testing.T) {
	const (
		// Simultaneous Cloud Control requests the account allows.
		accountQuota = 137
		// The host's own per-provider default. A declared ceiling REPLACES it,
		// so declaring this number would be a restatement — the SDK is explicit
		// that a guess dressed as a claim is worth less than silence.
		hostDefault = 8
	)
	got := NewPlugin().MaxConcurrency()

	if got <= hostDefault {
		t.Errorf("MaxConcurrency() = %d, which is not more than the host's own default of %d — "+
			"declaring it claims nothing the host did not already assume", got, hostDefault)
	}
	if got*4 > accountQuota {
		t.Errorf("MaxConcurrency() = %d, which is more than a quarter of the ~%d simultaneous "+
			"requests the account allows. The quota belongs to the account, not to this process: "+
			"a colleague's apply, a CI run and the console draw on the same budget",
			got, accountQuota)
	}
}
