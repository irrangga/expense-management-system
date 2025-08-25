<script setup lang="ts">
import { navigateTo } from "#app";
import { reactive } from "vue";

const state = reactive({
  amount_idr: null as number | null,
  description: "",
  receipt_url: "",
});

const submitExpense = async () => {
  if (!state.amount_idr || !state.description || !state.receipt_url) return;

  try {
    await useApi<Expense>("/api/expenses", {
      method: "POST",
      body: {
        amount_idr: state.amount_idr,
        description: state.description,
        receipt_url: state.receipt_url,
      },
    });

    alert("Expense submitted successfully!");
    navigateTo("/expenses");
  } catch (error) {
    alert("Failed to submit expense.");
  }
};
</script>

<template>
  <div class="p-4 max-w-md mx-auto mt-8">
    <h2 class="text-xl font-bold mb-4">Add Expense</h2>

    <form @submit.prevent="submitExpense" class="flex flex-col gap-4">
      <UFormField label="Amount" name="amount-idr">
        <UInput
          class="w-full"
          v-model.number="state.amount_idr"
          type="number"
          placeholder="Enter amount"
          required
        />
      </UFormField>

      <UFormField label="Description" name="description">
        <UInput
          class="w-full"
          v-model="state.description"
          type="text"
          placeholder="Enter description"
          required
        />
      </UFormField>

      <UFormField label="Receipt URL" name="receipt-url">
        <UInput
          class="w-full"
          v-model="state.receipt_url"
          type="text"
          placeholder="Enter receipt URL"
          required
        />
      </UFormField>

      <UButton
        type="submit"
        color="primary"
        class="w-full flex justify-center cursor-pointer"
      >
        Submit
      </UButton>
    </form>
  </div>
</template>
