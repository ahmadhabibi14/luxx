export const load = async (event) => {
  const session_token = event.cookies.get("session_token");
  const resp = await fetch(
    "http://127.0.0.1:1414/api/user/user-data",
    {
      method: "GET",
      headers: {
        "Authorization": `Bearer ${session_token}`
      },
    }
  );
  if (resp.ok) {
    const response = await resp.json();
    const responseData = await response;

    const user = responseData;
    console.log(response)
    return { user }
  } else {
    return {
      user: null
    }
  }
}