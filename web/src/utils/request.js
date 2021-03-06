import axios from 'axios'

var instance = axios.create({
    baseURL: process.env.VUE_APP_API_BASE_URL,
    timeout: 3000
})

export default instance