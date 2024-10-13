<script>
  import { Card, Button, Label, Input } from 'flowbite-svelte';
  import HoneyLemonLogo from '../assets/BankLogo.png';
  import { navigate } from 'svelte-routing';
  import { onMount } from 'svelte';
  import axios from 'axios';
 
  let loggedIn = false;

  let accountId = null;

  let pin = "";
  let newpin = "";
  let confirmnewpin = "";

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
    if (!loggedIn) {
      navigate("/");
    }

    const params = new URLSearchParams(window.location.search);
    accountId = params.get("accountId");
    console.log("Account ID:", accountId);
  });

  function changePin(event) {
    event.preventDefault();

    if (newpin.length != 6) {
      alert("Pin must be 6 digits");
      return;
    }

    if (newpin !== confirmnewpin) {
      alert("Pins do not match");
      return;
    }

    axios.patch(`http://127.0.0.1:4000/esb/accounts/update`, {
      id: accountId,
      oldPin: pin,
      newPin: newpin,
    }, {
      withCredentials: true,
    })
    .then((response) => {
      console.log(response);
      alert("Pin changed successfully");
      navigate("/mainaccount");
    })
    .catch((error) => {
      console.error(error);
      alert("Failed to change pin");
    });
  }
</script>

<Card class="w-full max-w-md mx-auto">
  <form class="flex flex-col space-y-6" on:submit={changePin}>
    <div class="flex items-center justify-between">
      <h3 class="text-4xl font-bold text-lime-900 dark:text-white">Change Pin</h3>
      <img src="{HoneyLemonLogo}" class="h-18 w-28" alt="HoneyLemonLogo" />
    </div>
    <Label class="space-y-2">
      <span class="text-gray-400">Your Selected Account</span>
      <Input type="text" name="selectedaccount" value="{accountId}" readonly />
    </Label>
    <Label class="space-y-2">
      <span class="text-gray-400">Enter your Old Pin</span>
      <Input type="password" name="oldpin" placeholder="Old PIN" bind:value={pin} binkrequired />
    </Label>
    <Label class="space-y-2">
      <span class="text-gray-400">Set a new 6 digit pin</span>
      <Input type="password" name="newpin" placeholder="New PIN" bind:value={newpin} required />
    </Label>
    <Label class="space-y-2">
      <span class="text-gray-400">Confirm new pin</span>
      <Input type="password" name="confirmnewpin" placeholder="Confirm new PIN" bind:value={confirmnewpin} required />
    </Label>
    <Button type="submit" class="w-full bg-green-400 hover:bg-green-500">Save Changes</Button>
  </form>
</Card>
