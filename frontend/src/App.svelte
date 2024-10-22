<script>
    import { Router, Route } from "svelte-routing";
    import Content from "./pages/content.svelte";
    import Login from "./pages/login.svelte";
    import Signup from "./pages/signup.svelte";
    import AccountRegister from "./pages/accountregister.svelte";
    import AccountAdd from "./pages/addnewaccount.svelte";
    import AccountSetUsername from "./pages/accountsetusername.svelte";
    import AccountSetPin from "./pages/accountsetpin.svelte";
    import EmailVerification from "./pages/emailverify.svelte";
    import ResetPassword from "./pages/resetpassword.svelte";
    import Account from "./pages/account.svelte";
    import ChangeProfile from "./pages/accountchangeprofile.svelte";
    import ChangePassword from "./pages/accountchangepassword.svelte";
    import ChangePin from "./pages/accountchangepin.svelte";
    import Transfer from "./pages/transfer.svelte";
    import Transfer2 from "./pages/transfer2.svelte";
    import Transfer3 from "./pages/transfer3.svelte";
    import Payment from "./pages/payment.svelte";
    import Payment2 from "./pages/payment2.svelte";
    import Payment3 from "./pages/payment3.svelte";
    import Loan from "./pages/loan.svelte";
    import Investment from "./pages/invest.svelte";
    import InvestmentBuy from "./pages/investmentbuy.svelte";
    import InvestmentSell from "./pages/investmentsell.svelte";
    import Statement from "./pages/statement.svelte";
    import Statement2 from "./pages/statement2.svelte";

    import MainLayout from "./layouts/mainlayout.svelte";
    import AccountLayout from "./layouts/accountlayout.svelte";
    import Account_Settings_MainLayout from "./layouts/accountmainlayout.svelte";
    import Login_Signup_Accregister_Layout from "./layouts/log_sign_accregis.svelte";
</script>

<Router>
    <Route path="/">
        <MainLayout>
            <Content />
        </MainLayout>
    </Route>
    <Route path="/login">
        <Login_Signup_Accregister_Layout>
            <Login />
        </Login_Signup_Accregister_Layout>
    </Route>
    <Route path="/signup">
        <Login_Signup_Accregister_Layout>
            <Signup />
        </Login_Signup_Accregister_Layout>
    </Route>
    <Route path="/accountregister">
        <Login_Signup_Accregister_Layout>
            <AccountRegister />
        </Login_Signup_Accregister_Layout>
    </Route>
    <Route path="/addaccount">
        <Account_Settings_MainLayout>
            <AccountLayout>
              <AccountAdd />
            </AccountLayout>
        </Account_Settings_MainLayout>
    </Route>
    <Route path="/changeprofile">
        <Account_Settings_MainLayout>
            <AccountLayout>
              <ChangeProfile />
            </AccountLayout>
        </Account_Settings_MainLayout>
    </Route>
    <Route path="/changepassword">
        <Account_Settings_MainLayout>
            <AccountLayout>
              <ChangePassword />
            </AccountLayout>
        </Account_Settings_MainLayout>
    </Route>
    <Route path="/changepin">
        <Account_Settings_MainLayout>
            <AccountLayout>
              <ChangePin />
            </AccountLayout>
        </Account_Settings_MainLayout>
    </Route>
    <Route path="/accountsetusername">
        <Login_Signup_Accregister_Layout>
            <AccountSetUsername />
        </Login_Signup_Accregister_Layout>
    </Route>
    <Route path="/accountsetpin">
        <Login_Signup_Accregister_Layout>
            <AccountSetPin />
        </Login_Signup_Accregister_Layout>
    </Route>
    <Route path="/emailverification">
        <Login_Signup_Accregister_Layout>
            <EmailVerification />
        </Login_Signup_Accregister_Layout>
    </Route>
    <Route path="/resetpassword">
        <Login_Signup_Accregister_Layout>
            <ResetPassword />
        </Login_Signup_Accregister_Layout>
    </Route>
    <Route path="/mainaccount">
        <Account_Settings_MainLayout>
            <AccountLayout>
                <Account />
            </AccountLayout>
        </Account_Settings_MainLayout>
    </Route>
    <Route path="/transfer">
        <Account_Settings_MainLayout>
            <AccountLayout>
                <Transfer />
            </AccountLayout>
        </Account_Settings_MainLayout>
    </Route>
    <Route path="/transfer2">
        <Account_Settings_MainLayout>
            <AccountLayout>
                <Transfer2 />
            </AccountLayout>
        </Account_Settings_MainLayout>
    </Route>
    <Route path="/transfer3">
        <Account_Settings_MainLayout>
            <AccountLayout>
                <Transfer3 />
            </AccountLayout>
        </Account_Settings_MainLayout>
    </Route>
    <Route path="/payment">
        <Account_Settings_MainLayout>
            <AccountLayout>
                <Payment />
            </AccountLayout>
        </Account_Settings_MainLayout>
    </Route>
    <Route path="/payment2">
        <Account_Settings_MainLayout>
            <AccountLayout>
                <Payment2 />
            </AccountLayout>
        </Account_Settings_MainLayout>
    </Route>
    <Route path="/payment3">
        <Account_Settings_MainLayout>
            <AccountLayout>
                <Payment3 />
            </AccountLayout>
        </Account_Settings_MainLayout>
    </Route>
    <Route path="/loan">
        <Account_Settings_MainLayout>
            <AccountLayout>
                <Loan />
            </AccountLayout>
        </Account_Settings_MainLayout>
    </Route>
    <Route path="/investment">
        <Account_Settings_MainLayout>
            <AccountLayout>
                <Investment />
            </AccountLayout>
        </Account_Settings_MainLayout>
    </Route>
    <Route path="/investmentbuy">
        <Account_Settings_MainLayout>
            <AccountLayout>
                <InvestmentBuy />
            </AccountLayout>
        </Account_Settings_MainLayout>
    </Route>
    <Route path="/investmentsell">
        <Account_Settings_MainLayout>
            <AccountLayout>
                <InvestmentSell />
            </AccountLayout>
        </Account_Settings_MainLayout>
    </Route>
    <Route path="/statement">
        <Account_Settings_MainLayout>
            <AccountLayout>
                <Statement />
            </AccountLayout>
        </Account_Settings_MainLayout>
    </Route>
    <Route path="/statement2">
        <Account_Settings_MainLayout>
            <AccountLayout>
                <Statement2 />
            </AccountLayout>
        </Account_Settings_MainLayout>
    </Route>
</Router>