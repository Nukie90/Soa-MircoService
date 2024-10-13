<script>
  import { Card, Button, Label, Input, Checkbox } from 'flowbite-svelte';
  import HoneyLemonLogo from '../assets/BankLogo.png';
</script>

<Card class="w-full max-w-lg mx-auto">
  <form class="flex flex-col space-y-6" action="/">
    <Label class="space-y-2">
      <span class="text-gray-400">Account number</span>
      <Input type="text" name="accnumber" placeholder="Account number" required />
    </Label>
    <Label class="space-y-2">
      <span class="text-gray-400">Picture of account book</span>
      <div class="flex items-center">
        <Input type="file" name="fileaccbook" required class="hidden" id="fileaccbook" />
        <label for="fileaccbook" class="cursor-pointer flex items-center justify-center w-full py-2 px-4 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50">
          Choose file
        </label>
      </div>
    </Label>
    <Button type="submit" class="w-full bg-green-400 hover:bg-green-500">Register</Button>
  </form>
</Card>

