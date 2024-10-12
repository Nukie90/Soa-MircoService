<script>
  import { Link } from "svelte-routing";
  import { Card, Button, Dropdown, DropdownItem, DropdownHeader, Select } from 'flowbite-svelte';
  import { Modal, Label, Input } from 'flowbite-svelte';
  import { ChevronDownOutline, ChevronRightOutline } from 'flowbite-svelte-icons';
  import { onMount } from "svelte";

  let userData = null;
  let accountData = null;
  let selectedAccount = null;
  let pin = ""; 
  let popupModal_login = false; 
  let popupModal_signup = false; 
  let popupModal_accountregister = false;
  let popupModal_topup = false;
  let popupModal_changepin = false;
  let popupModal_askdelete = false;

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
    popupModal_accountregister = true; 
    popupModal_signup = false; 
  }

  function openRegisterModalforLogin() {
    popupModal_accountregister = true; 
    popupModal_login = false; 
  }

  function closeRegisterModal() {
    popupModal_accountregister = false; 
  }

  function openTopupModal() {
    popupModal_topup = true;
  }

  function closeTopupModal() {
    popupModal_topup = false;
  }

  function openChangepinModal() {
    popupModal_changepin = true;
  }

  function closeChangepinModal() {
    popupModal_changepin = false;
  }

  function openAskDeleteModal() {
    popupModal_askdelete = true;
  }

  function closeAskDeleteModal() {
    popupModal_askdelete = false;
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

          <!-- if not login -->
          <!-- <Dropdown>
            <div slot="header" class="px-4 py-2 bg-blue-500 rounded-t-lg">
              <span class="block text-sm text-gray-900 dark:text-white font-semibold">John Doe</span>
            </div>
            <DropdownItem>
              <Button on:click={openLoginModal} class="w-full text-left">Login</Button> 
            </DropdownItem>
            <DropdownItem>
              <Button on:click={openSignupModal} class="w-full text-left">Signup</Button> 
            </DropdownItem>
          </Dropdown> -->

          <!-- if login -->
          <Dropdown>
            <div slot="header" class="px-4 py-2 bg-blue-500 rounded-t-lg">
              <span class="block text-sm text-gray-900 dark:text-white font-semibold">John Doe</span>
            </div>
            <DropdownItem>
              <Button on:click={openTopupModal} class="w-40 text-left">Top up</Button> 
            </DropdownItem>
            <DropdownItem class="flex items-center justify-between">
              Dropdown<ChevronRightOutline class="ms-2 text-primary-700 dark:text-white" style="height: 30px;" />
            </DropdownItem>
            <Dropdown placement="right-start">
              <DropdownItem>Acc1</DropdownItem>
              <DropdownItem>Acc2</DropdownItem>
              <DropdownItem>Acc3</DropdownItem>
            </Dropdown>
            <DropdownItem slot="footer">
              <Button class="w-full text-left">Log out</Button> 
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
        <Button class="w-40 h-9 bg-[#cccccc] hover:bg-slate-100 text-black" on:click={openChangepinModal}>Change PIN</Button>
        <Button class="w-40 h-9 bg-red-400 hover:bg-red-700 text-black" on:click={openAskDeleteModal}>Delete Account</Button>
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

    {#if popupModal_accountregister}
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

    {#if popupModal_topup}
      <div class="fixed inset-0 flex items-center z-50">
        <div class="bg-white rounded-lg shadow-lg p-6 w-96 relative" style="left: 400%;">
          <h3 class="text-3xl font-bold text-black text-center mb-4">Topup</h3>
          <form on:submit={registerAccount} class="flex flex-col space-y-4">
            <Label class="space-y-2">
              <span class="text-gray-400">Enter account no.</span>
              <Input type="text" name="topupacc" class="bg-blue-500 px-4 py-2" required />
            </Label>
            <Label class="space-y-2">
              <span class="text-gray-400">Amount</span>
              <Input type="text" name="topupamount" class="bg-blue-500 px-4 py-2" required />
            </Label>
            <Button type="submit" class="w-full bg-[#28A745] hover:bg-[#28A745] hover:opacity-70">Confirm</Button>
            <Button on:click={closeTopupModal} class="w-full bg-red-400 hover:bg-red-700 hover:opacity-70">Cancel</Button>
          </form>
        </div>
      </div>
    {/if}

    {#if popupModal_askdelete}
      <div class="fixed inset-0 flex items-center z-50">
        <div class="bg-white rounded-lg shadow-lg p-6 w-96 relative" style="left: 400%;">
          <h3 class="text-3xl font-bold text-black text-center mb-4">Delete account?</h3>
          <form on:submit={registerAccount} class="flex flex-col space-y-4">
            <Button type="submit" class="w-full bg-[#28A745] hover:bg-[#28A745] hover:opacity-70">Confirm</Button>
            <Button on:click={closeAskDeleteModal} class="w-full bg-red-400 hover:bg-red-700 hover:opacity-70">Cancel</Button>
          </form>
        </div>
      </div>
    {/if}

    {#if popupModal_changepin}
      <div class="fixed inset-0 flex items-center z-50">
        <div class="bg-white rounded-lg shadow-lg p-6 w-96 relative" style="right: 110%;">
          <h3 class="text-3xl font-bold text-black text-center mb-4">Change Pin</h3>
          <form on:submit={registerAccount} class="flex flex-col space-y-4">
            <Label class="space-y-2">
              <span class="text-gray-400">Your Selected Account</span>
              <Input type="text" name="selectedaccount" value="Lali" class="bg-blue-500 px-4 py-2" readonly />
            </Label>
            <Label class="space-y-2">
              <span class="text-gray-400">Enter your Old Pin</span>
              <Input type="text" name="topupacc" class="bg-blue-500 px-4 py-2" required />
            </Label>
            <Label class="space-y-2">
              <span class="text-gray-400">Set a new 6 digit pin</span>
              <Input type="text" name="topupamount" class="bg-blue-500 px-4 py-2" required />
            </Label>
            <Label class="space-y-2">
              <span class="text-gray-400">Confirm new pin</span>
              <Input type="text" name="topupamount" class="bg-blue-500 px-4 py-2" required />
            </Label>
            <Button type="submit" class="w-full bg-[#28A745] hover:bg-[#28A745] hover:opacity-70">Confirm</Button>
            <Button on:click={closeChangepinModal} class="w-full bg-red-400 hover:bg-red-700 hover:opacity-70">Cancel</Button>
          </form>
        </div>
      </div>
    {/if}
  </Card>
</div>
