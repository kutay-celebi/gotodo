import axios, { AxiosRequestConfig, AxiosResponse } from 'axios'

const service = axios.create({
  baseURL: 'http://localhost:8080', // todo @kcelebi env variable
  timeout: 30000
})

service.interceptors.request.use((config: AxiosRequestConfig) => {
  config.headers = {
    accept: 'application/json',
    'content-type': 'application/json'
  }
  return config
})

service.interceptors.response.use((value: AxiosResponse) => {
  if (process.env.NODE_ENV === 'development') {
    console.log(value)
  }
  return value
})

export default service
