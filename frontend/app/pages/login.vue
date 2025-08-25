<script setup lang="ts">
import type { FormSubmitEvent } from "@nuxt/ui";
import { reactive } from "vue";
import * as z from "zod";

const auth = useAuth();

const schema = z.object({
  email: z.email("Invalid email"),
});

type Schema = z.infer<typeof schema>;

const state = reactive<Partial<Schema>>({
  email: "",
});

async function onSubmit(event: FormSubmitEvent<Schema>) {
  try {
    await auth.login(event.data.email!);
    console.log("Login successful", auth.user);
    await navigateTo("/");
  } catch (err) {
    console.error("Login failed", err);
  }
}
</script>

<template>
  <div class="flex justify-center items-center min-h-screen bg-gray-50">
    <UCard class="w-full max-w-md p-6">
      <UForm
        :state="state"
        :schema="schema"
        @submit="onSubmit"
        class="space-y-4"
      >
        <UFormField label="Email" name="email">
          <UInput
            class="w-full"
            v-model="state.email"
            type="email"
            placeholder="Enter your email"
          />
        </UFormField>

        <UButton
          type="submit"
          color="primary"
          class="w-full flex justify-center cursor-pointer"
        >
          Login
        </UButton>
      </UForm>
    </UCard>
  </div>
</template>
