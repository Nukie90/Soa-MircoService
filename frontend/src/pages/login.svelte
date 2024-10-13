<script>
  import { navigate } from "svelte-routing";
  import { Card, Button, Label, Input } from "flowbite-svelte";
  import HoneyLemonLogo from "../assets/BankLogo.png";
  import axios from "axios";
  import { onMount } from "svelte";

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
    if (loggedIn) {
      navigate("/mainaccount");
    }
  });

  let idcard = "";
  let password = "";

  async function login(event) {
    event.preventDefault();
    try {
      const response = await axios.post(
        "http://127.0.0.1:4000/esb/login",
        {
          id: idcard,
          password: password,
        },
        {
          headers: {
            "Content-Type": "application/json",
          },
          withCredentials: true, // Enable this to include cookies
        }
      );

      if (response.data) {
        navigate("/");
      } else {
        alert("Login failed. Please check your credentials and try again.");
      }
    } catch (error) {
      console.error("Error during login:", error);
      alert("An error occurred during login. Please try again later.");
    }
  }
</script>

<Card>
  <form class="flex flex-col space-y-6" on:submit={login}>
    <div class="flex items-center justify-between">
      <h3 class="text-4xl font-bold text-lime-900 dark:text-white">Login</h3>
      <img src={HoneyLemonLogo} class="h-18 w-28" alt="HoneyLemonLogo" />
    </div>
    <Label class="space-y-2">
      <span class="text-gray-400">Idcard</span>
      <Input type="text" name="idcard" bind:value={idcard} required />
    </Label>
    <Label class="space-y-2">
      <span class="text-gray-400">Password</span>
      <Input type="password" name="password" bind:value={password} required />
    </Label>
    <div class="flex items-start">
      <a
        href="/resetpassword"
        class="ml-auto text-sm text-green-400 hover:text-green-500 dark:text-primary-500"
        >Forgot password?</a
      >
    </div>
    <Button
      type="submit"
      class="w-full bg-[#28A745] hover:bg-[#28A745] hover:opacity-70"
      >Login</Button
    >
    <div class="text-sm font-medium text-gray-500 dark:text-gray-300">
      Not registered? <a
        href="/signup"
        class="text-green-400 hover:text-green-500 dark:text-primary-500"
        >Create account</a
      >
    </div>
  </form>
</Card>
