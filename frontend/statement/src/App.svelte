<script>
  import { Link } from "svelte-routing";
  import { Card, Label, Button } from "flowbite-svelte";
</script>

<div class="flex justify-center mt-10">
  <Card class="w-full max-w-lg p-8 rounded-lg shadow-lg bg-white">
    <h2 class="text-2xl font-bold text-gray-700 mb-6 mt-4 text-center">Statement</h2>

    <!-- Transfer Form -->
    <form class="space-y-6">
      <!-- Recipient Account Field -->
      <div>
        <Label class="block text-lg font-semibold mt-4 text-gray-600">
          Select start date
        </Label>
        <input
          type="date"
          name="startdate"
          required
          class="mt-2 bg-yellow-500 w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 text-gray-700 placeholder-gray-400"
        />
      </div>

      <!-- Transfer Amount Field -->
      <div>
        <Label class="block text-lg font-semibold text-gray-600">
          Select end date
        </Label>
        <input
          type="date"
          name="enddate"
          required
          class="mt-2 bg-yellow-500 w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 text-gray-700 placeholder-gray-400"
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

        <!-- Next Button -->
        <Button class="w-36 h-12 bg-green-500 hover:bg-blue-600 text-white font-semibold rounded-lg shadow-md transition duration-150 ease-in-out">
          Submit
        </Button>
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
