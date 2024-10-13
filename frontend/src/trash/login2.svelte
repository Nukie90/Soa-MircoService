<script>
  import { Input, Label, Button } from 'flowbite-svelte';
  import HoneyLemonLogo from '../assets/BankLogo.png';
</script>

<div class="container">
  <div class="login-container">
    <div class="header">
      <h1 class="login-text">Login</h1>
      <img src="{HoneyLemonLogo}" class="logo" alt="HoneyLemonLogo" />
    </div>
    <form class="login-form">
      <div class="input-group">
        <Label for="id_card" class="mb-2 text-[#666666]">ID Card</Label>
        <Input type="text" id="id_card" placeholder="ID Card" required />
      </div>
      <div class="input-group">
        <Label for="password" class="mb-2 text-[#666666]">Password</Label>
        <Input type="password" id="password" placeholder="Password" required />
      </div>
      <div class="forgot-password-container">
        <a href="/forgot-password" class="forgot-password-link">Forgot Password?</a>
      </div>
      <div class="button-container">
        <Button type="submit" class="bg-[#28A745] hover:bg-[#28A745]-100 w-full">Login</Button>
      </div>
    </form>
    <div class="signup-container">
      <a href="/signup" class="signup-link">Create an account</a>
    </div>
  </div>
</div>

<style>
  .container {
    position: relative;
    width: 100%;
    height: 100vh; 
  }
  .login-container {
    width: 400px;
    height: 500px;
    margin: auto;
    margin-top: 100px;
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 8px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    background-color: white;
    position: relative;
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
  .login-text {
    font-weight: bold;
  }
  .login-form {
    display: flex;
    flex-direction: column;
    position: relative;
  }
  .input-group {
    margin-top: 20px;
  }
  .forgot-password-container {
    position: absolute;
    bottom: 100px; /* Adjust as needed */
    right: 0; 
    text-align: right; 
  }
  .forgot-password-link {
    color: #28A745;
    text-decoration: none !important; /* Force removal of underline */
    font-size: 14px;
  }
  .forgot-password-link:hover {
    text-decoration: underline;
  }
  .button-container {
    margin-top: 90px; 
    display: flex;
    justify-content: center;
  }
  .signup-container {
    position: absolute;
    bottom: -40px; 
    left: 50%;
    transform: translateX(-50%);
    text-align: center; 
    width: 100%; 
  }
  .signup-link {
    color: #28A745;
    text-decoration: none;
    font-size: 14px;
  }
</style>