<script>
  import { Link, navigate } from "svelte-routing";
  import { Card, Input, Label, Button, Modal } from "flowbite-svelte";
  import { ExclamationCircleOutline } from "flowbite-svelte-icons";
  import HoneyLemonLogo from "../assets/BankLogo.png";
  import CancelButton from "../assets/cancel.png";
  import ConfirmButton from "../assets/confirm.png";
  import BackButton from "../assets/back.png";
  import { currentUser, currentAccount } from "../lib/userstore.js";
  import { get } from "svelte/store";
  import { onMount } from "svelte";
  import axios from "axios";

  let popupModal_confirm = false;
  let popupModal_cancel = false;
  let enteredPin = "";
  let pinError = "";
  let token;

  // Get data from location state
  const { receiverAccountNumber, amount } = history.state || {};

  // Fetch currentUser from the store
  let user;
  let senderName;
  let receiverName;

  currentUser.subscribe((value) => {
    user = value;
  });

  onMount(() => {
    currentUser.subscribe((value) => {
      user = value;
    });

    token = getCookie("esb_token");
    if (token) {
      fetchData(token);
    }
  });

  async function fetchData(token) {
    try {
      // Set the token as a cookie
      document.cookie = `esb_token=${token}; path=/;`;

      const accResponse = await axios.get(
        "http://127.0.0.1:4000/esb/accounts/clientAcc",
        {
          withCredentials: true, // Ensure credentials are sent with the request
          headers: {
            esb_token: `Bearer ${token}`,
          },
        }
      );

      const userResponse = await axios.get(
        "http://127.0.0.1:4000/esb/clients/info",
        {
          withCredentials: true, // Ensure credentials are sent with the request
          headers: {
            esb_token: `Bearer ${token}`,
          },
        }
      );

      senderName = userResponse.data.name;

      const receiverAccount = await axios.get(
        "http://127.0.0.1:4000/esb/accounts/account/" + receiverAccountNumber,
        {
          withCredentials: true, // Ensure credentials are sent with the request
          headers: {
            esb_token: `Bearer ${token}`,
          },
        }
      );
      
      const receiverAccountName = await axios.get(
        "http://127.0.0.1:4000/esb/clients/" + receiverAccount.data.clientID,
        {
          withCredentials: true, // Ensure credentials are sent with the request
          headers: {
            esb_token: `Bearer ${token}`,
          },
        }
      );
      receiverName = receiverAccountName.data.name;
      //   const selectedAccount = accountData.length > 0 ? accountData[0] : null;
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  }

  function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(";").shift();
  }

  // Display the selected account number
  const fromAccountNumber = get(currentAccount);

  function handleConfirmClick() {
    popupModal_confirm = true;
  }

  function handleCancelClick() {
    popupModal_cancel = true;
  }

  async function handleTransactionConfirm() {
    try {
      const verifyResponse = await axios.post(
        "http://127.0.0.1:4000/esb/accounts/verifyPin",
        {
          accountID: fromAccountNumber,
          pin: enteredPin,
        },
        {
          withCredentials: true,
          headers: {
            esb_token: `Bearer ${token}`,
          },
        }
      );

      // Check the response status for successful PIN verification
      if (verifyResponse.status === 200) {
        // PIN verified successfully, proceed with transaction creation
        try {
          const transactionResponse = await axios.post(
            "http://127.0.0.1:4000/esb/transactions/create",
            {
              senderID: fromAccountNumber,
              receiverID: receiverAccountNumber,
              amount: parseFloat(amount),
            },
            {
              withCredentials: true,
              headers: {
                esb_token: `Bearer ${token}`,
              },
            }
          );

          // Handle transaction response if needed
          if (transactionResponse.status === 200) {
            let transaction = transactionResponse.data
            console.log(transaction);
            
            navigate("/transfer3", {
              state: {
                transaction,
                senderName,
                receiverName
              },
            });
          } else {
            alert(
              "Transaction failed with status: " + transactionResponse.status
            );
            enteredPin = "";
          }
        } catch (transactionError) {
          alert(
            "Error creating transaction: " +
              (transactionError.response?.data?.message ||
                transactionError.message)
          );
          enteredPin = "";
        }
      } else {
        alert("PIN verification failed with status: " + verifyResponse.status);
        enteredPin = "";
      }
    } catch (verifyError) {
      alert(
        "Error verifying PIN: " +
          (verifyError.response?.data?.message || verifyError.message)
      );
      enteredPin = "";
    }
  }

  function handleTransactionCancel() {
    navigate("/mainaccount");
  }
