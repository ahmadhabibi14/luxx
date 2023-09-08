<script>
  import { onMount } from "svelte";
  import { Route, Router, navigate } from "svelte-routing";
  import "./app.css";
  import Base from "./routes/base/Base.svelte";
  import NotFound from "./routes/error/NotFound.svelte";
  import Register from "./routes/auth/Register.svelte";
  import Login from "./routes/auth/Login.svelte";

  let url = "";

  function getCookie(name) {
    var match = document.cookie.match(new RegExp("(^| )" + name + "=([^;]+)"));
    if (match) return match[2];
  }
  onMount(() => {
    const auth = getCookie("auth");
    if (auth) {
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
