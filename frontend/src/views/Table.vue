<template>
  <h2>Table</h2>

  <table class="content-table">
    <thead>
    <tr>
      <th>Ticker</th>
      <th>Name</th>
      <th>Industry</th>
      <th>Sector</th>
      <th>Employees</th>
      <th>Country</th>
      <th>Total Revenue</th>
      <th>Total Cash</th>
      <th>Total Cash per share</th>
      <th>Ebitda</th>
      <th>Total Debt</th>
      <th></th>

    </tr>
    </thead>
    <tbody>
    <tr>
      <td>1</td>
      <td>Domenic</td>
      <td>88,110</td>
      <td>dcode</td>
      <td>dcodeewdwefw</td>
    </tr>
    <tr class="active-row">
      <td>2</td>
      <td>Sally</td>
      <td>72,400</td>
      <td>Students</td>
      <td>dcodeewdwefw</td>
    </tr>
    <tr>
      <td>3</td>
      <td>Nick</td>
      <td>52,300</td>
      <td>dcode</td>
      <td>dcodeewdwefw</td>
    </tr>
    </tbody>
  </table>

</template>

<script>

export default {
  components: {},
  data() {

  },
  methods: {

  }
}
</script>

<style scoped>
  .stock {
    margin: 0 auto;
    max-width: 600px;
  }
  .red {
    background-color: #f005;
  }
  .green {
    background-color: #0f05;
  }

  * {
    font-family: sans-serif; /* Change your font family */
  }

  .content-table {
    margin: 0 auto;
    border-collapse: collapse;
    /*margin: 25px 0;*/
    font-size: 0.9em;
    min-width: 400px;
    border-radius: 5px 5px 0 0;
    overflow: hidden;
    box-shadow: 0 0 20px rgba(0, 0, 0, 0.15);
  }

  .content-table thead tr {
    background-color: #009879;
    color: #ffffff;
    text-align: left;
    font-weight: bold;
  }

  .content-table th,
  .content-table td {
    padding: 12px 15px;
  }

  .content-table tbody tr {
    border-bottom: 1px solid #dddddd;
  }

  .content-table tbody tr:nth-of-type(even) {
    background-color: #f3f3f3;
  }

  .content-table tbody tr:last-of-type {
    border-bottom: 2px solid #009879;
  }

  .content-table tbody tr.active-row {
    font-weight: bold;
    color: #009879;
  }

</style>