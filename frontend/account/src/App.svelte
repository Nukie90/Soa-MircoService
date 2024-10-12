<script>
  import { Link } from "svelte-routing";
  import { Card, Button, Dropdown, DropdownItem, DropdownHeader, Select } from 'flowbite-svelte';
  import { Modal, Label, Input } from 'flowbite-svelte';
  import { ChevronDownOutline } from 'flowbite-svelte-icons';
  import { onMount } from "svelte";

  let userData = null;
  let accountData = null;
  let selectedAccount = null;
  let pin = ""; // State for PIN input
  let popupModal_login = false; // State for login modal
  let popupModal_signup = false; // State for signup modal
  let popupModal_register = false; // State for registration modal

  function openLoginModal() {
    popupModal_login = true;
    popupModal_signup = false;
  }

  function closeLoginModal() {
    popupModal_login = false;
  }

  function openSignupModal() {
    popupModal_signup = true;
    popupModal_login = false;
  }

  function closeSignupModal() {
    popupModal_signup = false;
  }

  function openRegisterModalforSignup() {
    popupModal_register = true; // Open the registration modal
    popupModal_signup = false; // Ensure signup modal is closed
  }

  function openRegisterModalforLogin() {
    popupModal_register = true; // Open the registration modal
    popupModal_login = false; // Ensure signup modal is closed
  }

  function closeRegisterModal() {
    popupModal_register = false; // Close the registration modal
  }

  function registerAccount(event) {
    event.preventDefault(); // Prevent form submission
    // Implement account registration logic here
  }

  onMount(() => {
    // Fetch or initialize userData and accountData here
  });
</script>

