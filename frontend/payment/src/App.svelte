<script>
  import { Link } from "svelte-routing";
  import { Card, Label, Button } from "flowbite-svelte";
  import { Modal, Input } from 'flowbite-svelte';

  let popupModal_confirm = false;

  function openConfirmModal() {
    popupModal_confirm = true; // Open modal
  }

  function closeConfirmModal() {
    popupModal_confirm = false; // Close modal
  }
</script>

<div class="flex justify-center mt-10">
  <Card class="w-full max-w-lg p-8 rounded-lg shadow-lg bg-white">
    <h2 class="text-2xl font-bold text-gray-700 mb-6 mt-4 text-center">Payment</h2>

    <!-- Transfer Form -->
    <form class="space-y-6">
      <!-- Recipient Account Field -->
      <div>
        <Label class="block text-lg font-semibold mt-4 text-gray-600">
          Contact account / Mobile no.
        </Label>
        <input
          type="text"
          name="accno"
          placeholder="Enter account or mobile number"
          required
          class="mt-2 bg-blue-500 w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 text-gray-700 placeholder-gray-400"
        />
      </div>

      <!-- Transfer Amount Field -->
      <div>
        <Label class="block text-lg font-semibold text-gray-600">
          Amount
        </Label>
        <input
          type="text"
          name="amount"
          placeholder="Enter amount"
          required
          class="mt-2 bg-blue-500 w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 text-gray-700 placeholder-gray-400"
        />
      </div>

      <!-- Action Buttons -->
      <div class="flex justify-between mt-8">
        <!-- Cancel Button -->
        <Link to="/">
          <Button class="w-36 h-12 bg-red-400 hover:bg-red-600 text-white font-semibold rounded-lg shadow-md transition duration-150 ease-in-out">
            Cancel
          </Button>
        </Link>

        <Button class="w-36 h-12 bg-green-500 hover:bg-blue-600 text-white font-semibold rounded-lg shadow-md transition duration-150 ease-in-out"
        on:click={openConfirmModal}>
          Confirm
        </Button>
        {#if popupModal_confirm}
          <div class="fixed inset-0 flex items-center z-50">
            <div class="bg-white rounded-lg shadow-lg p-6 w-96 relative mt-16" style="top: 150%; left: 400%;">
              <h3 class="text-3xl font-bold text-black text-center mb-4">Confirm for a pin</h3>
              <form class="flex flex-col space-y-6" action="#">
                <Label class="space-y-2">
                  <span>Enter a pin</span>
                  <Input type="password" class="bg-blue-500 px-4 py-2" required />
                </Label>
                <div class="flex justify-center gap-64">
                  <Button
                    class="w-20 h-9 bg-red-400 hover:bg-red-700 text-black"
                    on:click={closeConfirmModal}>Close</Button>
                  <Button 
                    class="w-20 h-9 bg-green-500 hover:bg-green-700 text-black"
                    >Enter</Button>
                </div>
              </form>
            </div>
          </div>
          {/if}
      </div>
    </form>
  </Card>
</div>

<style>
  /* Custom styles for input */
  input {
    transition: border-color 0.3s, box-shadow 0.3s;
  }

  input:focus {
    outline: none;
    box-shadow: 0 0 0 3px rgba(99, 102, 241, 0.5); /* Adding focus ring */
  }
</style>
