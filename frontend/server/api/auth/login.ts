export default defineEventHandler(async (event) => {
  const body = await readBody(event);
  const res = await $fetch<LoginResponse>(
    `${process.env.API_URL}/api/auth/login`,
    {
      method: "POST",
      body,
    }
  );
  return res;
});
