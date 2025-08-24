<script setup lang="ts">
import type { TableColumn } from "@nuxt/ui";

const { data } = await useFetch<PaginatedResponse<Expense[]>>("/api/expenses");

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
  <UTable :data="data?.data.flat()" :columns="columns" />
</template>
