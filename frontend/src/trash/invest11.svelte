<!-- src/components/InvestmentChart.svelte -->
<script>
  import { Line } from 'svelte-chartjs';
  import 'chart.js/auto';

  const labels = ['January', 'February', 'March', 'April', 'May', 'June', 'July'];
  
  const data = {
    labels,
    datasets: [
      {
        label: 'Apple',
        data: [65, 59, 80, 81, 56, 55, 40],
        fill: false,
        borderColor: 'rgb(75, 192, 192)',
        tension: 0.1,
      },
      {
        label: 'Tesla',
        data: [45, 79, 60, 91, 46, 35, 60],
        fill: false,
        borderColor: 'rgb(192, 75, 75)',
        tension: 0.1,
      },
      {
        label: 'Google',
        data: [30, 60, 70, 80, 90, 100, 110],
        fill: false,
        borderColor: 'rgb(75, 192, 75)',
        tension: 0.1,
      },
    ],
  };

  const options = {
    responsive: true,
    plugins: {
      legend: {
        position: 'top',
      },
      title: {
        display: true,
        text: 'Investment Growth Over Time',
      },
    },
  };
</script>

<Line {data} {options} />
