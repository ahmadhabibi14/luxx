import type { PageServerLoad } from "./$types";
import { redirect } from "@sveltejs/kit";
import type { User } from "../types/user";

export const load: PageServerLoad = async (event) => {
  const session_token = event.cookies.get("session_token");
  if (!session_token) {
    throw redirect(302, "/");
  }
  const resp: Response = await fetch(
    "http://127.0.0.1:1414/api/user/user-data",
    {
      method: "GET",
      headers: {
        "Authorization": `Bearer ${session_token}`
      },
    }
  );
  if (resp.ok) {
    const response: Promise<any> = await resp.json();
    const responseData = JSON.parse(await response);

    const user: User = responseData;
    return { user }
  }
}