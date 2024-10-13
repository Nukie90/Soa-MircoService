<script>
    import { Link, navigate } from "svelte-routing";
    import { Card, Button, Label, Input } from 'flowbite-svelte';
    import { ArrowRightOutline } from 'flowbite-svelte-icons';
    import { Dropdown, DropdownItem } from 'flowbite-svelte';
    import { DotsVerticalOutline } from 'flowbite-svelte-icons';
    import PaymentPic from '../assets/Payment.png';
    import Payment from '../assets/Pay.png';
    import Loan from '../assets/Loan.png';
    import Transfer from '../assets/Transfer.png';
    import Statement from '../assets/Statement.png';
    import Investment from '../assets/Invest.png';
    import HoneyLemonLogo from '../assets/BankLogo.png';
    import BackButton from '../assets/back.png';

    let tesst = '6666'
</script>

<div class="flex flex-col items-center">
  <Card class="bg-white mb-5" size="xs" padding="lg" style="width: 900px;">
    <div class="flex flex-col h-full">
        <div class="flex items-center justify-between">
            <Link to="/investment">
                <img src="{BackButton}" class="h-4 w-4" alt="HoneyLemonLogo" />
            </Link>
            <img src="{HoneyLemonLogo}" class="h-18 w-28" alt="HoneyLemonLogo" />
        </div>
        <div class="flex items-center justify-between">
            <div class="text-2xl text-black font-bold mt-4 mb-3">
                Buy
            </div>
            <div class="text-base text-black font-semibold mt-4 mb-3">
                
                In port: {tesst}
            </div>
        </div>
        <div>
            <Card class="bg-white border-black mb-5" size="sm" padding="md">
                <div class="flex flex-col h-full">
                    <div class="text-base text-black">
                        Name
                    </div>
                    <Input type="text" name="investbuyname" class="w-full mt-2"/>
                <div class="flex flex-col h-full">
                    <div class="text-base text-black mt-2">
                        Volume
                    </div>
                    <Input type="text" name="investbuyvolume" class="w-full mt-2"/>
                </div>
                <div class="flex flex-col h-full">
                    <div class="text-base text-black mt-2">
                        Price
                    </div>
                    <Input type="text" name="investbuyprice" class="w-full mt-2"/>
                </div>
                <div class="flex flex-col h-full">
                    <div class="text-base text-black mt-2">
                        Pin
                    </div>
                    <Input type="text" name="pin" required class="w-full mt-2"/>
                </div>  
            </Card>
        </div>
        <Button type="submit" class="w-full bg-[#218838] hover:bg-[#28A745]">
            Buy
        </Button>
    </div>
  </Card>
</div>