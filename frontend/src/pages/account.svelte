<script>
  import {
    Card,
    Input,
    Label,
    Modal,
    Button,
    Dropdown,
    DropdownItem,
  } from "flowbite-svelte";
  import { Link, navigate } from "svelte-routing";
  import { DotsVerticalOutline } from "flowbite-svelte-icons";
  import { currentAccount } from "../lib/userstore.js";
  import { onMount } from "svelte";
  import HoneyLemonLogo from "../assets/BankLogo.png";
  import Transfer from "../assets/Transfer.png";
  import Payment from "../assets/Pay.png";
  import Loan from "../assets/Loan.png";
  import Investment from "../assets/Invest.png";
  import Statement from "../assets/Statement.png";
  import axios from "axios";
  import { get } from "svelte/store";

  let user = null;
  let enteredPin = "";
  let userData = null;
  let accountData = null;
  let selectedAccount = null;
  let loggedIn = false; // Default state for logged in status

  function checkLoginStatus() {
    // Check for the presence of a specific cookie
    const cookies = document.cookie.split(";").map((cookie) => cookie.trim());
    const authCookie = cookies.find((cookie) =>
      cookie.startsWith("esb_token=")
    );

    loggedIn = !!authCookie;
    console.log(authCookie);
  }

  function goToChangePin() {
    if (!selectedAccount) {
      alert("You have no account");
      return;
    }
    navigate(`/changepin?accountId=${selectedAccount.id}`);
  }

  onMount(() => {
    checkLoginStatus();
    if (!loggedIn) {
      navigate("/");
    }

    const token = getCookie("esb_token");
    if (token) {
      fetchData(token);
    }
  });

  function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(";").shift();
  }

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

      accountData = accResponse.data;
      // console.log('accountData:', accountData);

      userData = userResponse.data;
      // console.log('userData:', userData);

      if (accountData.length > 0){
        selectedAccount = accountData[0]
        currentAccount.set(accountData[0].id)
      } else {
        selectedAccount = null
        currentAccount.set(null)
      }
      console.log("Selected account:", selectedAccount);
      
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  }

  function switchAccount(accountNumber) {
    selectedAccount = accountData.find((account) => account.id === accountNumber);
    console.log("switchAccout called: " + selectedAccount.id);
    
    currentAccount.set(selectedAccount.id);
    console.log("currentAccount: " + get(currentAccount));
  }

  let popupModal_deleteacc = false;

  function handleDeleteAcc() {
    popupModal_deleteacc = true;
  }

  function handleDeleteAccConfirm(event) {
    event.preventDefault(); // Prevent default form submission

    if (!selectedAccount) {
      alert("No account selected.");
      return;
    }

    if (!enteredPin) {
      alert("Please enter your PIN.");
      return;
    }

    axios
      .delete("http://127.0.0.1:4000/esb/accounts/delete", {
        withCredentials: true, // Ensure credentials are sent with the request
        data: {
          id: selectedAccount.id, // Pass the selected account ID
          pin: enteredPin, // Pass the entered PIN
        },
      })
      .then((response) => {
        console.log("Account deleted successfully:", response.data);
        alert("Account deleted successfully.");
        location.reload(); // Reload the page to reflect the changes
      })
      .catch((error) => {
        console.error("Error deleting account:", error);
        alert("Failed to delete account. Please check your PIN.");
      });
  }
</script>

