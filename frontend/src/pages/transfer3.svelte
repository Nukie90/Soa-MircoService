<script>
  import { Card, Button, Label } from "flowbite-svelte";
  import { navigate } from "svelte-routing";

  // Get data from location state
  const { transaction, senderName, receiverName } = history.state || {};

  let tx = transaction.transaction
  // Function to format date to the desired format
  function formatDate(dateString) {
    const date = new Date(dateString);
    // Convert to ISO string and remove milliseconds and timezone info
    let newDate = date.toISOString().split(".")[0].split("T");
    return newDate[1] + " - " + newDate[0]
  }

  console.log(tx)

  function handleBackClick() {
    navigate("/mainaccount");
  }
</script>

<Card>
  <form class="flex flex-col space-y-6">
    <div class="flex flex-col items-center">
      <span class="text-black text-xl mt-2">Successful</span>
      <span class="text-gray-400 text-sm mt-2">{formatDate(tx.CreatedAt)}</span>
    </div>
    <div class="flex items-center justify-between">
      <Label class="space-y-2">
        <span class="text-black text-xl">From:</span>
      </Label>
      <Label class="space-y-2 flex flex-col mt-8">
        <span class="text-xl text-[#28A745]">{senderName}</span>
        <span class="text-base text-[#666666]">{tx.SenderID}</span>
      </Label>
    </div>
    <div class="flex items-center justify-between">
      <Label class="space-y-2">
        <span class="text-black text-xl">To:</span>
      </Label>
      <Label class="space-y-2 flex flex-col">
        <span class="text-xl text-[#28A745]">{receiverName}</span>
        <span class="text-base text-[#666666]">{tx.ReceiverID}</span>
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
    <Button
      type="button"
      class="w-full bg-[#218838] hover:bg-[#28A745]"
      on:click={handleBackClick}
    >
      Back to Account page
    </Button>
  </form>
</Card>