<div class="flex flex-col items-center">
  <Card class="bg-white mb-4" size="lg" padding="xl" style="width: 950px; height: 350px;">
    <div class="flex flex-col h-full">
      <div class="flex items-center justify-between mb-6">
        <div class="flex flex-col">
          <h5 class="text-[#004D00] mb-1 text-3xl font-bold tracking-tight text-gray-900 dark:text-white">
            {#if userData}
              {userData.name}
            {/if}
          </h5>
          <h6 class="mb-3 font-normal text-lg text-gray-700 dark:text-gray-400 leading-tight">
            {userData && selectedAccount ? selectedAccount.id : "xxxxx"}
          </h6>
        </div>

        <div class="flex items-center space-x-3">
          <Button class="flex items-center">
            ... <ChevronDownOutline class="w-4 h-4 ms-2 text-gray-700 dark:text-gray-400" />
          </Button>
          <Dropdown>
            <div slot="header" class="px-4 py-2 bg-blue-500 rounded-t-lg">
              <span class="block text-sm text-gray-900 dark:text-white font-semibold">Bonnie Green</span>
            </div>
            <DropdownItem>
              <Button on:click={openLoginModal} class="w-full text-left">Login</Button> 
            </DropdownItem>
            <DropdownItem>
              <Button on:click={openSignupModal} class="w-full text-left">Signup</Button> 
            </DropdownItem>
          </Dropdown>
        </div>
      </div>
      <div class="flex flex-1 items-center justify-center">
        <div class="flex flex-col items-center mt-[-70px]">
          <div class="flex flex-col items-center justify-center w-60 h-60 bg-[#28A745] rounded-full border-2 border-gray-300">
            <span class="text-xl font-bold text-white block">Available Balance</span>
            <span class="text-xl font-medium text-white block mt-4">
              {userData && selectedAccount ? selectedAccount.balance || 0 : "0 "} ฿
            </span>
          </div>
        </div>
      </div>

      <div class="flex flex-1 items-center justify-between">
        <Button class="w-40 h-9 bg-[#cccccc] hover:bg-slate-100 text-black">Change PIN</Button>
        <Button class="w-40 h-9 bg-red-400 hover:bg-red-700 text-black">Delete Account</Button>
      </div>
    </div>

    {#if popupModal_login}
      <div class="fixed inset-0 flex items-center z-50">
        <div class="fixed inset-0 bg-black opacity-50" on:click={closeLoginModal}></div>
        <div class="bg-white rounded-lg shadow-lg p-6 w-96 relative" style="left: 400%;">
          <button class="absolute top-2 right-2 text-gray-500" style="left: 78%;" on:click={closeLoginModal}>&times;</button>
          <h3 class="text-3xl font-bold text-black text-center mb-4">Login</h3>
          <form class="flex flex-col space-y-4">
            <Label class="space-y-1">
              <span class="text-gray-400">Id card</span>
              <Input type="text" name="idcard" required class="bg-blue-500 px-4 py-2" />
            </Label>
            <Label class="space-y-1">
              <span class="text-gray-400">Password</span>
              <Input type="password" name="password" required class="bg-blue-500 px-4 py-2" />
            </Label>
            <Button type="submit" class="w-full bg-[#28A745] hover:bg-[#28A745] hover:opacity-70" on:click={openRegisterModalforLogin}>Login</Button>
            <div class="text-sm font-medium text-gray-500 dark:text-gray-300 text-center">
              Not registered? <span class="text-green-400 hover:text-green-500 dark:text-primary-500 cursor-pointer" on:click={openSignupModal}>Create account</span>
            </div>
          </form>
        </div>
      </div>
    {/if}

    {#if popupModal_signup}
      <div class="fixed inset-0 flex items-center z-50">
        <div class="fixed inset-0 bg-black opacity-50" on:click={closeSignupModal}></div>
        <div class="bg-white rounded-lg shadow-lg p-6 w-96 relative" style="left: 400%;">
          <button class="absolute top-2 right-2 text-gray-500" style="left: 78%;" on:click={closeSignupModal}>&times;</button>
          <h3 class="text-3xl font-bold text-black text-center mb-4">Signup</h3>
          <form class="flex flex-col space-y-6">
            <Label class="space-y-2">
              <span class="text-gray-400">Id card</span>
              <Input type="text" name="idcard" class="bg-blue-500 px-4 py-2" required />
            </Label>
            <Label class="space-y-2">
              <span class="text-gray-400">Full Name</span>
              <Input type="text" name="fullname" class="bg-blue-500 px-4 py-2" required />
            </Label>
            <Label class="space-y-2">
              <span class="text-gray-400">Birthdate</span>
              <Input type="date" name="birthdate" class="bg-blue-500 px-4 py-2" required />
            </Label>
            <Label class="space-y-2">
              <span class="text-gray-400">Address</span>
              <Input type="text" name="address" class="bg-blue-500 px-4 py-2" required />
            </Label>
            <Label class="space-y-2">
              <span class="text-gray-400">Password</span>
              <Input type="password" name="password" class="bg-blue-500 px-4 py-2" required />
            </Label>
            <Label class="space-y-2">
              <span class="text-gray-400">Confirm password</span>
              <Input type="password" name="confirmpassword" class="bg-blue-500 px-4 py-2" required />
            </Label>
            <Button type="button" class="w-full bg-[#28A745] hover:bg-[#28A745] hover:opacity-70" on:click={openRegisterModalforSignup}>Register Account</Button>
            <div class="text-sm font-medium text-gray-500 dark:text-gray-300 text-center">
              Already registered? <span class="text-green-400 hover:text-green-500 dark:text-primary-500 cursor-pointer" on:click={openLoginModal}>Login</span>
            </div>
          </form>
        </div>
      </div>
    {/if}

    {#if popupModal_register}
      <div class="fixed inset-0 flex items-center z-50">
        <div class="bg-white rounded-lg shadow-lg p-6 w-96 relative" style="left: 400%;">
          <h3 class="text-3xl font-bold text-black text-center mb-4">Register Account</h3>
          <form on:submit={registerAccount} class="flex flex-col space-y-4">
            <Label class="space-y-2">
              <div class="text-xs text-black mb-2 mr-4 w-full">
                Type
                <Select name="acctype" class="bg-blue-500 px-4 py-2">
                    <option value="Saving">Saving</option>
                    <option value="Credit">Credit</option>
                    <option value="Interest">Interest</option>
                    <option value="Loan">Loan</option>
                </Select>
              </div>
            </Label>
            <Label class="space-y-2">
              <span class="text-gray-400">Set a 6 digit pin</span>
              <Input type="password" name="pin" bind:value={pin} class="bg-blue-500 px-4 py-2" required />
            </Label>
            <Label class="space-y-2">
              <span class="text-gray-400">Confirm Pin</span>
              <Input type="password" name="confirmpin" class="bg-blue-500 px-4 py-2" required />
            </Label>
            <Button type="submit" class="w-full bg-[#28A745] hover:bg-[#28A745] hover:opacity-70">Register</Button>
          </form>
        </div>
      </div>
    {/if}
  </Card>
</div>
