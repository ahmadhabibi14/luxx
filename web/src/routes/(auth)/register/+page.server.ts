import { redirect, type Actions } from "@sveltejs/kit";

export const actions: Actions = {
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
    const resp: Response = await fetch(
      "http://127.0.0.1:1414/api/auth/register",
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body
      }
    );
    // Check if response is ok
    if (resp.ok) {
      const response: Promise<any> = await resp.json();
      const responseData = JSON.parse(await response);
      cookies.set(
        "session_token",
        responseData.token,
        {
          path: "/",
          httpOnly: true,
          maxAge: 30 * 2 * 24 * 60 * 60,
          secure: false,
          sameSite: "strict",
        }
      )
      throw redirect(302, "/");
    } else {
      const response: Promise<any> = await resp.json();
      const responseData = JSON.parse(await response);
      return {
        error: responseData.error
      }
    };
  }
};