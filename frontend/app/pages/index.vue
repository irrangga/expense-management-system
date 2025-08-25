<script setup lang="ts">
import type { TableColumn } from "@nuxt/ui";
import { useAuth } from "~/composables/use-auth";

const { logout } = useAuth();

const handleLogout = () => {
  logout();
  navigateTo("/login");
};

const { data } = await useApi<PaginatedResponse<Expense[]>>("/api/expenses");

const columns: TableColumn<Expense>[] = [
  {
    accessorKey: "user_id",
    header: "User ID",
    cell: ({ row }) => row.getValue("user_id"),
  },
  {
    accessorKey: "amount_idr",
    header: "Amount IDR",
    cell: ({ row }) => row.getValue("amount_idr"),
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
    cell: ({ row }) => new Date(row.getValue("processed_at")).toLocaleString(),
  },
];
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
    <UTable :data="data?.data.flat()" :columns="columns" />
  </div>
</template>
