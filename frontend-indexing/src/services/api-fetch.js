import axios from "axios";

export const api = axios.create({
  baseURL: 'http://localhost:8080'
})

export const searchEmails = (searchTerm) => {
  // return [
  //   {
  //     id: 1,
  //     subject: "subject 1",
  //     from: "Juan",
  //     to: "Crhis",
  //     body: "some random text"
  //   },
  //   {
  //     id: 2,
  //     subject: "subject 2",
  //     from: "Fer",
  //     to: "Maria",
  //     body: "Todas las mujeres mienten"
  //   },
  //   {
  //     id: 3,
  //     subject: "subject 3",
  //     from: "Jose",
  //     to: "crhis",
  //     body: "Hello"
  //   },
  // ]
  return api.get(`/search?value=${searchTerm}`)
}
