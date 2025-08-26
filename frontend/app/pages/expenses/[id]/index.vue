<script setup lang="ts">
import { useRoute } from "vue-router";

const route = useRoute();

const { data } = await useApi<Expense>(`/api/expenses/${route.params.id}`);

const formattedDate = (value: string) => {
  if (!value) return "-";
  return new Date(value).toLocaleString();
};
</script>

<template>
  <UContainer class="p-4 max-w-lg mx-auto mt-8">
    <nuxt-link
      to="/expenses"
      class="inline-block mb-4 text-primary hover:underline"
    >
      ← Back to Expenses
    </nuxt-link>

    <h2 class="text-xl font-bold mb-4">Expense Details</h2>

    <UCard class="p-4 flex flex-col gap-4">
      <div v-if="data" class="flex flex-col gap-4">
        <UFormField label="Amount">
          <p>
            {{
              new Intl.NumberFormat("id-ID", {
                style: "currency",
                currency: "IDR",
                minimumFractionDigits: 0,
              }).format(data.amount_idr)
            }}
          </p>
        </UFormField>

        <UFormField label="Description">
          <p>{{ data.description }}</p>
        </UFormField>

        <UFormField label="Receipt URL">
          <p>{{ data.receipt_url }}</p>
        </UFormField>

        <UFormField label="Status">
          <p>{{ data.status }}</p>
        </UFormField>

        <UFormField label="Submitted At">
          <p>{{ new Date(data.submitted_at).toLocaleString() }}</p>
        </UFormField>

        <UFormField label="Processed At">
          <p>{{ formattedDate(data.processed_at) }}</p>
        </UFormField>
      </div>

      <div v-else class="text-center text-gray-500">
        Loading expense details...
      </div>
    </UCard>
  </UContainer>
</template>