<div class="flex flex-col items-center">
  <Card class="bg-white mb-5" size="lg" padding="xl" style="width: 900px;">
    <div class="flex flex-col h-full">
      <div class="flex items-center justify-between mb-4">
        <div class="flex flex-col">
          <h5
            class="text-[#004D00] mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white"
          >
            <!-- get name from localstorage -->
            {#if userData}
              {userData.name}
            {/if}
          </h5>
          <h6
            class="mb-3 font-normal text-base text-gray-700 dark:text-gray-400 leading-tight"
          >
            {userData && selectedAccount
              ? `${selectedAccount.id || "xxxxx"}`
              : "xxxxx"}
          </h6>
        </div>
        <div class="flex items-center">
          <DotsVerticalOutline
            class="dots-menu dark:text-white cursor-pointer mb-10"
          />
          <Dropdown triggeredBy=".dots-menu" class="bg-slate-100 rounded shadow-lg">
            {#if accountData && accountData.length > 0}
              {#each accountData as account}
                <DropdownItem
                  class="bg-white hover:bg-slate-50 text-gray-700"
                  on:click={() => switchAccount(account.id)}
                >
                  {account.id}
                </DropdownItem>
              {/each}
            {:else}
              <!-- No accounts, show only "Add new account" -->
              <DropdownItem class="bg-white text-gray-500">
                No accounts available
              </DropdownItem>
            {/if}
            <DropdownItem slot="footer" class="bg-[#28A745] hover:bg-[#03C04A]">
              <Link
                to="/addaccount"
                class="w-full text-left block text-white dark:text-gray-400"
              >
                Add new account
              </Link>
            </DropdownItem>
          </Dropdown>
        </div>
      </div>
      <div class="flex flex-1 items-center justify-center">
        <div class="flex flex-col items-center mt-[-80px]">
          <div
            class="flex flex-col items-center justify-center w-60 h-60 bg-[#28A745] rounded-full border-2 border-gray-300"
          >
            <span class="text-xl font-bold text-white block"
              >Available Balance</span
            >
            <span class="text-xl font-medium text-white block mt-4">
              {userData && selectedAccount
                ? `${selectedAccount.balance || 0} ฿`
                : "0 ฿"}
            </span>
          </div>
        </div>
      </div>
      <div class="flex flex-1 items-center justify-between">
        <div>
          <Button
            class="w-40 h-9 bg-[#cccccc] hover:bg-slate-100 text-black flex items-center justify-center space-x-2"
            on:click={goToChangePin}
          >
            Change pin
          </Button>
        </div>
        <div>
          <Button
            class="w-40 h-9 bg-red-400 hover:bg-red-700 text-black flex items-center justify-center space-x-2"
            on:click={handleDeleteAcc}>Delete Account</Button
          >
        </div>
        <Modal bind:open={popupModal_deleteacc} size="xs" autoclose={false}>
          <form class="flex flex-col space-y-6" action="#">
            <h3 class="mb-4 text-xl font-medium text-gray-900 dark:text-white">
              Enter PIN to Confirm
            </h3>
            <Label class="space-y-2">
              <span>Enter your PIN</span>
              <Input type="password" bind:value={enteredPin} required />
            </Label>
            <div class="flex justify-center gap-4">
              <Button color="red" on:click={handleDeleteAccConfirm}
                >Enter</Button
              >
              <Button
                color="alternative"
                on:click={() => (popupModal_deleteacc = false)}>Cancel</Button
              >
            </div>
          </form>
        </Modal>
      </div>
    </div>
  </Card>
  <div class="flex justify-center space-x-7 mt-8">
    <Link to="/transfer">
      <Button
        class="w-40 h-16 bg-white hover:bg-slate-50 text-black flex items-center justify-center space-x-2"
      >
        <img src={Transfer} alt="Transfer" class="w-8 h-8" />
        <span>Transfer</span>
      </Button>
    </Link>
    <Link to="/payment">
      <Button
        class="w-40 h-16 bg-white hover:bg-slate-50 text-black flex items-center justify-center space-x-2"
      >
        <img src={Payment} alt="Payment" class="w-8 h-8" />
        <span>Payment</span>
      </Button>
    </Link>
    <Link to="/loan">
      <Button
        class="w-40 h-16 bg-white hover:bg-slate-50 text-black flex items-center justify-center space-x-2"
      >
        <img src={Loan} alt="Loan" class="w-8 h-8" />
        <span>Loan</span>
      </Button>
    </Link>
    <Link to="/investment">
      <Button
        class="w-40 h-16 bg-white hover:bg-slate-50 text-black flex items-center justify-center space-x-2"
      >
        <img src={Investment} alt="Investment" class="w-8 h-8" />
        <span>Investment</span>
      </Button>
    </Link>
    <Link to="/statement">
      <Button
        class="w-40 h-16 bg-white hover:bg-slate-50 text-black flex items-center justify-center space-x-2"
      >
        <img src={Statement} alt="Statement" class="w-8 h-8" />
        <span>Statement</span>
      </Button>
    </Link>
  </div>
</div>
