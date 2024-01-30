import axios from "axios";

export const api = axios.create({
  baseURL: 'http://localhost:8080/v1'
})

export const searchEmails = (searchTerm) => {
  return api.get(`/search?value=${searchTerm}`)
}
