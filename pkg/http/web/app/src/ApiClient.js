import axios from 'axios'

const BASE_URI = 'http://localhost:8081/api';
const client = axios.create({
    baseURL: BASE_URI,
    json: true,
    headers: {
        // Authorization: `Bearer ${accessToken}`
        "Access-Control-Allow-Origin": "*",
        "Content-Type": "application/json"
    }
});

const ApiClient = {
    async perform(method, url, data) {
        return client({
            method,
            url,
            data
        }).then(req => {
            return req.data
        })
    }
}
export default ApiClient;