</script>

<Card>
  <form class="flex flex-col space-y-6" action="/">
    <div class="flex items-center justify-between">
      <Link to="/transfer">
        <img src={BackButton} class="h-4 w-4" alt="Back" />
      </Link>
      <img src={HoneyLemonLogo} class="h-18 w-28" alt="HoneyLemonLogo" />
    </div>
    <div class="flex items-center justify-between">
      <Label class="space-y-2">
        <span class="text-black text-xl">From:</span>
      </Label>
      <Label class="space-y-2 flex flex-col mt-8">
        <span class="text-xl text-[#28A745]">{senderName}</span>
        <span class="text-base text-[#666666]">{fromAccountNumber}</span>
      </Label>
    </div>
    <div class="flex items-center justify-between">
      <Label class="space-y-2">
        <span class="text-black text-xl">To:</span>
      </Label>
      <Label class="space-y-2 flex flex-col">
        <span class="text-xl text-[#28A745]">{receiverName}</span>
        <span class="text-base text-[#666666]"
          >{receiverAccountNumber}</span
        >
      </Label>
    </div>
    <div class="flex items-center justify-between">
      <Label class="space-y-2">
        <span class="text-black text-xl">Amount:</span>
      </Label>
      <Label class="space-y-2">
        <span class="text-black text-xl">{amount || "N/A"}</span>
      </Label>
    </div>
    <div class="flex items-center justify-between">
      <Label class="space-y-2 flex flex-col">
        <span class="text-red-500">Beware of scammers!</span>
        <span class="text-gray-400"
          >Please check the recipient's name and account number every time
          before pressing the confirm button.</span
        >
      </Label>
    </div>
    <div class="flex items-center justify-between">
      <div class="flex items-center">
        <img
          src={CancelButton}
          class="h-12 w-12 mt-7 cursor-pointer"
          alt="Cancel"
          on:click={handleCancelClick}
        />
        <span class="text-black ml-1 mt-7">Cancel</span>
      </div>
      <div class="flex items-center">
        <span class="text-black mr-1 mt-7">Confirm</span>
        <img
          src={ConfirmButton}
          class="h-12 w-12 mt-7 cursor-pointer"
          alt="Confirm"
          on:click={handleConfirmClick}
        />
      </div>
    </div>
    <Modal bind:open={popupModal_cancel} size="xs" autoclose>
      <div class="text-center">
        <ExclamationCircleOutline
          class="mx-auto mb-4 text-gray-400 w-12 h-12"
        />
        <h3 class="mb-5 text-lg font-normal text-gray-500">
          Are you sure you want to cancel?
        </h3>
        <div class="flex justify-center gap-2">
          <Button color="red" on:click={handleTransactionCancel}
            >Yes, I'm sure</Button
          >
          <Button
            color="alternative"
            on:click={() => (popupModal_cancel = false)}>No, cancel</Button
          >
        </div>
      </div>
    </Modal>
    <Modal bind:open={popupModal_confirm} size="xs" autoclose>
      <form class="flex flex-col space-y-6" action="#">
        <h3 class="mb-4 text-xl font-medium text-gray-900 dark:text-white">
          Enter PIN to Confirm
        </h3>
        <Label class="space-y-2">
          <span>Enter your PIN</span>
          <Input type="password" bind:value={enteredPin} required />
        </Label>
        {#if pinError}
          <p class="text-red-500">{pinError}</p>
        {/if}
        <div class="flex justify-center gap-4">
          <Button color="red" on:click={handleTransactionConfirm}>Enter</Button>
          <Button
            color="alternative"
            on:click={() => (popupModal_confirm = false)}>Cancel</Button
          >
        </div>
      </form>
    </Modal>
  </form>
</Card>
