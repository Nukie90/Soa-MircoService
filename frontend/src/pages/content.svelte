<!-- src/pages/Home.svelte -->
<script>
  // Home-specific logic
</script>

<main>
  <div class="bg-[#F0F0F0]">
    <h1 class="mt-80">Your trusted partner in financial services</h1>
  </div>
  <div class="bg-[#F0F0F0]">
    <h3 class="mt-36">Our services</h3>
  </div>
  <div class="bg-[#F0F0F0] flex justify-around py-5">
    <div class="box bg-white p-5 rounded-lg">
        <h2>Transfer</h2>
        <p>Transfer funds easily between accounts or to other banks with our secure transfer services.</p>
    </div>
    <div class="box bg-white p-5 rounded-lg">
      <h2>Payment</h2>
      <p>Get personal loans at attractive interest rates with flexible repayment options.</p>
    </div>
    <div class="box bg-white p-5 rounded-lg">
      <h2>Personal loans</h2>
      <p>Pay your telephone, water, and electricity bills easily and securely through our platform.</p>
    </div>
    <div class="box bg-white p-5 rounded-lg">
      <h2>Investment</h2>
      <p>Invest in a wide range of financial products to grow your wealth.</p>
    </div>
    <div class="box bg-white p-5 rounded-lg">
      <h2>Statement</h2>
      <p>Access and download your account statements anytime for easy tracking of your financial activities.</p>
    </div>
  </div>
  <div class="bg-[#28A745] p-5">
    <h4>© 2024 HoneyLemon Bank. All rights reserved.</h4>
  </div>
</main>

<style>
  .box {
    width: 18%; /* Adjusted to fit five boxes */
    height: 250px;
    padding: 20px;
  }
  .box h2 {
    font-size: 22px;
    color: #28A745;
    text-align: center;
    margin-top: 10px;
  }
  .box p {
    font-size: 16px;
    color: #000000;
    text-align: center;
    margin-top: 50px;
  }
</style>
