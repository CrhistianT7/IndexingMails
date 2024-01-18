<script setup>
import { onMounted, ref } from "vue"
import axios, { AxiosError } from "axios"

import Header from "../components/Header.vue"
import Results from "../components/Results.vue"
import { searchEmails } from "../services/api-fetch"

const loading = ref(false)
//
const termToSearch = ref({
  term: "",
})
const matchingEmails = ref([])

// onMounted(() => {
//   matchingEmails.value = searchEmails("algo")
// })

const searchTerm = async () => {
  console.log(termToSearch.value)
  console.log("searching for this thing")
  try {
    const response = await searchEmails()
    //const response = await axios.get('https://jsonplaceholder.org/users/1')
    console.log(response)
    //matchingEmails.value = response
  } catch (error) {
    if (axios.isAxiosError(error)) {
      console.log("API error: ", error.message)
    } else {
      console.log("Unknown error: ", error)
    }
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <Header v-model:term="termToSearch.term" @search-term="searchTerm" />
  <Results v-if="matchingEmails.length" :matching-emails="matchingEmails" />
  <div v-else class="mt-8 gap-8 bg-white p-4 rounded-md font-semibold shadow-md text-center">No matching emails</div>
</template>
