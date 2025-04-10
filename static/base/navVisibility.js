function getCookie(name) {
	const value = `; ${document.cookie}`
	const parts = value.split(`; ${name}=`)
	if (parts.length === 2) return parts.pop().split(";").shift()
	return null
}

function toggleNavElements(hasToken) {
	const guestNav = document.getElementById("guest-nav")
	const loggedInNav = document.getElementById("logged-in-nav")

	const mobileGuestNav = document.getElementById("mobile-guest-nav")
	const mobileLoggedInNav = document.getElementById("mobile-logged-in-nav")

	if (guestNav && loggedInNav) {
		guestNav.style.display = hasToken ? "none" : "block"
		loggedInNav.style.display = hasToken ? "block" : "none"
	}

	if (mobileGuestNav && mobileLoggedInNav) {
		mobileGuestNav.style.display = hasToken ? "none" : "block"
		mobileLoggedInNav.style.display = hasToken ? "block" : "none"
	}
}

function checkLoginStatus() {
	const token = getCookie("auth_token")
	toggleNavElements(!!token)
}

function logout() {
	document.cookie =
		"auth_token=; expires=Thu, 01 Jan 1970 00:00:00 GMT; path=/"
	checkLoginStatus()
	window.location.href = "/login"
}

document.addEventListener("DOMContentLoaded", checkLoginStatus)
