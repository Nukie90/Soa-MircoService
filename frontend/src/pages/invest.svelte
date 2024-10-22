<script>
  import { Link } from "svelte-routing";
  import { Card, Button, Input, Select } from 'flowbite-svelte';
  import HoneyLemonLogo from '../assets/BankLogo.png';
  import BackButton from '../assets/back.png';
  import CancelButton from '../assets/cancel.png';
  import ConfirmButton from '../assets/confirm.png';

  import { Pie } from 'svelte-chartjs';
  import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js';

  import { Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell, TableSearch } from 'flowbite-svelte';

  ChartJS.register(ArcElement, Tooltip, Legend);

  const data = {
    labels: ['Com7', 'Tesla', 'Apple', 'Amazon'],
    datasets: [
      {
        data: [50, 300, 250, 1000],
        backgroundColor: ['red', 'blue', 'green', 'pink'],
        hoverOffset: 4,
      },
    ],
  };

  const options = {
    responsive: true,
    plugins: {
      legend: {
        display: false,
        position: 'top',
      },
      title: {
        display: true,
        text: 'Chart.js Pie Chart',
      },
    },
  };

  let searchTerm = '';
  let items = [
    { id: 1, name: 'Com7', volume: 50, vp1: 15.00, marget: 750.00},
    { id: 2, name: 'Tesla', volume: 300, vp1: 6.00, marget: 1800.00},
    { id: 3, name: 'Apple', volume: 250, vp1: 7.00, marget: 1750.00},
    { id: 4, name: 'Amazon', volume: 1000, vp1: 10.00,  marget: 10000.00}
  ];
  $: filteredItems = items.filter((item) => item.name.toLowerCase().indexOf(searchTerm.toLowerCase()) !== -1);

  export let test = '11';
</script>

<div class="flex flex-col items-center">
  <Card class="bg-white mt-24 mb-5" size="2xl" padding="lg" style="width: 1400px; height: 670px;">
    <div class="flex flex-col h-full">
      <div class="flex items-center justify-between">
        <Link to="/mainaccount">
          <img src="{BackButton}" class="h-4 w-4" alt="Back Button" />
        </Link>
        <img src="{HoneyLemonLogo}" class="h-18 w-28" alt="HoneyLemonLogo" />
      </div>
      <div class="flex flex-col gap-4 ml-4 mt-6">
        <div class="flex items-center justify-between">
          <Card class="bg-white border-black items-center" size="xl" padding="xs" style="width: 350px;">
            <div class="text-black">
              Total Cost
            </div>
            <div class="text-black">
              {test}
            </div>
          </Card>
          <Card class="bg-white border-black items-center" size="xl" padding="xs" style="width: 350px;">
            <div class="text-black">
              Total Asset
            </div>
            <div class="text-black">
              {test}
            </div>
          </Card>
          <Card class="bg-white border-black items-center" size="xl" padding="xs" style="width: 350px;">
            <div class="text-black">
              P/L
            </div>
            <div class="text-black">
              {test}
            </div>
          </Card>
        </div>
        <div class="flex items-start justify-between h-full gap-4">
          <Card class="bg-white items-center border-black flex-grow h-full" size="xl" padding="md">
            <div class="text-black mb-4">
              My Portfolio
            </div>
            <div class="w-5/6 h-5/6 flex items-center justify-center">
              <Pie {data} {options} />
            </div>
          </Card>
          <Card class="bg-white border-black mb-3 flex-grow w-full h-full" size="xl" padding="md">
            <TableSearch placeholder="Search by Name" hoverable={true} bind:inputValue={searchTerm}>
              <TableHead>
                <TableHeadCell>#</TableHeadCell>
                <TableHeadCell>Name</TableHeadCell>
                <TableHeadCell>Volume</TableHeadCell>
                <TableHeadCell>Buy Value</TableHeadCell>
                <TableHeadCell>Buy Price</TableHeadCell>
                <TableHeadCell>Marget Value</TableHeadCell>
                <TableHeadCell>Marget Price</TableHeadCell>
                <TableHeadCell>P/L</TableHeadCell>
              </TableHead>
              <TableBody tableBodyClass="divide-y">
                {#each filteredItems as item}
                  <TableBodyRow>
                    <TableBodyCell>{item.id}</TableBodyCell>
                    <TableBodyCell>{item.name}</TableBodyCell>
                    <TableBodyCell>{item.volume}</TableBodyCell>
                    <TableBodyCell>{item.vp1}</TableBodyCell>
                    <TableBodyCell>{item.marget}</TableBodyCell>
                  </TableBodyRow>
                {/each}
              </TableBody>
            </TableSearch>
          </Card>
        </div>
        <div class="flex items-center justify-between mt-5">
          <div class="ml-auto flex items-center justify-between gap-4 mr-4">
            <Link to="/investmentbuy">
              <Button type="submit" class="bg-green-500 hover:bg-green-600 w-24">Buy</Button>
            </Link>
            <Link to="/investmentsell">
              <Button type="submit" class="bg-red-400 hover:bg-red-700 w-24">Sell</Button>
            </Link>
          </div>
        </div>        
      </div>
    </div>
  </Card>
</div>