export default defineEventHandler(async () => {
  const expenses: PaginatedResponse<Expense[]> = await $fetch(
    `${process.env.API_URL}/api/expenses`
  );
  return expenses;
});
