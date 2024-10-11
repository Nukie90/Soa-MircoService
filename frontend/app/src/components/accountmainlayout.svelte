<script>
  import { currentUser } from '../lib/userstore.js';
  import { Link, navigate } from "svelte-routing";
  import { Modal } from 'flowbite-svelte';
  import { ExclamationCircleOutline } from 'flowbite-svelte-icons';
  import HoneyLemonLogo from "../assets/BankLogo.png";
  import AccountLogo from "../assets/AccountLogo.png";
  import { Button, Dropdown, DropdownItem, Avatar, DropdownHeader, DropdownDivider } from 'flowbite-svelte';
  import { BellSolid, EyeSolid } from 'flowbite-svelte-icons';
  import axios from 'axios';

  let popupModal = false;
  let userData = null;
  let user = null;

  // Subscribe to the currentUser store
  currentUser.subscribe(value => {
    user = value;

  const token = getCookie('esb_token');
  if (token) {
    fetchData(token);
    }
  });

  function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
  }
  
  async function fetchData(token) {
    try {
      // Set the token as a cookie
      document.cookie = `esb_token=${token}; path=/;`;

      const userResponse = await axios.get('http://127.0.0.1:4000/esb/clients/info', {
        withCredentials: true, // Ensure credentials are sent with the request
        headers: {
          'esb_token': `Bearer ${token}`
        }
      });

      userData = userResponse.data;

    } catch (error) {
      console.error('Error fetching data:', error);
    }
  }

  async function handleLogout() {
    try {
      const response = await axios.get("http://localhost:4000/esb/logout", {
        withCredentials: true, // Include cookies with the request
      });

      // Clear the token cookie
      document.cookie =
        "esb_token=; expires=Thu, 01 Jan 1970 00:00:00 GMT; path=/";

      if (response.status === 200) {
        console.log("Logout successful:", response.data);
        navigate("/");
      } else {
        console.error("Logout failed:", response.status);
      }
    } catch (error) {
      console.error("Logout failed:", error);
    }
  }
</script>

<div class="navbar bg-[#F0F0F0] h-24 flex items-center justify-between px-5 fixed top-0 left-0 right-0 shadow-lg z-50">
  <div class="flex flex-row items-center">
    <div class="ml-1">
      <Link to="/">
        <img src={HoneyLemonLogo} alt="HoneyLemon Logo" class="Logo" />
      </Link>
    </div>
  </div>
  <div>
    <!-- mount to /mainaccou -->
    <Avatar class="acs mr-12" src="{AccountLogo}" />
    <Dropdown class="bg-[#F0F0F0] rounded shadow-lg" triggeredBy=".acs">
      <DropdownHeader class="bg-[#28A745] p-4 rounded-t-lg">
        <span class="block text-sm text-white dark:text-white">{userData ? userData.name : "User"}</span>
      </DropdownHeader>
      <DropdownDivider class="my-2 border-t border-gray-300" />
      <DropdownItem class="bg-white hover:bg-slate-50 px-4 py-2 text-gray-700">
        <Link to="/mainaccount" class="w-full text-left block text-gray-700 dark:text-gray-400 hover:bg-slate-50">
        Home
        </Link>
      </DropdownItem>
      <DropdownItem class="bg-white hover:bg-slate-50 px-4 py-2 text-gray-700">
        <Link to="/changeprofile" class="w-full text-left block text-gray-700 dark:text-gray-400 hover:bg-slate-50">
        Edit profile
        </Link>
      </DropdownItem>
      <DropdownDivider class="my-2 border-t border-gray-300" />
      <DropdownItem class="bg-white hover:bg-slate-50 px-4 py-2 text-gray-700">
        <div class="w-full text-left block text-gray-700 dark:text-gray-400 hover:bg-slate-50" on:click={()=>
          (popupModal = true)}>
          Sign out
        </div>
      </DropdownItem>
      <Modal bind:open={popupModal} size="xs" autoclose>
        <div class="text-center">
          <ExclamationCircleOutline class="mx-auto mb-4 text-gray-400 w-12 h-12" />
          <h3 class="mb-5 text-lg font-normal text-gray-500">Are you sure you want to sign out?</h3>
          <div class="flex justify-center gap-2">
            <Button color="red" on:click={handleLogout}>Yes, I'm sure</Button>
            <Button color="alternative" on:click={()=> (popupModal = false)}>No, cancel</Button>
          </div>
        </div>
      </Modal>
    </Dropdown>
  </div>
</div>

<div class="content-container pt-24">
  <slot></slot>
</div>

<style>
  .Logo {
    height: 36px;
  }

  .LoginButton {
    background-color: #ffffff;
    color: #000000;
    font-size: 16px;
    transition: background-color 0.3s ease, color 0.3s ease;
  }

  .LoginButton:hover {
    opacity: 70%;
  }

  .content-container {
    padding-top: 96px;
  }
</style>