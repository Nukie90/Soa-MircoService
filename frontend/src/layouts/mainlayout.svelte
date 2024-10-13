<script>
  import { onMount } from "svelte";
  import { Link, navigate } from "svelte-routing";
  import HoneyLemonLogo from "../assets/BankLogo.png";

  let loggedIn = false; // Default state for logged in status

  function checkLoginStatus() {
    // Check for the presence of a specific cookie
    const cookies = document.cookie.split(";").map((cookie) => cookie.trim());
    const authCookie = cookies.find((cookie) =>
      cookie.startsWith("esb_token=")
    );

    loggedIn = !!authCookie;
  }

  onMount(() => {
    checkLoginStatus();
    if (loggedIn){
      navigate("/mainaccount")
    }
  });
</script>

<div
  class="navbar bg-[#F0F0F0] h-24 flex items-center justify-between px-5 fixed top-0 left-0 right-0 shadow-lg z-50"
>
  <div class="flex flex-row items-center">
    <div class="ml-1 mr-5">
      <Link to="/">
        <img src={HoneyLemonLogo} class="h-9" alt="HoneyLemonLogo" />
      </Link>
    </div>
  </div>
  <div class="flex space-x-3 items-center">
    <Link to="/login">
      <button
        class="bg-white text-black px-4 py-2 rounded transition-all duration-300 ease-in-out hover:opacity-70"
        >Log in</button
      >
    </Link>
    <Link to="/signup">
      <button
        class="bg-[#28A745] text-white px-4 py-2 rounded mr-1 transition-all duration-300 ease-in-out hover:opacity-70"
        >Sign up</button
      >
    </Link>
  </div>
</div>

<div class="pt-24">
  <slot></slot>
</div>
