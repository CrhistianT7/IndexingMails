import axios from "axios";

// export type Email = {
//   id: number
//   message_id: string
//   date: string
//   from: string
//   to: string
//   subject: string
// }

export const api = axios.create({
  baseURL: 'http://localhost:8080'
})

export const searchEmails = (searchTerm) => {
  if (!searchTerm)
  return api.get(`/search?value=${searchTerm}`)
}
