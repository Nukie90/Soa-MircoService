<script>
  import { Input, Label, Button } from 'flowbite-svelte';
  import HoneyLemonLogo from '../assets/BankLogo.png';
</script>

<div class="container">
  <div class="signup-container">
    <div class="header">
      <h1 class="signup-text">Signup</h1>
      <img src="{HoneyLemonLogo}" class="logo" alt="HoneyLemonLogo" />
    </div>
    <form class="signup-form">
      <div class="input-group">
        <Label for="id_card" class="mb-2 text-[#666666] ">ID Card</Label>
        <Input type="text" id="id_card" placeholder="ID Card" required />
      </div>
      <div class="input-group">
        <Label for="password" class="mb-2 text-[#666666]">Password</Label>
        <Input type="password" id="password" placeholder="Password" required />
      </div>
      <div class="input-group">
        <Label for="confirm_password" class="mb-2 text-[#666666]">Confirm Password</Label>
        <Input type="password" id="confirm_password" placeholder="Confirm Password" required />
      </div>
      <div class="button-container">
          <Button type="submit" class="bg-[#28A745] hover:bg-[#28A745]-100 w-full">Signup</Button>
      </div>
    </form>
  </div>
  <div class="login-container">
      <a href="/login" class="login-link">Login</a>
  </div>
</div>

<style>
  .container {
    position: relative;
    width: 100%;
    height: 100vh; 
  }
  .signup-container {
    width: 400px;
    height: 500px;
    margin: auto;
    margin-top: 100px;
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 8px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    background-color: white;
  }
  .header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
  }
  .header h1 {
    font-size: 40px;
  }
  .logo {
    height: 18px;
  }
  .signup-text {
    font-weight: bold;
  }
  .signup-form {
    display: flex;
    flex-direction: column;
  }
  .input-group {
    margin-top: 20px;
  }
  .button-container {
    margin-top: 50px;
    display: flex;
    justify-content: center;
  }
  .login-container {
    position: absolute;
    bottom: 285px;
    left: 50%;
    transform: translateX(-50%);
    text-align: center; 
    width: 100%; 
  }
  .login-link {
    color: #28A745;
    text-decoration: none;
    font-size: 14px;
  }
</style>
