<script>
  import { Link } from "svelte-routing";
  import { Card, Button, Dropdown, DropdownItem } from "flowbite-svelte";
  import { DotsVerticalOutline } from "flowbite-svelte-icons";
  import { onMount } from "svelte";

  let userData = null;
  let accountData = null;
  let selectedAccount = null;

  // Function to switch accounts
  function switchAccount(accountId) {
      // logic for switching accounts
  }

  // Fetch user data and account data when component mounts
  onMount(() => {
      // Fetch or initialize userData and accountData here
  });
</script>

<div class="flex flex-col items-center">
  <Card class="bg-white mb-4" size="lg" padding="xl" style="width: 950px; height: 350px;">
    <div class="flex flex-col h-full">
      <div class="flex items-center justify-between mb-4">
        <div class="flex flex-col">
          <h5 class="text-[#004D00] mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
            {#if userData}
              {userData.name}
            {/if}
          </h5>
          <h6 class="mb-3 font-normal text-base text-gray-700 dark:text-gray-400 leading-tight">
            {userData && selectedAccount ? selectedAccount.id : "xxxxx"}
          </h6>
        </div>
        
        <!-- Sign In and Register Buttons -->
        <div class="ml-16 flex space-x-2">
          <Link to="/signin">
            <Button class="bg-blue-500 hover:bg-blue-600 text-white text-sm">Sign In</Button>
          </Link>
          <Link to="/register">
            <Button class="bg-green-500 hover:bg-green-600 text-white text-sm">Register</Button>
          </Link>
        </div>

        <div class="flex items-center">
          <DotsVerticalOutline class="dots-menu dark:text-white cursor-pointer mb-10" />
          <Dropdown triggeredBy=".dots-menu" class="bg-slate-100 rounded shadow-lg">
            {#if accountData && accountData.length > 0}
              {#each accountData as account}
                <DropdownItem
                  class="bg-white hover:bg-slate-50 text-gray-700"
                  on:click={() => switchAccount(account.id)}
                >
                  {account.id}
                </DropdownItem>
              {/each}
            {:else}
              <DropdownItem class="bg-white text-gray-500">No accounts available</DropdownItem>
            {/if}
            <DropdownItem slot="footer" class="bg-[#28A745] hover:bg-[#03C04A]">
              <Link to="/addaccount" class="w-full text-left block text-white dark:text-gray-400">Add new account</Link>
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
  </Card>
</div>
