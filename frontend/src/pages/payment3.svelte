<script>
  import { Link, navigate } from "svelte-routing";
  import { Card, Button, Label, Input, Checkbox } from "flowbite-svelte";
  import Succ from "../assets/successful.png";

  // Get data from location state
  const { payment, userFullName } = history.state || {};

  let tx = payment;

  function CurrentTime() {
    let date = new Date();
    let year = date.getFullYear();
    let month = date.getMonth() + 1;
    let day = date.getDate();
    let hour = date.getHours();
    let minute = date.getMinutes();
    let second = date.getSeconds();
    let time = `${year}-${month}-${day} ${hour}:${minute}:${second}`;
    return time;
  }

  console.log(tx);

  function handleBackClick() {
    navigate("/mainaccount");
  }
</script>

<Card>
  <form class="flex flex-col space-y-6" action="/">
    <div class="flex flex-col items-center">
      <img src={Succ} class="h-18 w-28" alt="HoneyLemonLogo" />
      <span class="text-black text-xl mt-2">Successful</span>
      <span class="text-gray-400 text-sm mt-2">{CurrentTime()}</span>
    </div>
    <div class="flex items-center justify-between">
      <Label class="space-y-2">
        <span class="text-black text-xl">From:</span>
      </Label>
      <Label class="space-y-2 flex flex-col mt-8">
        <span class="text-xl text-[#28A745]">{userFullName}</span>
        <span class="text-base text-[#666666]">{tx.SourceAccountID}</span>
      </Label>
    </div>
    <div class="flex items-center justify-between">
      <Label class="space-y-2">
        <span class="text-black text-xl">RefCode:</span>
      </Label>
      <Label class="space-y-2 flex flex-col mt-8">
        <span class="text-base text-[#666666]">{tx.ReferenceCode}</span>
      </Label>
    </div>
    <div class="flex items-center justify-between">
      <Label class="space-y-2">
        <span class="text-black text-xl">Amount:</span>
      </Label>
      <Label class="space-y-2 ">
        <span class="text-black text-xl">{tx.Amount}</span>
      </Label>
    </div>
    <Link to="/mainaccount">
      <Button type="submit" class="w-full bg-[#218838] hover:bg-[#28A745]" on:click={handleBackClick}>
        Back to Account page
      </Button>
    </Link>
  </form>
</Card>
