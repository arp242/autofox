package main

var shortcuts = map[string][][]string{
	"disable-onboarding": {
		[]string{"browser.rights.3.shown", "true"},
		[]string{"browser.onboarding.hidden", "true"},
		[]string{"general.warnOnAboutConfig", "false"},
		[]string{"browser.onboarding.notification.finished", "true"},
		[]string{"browser.onboarding.tour.onboarding-tour-addons.completed", "true"},
		[]string{"browser.onboarding.tour.onboarding-tour-customize.completed", "true"},
		[]string{"browser.onboarding.tour.onboarding-tour-default-browser.completed", "true"},
		[]string{"browser.onboarding.tour.onboarding-tour-private-browsing.completed", "true"},
		[]string{"browser.onboarding.tour.onboarding-tour-search.completed", "true"},
		[]string{"browser.onboarding.tour.onboarding-tour-sync.completed", "true"},
	},

	// TODO: do we need this?
	// privacy.history.custom                 true      # Custom privacy settings
	"reasonable-privacy": {
		[]string{"dom.battery.enabled", "false"},                // Serves no purpose except to track
		[]string{"network.IDN_show_punycode", "true"},           // Show punycode
		[]string{"network.cookie.cookieBehavior", "1"},          // Reject 3rd party cookies
		[]string{"network.http.referer.userControlPolicy", "1"}, // Send Referrer header only for same-origin
		[]string{"privacy.donottrackheader.enabled", "true"},    // Send Do-Not-Track header
	},

	"disable-pocket": {
		[]string{"extensions.pocket.enabled", "false"},
	},
}
