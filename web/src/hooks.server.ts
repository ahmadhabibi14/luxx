import type { Handle } from "@sveltejs/kit";

export const handle: Handle = async ({ event, resolve }) => {
  const session_token = event.cookies.get("session_token");
  try {
    const resp: Response = await fetch(
      "localhost:1414/api/user/user-data",
      {
        method: "GET",
        headers: {
          "Authorization": `Bearer ${session_token}`
        },
      }
    );
    // Check if response is ok
    if (resp.ok) {
      const response: Promise<any> = await resp.json();
      const responseData = JSON.parse(await response);

      event.locals.authedUser = responseData;
    } else {
      event.locals.authedUser = undefined;
    }
  } finally {
    const response = await resolve(event);
    return response;
  }
}