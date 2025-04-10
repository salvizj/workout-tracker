document.addEventListener("DOMContentLoaded", () => {
	const form =
		document.getElementById("signUp-form") ||
		document.getElementById("login-form")

	if (form) {
		const isSignUp = form.id === "signUp-form"

		form.addEventListener("submit", (e) => {
			handleAuthSubmit(e, isSignUp)
		})

		form.querySelectorAll("input").forEach((input) => {
			input.addEventListener("input", (e) => {
				clearError(e.target)
			})
		})
	}
})

document.querySelectorAll("input").forEach((input) => {
	input.addEventListener("input", (e) => {
		clearError(e.target)
	})
})

function handleAuthSubmit(e, isSignUp = false) {
	e.preventDefault()
	const form = e.target

	const formData = new FormData(form)
	const email = formData.get("email")?.trim() || ""
	const password = formData.get("password") || ""
	const confirmPassword = isSignUp
		? formData.get("confirmPassword") || ""
		: undefined

	const error = validateAuthForm({
		email,
		password,
		confirmPassword,
		isSignUp,
	})
	if (error) {
		showFormError(form, error)
		return false
	}
	form.submit()
}

function validateAuthForm({ email, password, confirmPassword, isSignUp }) {
	const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/

	if (!emailRegex.test(email)) {
		return "Invalid email format."
	}

	if (password.length < 8) {
		return "Password must be at least 8 characters."
	}

	if (isSignUp) {
		if (confirmPassword.length < 8) {
			return "Confirm password must be at least 8 characters."
		}

		if (password !== confirmPassword) {
			return "Password and confirm password must match."
		}
	}

	return null
}

function showFormError(form, message) {
	const errorBlock = form.querySelector(".error-block")
	if (!errorBlock) return

	errorBlock.textContent = message
	errorBlock.style.display = "block"
}

function clearError(inputElement) {
	const form = inputElement.closest("form")
	const errorBlock = form?.querySelector(".error-block")
	if (!errorBlock) return

	errorBlock.textContent = ""
	errorBlock.style.display = "none"
}
