import { redirect, error } from "@sveltejs/kit";

export const actions = {
	default: async ({ cookies, request, fetch }) => {
		const formData = await request.formData();
		const email = formData.get("email");
		const password = formData.get("password");
		const username = formData.get("username");
		const fullname = formData.get("fullname");
		const body = await JSON.stringify({
			email,
			password,
			username,
			fullname
		});
		const resp = await fetch("http://127.0.0.1:1414/api/auth/register", {
			method: "POST",
			headers: {
				"Content-Type": "application/json"
			},
			body
		});
		// Check if response is ok
		try {
			if (resp.ok) {
				const response = await resp.json();
				const responseData = await response;
				console.log(responseData);
				cookies.set("session_token", responseData.token, {
					path: "/",
					httpOnly: true,
					maxAge: 30 * 2 * 24 * 60 * 60,
					secure: false,
					sameSite: "strict"
				});
				throw redirect(302, "/");
			} else {
				const response = await resp.json();
				const responseData = await response;
				return {
					error: responseData.error
				};
			}
		} catch (err) {
			throw error(502, "Something went wrong");
		}
	}
};
