<script setup>
defineEmits(["update:term", "search-term"])

const props = defineProps({
  term: {
    type: String,
    required: true,
  },
})
</script>

<template>
  <header>
    <h1 class="font-black text-5xl text-center">Email Searching</h1>
    <div class="mt-12 lg:flex justify-center">
      <form
        class="bg-white shadow-md rounded-lg py-2 px-3 w-96"
        @submit.prevent="$emit('search-term')"
      >
        <div class="flex gap-4 items-center">
          <input
            id="term"
            type="text"
            class="border-2 w-full p-2 placeholder-gray-400 rounded-md"
            placeholder="Type the term you want to look for..."
            :value="term"
            @input="$emit('update:term', $event.target.value)"
          />
          <input
            type="submit"
            class="cursor-pointer h-fit bg-indigo-500 p-1 text-white uppercase font-bold hover:bg-indigo-700 transition-colors rounded-md"
            value="Buscar"
          />
        </div>
      </form>
    </div>
  </header>
</template>
