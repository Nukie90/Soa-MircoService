<script>
  import { Router, Route, Link } from "svelte-routing";
  import Account from "../../account/src/App.svelte"; // Account frontend
  import Transfer from "../../transfer/src/App.svelte"; // Transfer frontend
  import Payment from "../../payment/src/App.svelte"; // Payment frontend
  import Statement from "../../statement/src/App.svelte"; // Statement frontend
</script>

<Router>
  <div class="flex justify-center">
    <div class="w-full max-w-4xl relative">
      <Account />
      
      <div class="flex flex-col space-y-4 absolute right-0 top-24">
        <Link to="/">
          <button class="bg-green-500 hover:bg-green-700 text-white py-2 px-4 rounded">
            Transfer
          </button>
        </Link>
        <Link to="/payment">
          <button class="bg-blue-500 hover:bg-blue-700 text-white py-2 px-4 rounded">
            Payment
          </button>
        </Link>
        <Link to="/statement">
          <button class="bg-yellow-500 hover:bg-yellow-700 text-white py-2 px-4 rounded">
            Statement
          </button>
        </Link>
      </div>

      <div class="mt-6">
        <Route path="/" component={Transfer} />
        <Route path="/payment" component={Payment} />
        <Route path="/statement" component={Statement} />
      </div>
    </div>
  </div>
</Router>

<style>
  button {
    width: 120px;
    height: 40px;
  }
</style>
