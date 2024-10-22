<script>
    import { Link } from "svelte-routing";
    import { Card, Button, Label, Input, Checkbox } from 'flowbite-svelte';
    import HoneyLemonLogo from '../assets/BankLogo.png';
    import CancelButton from '../assets/cancel.png';
    import ConfirmButton from '../assets/confirm.png';
    import NextButton from '../assets/next.png';
    import BackButton from '../assets/back.png';
</script>

<Card>
  <form class="flex flex-col space-y-6" action="/">
    <div class="flex items-center justify-between">
        <Link to="/payment">
            <img src="{BackButton}" class="h-4 w-4" alt="HoneyLemonLogo" />
        </Link>
        <img src="{HoneyLemonLogo}" class="h-18 w-28" alt="HoneyLemonLogo" />
    </div>
    <div class="flex items-center justify-between">
        <Label class="space-y-2">
            <span class="text-black text-xl">From:</span>
        </Label>
        <Label class="space-y-2 flex flex-col mt-8">
            <span class="text-xl text-[#28A745]">Shogun</span>
            <span class="text-base text-[#666666]">xxx-xxx-xxx</span>
        </Label>
    </div>
    <div class="flex items-center justify-between">
        <Label class="space-y-2">
            <span class="text-black text-xl">To:</span>
        </Label>
        <Label class="space-y-2 flex flex-col mt-8">
            <span class="text-xl text-[#28A745]">AIS</span>
            <span class="text-base text-[#666666]">xxx-xxx-xxx</span>
        </Label>
    </div>
    <div class="flex items-center justify-between">
        <Label class="space-y-2">
            <span class="text-black text-xl">Amount:</span>
        </Label>
        <Label class="space-y-2 ">
            <span class="text-black text-xl ">5000.00</span>
        </Label>
    </div>
    <div class="flex items-center justify-between">
        <Label class="space-y-2 flex flex-col">
            <span class="text-red-500">Beware of scammers !</span>
            <span class="text-gray-400">Please check the recipient's name and account number every time before pressing the confirm button.</span>
        </Label>
    </div>
    <div class="flex items-center justify-between">
        <div class="flex items-center justify-between">
            <Link to="/">
                <img src="{CancelButton}" class="h-12 w-12 mt-7" alt="HoneyLemonLogo" />
            </Link>
            <span class="text-black ml-1 mt-7">Cancel</span>
        </div>
        <div class="flex items-center justify-between">
            <span class="text-gray-400 mr-1 mt-7">Next</span>
            <Link to="/payment3">
                <img src="{ConfirmButton}" class="h-12 w-12 mt-7" alt="HoneyLemonLogo" />
            </Link>
        </div>
    </div>
  </form>
</Card>
