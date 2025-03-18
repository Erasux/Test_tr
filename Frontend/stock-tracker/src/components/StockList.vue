<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useStockStore } from "../stores/stocks";

const store = useStockStore();
const searchTicker = ref("");
const searchCompany = ref("");
const searchBrokerage = ref("");

// Load stocks when the component mounts
onMounted(() => {
  store.fetchStocks();
});

// Apply filters
const applyFilters = () => {
  const filters: Record<string, string> = {};
  if (searchTicker.value) filters["ticker"] = searchTicker.value;
  if (searchCompany.value) filters["company"] = searchCompany.value;
  if (searchBrokerage.value) filters["brokerage"] = searchBrokerage.value;
  
  store.fetchStocks(filters);
};
</script>

<template>
  <div class="container mt-4">
    <h1 class="text-center">📊 Stock Tracker</h1>

    <div class="row my-4">
      <div class="col-md-3">
        <input v-model="searchTicker" placeholder="Search by Ticker" class="form-control" />
      </div>
      <div class="col-md-3">
        <input v-model="searchCompany" placeholder="Search by Company" class="form-control" />
      </div>
      <div class="col-md-3">
        <input v-model="searchBrokerage" placeholder="Search by Brokerage" class="form-control" />
      </div>
      <div class="col-md-3">
        <button @click="applyFilters" class="btn btn-primary w-100">Filter</button>
      </div>
    </div>

    <div v-if="store.loading" class="text-center">Loading data...</div>
    <div v-else-if="store.error" class="text-center text-danger">{{ store.error }}</div>
    
    <table v-else class="table table-bordered">
      <thead class="table-light">
        <tr>
          <th>Ticker</th>
          <th>Company</th>
          <th>Brokerage</th>
          <th>Target Price</th>
          <th>Rating</th>
          <th>Action</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="stock in store.stocks" :key="stock.id">
          <td>{{ stock.ticker }}</td>
          <td>{{ stock.company }}</td>
          <td>{{ stock.brokerage }}</td>
          <td>${{ stock.target_from }} → ${{ stock.target_to }}</td>
          <td>{{ stock.rating_from }} → {{ stock.rating_to }}</td>
          <td>{{ stock.action }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
