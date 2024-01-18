<script setup>
import { computed, ref } from "vue"

const props = defineProps({
  matchingEmails: {
    type: Array,
    required: true
  }
})

const emailSelected = ref(undefined)

const isSelectedd = id => id === emailSelected.value && emailSelected.value !== undefined

const handleEmailSelect = (id) => {
  emailSelected.value = id
}

const emailBody = computed(() => {
  const email = props.matchingEmails.filter(x => x.id === emailSelected?.value)
  return email[0].body
})
</script>

<template>
  <div class="md:flex mt-8 gap-8 bg-white p-4 rounded-md">
    <div class="md:w-1/2 border rounded-xl shadow-sm p-4">
      <table class="table-auto w-full">
        <thead class="rounded-xl">
          <tr class="bg-indigo-500 text-white">
            <th class="font-bold p-2 border-b text-left rounded-tl-xl">
              subject
            </th>
            <th class="font-bold p-2 border-b text-left">From</th>
            <th class="font-bold p-2 border-b text-left rounded-tr-xl">To</th>
          </tr>
        </thead>
        <tbody>
          <tr
            :class="[isSelectedd(email.id) ? 'bg-indigo-300' : 'odd:bg-gray-200']"
            class="hover:bg-indigo-600 hover:text-white cursor-pointer"
            @click="handleEmailSelect(email.id)"
            v-for="email in matchingEmails"
          >
            <td class="w-1/2">
              {{email.subject}}
            </td>
            <td class="w-1/4">{{email.from}}</td>
            <td class="w-1/4">{{email.to}}</td>
          </tr>
        </tbody>
      </table>
    </div>
    <div class="md:w-1/2 border rounded-xl shadow-sm p-4 flex justify-center items-center font-semibold">
      <div v-if="!emailSelected">
        Click in one mail to see details
      </div>
      <div v-else>
        {{emailBody}}
      </div>
    </div>
  </div>
</template>
<CarritoCard v-for="carritoItem in carrito" :carrito-item="carritoItem" @incrementar-cantidad="incrementarCantidad" @decrementar-cantidad="decrementarCantidad" @eliminar-plato="eliminarPlato"/>