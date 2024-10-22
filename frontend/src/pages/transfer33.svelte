<script>
    import { Card, Button, Label } from 'flowbite-svelte';
    import { navigate } from 'svelte-routing';
    import { currentUser } from '../lib/userstore.js';
  
    let user;
    currentUser.subscribe(value => {
      user = value;
    });
  
    // Get data from location state
    const { receiverAccountNumber, amount, pin } = history.state || {};
  
    // Find the selected account
    const selectedAccount = user.accounts.find(acc => acc.accountNumber === user.selectedAccount);
    if (selectedAccount) {
      // Deduct the amount from the selected account's balance
      selectedAccount.balance -= parseFloat(amount);
  
      // Update the user's accounts in local storage
      const updatedUsers = localStorage.getItem('users');
      const users = JSON.parse(updatedUsers).map(u => u.idcard === user.idcard ? user : u);
      localStorage.setItem('users', JSON.stringify(users));
  
      // Update the currentUser store
      currentUser.set(user);
    }
  
    const userFullName = user?.fullname || 'N/A';
    const fromAccountNumber = user?.selectedAccount || 'N/A';
    const newBalance = selectedAccount ? selectedAccount.balance : '0';
  
    function handleBackClick() {
      navigate('/mainaccount');
    }
  </script>
  
  <Card>
    <form class="flex flex-col space-y-6">
      <div class="flex flex-col items-center">
        <span class="text-black text-xl mt-2">Successful</span>
        <span class="text-gray-400 text-sm mt-2">01 August 2024 - 11:34</span>
      </div>
      <div class="flex items-center justify-between">
        <Label class="space-y-2">
          <span class="text-black text-xl">From:</span>
        </Label>
        <Label class="space-y-2 flex flex-col mt-8">
          <span class="text-xl text-[#28A745]">{userFullName}</span>
          <span class="text-base text-[#666666]">{fromAccountNumber}</span>
        </Label>
      </div>
      <div class="flex items-center justify-between">
        <Label class="space-y-2">
          <span class="text-black text-xl">To:</span>
        </Label>
        <Label class="space-y-2 flex flex-col">
          <span class="text-base text-[#666666]">{receiverAccountNumber || 'N/A'}</span>
        </Label>
      </div>
      <div class="flex items-center justify-between">
        <Label class="space-y-2">
          <span class="text-black text-xl">Amount:</span>
        </Label>
        <Label class="space-y-2 ">
          <span class="text-black text-xl ">{amount || 'N/A'}</span>
        </Label>
      </div>
      <div class="flex items-center justify-between">
        <Label class="space-y-2">
          <span class="text-black text-xl">Balance:</span>
        </Label>
        <Label class="space-y-2 ">
          <span class="text-black text-xl ">{newBalance}</span>
        </Label>
      </div>
      <Button type="button" class="w-full bg-[#218838] hover:bg-[#28A745]" on:click={handleBackClick}>
        Back to Account page
      </Button>
    </form>
  </Card>
  