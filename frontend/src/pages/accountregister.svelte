<script>
  import { users, currentUser } from '../lib/userstore.js';
  import { navigate } from 'svelte-routing';
  import { Card, Button, Label, Input, Select } from 'flowbite-svelte';
  import { onMount } from 'svelte';

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
    if (!loggedIn) {
      navigate("/");
    }
  });

  let type = "";
  let pin = "";
  let confirmpin = "";

  function registerAccount(event) {
    event.preventDefault();

    // Validate pin length
    if (pin.length != 6) {
      alert("Pin must be 6 digits");
      return;
    }

    // Validate pin and confirm pin match
    if (pin !== confirmpin) {
      alert("Pins do not match");
      return;
    }

    console.log("type", event.target.acctype.value, "pin", pin);
    // Register account
    alert("Account registered successfully");
  }
</script>

<Card class="w-full max-w-lg mx-auto">
  <form class="flex flex-col space-y-6" on:submit={registerAccount}>
    <Label class="space-y-2">
      <div class="text-xs text-black mb-2 mr-4 w-full">
        Type
        <Select name="acctype" class="w-full mt-1 text-xs mt-2">
            <option value="Saving">Saving</option>
            <option value="Credit">Credit</option>
            <option value="Interest">Interest</option>
            <option value="Loan">Loan</option>
        </Select>
      </div>
    </Label>
    <Label class="space-y-2">
      <span class="text-gray-400">Set a 6 digit pin</span>
      <Input type="password" name="pin" bind:value={pin} required />
    </Label>
    <Label class="space-y-2">
      <span class="text-gray-400">Confirm Pin</span>
      <Input type="password" name="confirmpin" required />
    </Label>
    <Button type="submit" class="w-full bg-green-400 hover:bg-green-500">Register</Button>
  </form>
</Card>
