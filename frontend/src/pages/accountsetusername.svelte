<script>
  import { Card, Button, Label, Input, Checkbox } from 'flowbite-svelte';
  import HoneyLemonLogo from '../assets/BankLogo.png';
</script>

<Card class="w-full max-w-md mx-auto">
  <form class="flex flex-col space-y-6" action="/">
    <div class="flex items-center justify-between">
      <h3 class="text-4xl font-bold text-lime-900 dark:text-white">Set username</h3>
      <img src="{HoneyLemonLogo}" class="h-18 w-28" alt="HoneyLemonLogo" />
    </div>
    <Label class="space-y-2">
      <span class="text-gray-400">Username</span>
      <Input type="text" name="username" placeholder="Username" required />
    </Label>
    <Button type="submit" class="w-full bg-green-400 hover:bg-green-500">Next</Button>
  </form>
</Card>
