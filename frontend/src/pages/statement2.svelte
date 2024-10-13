<script>
  import { onMount } from "svelte";
  import { Card } from "flowbite-svelte";
  import { Link, navigate } from "svelte-routing";
  import BackButton from "../assets/back.png";
  import axios from "axios";
  import { get } from "svelte/store";
  import { currentAccount } from "../lib/userstore";

  // Get data from location state
  const { startdate, enddate } = history.state || {};
  let filteredActivity = [];
  let clientName;
  let accountID;
  let clientID;
  let start;
  let end;
  let activity;
  let token;

  function formatDate(dateString) {
    const date = new Date(dateString);
    // Convert to ISO string and remove milliseconds and timezone info
    return date.toISOString().split(".")[0] + "Z";
  }

  onMount(() => {
    token = getCookie("esb_token");
    if (token) {
      fetchData(token);
    }

    fetchStatement();
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

      const statementResponse = await axios.get(
        "http://127.0.0.1:4000/esb/clients/info",
        {
          withCredentials: true, // Ensure credentials are sent with the request
          headers: {
            esb_token: `Bearer ${token}`,
          },
        }
      );
    } catch (error) {
      console.error("Error fetching data:", error);
      alert("Error fetching data: " + error);
      navigate("/statement")
    }
  }

  function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(";").shift();
  }

  async function fetchStatement() {
    try {
      const formattedStartDate = formatDate(startdate);
      const formattedEndDate = formatDate(enddate);

      const statementResponse = await axios.get(
        `http://127.0.0.1:4000/esb/statements?accountID=${get(currentAccount)}&start=${formattedStartDate}&end=${formattedEndDate}`,
        {
          withCredentials: true,
          headers: {
            esb_token: `Bearer ${token}`,
          },
        }
      );

      let data = statementResponse.data;
      console.log("Statement Response:", data);
      clientName = data.clientName;
      accountID = data.accountID;
      clientID = data.clientID;
      start = data.start;
      end = data.end;
      activity = data.activity;

      filteredActivity = activity.filter(
        (tx) =>
          new Date(tx.timestamp) >= new Date(start) &&
          new Date(tx.timestamp) <= new Date(end)
      );
    } catch (error) {
      console.error("Error fetching statements:", error);
      alert("Error fetching data: " + error);
      navigate("/statement")
    }
  }
</script>

<div class="flex flex-col items-center">
  <Card class="bg-white mb-5" size="lg" padding="xl" style="width: 900px;">
    <div class="flex flex-col h-full">
      <div class="flex items-center justify-between mb-4">
        <Link to="/statement">
          <img src={BackButton} class="h-4 w-4" alt="Back" />
        </Link>
      </div>
      <div class="mb-4 text-gray-700 dark:text-gray-400">
        <h5
          class="text-[#004D00] mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white"
        >
          Statement for {clientName}
        </h5>
        <p><strong>Account ID:</strong> {accountID}</p>
        <p><strong>Client ID:</strong> {clientID}</p>
        <p>
          <strong>Period:</strong>
          {new Date(start).toLocaleDateString()} to {new Date(
            end
          ).toLocaleDateString()}
        </p>
      </div>
      <div class="overflow-x-auto">
        <table class="min-w-full text-left text-sm">
          <thead>
            <tr class="border-b border-gray-200">
              <th class="py-2">Transaction ID</th>
              <th class="py-2">From</th>
              <th class="py-2">To</th>
              <th class="py-2">Amount</th>
              <th class="py-2">Type</th>
              <th class="py-2">Timestamp</th>
            </tr>
          </thead>
          <tbody>
            {#each filteredActivity as tx}
              <tr class="border-b border-gray-200">
                <td class="py-2">{tx.txID}</td>
                <td class="py-2">{tx.from}</td>
                <td class="py-2">{tx.to}</td>
                <td class="py-2">${tx.amount}</td>
                <td class="py-2 capitalize">{tx.type}</td>
                <td class="py-2">{new Date(tx.timestamp).toLocaleString()}</td>
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
    </div>
  </Card>
</div>
