<script>
    import { Link } from "svelte-routing";
    import { Card, Button, Input, Select } from 'flowbite-svelte';
    import HoneyLemonLogo from '../assets/BankLogo.png';
    import BackButton from '../assets/back.png';
    import CancelButton from '../assets/cancel.png';
    import ConfirmButton from '../assets/confirm.png';
</script>

<div class="flex flex-col items-center">
  <Card class="bg-white mt-24 mb-5" size="2xl" padding="lg" style="width: 1400px; height: 770px;">
    <div class="flex flex-col h-full">
        <div class="flex items-center justify-between">
            <Link to="/mainaccount">
                <img src="{BackButton}" class="h-4 w-4" alt="Back Button" />
            </Link>
            <img src="{HoneyLemonLogo}" class="h-18 w-28" alt="HoneyLemonLogo" />
        </div>
        <div class="text-lg text-black mt-4 ml-4 mb-3">
            Fill in the Information
        </div>
        <div class="flex flex-col gap-4 ml-4">
            <!-- Personal Information Card -->
            <Card class="bg-white border-black" size="2xl" padding="lg" style="width: 1300px;">
                <div class="flex flex-wrap h-full">
                    <div class="text-xs text-black mb-2 mr-4 w-1/4">
                        Full Name
                        <Input type="text" name="fullname" placeholder="Full Name" required class="w-full mt-1 text-xs"/>
                    </div>
                    <div class="text-xs text-black mb-2 mr-4 w-1/4">
                        Date of Birth
                        <Input type="date" name="dob" placeholder="yyyy-mm-dd" required class="w-full mt-1 text-xs"/>
                    </div>
                    <div class="text-xs text-black mb-2 mr-4 w-1/4">
                        Gender
                        <Select name="gender" class="w-full mt-1 text-xs">
                            <option value="annually">Male</option>
                            <option value="quarterly">Female</option>
                            <option value="others">Others</option>
                        </Select>
                    </div>
                    <div class="text-xs text-black mb-2 mr-4 w-1/4">
                        Tel.
                        <Input type="tel" name="tel" placeholder="Telephone Number" required class="w-full mt-1 text-xs"/>
                    </div>
                    <div class="text-xs text-black mb-2 w-1/2">
                        Address
                        <Input type="text" name="address" placeholder="Address" required class="w-full mt-1 text-xs"/>
                    </div>
                </div>
            </Card>
            
            <!-- Financial Information Card -->
            <Card class="bg-white border-black" size="2xl" padding="lg" style="width: 1300px;">
                <div class="flex flex-wrap h-full">
                    <div class="text-xs text-black mb-2 mr-4 w-1/6">
                        Employment Status
                        <Input type="text" name="employment_status" placeholder="Employment Status" required class="w-full mt-1 text-xs"/>
                    </div>
                    <div class="text-xs text-black mb-2 mr-4 w-1/6">
                        Employer Name
                        <Input type="text" name="employer_name" placeholder="Employer Name" required class="w-full mt-1 text-xs"/>
                    </div>
                    <div class="text-xs text-black mb-2 mr-4 w-1/6">
                        Company Name
                        <Input type="text" name="company_name" placeholder="Company Name" required class="w-full mt-1 text-xs"/>
                    </div>
                    <div class="text-xs text-black mb-2 mr-4 w-1/6">
                        Monthly Income
                        <Input type="text" name="monthly_income" placeholder="Monthly Income" required class="w-full mt-1 text-xs"/>
                    </div>
                    <div class="text-xs text-black mb-2 w-1/6">
                        Additional Sources of Income
                        <Input type="text" name="additional_income" placeholder="Additional Sources of Income" required class="w-full mt-1 text-xs"/>
                    </div>
                </div>
            </Card>
            
            <!-- Loan Details Card -->
            <Card class="bg-white border-black mb-3" size="2xl" padding="lg" style="width: 1300px;">
                <div class="flex flex-wrap h-full">
                    <div class="text-xs text-black mb-2 mr-4 w-1/6">
                        Loan Amount
                        <Input type="text" name="loan_amount" placeholder="Loan Amount" required class="w-full mt-1 text-xs"/>
                    </div>
                    <div class="text-xs text-black mb-2 mr-4 w-1/6">
                        Loan Type
                        <Select name="loan_type" class="w-full mt-1 text-xs">
                            <option value="personal">Personal</option>
                            <option value="emergency">Emergency</option>
                            <option value="others">Others</option>
                        </Select>
                    </div>
                    <div class="text-xs text-black mb-2 mr-4 w-1/6">
                        Purpose of Loan
                        <Input type="text" name="purpose_loan" placeholder="Purpose of Loan" required class="w-full mt-1 text-xs"/>
                    </div>
                    <div class="text-xs text-black mb-2 mr-4 w-1/6">
                        Loan Term
                        <Select name="loan_term" class="w-full mt-1 text-xs">
                            <option value="6_months">6 Months</option>
                            <option value="1_year">1 Year</option>
                            <option value="2_years">2 Years</option>
                            <option value="others">Others</option>
                        </Select>
                    </div>
                    <div class="text-xs text-black mb-2 w-1/6">
                        Repayment Frequency
                        <Select name="repayment_frequency" class="w-full mt-1 text-xs">
                            <option value="annually">Annually</option>
                            <option value="quarterly">Quarterly</option>
                            <option value="others">Others</option>
                        </Select>
                    </div>
                </div>
            </Card>
            
            <!-- Supporting Documents Card -->
            <Card class="bg-white border-black mb-3" size="2xl" padding="lg" style="width: 1300px;">
                <div class="flex flex-wrap h-full">
                    <div class="text-xs text-black mb-2 mr-4 w-1/4">
                        Proof of Id Card
                        <div class="relative inline-block w-full h-10 cursor-pointer mt-1">
                            <input type="file" name="proof_identity" required class="absolute top-0 left-0 w-full h-full opacity-0 cursor-pointer"/>
                            <div class="bg-white border border-gray-300 rounded-lg h-full w-full flex items-center justify-center">
                                <span class="text-sm text-gray-600">Choose file</span>
                            </div>
                        </div>
                    </div>
                    <div class="text-xs text-black mb-2 mr-4 w-1/4">
                        Proof of Income
                        <div class="relative inline-block w-full h-10 cursor-pointer mt-1">
                            <input type="file" name="proof_income" required class="absolute top-0 left-0 w-full h-full opacity-0 cursor-pointer"/>
                            <div class="bg-white border border-gray-300 rounded-lg h-full w-full flex items-center justify-center">
                                <span class="text-sm text-gray-600">Choose file</span>
                            </div>
                        </div>
                    </div>
                    <div class="text-xs text-black mb-2 w-1/4">
                        Proof of Address
                        <div class="relative inline-block w-full h-10 cursor-pointer mt-1">
                            <input type="file" name="proof_address" required class="absolute top-0 left-0 w-full h-full opacity-0 cursor-pointer"/>
                            <div class="bg-white border border-gray-300 rounded-lg h-full w-full flex items-center justify-center">
                                <span class="text-sm text-gray-600">Choose file</span>
                            </div>
                        </div>
                    </div>
                </div>
            </Card>
        </div>
        <div class="flex items-center justify-between">
            <div class="flex items-center justify-between ml-4">
                <Link to="/mainaccount">
                    <img src="{CancelButton}" class="h-10 w-10 mt-2 " alt="HoneyLemonLogo" />
                </Link>
                <span class="text-black ml-1 mt-2">Clear</span>
            </div>
            <div class="flex items-center justify-between mr-4">
                <span class="text-black mr-1 mt-2">Confirm</span>
                <Link to="/mainaccount">
                    <img src="{ConfirmButton}" class="h-10 w-10 mt-2" alt="HoneyLemonLogo" />
                </Link>
            </div>
        </div>
    </div>
  </Card>
</div>
