<script>
  import { onMount } from "svelte";
  import { Route, Router, Link, navigate } from "svelte-routing";
  import "./app.css";
  import Base from "./routes/base/Base.svelte";
  import NotFound from "./routes/error/NotFound.svelte";
  import Register from "./routes/auth/Register.svelte";
  import Login from "./routes/auth/Login.svelte";

  let url = "";

  function getCookie(name) {
    const cookies = document.cookie.split(";");
    for (const cookie of cookies) {
      const [cookieName, cookieVal] = cookie.split("=");
      if (cookieName.trim() === name) {
        return decodeURIComponent(cookieVal);
      }
    }
  }
  onMount(() => {
    const cookie_value = getCookie("token");
    if (cookie_value) {
      navigate("/register");
    }
  });
</script>

<Router {url}>
  <Route path="/" component={Base} />
  <Route path="/register" component={Register} />
  <Route path="/login" component={Login} />
  <Route component={NotFound} />
</Router>