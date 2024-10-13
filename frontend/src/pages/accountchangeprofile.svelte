<script>
  import { onMount } from 'svelte';
  import { navigate } from 'svelte-routing';
  import { Card, Button, Label, Input } from 'flowbite-svelte';
  import { currentUser, users } from '../lib/userstore.js';
  import HoneyLemonLogo from '../assets/BankLogo.png';

  let user = null;
  let localUsers = [];

  // Subscribe to user and users store
  onMount(() => {
    currentUser.subscribe(value => {
      user = value;
    });

    users.subscribe(value => {
      localUsers = value;
    });
  });

  function saveChanges(event) {
    event.preventDefault();

    // Get updated values from form
    const fullname = event.target.fullname.value;
    const birthdate = event.target.birthdate.value;
    const address = event.target.address.value;
    const email = event.target.email.value;

    // Update user object
    if (user) {
      user.fullname = fullname;
      user.birthdate = birthdate;
      user.address = address;
      user.email = email;

      // Update the users store with the updated user
      const updatedUsers = localUsers.map(u => u.idcard === user.idcard ? user : u);
      users.set(updatedUsers);

      // Update currentUser store
      currentUser.set(user);

      alert("Profile updated successfully");
      navigate('/mainaccount')
    } else {
      alert("User not found");
    }
  }
</script>

<Card class="w-full max-w-lg mx-auto mt-20">
  <form class="flex flex-col space-y-6" on:submit={saveChanges}>
    <div class="flex items-center justify-between">
      <h3 class="text-4xl font-bold text-lime-900 dark:text-white">Edit Profile</h3>
      <img src="{HoneyLemonLogo}" class="h-18 w-28" alt="HoneyLemonLogo" />
    </div>
    <Label class="space-y-2">
      <span class="text-gray-400">ID Card</span>
      <Input type="text" name="idcard" value={user ? user.idcard : ''} readonly />
    </Label>
    <Label class="space-y-2">
      <span class="text-gray-400">Full Name</span>
      <Input type="text" name="fullname" value={user ? user.fullname : ''} />
    </Label>
    <Label class="space-y-2">
      <span class="text-gray-400">Birthdate</span>
      <Input type="date" name="birthdate" value={user ? user.birthdate : ''} />
    </Label>
    <Label class="space-y-2">
      <span class="text-gray-400">Address</span>
      <Input type="text" name="address" value={user ? user.address : ''} />
    </Label>
    <Label class="space-y-2">
      <span class="text-gray-400">Email</span>
      <Input type="email" name="email" value={user ? user.email : ''} />
    </Label>
    <Button type="submit" class="w-full bg-green-400 hover:bg-green-500">Save Changes</Button>
  </form>
</Card>
