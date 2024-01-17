import axios from "axios";

export const api = axios.create({
  baseURL: 'http://localhost:8080'
})

export const searchEmails = (searchTerm) => {
  return [
    {
      subject: "subject 1",
      from: "Juan",
      to: "Crhis",
      body: "some random text"
    },
    {
      subject: "subject 2",
      from: "Fer",
      to: "Maria",
      body: "Todas las mujeres mienten"
    },
  ]
  //return api.get(`/search?value=${searchTerm}`)
}
