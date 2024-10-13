<script>
  import { Link, navigate } from "svelte-routing";
  import { Card, Button, Label, Input } from "flowbite-svelte";
  import HoneyLemonLogo from "../assets/BankLogo.png";
  import BackButton from "../assets/back.png";
  import { onMount } from "svelte";
  import { get } from "svelte/store";
  import { currentAccount } from "../lib/userstore";

  let startdate = "";
  let enddate = "";

  onMount(() => {
    console.log("currentAccount: " + get(currentAccount));
  });

  function handleSubmit() {
    console.log("Start Date:", startdate);
    console.log("End Date:", enddate);
    if (!startdate || !enddate) {
      alert("Please fill in the information.");
    } else {
      navigate("/statement2", {
        state: {
            startdate,
            enddate,
        },
      });
    }
  }
</script>

<div class="flex flex-col items-center">
  <Card class="bg-white mb-5" size="md" padding="xl" style="width: 900px;">
    <div class="flex flex-col h-full">
      <div class="flex items-center justify-between">
        <Link to="/mainaccount">
          <img src={BackButton} class="h-4 w-4" alt="Back Button" />
        </Link>
        <img src={HoneyLemonLogo} class="h-18 w-28" alt="HoneyLemonLogo" />
      </div>
      <div class="text-base text-black mt-4 ml-4 mb-3">Request Statement</div>
      <div class="ml-4">
        <Card class="bg-white border-black mb-5" size="md" padding="md">
          <div class="flex flex-col h-full">
            <div class="text-base text-black mt-2">Start date</div>
            <Input
              type="date"
              name="startdate"
              placeholder="yyyy-mm-dd"
              required
              class="w-2/3 mt-2"
              bind:value={startdate}
            />
          </div>
          <div class="flex flex-col h-full">
            <div class="text-base text-black mt-2">End date</div>
            <Input
              type="date"
              name="enddate"
              placeholder="yyyy-mm-dd"
              required
              class="w-2/3 mt-2"
              bind:value={enddate}
            />
          </div>
        </Card>
      </div>
      <Button
        type="button"
        class="w-full bg-[#218838] hover:bg-[#28A745]"
        on:click={handleSubmit}
      >
        Submit request
      </Button>
    </div>
  </Card>
</div>
