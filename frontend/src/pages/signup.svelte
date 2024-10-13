<script>
  import { users, currentUser } from '../lib/userstore.js';
  import { navigate } from 'svelte-routing';
  import { Card, Button, Label, Input } from 'flowbite-svelte';
  import HoneyLemonLogo from '../assets/BankLogo.png';
  import { onMount } from "svelte";
  import axios from 'axios';

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
  let localUsers = [];

  users.subscribe(value => {
    localUsers = value;
  });

  let idcard = "";
  let fullname = "";
  let birthdate = "";
  let address = "";
  let password = "";
  let confirmpassword = "";

  async function signup(event) {
    event.preventDefault();

    // Validate ID card length
    if (idcard.length != 13) {
      alert("ID card is not correct");
      return;
    }

    // Validate password and confirm password match
    if (password !== confirmpassword) {
      alert("Passwords do not match");
      return;
    }

    // Check if ID card is already in use
    if (localUsers.some(user => user.idcard === idcard)) {
      alert("This ID card has already registered");
      return;
    }

    try {
      const response = await axios.post(
        "http://127.0.0.1:4000/esb/register", // Corrected URL
        {
          id: idcard,
          name: fullname,
          address: address,
          birthDate: birthdate,
          password: password,
        },
        {
          headers: {
            "Content-Type": "application/json",
          },
        },
      );

      if (response.data) {
        navigate("/login");
      } else {
        alert("Signup failed. Please try again.");
      }
    } catch (error) {
      console.error("Error during signup:", error);
      alert("An error occurred during signup. Please try again later.");
    }
  }
</script>

<Card>
  <form class="flex flex-col space-y-6" on:submit={signup}>
    <!-- Form fields -->
    <Label class="space-y-2">
      <span class="text-gray-400">Idcard</span>
      <Input type="text" name="idcard" bind:value={idcard} required />
    </Label>
    <Label class="space-y-2">
      <span class="text-gray-400">Full Name</span>
      <Input type="text" name="fullname" bind:value={fullname} required />
    </Label>
    <Label class="space-y-2">
      <span class="text-gray-400">Birthdate</span>
      <Input type="date" name="birthdate" bind:value={birthdate} required />
    </Label>
    <Label class="space-y-2">
      <span class="text-gray-400">Address</span>
      <Input type="text" name="address" bind:value={address} required />
    </Label>
    <!-- <Label class="space-y-2">
      <span class="text-gray-400">Email</span>
      <Input type="email" name="email" required />
    </Label> -->
    <Label class="space-y-2">
      <span class="text-gray-400">Password</span>
      <Input type="password" name="password" bind:value={password} required />
    </Label>
    <Label class="space-y-2">
      <span class="text-gray-400">Confirm password</span>
      <Input type="password" name="confirmpassword" bind:value={confirmpassword} required />
    </Label>
    <Button type="submit" class="w-full bg-[#28A745] hover:bg-[#28A745] hover:opacity-70">Signup</Button>
    <div class="text-sm font-medium text-gray-500 dark:text-gray-300">
      Already registered? <a href="/login" class="text-green-400 hover:text-green-500 dark:text-primary-500">Login</a>
    </div>
  </form>
</Card>