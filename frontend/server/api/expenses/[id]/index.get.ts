export default defineEventHandler(async (event) => {
  const authHeader = getHeader(event, "authorization");

  const { id } = event.context.params as { id: string };

  const response = await $fetch<Expense>(
    `${process.env.API_URL}/api/expenses/${id}`,
    {
      headers: {
        Authorization: authHeader || "",
      },
    }
  );

  return response;
});
