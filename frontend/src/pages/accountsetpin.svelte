<script>
  import { Card, Button, Label, Input } from 'flowbite-svelte';
  import { navigate } from 'svelte-routing';
  import { currentUser } from '../lib/userstore.js';
  import HoneyLemonLogo from '../assets/BankLogo.png';

  let pin = '';
  let confirmPin = '';

  function setPin(event) {
    event.preventDefault();

    // Validate that the pins match
    if (pin !== confirmPin) {
      alert("Pins do not match");
      return;
    }

    // Add the pin to the currentUser
    currentUser.update(user => {
      user.pin = pin;
      return user;
    });

    alert("Pin set successfully");

    // Navigate to the main account page
    navigate('/mainaccount');
  }
</script>

<Card class="w-full max-w-md mx-auto">
  <form class="flex flex-col space-y-6" on:submit={setPin}>
    <div class="flex items-center justify-between">
      <h3 class="text-4xl font-bold text-lime-900 dark:text-white">Set Pin</h3>
      <img src="{HoneyLemonLogo}" class="h-18 w-28" alt="HoneyLemonLogo" />
    </div>
    <Label class="space-y-2">
      <span class="text-gray-400">Set a 6 digit pin</span>
      <Input type="password" bind:value={pin} placeholder="xxxxxx" required />
    </Label>
    <Label class="space-y-2">
      <span class="text-gray-400">Confirm Pin</span>
      <Input type="password" bind:value={confirmPin} placeholder="xxxxxx" required />
    </Label>
    <Button type="submit" class="w-full bg-green-400 hover:bg-green-500">Next</Button>
  </form>
</Card>
