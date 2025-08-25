export default defineEventHandler(async (event) => {
  const authHeader = getHeader(event, "authorization");

  const body = await readBody(event);

  if (!body.amount_idr || !body.description) {
    throw createError({ statusCode: 400, statusMessage: "Missing fields" });
  }

  const response = await $fetch<Expense>(
    `${process.env.API_URL}/api/expenses`,
    {
      method: "POST",
      headers: {
        Authorization: authHeader || "",
        "Content-Type": "application/json",
      },
      body,
    }
  );

  return response;
});
