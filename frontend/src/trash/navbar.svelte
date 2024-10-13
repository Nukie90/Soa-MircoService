<div class="navbar bg-[#F0F0F0] h-24 flex items-center justify-between px-5 fixed top-0 left-0 right-0">
    <div class="flex flex-row items-center">
      <div class="mr-5">
        <img src="{HoneyLemonLogo}" class="Logo" alt="HoneyLemonLogo" />
      </div>
      <div class="flex space-x-5">
        <a href="https://pnpm.io/installation" class="text-black hover:text-green-500 no-underline">Personal loans</a>
        <a href="https://pnpm.io/installation" class="text-black hover:text-green-500 no-underline">Payment services</a>
        <a href="https://pnpm.io/installation" class="text-black hover:text-green-500 no-underline">Investment</a>
      </div>
    </div>
    <div class="flex space-x-3 items-center">
      <Link to="/login">
        <button class="LoginButton px-4 py-2 rounded">Log in</button>
      </Link>
      <button class="SignupButton px-4 py-2 rounded">Sign up</button>
    </div>
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
  .SignupButton {
    background-color: #28A745; 
    color: #ffffff;
    font-size: 16px;
  }
  .SignupButton:hover {
    opacity: 70%;
  }
</style>