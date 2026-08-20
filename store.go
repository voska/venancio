// Package venancio describes the Drogaria Venancio storefront
// (www.drogariavenancio.com.br), a pharmacy chain in Rio de Janeiro.
package venancio

import "github.com/voska/vtexkit/store"

// Store is the Drogaria Venancio descriptor.
//
// Three fields is the whole thing: this is a stock VTEX store and everything
// else is discovered at runtime. Probed live on 2026-08-20 —
//
//   - Account derives correctly from the base URL. Unlike Mantiqueira, whose
//     account is grupomantiqueira, this store's account really is
//     drogariavenancio, so no override is needed. The CLI is still named
//     "venancio" because Name drives the binary, not the account.
//   - Classic password and emailed access code are both enabled, and the OAuth
//     providers (Google, Facebook) are ones VTEX ID handles generically, so no
//     OAuthDriver is needed.
//   - Intelligent Search REST and the catalog API both return results, so
//     Search stays at the default SearchAuto.
//   - Subscriptions are enabled (vtex.subscriptions-graphql@3.x installed, RNS
//     answers 401 rather than 404), so `venancio subs` works with no extra
//     configuration — recurring medication is the obvious use.
//
// No MinOrder and no Quirks: there is no evidence for either at this store.
//
// No Wishlist hashes. They must be captured per store from a live browser
// session; the zero value makes `fav` report the wishlist as unreachable
// rather than failing strangely.
var Store = store.Store{
	Name:        "venancio",
	DisplayName: "Drogaria Venancio",
	BaseURL:     "https://www.drogariavenancio.com.br",
}
