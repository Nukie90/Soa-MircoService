<script>
    import { Link, navigate } from "svelte-routing";
    import { Card, Input, Label, Button, Modal } from 'flowbite-svelte';
    import { ExclamationCircleOutline } from 'flowbite-svelte-icons';
    import HoneyLemonLogo from '../assets/BankLogo.png';
    import CancelButton from '../assets/cancel.png';
    import NextButton from '../assets/next.png';
    import BackButton from '../assets/back.png';
    import { currentUser } from '../lib/userstore.js';

    let receiverAccountNumber = '';
    let amount = '';
    let popupModal_confirm = false;
    let popupModal_cancel = false;
    let user;

    // Fetch currentUser from the store
    currentUser.subscribe(value => {
        user = value;
    });

    function handleConfirmClick() {
        popupModal_confirm = true;
    }

    function handleCancelClick() {
        popupModal_cancel = true;
    }

    function handleTransactionConfirm() {
        navigate('/transfer3');
    }

    function handleTransactionCancel() {
        navigate('/mainaccount');
    }

    function handleNextClick() {
        navigate('/transfer2', {
            state: {
                receiverAccountNumber,
                amount,
                fromAccountNumber: user?.selectedAccount
            }
        });
    }
</script>

<Card>
  <form class="flex flex-col space-y-6" action="/">
    <div class="flex items-center justify-between">
        <Link to="/transfer">
            <img src="{BackButton}" class="h-4 w-4" alt="Back" />
        </Link>
        <img src="{HoneyLemonLogo}" class="h-18 w-28" alt="HoneyLemonLogo" />
    </div>
    <Label class="space-y-2">
      <span class="text-black text-xl">To:</span>
    </Label>
    <Label class="space-y-2">
      <span class="text-gray-400 ml-4">Account No.</span>
      <Input type="text" name="accno" bind:value={receiverAccountNumber} required class="ml-4 w-11/12"/>
    </Label>
    <Label class="space-y-2">
      <span class="text-gray-400 ml-4">Amount</span>
      <Input type="text" name="amount" bind:value={amount} required class="ml-4 w-11/12"/>
    </Label>
    <div class="flex items-center justify-between">
        <div class="flex items-center justify-between">
            <img src="{CancelButton}" class="h-12 w-12 mt-7 cursor-pointer" alt="Cancel" on:click={handleCancelClick} />
            <span class="text-black ml-1 mt-7">Cancel</span>
        </div>
        <div class="flex items-center justify-between">
            <span class="text-gray-400 mr-1 mt-7">Next</span>
            <img src="{NextButton}" class="h-12 w-12 mt-7" alt="NextButton" on:click={handleNextClick} />
        </div>
        <Modal bind:open={popupModal_cancel} size="xs" autoclose>
            <div class="text-center">
                <ExclamationCircleOutline class="mx-auto mb-4 text-gray-400 w-12 h-12" />
                <h3 class="mb-5 text-lg font-normal text-gray-500">Are you sure you want to cancel?</h3>
                <div class="flex justify-center gap-2">
                    <Button color="red" on:click={handleTransactionCancel}>Yes, I'm sure</Button>
                    <Button color="alternative" on:click={() => (popupModal_cancel = false)}>No, cancel</Button>
                </div>
            </div>
        </Modal>
    </div>
  </form>
</Card>
