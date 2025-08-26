<script setup lang="ts">
import type { TableColumn } from "@nuxt/ui";
import { useAuth } from "~/composables/use-auth";

const { user, logout } = useAuth();

const handleLogout = () => {
  logout();
  navigateTo("/login");
};

const handleAddExpense = () => {
  navigateTo("/expenses/new");
};

const page = ref(1);
const pageSize = ref(10);

const url = computed(
  () => `/api/expenses?page=${page.value}&limit=${pageSize.value}`
);

const { data } = await useApi<PaginatedResponse<Expense[]>>(url);

const allColumns: TableColumn<Expense>[] = [
  {
    id: "user_name",
    accessorKey: "user_name",
    header: "Name",
    cell: ({ row }) => row.getValue("user_name"),
  },
  {
    accessorKey: "amount_idr",
    header: "Amount (IDR)",
    cell: ({ row }) => {
      const amount: number = row.getValue("amount_idr");

      return new Intl.NumberFormat("id-ID", {
        style: "currency",
        currency: "IDR",
        minimumFractionDigits: 0,
      }).format(amount);
    },
  },
  {
    accessorKey: "description",
    header: "Description",
  },
  {
    accessorKey: "status",
    header: "Status",
  },
  {
    accessorKey: "submitted_at",
    header: "Submitted At",
    cell: ({ row }) => new Date(row.getValue("submitted_at")).toLocaleString(),
  },
  {
    accessorKey: "processed_at",
    header: "Processed At",
    cell: ({ row }) => {
      const value: string | undefined = row.getValue("processed_at");

      if (!value) return "-";

      return new Date(value).toLocaleString();
    },
  },
  {
    id: "action",
  },
];

const columns: TableColumn<Expense>[] =
  user.value?.role === "employee"
    ? allColumns.filter((col) => col.id !== "user_name")
    : allColumns;
</script>

<template>
  <div class="p-4 space-y-4 flex flex-col">
    <UButton
      color="error"
      @click="handleLogout"
      class="self-end cursor-pointer"
    >
      Logout
    </UButton>

    <UButton
      color="primary"
      @click="handleAddExpense"
      class="flex justify-center cursor-pointer"
    >
      Add New Expense
    </UButton>

    <UTable :data="data?.data" :columns="columns">
      <template #action-cell="{ row }">
        <UButton
          type="submit"
          color="neutral"
          class="cursor-pointer"
          @click="navigateTo(`/expenses/${row.original.id}`)"
        >
          Details
        </UButton>
      </template>
    </UTable>

    <UPagination
      class="self-center cursor-pointer"
      v-model:page="page"
      :total="data?.total"
      :items-per-page="pageSize"
      :show-edges="true"
      :sibling-count="1"
      color="primary"
      size="md"
      :ui="{
        first: 'cursor-pointer',
        prev: 'cursor-pointer',
        item: 'cursor-pointer',
        next: 'cursor-pointer',
        last: 'cursor-pointer',
      }"
    />
  </div>
</template>
