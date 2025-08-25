export default defineEventHandler(async (event) => {
  const authHeader = getHeader(event, "authorization");

  const response: PaginatedResponse<Expense[]> = await $fetch(
    `${process.env.API_URL}/api/expenses`,
    {
      headers: {
        Authorization: authHeader || "",
      },
    }
  );

  return response;
});
